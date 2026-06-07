package main

import "fmt"

type BloomFilterBuilder struct {
	config            *BloomFilterConfig
	expectedElements  int
	falsePositiveRate float64
	hashStrategy      IHashStrategy
}

func NewBloomFilterBuilder(expectedElements int) *BloomFilterBuilder {
	return &BloomFilterBuilder{
		config:            nil,
		expectedElements:  expectedElements,
		falsePositiveRate: 0.01,
		hashStrategy:      &MurmurHash3Strategy{},
	}
}

func (b *BloomFilterBuilder) FalsePositiveRate(rate float64) *BloomFilterBuilder {
	b.falsePositiveRate = rate
	return b
}

func (b *BloomFilterBuilder) WithHashStrategy(strategy IHashStrategy) *BloomFilterBuilder {
	b.hashStrategy = strategy
	return b
}

func (b *BloomFilterBuilder) withConfig() *BloomFilterBuilder {
	b.config = NewBloomFilterConfig(b.expectedElements, b.falsePositiveRate)
	return b
}

func (b *BloomFilterBuilder) Build() (*BloomFilter, error) {
	if b.expectedElements <= 0 {
		return nil, fmt.Errorf("expected elements must be positive")
	}
	if b.falsePositiveRate <= 0 || b.falsePositiveRate >= 1 {
		return nil, fmt.Errorf("false positive rate must be between 0 and 1 (exclusive)")
	}
	b.withConfig()
	return newBloomFilter(b), nil
}
