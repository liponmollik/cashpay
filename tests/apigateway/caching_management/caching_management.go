package caching_management

import (
	"testing"
)

func TestCacheManager_Store(t *testing.T) {
	// Initialize a CacheManager instance
	cacheManager := NewCacheManager()

	// Define test data
	key := "test_key"
	value := "test_value"

	// Store the test data in the cache
	cacheManager.Store(key, value)

	// Retrieve the stored value from the cache
	storedValue, found := cacheManager.Retrieve(key)

	// Check if the stored value matches the expected value
	if !found || storedValue != value {
		t.Errorf("Failed to store or retrieve value from cache. Expected: %s, Got: %s", value, storedValue)
	}
}

func TestCacheManager_Retrieve(t *testing.T) {
	// Initialize a CacheManager instance
	cacheManager := NewCacheManager()

	// Define test data
	key := "test_key"
	value := "test_value"

	// Store test data in the cache
	cacheManager.Store(key, value)

	// Retrieve the stored value from the cache
	storedValue, found := cacheManager.Retrieve(key)

	// Check if the stored value matches the expected value
	if !found || storedValue != value {
		t.Errorf("Failed to retrieve value from cache. Expected: %s, Got: %s", value, storedValue)
	}
}

func TestCacheEviction_EvictLeastRecentlyUsed(t *testing.T) {
	// Initialize a CacheEviction instance
	cacheEviction := NewCacheEviction()

	// Define test data
	key := "test_key"
	value := "test_value"

	// Store test data in the cache
	cacheEviction.Store(key, value)

	// Simulate cache eviction
	cacheEviction.EvictLeastRecentlyUsed()

	// Retrieve the evicted value from the cache
	_, found := cacheEviction.Retrieve(key)

	// Check if the value is evicted as expected
	if found {
		t.Errorf("Value was not evicted from cache as expected.")
	}
}

func TestCacheEviction_EvictLeastFrequentlyUsed(t *testing.T) {
	// Initialize a CacheEviction instance
	cacheEviction := NewCacheEviction()

	// Define test data
	key := "test_key"
	value := "test_value"

	// Store test data in the cache
	cacheEviction.Store(key, value)

	// Simulate cache eviction
	cacheEviction.EvictLeastFrequentlyUsed()

	// Retrieve the evicted value from the cache
	_, found := cacheEviction.Retrieve(key)

	// Check if the value is evicted as expected
	if found {
		t.Errorf("Value was not evicted from cache as expected.")
	}
}
