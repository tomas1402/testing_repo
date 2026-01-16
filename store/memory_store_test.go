package store

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestSaveAndFindUser validates basic save and find operations
func TestSaveAndFindUser(t *testing.T) {
	store := NewMemoryStore()
	user := User{
		ID:        "123",
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	store.Save(user)
	found, err := store.FindByID("123")

	assert.NoError(t, err)
	assert.Equal(t, user, found)
}

// TestFindNonExistentUser validates custom error handling
func TestFindNonExistentUser(t *testing.T) {
	store := NewMemoryStore()
	user, err := store.FindByID("non-existent")

	assert.Nil(t, user)
	assert.Equal(t, ErrUserNotFound, err)
}

// TestConcurrentOperations validates thread-safety with multiple goroutines
func TestConcurrentOperations(t *testing.T) {
	store := NewMemoryStore()
	var wg sync.WaitGroup
	const numGoroutines = 100

	// Concurrent writes
	for i := 0; i < numGoroutines; i++ {
		id := fmt.Sprintf("user-%d", i)
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			user := User{
				ID:        id,
				Name:      "Concurrent User",
				Email:     fmt.Sprintf("%s@example.com", id),
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			}
			store.Save(user)
		}(id)
	}
	wg.Wait()

	// Concurrent reads after writes
	var readWg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		id := fmt.Sprintf("user-%d", i)
		readWg.Add(1)
		go func(id string) {
			defer readWg.Done()
			user, err := store.FindByID(id)
			assert.NoError(t, err)
			assert.Equal(t, "Concurrent User", user.Name)
			assert.Equal(t, fmt.Sprintf("%s@example.com", id), user.Email)
		}(id)
	}
	readWg.Wait()
}