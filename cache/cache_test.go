package cache_test

import (
	"github.com/agirenko/gocache/cache"
	"testing"
)

// ChatGPT generated tests
func TestCc_SetAndGet(t *testing.T) {
	mycache := cache.New()

	// Set a key-value pair
	mycache.Set("key", "value")

	// Get the value of the key
	value := mycache.Get("key")

	// Check if the value is the expected value
	if value != "value" {
		t.Errorf("Expected value %v, but got %v", "value", value)
	}
}

func TestCc_Delete(t *testing.T) {
	mycache := cache.New()

	// Set a key-value pair
	mycache.Set("key", "value")

	// Delete the key
	deleted := mycache.Delete("key")

	// Check if the key was deleted
	if !deleted {
		t.Error("Expected key to be deleted, but it wasn't")
	}

	// Try to get the value of the deleted key
	value := mycache.Get("key")

	// Check if the value is nil
	if value != nil {
		t.Error("Expected value to be nil, but it wasn't")
	}

	// Try to delete the key again
	deleted = mycache.Delete("key")

	// Check if the key was not deleted again
	if deleted {
		t.Error("Expected key to not be deleted again, but it was")
	}
}
