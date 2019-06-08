package pr

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gotItMemoized/FullStackEngineerChallenge/backend/user"

	"github.com/go-chi/chi"
)

func TestReviewHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		endpoint       string
		expected       string
		expectedStatus int
	}{
		{name: "get all", method: "GET", endpoint: "/review/all", expected: `[]`, expectedStatus: 200},
		{name: "get single", method: "GET", endpoint: "/review/1", expected: `Not found`, expectedStatus: 404},
		{name: "update single fail", method: "PUT", endpoint: "/review/2", expected: `Need more information to update`, expectedStatus: 400},
		{name: "delete single fail", method: "DELETE", endpoint: "/review/2", expected: ``, expectedStatus: 405},
	}

	ud := &user.MapData{Seed: true}
	p := &ReviewHandler{
		Data: &MapData{UserData: ud},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ud.Start()
			req, err := http.NewRequest(tt.method, tt.endpoint, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := getTestHandler(p)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v but expected %v",
					status, tt.expectedStatus)
			}
			if strings.TrimSpace(rr.Body.String()) != tt.expected {
				t.Errorf("handler returned unexpected body: got %v but expected %v",
					rr.Body.String(), tt.expected)
			}
		})
	}
}

func getTestHandler(p *ReviewHandler) http.Handler {
	r := chi.NewRouter()
	r.Route("/review", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/all", p.All)
			r.Get("/{id}", p.Get)
			r.Put("/{id}", p.Update)
		})
	})

	return r
}
