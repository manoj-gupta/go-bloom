package bloomfilter

import (
	"hash"
	"hash/fnv"
)

// DefaultHashFunctions for BloomFilter
var DefaultHashFunctions = []hash.Hash64{fnv.New64(), fnv.New64a()}

type BloomFilter struct {
	bitset  []bool        // The bloom-filter bitset
	m       uint          // Size of the bloom filter
	hashfns []hash.Hash64 // The hash functions
}

// Returns a new BloomFilter object,
func New(size uint, hashes []hash.Hash64) *BloomFilter {
	return &BloomFilter{
		bitset:  make([]bool, size),
		m:       size,
		hashfns: hashes,
	}
}

// Add ... add 'data' into the bloom filter
func (bf *BloomFilter) Add(data []byte) {
	for _, v := range bf.hashValues(data) {
		position := uint(v) % bf.m
		bf.bitset[position] = true
	}
}

// Search ... check if 'data' is in bloom filter
func (bf *BloomFilter) Search(data []byte) bool {
	for _, v := range bf.hashValues(data) {
		position := uint(v) % bf.m
		if !bf.bitset[uint(position)] {
			return false
		}
	}
	return true
}

// hashValues .. return hash values by hashing over hash functions
func (bf *BloomFilter) hashValues(item []byte) []uint64 {
	var result []uint64

	for _, hashFunc := range bf.hashfns {
		hashFunc.Write(item)
		result = append(result, hashFunc.Sum64())
		hashFunc.Reset()
	}

	return result
}
