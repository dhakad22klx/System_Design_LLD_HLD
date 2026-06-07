package main

import "math"

type BloomFilterConfig struct {
	expectedElements  int
	falsePositiveRate float64
	bitArraySize      int
	numHashFunctions  int
}

func NewBloomFilterConfig(expectedElements int, falsePositiveRate float64) *BloomFilterConfig {
	// m = -(n * ln(p)) / (ln(2))^2
	bitArraySize := int(math.Ceil(
		-(float64(expectedElements) * math.Log(falsePositiveRate)) / (math.Log(2) * math.Log(2)),
	))

	// k = (m / n) * ln(2)
	numHashFunctions := int(math.Round(
		(float64(bitArraySize) / float64(expectedElements)) * math.Log(2),
	))
	if numHashFunctions < 1 {
		numHashFunctions = 1
	}

	return &BloomFilterConfig{
		expectedElements:  expectedElements,
		falsePositiveRate: falsePositiveRate,
		bitArraySize:      bitArraySize,
		numHashFunctions:  numHashFunctions,
	}
}

func (c *BloomFilterConfig) ExpectedElements() int      { return c.expectedElements }
func (c *BloomFilterConfig) FalsePositiveRate() float64 { return c.falsePositiveRate }
func (c *BloomFilterConfig) BitArraySize() int          { return c.bitArraySize }
func (c *BloomFilterConfig) NumHashFunctions() int      { return c.numHashFunctions }
