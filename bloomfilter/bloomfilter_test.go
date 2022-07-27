package bloomfilter

import (
	"crypto/rand"
	"testing"
)

// 100 Million
var filterSize uint = 100000000

// 100 Thousand
var sampleSize int = 100000

// Test items if items may exist into set
func TestExistance(t *testing.T) {
	bf := New(filterSize, DefaultHashFunctions)

	for i := 0; i < sampleSize; i++ {
		d1 := randomBytes(10)

		bf.Add(d1)

		if bf.Search(d1) != true {
			t.Errorf("'%q' not found", d1)
		}

		// create some data that don't exist
		d2 := append(d1, randomBytes(10)...)

		// Test that data does NOT exist
		if bf.Search(d2) == true {
			t.Errorf("'%q' should not be found", d2)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	bf := New(filterSize, DefaultHashFunctions)
	for i := 0; i < b.N; i++ {
		bf.Add(randomBytes(10))
	}
}

func BenchmarkSearch(b *testing.B) {
	bf := New(filterSize, DefaultHashFunctions)
	for i := 0; i < b.N; i++ {
		bf.Add(randomBytes(10))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Search(randomBytes(10))
	}
}

func randomBytes(size int) []byte {
	b := make([]byte, size)
	rand.Read(b)
	return b
}
