package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockUserStore simulates the user storage layer
type MockUserStore struct {
	users map[int]User
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m *MockUserStore) List() ([]User, error) {
	list := make([]User, 0, len(m.users))
	for _, user := range m.users {
		list = append(list, user)
	}
	return list, nil
}

func (m *MockUserStore) Delete(id int) error {
	if _, exists := m.users[id]; !exists {
		return &UserNotFoundError{ID: id}
	}
	delete(m.users, id)
	return nil
}

type UserNotFoundError struct {
	ID int
}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}

// SetupTestServer creates a test HTTP server with mock handlers
func SetupTestServer() (*httptest.Server, *MockUserStore) {
	store := &MockUserStore{
		users: map[int]User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users, _ := store.List()
		json.NewEncoder(w).Encode(users)
	})

	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Extract ID from path (simplified for example)
		idStr := r.PathValue("id")
		var id int
		_, _ = sscanf(idStr, "%d", &id)

		if err := store.Delete(id); err != nil {
			if _, ok := err.(*UserNotFoundError); ok {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	return httptest.NewServer(mux), store
}

func TestListUsers(t *testing.T) {
	server, _ := SetupTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/users")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var users []User
	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &users)

	assert.Len(t, users, 2)
	assert.Contains(t, users, User{ID: 1, Name: "Alice"})
	assert.Contains(t, users, User{ID: 2, Name: "Bob"})
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		expectedStatus int
	}{
		{"Delete existing user", 1, http.StatusNoContent},
		{"Delete non-existing user", 999, http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, store := SetupTestServer()
			defer server.Close()

			// Save initial state for validation
			initialUsers := make(map[int]User)
			for k, v := range store.users {
				initialUsers[k] = v
			}

			req, _ := http.NewRequest("DELETE", server.URL+"/users/"+string(tt.id), nil)
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Validate state changes for existing user deletion
			if tt.id == 1 {
				assert.NotContains(t, store.users, 1)
				assert.Contains(t, store.users, 2)
			}
		})
	}
}

func TestRoutingRefactor(t *testing.T) {
	server, _ := SetupTestServer()
	defer server.Close()

	// Test GET endpoint
	resp, err := http.Get(server.URL + "/users")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Test DELETE endpoint
	req, _ := http.NewRequest("DELETE", server.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Test invalid method
	req, _ = http.NewRequest("POST", server.URL+"/users", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
}