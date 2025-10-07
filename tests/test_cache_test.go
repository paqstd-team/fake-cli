package tests

import (
	"testing"

	"github.com/paqstd-team/fake-cli/cache"
)

func TestCache_BasicAndEviction(t *testing.T) {
	c := cache.NewCache(2)
	c.Set("a", 1)
	if v, ok := c.Get("a"); !ok || v.(int) != 1 {
		t.Fatalf("get1: %v %v", v, ok)
	}
	if v, ok := c.Get("a"); !ok || v.(int) != 1 {
		t.Fatalf("get2: %v %v", v, ok)
	}
	if _, ok := c.Get("a"); ok {
		t.Fatalf("expected eviction after 3rd get")
	}
}

func TestCache_Infinite(t *testing.T) {
	c := cache.NewCache(-1)
	c.Set("x", "y")
	for i := 0; i < 100; i++ {
		if v, ok := c.Get("x"); !ok || v.(string) != "y" {
			t.Fatalf("miss at %d", i)
		}
	}
}
