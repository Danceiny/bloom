package bloom

type StorageBloomFilter struct {
	m uint
	k uint
	b Storage
}

func (f *StorageBloomFilter) Add(data []byte) BF {
	h := baseHashes(data)
	_ = f.b.PrepareSet()
	for i := uint(0); i < f.k; i++ {
		f.b.Set(f.Location(h, i))
	}
	_ = f.b.FlushSet()
	return f
}

// Location returns the ith hashed Location using the four base hash values
func (f *StorageBloomFilter) Location(h [4]uint64, i uint) uint {
	return uint(location(h, i) % uint64(f.m))
}

// Test returns true if the data is in the MemoryBloomFilter, false otherwise.
// If true, the result might be a false positive. If false, the data
// is definitely not in the set.
func (f *StorageBloomFilter) Test(data []byte) bool {
	h := baseHashes(data)
	for i := uint(0); i < f.k; i++ {
		var b, err = f.b.Test(f.Location(h, i))
		if err != nil || !b {
			// we view it as false while backend error occurred
			return false
		}
	}
	return true
}
func (f *StorageBloomFilter) TestAndAdd(data []byte) bool {
	present := true
	h := baseHashes(data)
	_ = f.b.PrepareSet()
	for i := uint(0); i < f.k; i++ {
		l := f.Location(h, i)
		var v, err = f.b.Test(l)
		if err != nil || !v {
			present = false
		}
		f.b.Set(l)
	}
	_ = f.b.FlushSet()
	return present
}
