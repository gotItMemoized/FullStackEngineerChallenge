package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gotItMemoized/FullStackEngineerChallenge/backend/pr"

	"github.com/gotItMemoized/FullStackEngineerChallenge/backend/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// tokenAuth checker
var tokenAuth *jwtauth.JWTAuth

// flags
var resetDatabase, seedDatabase, usePostgres *bool

func init() {
	// set up flags
	resetDatabase = flag.Bool("resetPostgres", false, "reset the dev database")
	seedDatabase = flag.Bool("seedDatabase", false, "seed the dev database")
	usePostgres = flag.Bool("usePostgres", false, "use a postgres database")
	flag.Parse()

	// set up JWT auth
	secret := os.Getenv("JWT_SECRET")
	if len(secret) == 0 {
		log.Print("Could not load JWT secret so we'll use a default")
		secret = "thisIsAnInsecureSecret"
	}
	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	userHandler := &user.UserHandler{Auth: tokenAuth}
	reviewHandler := &pr.ReviewHandler{}

	// this just lets us safely ctrl-c out of the app while running easily
	// runs in a goroutine and just listens for an interrupt
	go func() {
		sig := <-sigs
		log.Println(sig)
		log.Println("Shutting down")
		userHandler.Data.Stop()
		reviewHandler.Data.Stop()
		os.Exit(0)
	}()

	// setup the services
	us, rs := setupDatasource()
	userHandler.Data = *us
	reviewHandler.Data = *rs

	userHandler.Data.Start()
	reviewHandler.Data.Start()

	// set up api routes
	router := getRouter(tokenAuth, userHandler, reviewHandler)
	// serve http routes, you'd want to set up local certs and https for security though
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func getRouter(auth *jwtauth.JWTAuth, userHandler *user.UserHandler, reviewHandler *pr.ReviewHandler) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	if os.Getenv("ENV") != "DEV" {
		// this is good, but annoying when debugging
		r.Use(middleware.Timeout(2 * time.Second))
	}

	r.Route("/user", func(r chi.Router) {
		r.Post("/login", userHandler.Login)
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth))
			r.Use(jwtauth.Authenticator)
			r.Get("/all", userHandler.All)
			r.Get("/{id}", userHandler.Get)
		})
		// admin only stuff
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth))
			r.Use(adminAuthenticator)
			r.Put("/{id}", userHandler.Update)
			r.Delete("/{id}", userHandler.Delete)
			r.Post("/", userHandler.Create)
		})
	})

	r.Route("/review", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth))
			r.Use(jwtauth.Authenticator)
			r.Get("/{id}", reviewHandler.Get)
		})

		// admin only stuff
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth))
			r.Use(adminAuthenticator)
			r.Get("/all", reviewHandler.All)
			r.Put("/{id}", reviewHandler.Update)
			r.Post("/", reviewHandler.Create)
		})
	})

	r.Route("/feedback", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth))
			r.Use(jwtauth.Authenticator)
			r.Get("/all", reviewHandler.GetPendingFeedbackForReviewer)
			r.Get("/{id}", reviewHandler.GetFeedback)
			r.Put("/{id}", reviewHandler.GiveFeedback)
		})
	})

	return r
}

// Spin on jwtauth.Authenticator middleware to include an isAdmin check on the jwt claim
func adminAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		isAdmin, hasIsAdminClaim := claims["isAdmin"]
		if token == nil || !token.Valid || !(hasIsAdminClaim && isAdmin == true) {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

// does the postgres connection or sets up an in-memory 'database'
func setupDatasource() (*user.Data, *pr.Data) {
	var userData user.Data
	var reviewData pr.Data
	if *usePostgres {
		log.Println("Connecting to database")
		pqConn := os.Getenv("POSTGRES_CONNECTION")
		var connStr string
		if len(pqConn) != 0 {
			connStr = pqConn
		} else {
			log.Println("Connection string not found, using default")
			// You'd want to replace this with an env variable for prod, and use SSL
			connStr = "dbname=paypay sslmode=disable"
		}
		db, err := sqlx.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(" - Success")

		// database set up based on flag
		if (*resetDatabase) && os.Getenv("ENV") == "DEV" {
			log.Println("Resetting database")
			// load the .sql files
			clearFile, err := ioutil.ReadFile("./sql/clear.sql")
			if err != nil {
				log.Fatal(err)
			}
			initFile, err := ioutil.ReadFile("./sql/init.sql")
			if err != nil {
				log.Fatal(err)
			}
			// run the .sql files
			_, err = db.Exec(string(clearFile))
			if err != nil {
				log.Fatal(err)
			}
			_, err = db.Exec(string(initFile))
			if err != nil {
				log.Fatal(err)
			}
			log.Println(" - Success")
		}

		// seed with initial data if we have the seedDatabase flag
		if (*seedDatabase) && os.Getenv("ENV") == "DEV" {
			log.Println("Seeding database")
			// load the sql file
			seedFile, err := ioutil.ReadFile("./sql/seed.sql")
			if err != nil {
				log.Fatal(err)
			}
			// run it
			_, err = db.Exec(string(seedFile))
			if err != nil {
				log.Fatal(err)
			}
			log.Println(" - Success")
		}
		userData = &user.SqlData{DB: db}
		reviewData = &pr.SqlData{DB: db}
	} else {
		log.Println("Setting up an in-memory database")
		us := &user.MapData{Seed: true}
		userData = us
		reviewData = &pr.MapData{UserData: us}
		log.Println(" - Success")
	}

	if userData == nil || reviewData == nil {
		log.Fatal("Could not start up without services")
	}
	return &userData, &reviewData
}
