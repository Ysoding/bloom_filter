package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
)

type CountingBloomFilter struct {
	m    int
	k    int
	h    hash.Hash32
	bits []int
}

func NewCountingBloomFilter(numbers uint32, falsePositiveProbability float64) *CountingBloomFilter {
	cbf := &CountingBloomFilter{h: fnv.New32()}
	cbf.estimate(numbers, falsePositiveProbability)
	cbf.bits = make([]int, cbf.m)
	return cbf
}

func (cbf *CountingBloomFilter) estimate(n uint32, p float64) {
	// m = -1 * (n * ln(P)) / (ln2)^2
	cbf.m = int(-1 * (float64(n) * math.Log(p)) / math.Pow(math.Log(2), 2))

	// k = -1 * (lnP) / (ln2)
	// k = m / n * ln2
	cbf.k = int(float64(cbf.m) / float64(n) * math.Log(2))
}

func (cbf *CountingBloomFilter) hash(idx int, element []byte) int {
	cbf.h.Reset()
	cbf.h.Write(element)
	d := int(cbf.h.Sum32())
	return (d + idx) % cbf.m
}

func (cbf *CountingBloomFilter) Add(element []byte) {
	for i := 0; i < cbf.k; i++ {
		idx := cbf.hash(i, element)
		cbf.bits[idx]++
	}
}

func (cbf *CountingBloomFilter) Exist(element []byte) bool {
	for i := 0; i < cbf.k; i++ {
		idx := cbf.hash(i, element)
		if cbf.bits[idx] == 0 {
			return false
		}
	}
	return true
}

func (cbf *CountingBloomFilter) Remove(element []byte) {
	if !cbf.Exist(element) {
		return
	}

	for i := 0; i < cbf.k; i++ {
		idx := cbf.hash(i, element)
		if cbf.bits[idx] > 0 {
			cbf.bits[idx]--
		}
	}
}
