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
var resetDatabase, seedDatabase *bool

func init() {
	// set up flags
	resetDatabase = flag.Bool("resetDatabase", false, "reset the dev database")
	seedDatabase = flag.Bool("seedDatabase", false, "seed the dev database")
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

	userService := &user.UserService{}
	reviewService := &pr.ReviewService{}

	// this just lets us safely ctrl-c out of the app while running easily
	// runs in a goroutine and just listens for an interrupt
	go func() {
		sig := <-sigs
		log.Println(sig)
		log.Println("Shutting down")
		userService.Stop()
		reviewService.Stop()
		os.Exit(0)
	}()

	// TODO: swap this whole bit out for setting up an interface
	log.Println("Connecting to database")
	// You'd want to replace this with an env variable for prod, and use SSL
	connStr := "dbname=paypay sslmode=disable"
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

	// setup the service
	userService.Start(db, tokenAuth)
	reviewService.Start(db)

	// set up api routes
	router := getRouter(tokenAuth, userService, reviewService)
	// serve http routes, you'd want to set up local certs and https for security though
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func getRouter(auth *jwtauth.JWTAuth, userService *user.UserService, reviewService *pr.ReviewService) http.Handler {
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
		r.Post("/login", userService.Login)
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth))
			r.Use(jwtauth.Authenticator)
			r.Get("/all", userService.All)
			r.Get("/me", userService.CurrentUser)
			r.Get("/{id}", userService.Get)
		})
		// admin only stuff
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth))
			r.Use(adminAuthenticator)
			r.Post("/{id}", userService.Update)
			r.Delete("/{id}", userService.Delete)
			r.Post("/", userService.Create)
		})
	})

	// TODO: wire up reviewService
	r.Route("/review", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth))
			r.Use(jwtauth.Authenticator)
			// r.Get("/me", reviewService.MyFeedback)
			r.Get("/{id}", reviewService.Get)
			// r.Post("/{id}", reviewService.Update)
		})

		// admin only stuff
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth))
			r.Use(adminAuthenticator)
			r.Get("/all", reviewService.All)
			// r.Get("/all/{id}", reviewService.AllForUser)
			r.Post("/{id}", reviewService.Update)
			r.Post("/", reviewService.Create)
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
