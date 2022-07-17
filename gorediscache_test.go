package gorediscache

import (
	"reflect"
	"testing"
)

func TestCacheString(t *testing.T) {
	cache := NewCache("redis://localhost:6379/2", 0)
	defer cache.Close()
	err := cache.SetString("hello", "world", 0)
	if err != nil {
		t.Error(err)
	}

	value, err := cache.GetString("hello")
	if err != nil {
		t.Error(err)
	}
	if value != "world" {
		t.Errorf(`wanted "%v", got "%v"`, "world", value)
	}

	err = cache.Delete("hello")
	if err != nil {
		t.Error(err)
	}

	value, err = cache.GetString("hello")
	if err != nil {
		t.Error(err)
	}
	if value != "" {
		t.Errorf(`wanted "%v", got "%v"`, "", value)
	}
}

func TestCacheMap(t *testing.T) {
	cache := NewCache("redis://localhost:6379/2", 0)
	defer cache.Close()
	err := cache.SetMap("m", map[string]string{"a": "b", "c": "d"}, 0)
	if err != nil {
		t.Error(err)
	}

	value, err := cache.GetMap("m")
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(value).String() != "map[string]string" {
		t.Errorf(`wanted "%s", got "%s"`, "map[string]string", reflect.TypeOf(value).String())
	}
	if value["a"] != "b" {
		t.Errorf(`wanted "%s", got "%s"`, "b", value["a"])
	}
	if value["c"] != "d" {
		t.Errorf(`wanted "%s", got "%s"`, "d", value["c"])
	}

	err = cache.Delete("m")
	if err != nil {
		t.Error(err)
	}

	value, err = cache.GetMap("m")
	if err != nil {
		t.Error(err)
	}
	if len(value) != 0 {
		t.Errorf(`wanted "%v", got "%v"`, 0, len(value))
	}
}
