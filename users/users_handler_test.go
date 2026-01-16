package users_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"your_project/users"
)

type MockStore struct {
	CreateFunc   func(user users.User) error
	CreateCalled bool
	GetFunc      func(id string) (users.User, error)
}

func (m *MockStore) Create(user users.User) error {
	m.CreateCalled = true
	return m.CreateFunc(user)
}

func (m *MockStore) Get(id string) (users.User, error) {
	return m.GetFunc(id)
}

func TestCreateUserHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid email",
			requestBody:    `{"email": "test@example.com"}`,
			expectedStatus: http.StatusCreated,
			expectedBody:   "",
		},
		{
			name:           "invalid email",
			requestBody:    `{"email": "testexample.com"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid email",
		},
		{
			name:           "invalid JSON",
			requestBody:    `{"email": test@example.com}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			mockStore := &MockStore{
				CreateFunc: func(user users.User) error {
					return nil
				},
			}

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				users.CreateUserHandler(w, r, mockStore)
			})
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code, "Status code mismatch")

			if tt.expectedBody != "" {
				body, _ := io.ReadAll(rr.Body)
				assert.Contains(t, string(body), tt.expectedBody, "Response body mismatch")
			}

			if tt.name == "valid email" {
				assert.True(t, mockStore.CreateCalled, "Create should be called for valid email")
			} else {
				assert.False(t, mockStore.CreateCalled, "Create should not be called for invalid cases")
			}
		})
	}
}

func TestGetUserHandler(t *testing.T) {
	mockStore := &MockStore{
		GetFunc: func(id string) (users.User, error) {
			return users.User{}, fmt.Errorf("user not found")
		},
	}

	req, _ := http.NewRequest("GET", "/users/123", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users.GetUserHandler(w, r, mockStore)
	})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Status code mismatch")

	body, _ := io.ReadAll(rr.Body)
	assert.Contains(t, string(body), "User not found", "Response body mismatch")
}