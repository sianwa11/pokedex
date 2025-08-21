package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{"key1", []byte("value1")},
		{"key2", []byte("value2")},
		{"key3", []byte("value3")},	
		{"key4", []byte("value4")},
	}

	for _, c := range cases {
		cache := NewCache(interval)
		cache.Add(c.key, c.val)

		cached, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key")
			return
		}

		if string(c.val) != string(cached) {
			t.Errorf("values are not the same")
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("key1", []byte("value1"))

	_, ok := cache.Get("key1")
	if !ok {
		t.Errorf("Expected to get value 1")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("key1")
	
	if ok {
		t.Errorf("Not expected to get value 1")
		return
	}
}