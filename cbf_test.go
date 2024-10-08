package bloomfilter

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	cbf := NewCountingBloomFilter(100, 0.01)
	cbf.Add([]byte("test1"))
	cbf.Add([]byte("aaa"))
	cbf.Add([]byte("another"))

	assert.True(t, cbf.Exist([]byte("aaa")))

	cbf.Remove([]byte("another"))

	assert.False(t, cbf.Exist([]byte("another")))
}

func BenchmarkAdd(b *testing.B) {
	cbf := NewCountingBloomFilter(100000, 0.01)

	for i := 0; i < b.N; i++ {
		element := []byte(strconv.Itoa(i))
		cbf.Add(element)
	}
}

func BenchmarkExist(b *testing.B) {
	cbf := NewCountingBloomFilter(100000, 0.01)

	for i := 0; i < 100000; i++ {
		element := []byte(strconv.Itoa(i))
		cbf.Add(element)
	}

	for i := 0; i < b.N; i++ {
		element := []byte(strconv.Itoa(i % 100000))
		cbf.Exist(element)
	}
}

func BenchmarkRemove(b *testing.B) {
	cbf := NewCountingBloomFilter(100000, 0.01)

	for i := 0; i < 100000; i++ {
		element := []byte(strconv.Itoa(i))
		cbf.Add(element)
	}

	for i := 0; i < b.N; i++ {
		element := []byte(strconv.Itoa(i % 100000))
		cbf.Remove(element)
	}
}
