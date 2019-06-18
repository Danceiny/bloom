package bloom

import "github.com/go-redis/redis"

type Storage interface {
	Set(offset uint)
	FlushSet() error
	PrepareSet() error
	Test(offset uint) (bool, error)
}

type RedisStorage struct {
	r           *redis.Client
	transaction *redis.Pipeline
	key         string
}

//// Set bit i to 1
//func (b *BitSet) Set(i uint) *BitSet {
//	b.extendSetMaybe(i)
//	b.set[i>>log2WordSize] |= 1 << (i & (wordSize - 1))
//	return b
//}
func (rs *RedisStorage) Set(offset uint) {
	rs.transaction.SetBit(rs.key, int64(offset), 1)
}

func (rs *RedisStorage) FlushSet() (err error) {
	_, err = rs.transaction.Exec()
	return
}

func (rs *RedisStorage) PrepareSet() (err error) {
	rs.transaction = rs.r.Pipeline().(*redis.Pipeline)
	return
}

func (rs *RedisStorage) Test(offset uint) (bool, error) {
	var bitValue, err = rs.r.GetBit(rs.key, int64(offset)).Result()
	return bitValue == 1, err
}
