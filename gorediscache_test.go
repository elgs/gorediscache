package gorediscache

import (
	"testing"
)

func TestCache(t *testing.T) {

	cache := NewCache("redis://localhost:6379/2", 0)
	err := cache.Set("hello", "world", 0)
	if err != nil {
		t.Error(err)
	}

	value, err := cache.Get("hello")
	if err != nil {
		t.Error(err)
	}
	if value != "world" {
		t.Errorf(`wanted "%s", got "%s"`, "world", value)
	}

	err = cache.Delete("hello")
	if err != nil {
		t.Error(err)
	}

	value, err = cache.Get("hello")
	if err != nil {
		t.Error(err)
	}
	if value != "" {
		t.Errorf(`wanted "%s", got "%s"`, "", value)
	}
}
