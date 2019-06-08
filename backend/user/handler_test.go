package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

func TestUserHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		endpoint       string
		expected       string
		expectedStatus int
	}{
		{name: "get all", method: "GET", endpoint: "/user/all", expected: `[{"id":"1","name":"James","username":"admin","password":"ssh","newPassword":"ssh","isAdmin":true},{"id":"2","name":"Jamie","username":"user","password":"ssh","newPassword":"ssh"}]`, expectedStatus: 200},
		{name: "get single", method: "GET", endpoint: "/user/1", expected: `{"id":"1","name":"James","username":"admin","password":"ssh","newPassword":"ssh","isAdmin":true}`, expectedStatus: 200},
		{name: "update single fail", method: "PUT", endpoint: "/user/2", expected: `Need more information to update`, expectedStatus: 400},
		{name: "delete single fail", method: "DELETE", endpoint: "/user/2", expected: `Could not determine current user`, expectedStatus: 401},
	}

	u := &UserHandler{
		Data: &MapData{Seed: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u.Data.Start()
			req, err := http.NewRequest(tt.method, tt.endpoint, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := getTestHandler(u)
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

func getTestHandler(u *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/all", u.All)
			r.Get("/{id}", u.Get)
			r.Put("/{id}", u.Update)
			r.Delete("/{id}", u.Delete)
		})
	})

	return r
}
