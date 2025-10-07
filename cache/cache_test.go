package cache

import (
	"testing"
)

func TestCache_BasicOperations(t *testing.T) {
	tests := []struct {
		name       string
		maxSize    int
		operations []struct {
			op            string
			key           string
			value         interface{}
			expectedValue interface{}
			expectedOk    bool
		}
	}{
		{
			name:    "basic_set_get",
			maxSize: 5,
			operations: []struct {
				op            string
				key           string
				value         interface{}
				expectedValue interface{}
				expectedOk    bool
			}{
				{"set", "user_id", "12345", nil, false},
				{"get", "user_id", nil, "12345", true},
				{"set", "session_token", "abc123def456", nil, false},
				{"get", "session_token", nil, "abc123def456", true},
			},
		},
		{
			name:    "cache_miss",
			maxSize: 3,
			operations: []struct {
				op            string
				key           string
				value         interface{}
				expectedValue interface{}
				expectedOk    bool
			}{
				{"get", "nonexistent", nil, nil, false},
				{"set", "product_id", 98765, nil, false},
				{"get", "product_id", nil, 98765, true},
			},
		},
		{
			name:    "different_data_types",
			maxSize: 10,
			operations: []struct {
				op            string
				key           string
				value         interface{}
				expectedValue interface{}
				expectedOk    bool
			}{
				{"set", "string_val", "hello world", nil, false},
				{"set", "int_val", 42, nil, false},
				{"set", "float_val", 3.14159, nil, false},
				{"set", "bool_val", true, nil, false},
				{"get", "string_val", nil, "hello world", true},
				{"get", "int_val", nil, 42, true},
				{"get", "float_val", nil, 3.14159, true},
				{"get", "bool_val", nil, true, true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewCache(tt.maxSize)

			for _, op := range tt.operations {
				switch op.op {
				case "set":
					cache.Set(op.key, op.value)
				case "get":
					value, ok := cache.Get(op.key)
					if ok != op.expectedOk {
						t.Errorf("Expected ok=%v, got %v", op.expectedOk, ok)
					}
					if ok && value != op.expectedValue {
						t.Errorf("Expected value=%v, got %v", op.expectedValue, value)
					}
				}
			}
		})
	}
}

func TestCache_Eviction(t *testing.T) {
	tests := []struct {
		name       string
		maxSize    int
		operations []struct {
			op            string
			key           string
			value         interface{}
			expectedValue interface{}
			expectedOk    bool
		}
	}{
		{
			name:    "eviction_after_max_requests",
			maxSize: 2,
			operations: []struct {
				op            string
				key           string
				value         interface{}
				expectedValue interface{}
				expectedOk    bool
			}{
				{"set", "key1", "value1", nil, false},
				{"get", "key1", nil, "value1", true},
				{"get", "key1", nil, "value1", true},
				{"get", "key1", nil, nil, false},
			},
		},
		{
			name:    "multiple_keys_eviction",
			maxSize: 1,
			operations: []struct {
				op            string
				key           string
				value         interface{}
				expectedValue interface{}
				expectedOk    bool
			}{
				{"set", "user_1", "john_doe", nil, false},
				{"get", "user_1", nil, "john_doe", true},
				{"get", "user_1", nil, nil, false},
				{"set", "user_2", "jane_smith", nil, false},
				{"get", "user_2", nil, "jane_smith", true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewCache(tt.maxSize)

			for _, op := range tt.operations {
				switch op.op {
				case "set":
					cache.Set(op.key, op.value)
				case "get":
					value, ok := cache.Get(op.key)
					if ok != op.expectedOk {
						t.Errorf("Expected ok=%v, got %v", op.expectedOk, ok)
					}
					if ok && value != op.expectedValue {
						t.Errorf("Expected value=%v, got %v", op.expectedValue, value)
					}
				}
			}
		})
	}
}

func TestCache_InfiniteCache(t *testing.T) {
	cache := NewCache(-1)

	cache.Set("persistent_key", "persistent_value")

	for i := 0; i < 1000; i++ {
		value, ok := cache.Get("persistent_key")
		if !ok || value != "persistent_value" {
			t.Fatalf("Cache miss at iteration %d: value=%v, ok=%v", i, value, ok)
		}
	}
}

func TestCache_ZeroCache(t *testing.T) {
	cache := NewCache(0)

	cache.Set("temp_key", "temp_value")

	value, ok := cache.Get("temp_key")
	if ok {
		t.Errorf("Expected cache miss with zero cache, got value=%v", value)
	}
}

func TestCache_ConcurrentAccess(t *testing.T) {
	cache := NewCache(100)

	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer func() { done <- true }()

			key := "key_" + string(rune(id))
			value := "value_" + string(rune(id))

			cache.Set(key, value)

			for j := 0; j < 100; j++ {
				if v, ok := cache.Get(key); !ok || v != value {
					t.Errorf("Concurrent access failed for goroutine %d", id)
					return
				}
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
