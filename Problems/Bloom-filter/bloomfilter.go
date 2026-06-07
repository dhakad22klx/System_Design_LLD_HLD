package main

import "sync"

type BloomFilter struct {
	config       *BloomFilterConfig
	bitArray     *BitArray
	hashStrategy IHashStrategy
	mu           sync.RWMutex
}

func newBloomFilter(b *BloomFilterBuilder) *BloomFilter {
	return &BloomFilter{
		config:       b.config,
		bitArray:     NewBitArray(b.config.BitArraySize()),
		hashStrategy: b.hashStrategy,
	}
}

// Lock ensures all k bits are set atomically
func (bf *BloomFilter) Add(element string) {
	if element == "" {
		panic("element cannot be empty")
	}

	bf.mu.Lock()
	defer bf.mu.Unlock()

	for i := 0; i < bf.config.NumHashFunctions(); i++ {
		position := bf.hashStrategy.Hash(element, i, bf.config.BitArraySize())
		bf.bitArray.Set(position)
	}
}

// RLock ensures consistent reads across all k positions
func (bf *BloomFilter) MightContain(element string) bool {
	if element == "" {
		panic("element cannot be empty")
	}

	bf.mu.RLock()
	defer bf.mu.RUnlock()

	for i := 0; i < bf.config.NumHashFunctions(); i++ {
		position := bf.hashStrategy.Hash(element, i, bf.config.BitArraySize())
		if !bf.bitArray.Get(position) {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) Clear() {
	bf.mu.Lock()
	defer bf.mu.Unlock()

	bf.bitArray.Clear()
}

func (bf *BloomFilter) Config() *BloomFilterConfig {
	return bf.config
}
