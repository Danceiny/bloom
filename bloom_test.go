package bloom

import (
	"github.com/go-redis/redis"
	"testing"
)

func TestBasicRedisBloomFilter(t *testing.T) {
	rc := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	f := NewRedisBloomFilter("t_redis_bf", 1000, 4, rc)
	n1 := []byte("Bess")
	n2 := []byte("Jane")
	n3 := []byte("Emma")
	f.Add(n1)

	n3a := f.TestAndAdd(n3)
	n1b := f.Test(n1)
	n2b := f.Test(n2)
	n3b := f.Test(n3)
	if !n1b {
		t.Errorf("%v should be in.", n1)
	}
	if n2b {
		t.Errorf("%v should not be in.", n2)
	}
	if n3a {
		t.Errorf("%v should not be in the first time we look.", n3)
	}
	if !n3b {
		t.Errorf("%v should be in the second time we look.", n3)
	}
}
