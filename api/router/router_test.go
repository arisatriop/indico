package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRouter(t *testing.T) {
	r := gin.Default()
	Init(r)

	t.Run("Test routes", func(t *testing.T) {
		if len(r.Routes()) != 15 {
			t.Errorf("Expected 15 routes, got %d", len(r.Routes()))
		}
	})

	t.Run("Auth", func(t *testing.T) {
		ts := httptest.NewServer(r)
		defer ts.Close()

		resp, err := http.Get(ts.URL + "/api/users/me")
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("Expected status code 401, got %d", resp.StatusCode)

		}
	})

	t.Run("Admin", func(t *testing.T) {
		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest("GET", ts.URL+"/api/users/", nil)
		if err != nil {
			t.Error(err)
		}

		t.Run("Invalid token", func(t *testing.T) {
			req.Header.Set("Authorization", "Bearer some-invalid-token")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Error(err)
			}

			if resp.StatusCode != http.StatusUnauthorized {
				t.Errorf("Expected status code 401, got %d", resp.StatusCode)

			}
		})
	})
}
