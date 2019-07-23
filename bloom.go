package bloom

import "github.com/go-redis/redis"

type BF interface {
	Add([]byte) BF
	Test([]byte) bool                  // test whether existed or not
	TestAndAdd(data []byte) bool       // test and add
	Location(h [4]uint64, i uint) uint // do hash
}

type BackendType int

const (
	MEMORY BackendType = iota
	REDIS
)

func NewStorageBloomFilter(m, k uint, storage Storage) BF {
	return &StorageBloomFilter{
		m: m,
		k: k,
		b: storage,
	}
}

func NewRedisBloomFilter(name string, m, k uint, r *redis.Client) BF {
	var rs = RedisStorage{
		r:   r,
		key: name,
	}
	return &StorageBloomFilter{
		m: m,
		k: k,
		b: &rs,
	}
}
