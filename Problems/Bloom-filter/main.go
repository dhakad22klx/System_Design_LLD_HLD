package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("=== Bloom Filter Demo ===")
	fmt.Println()

	fmt.Println("1. Creating Bloom filter (expected: 1000, FP rate: 1%)")
	filter, err := NewBloomFilterBuilder(1000).
		FalsePositiveRate(0.01).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	config := filter.Config()
	fmt.Printf("   Bit array size: %d\n", config.BitArraySize())
	fmt.Printf("   Hash functions: %d\n", config.NumHashFunctions())
	fmt.Println("   Using default MurmurHash3Strategy")

	filter.Add("apple")
	filter.Add("banana")
	filter.Add("cherry")
	fmt.Println("   Added: apple, banana, cherry")

	fmt.Println()
	fmt.Println("2. Checking membership")
	fmt.Printf("   mightContain('apple')  = %v\n", filter.MightContain("apple"))
	fmt.Printf("   mightContain('banana') = %v\n", filter.MightContain("banana"))
	fmt.Printf("   mightContain('cherry') = %v\n", filter.MightContain("cherry"))
	fmt.Printf("   mightContain('grape')  = %v\n", filter.MightContain("grape"))
	fmt.Printf("   mightContain('mango')  = %v\n", filter.MightContain("mango"))

	fmt.Println()
	fmt.Println("3. Creating filter with FNV hash strategy")
	fnvFilter, err := NewBloomFilterBuilder(1000).
		FalsePositiveRate(0.01).
		WithHashStrategy(&FNV1aHashStrategy{}).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	fnvFilter.Add("hello")
	fnvFilter.Add("world")
	fmt.Println("   Added: hello, world")
	fmt.Printf("   mightContain('hello') = %v\n", fnvFilter.MightContain("hello"))
	fmt.Printf("   mightContain('world') = %v\n", fnvFilter.MightContain("world"))
	fmt.Printf("   mightContain('foo')   = %v\n", fnvFilter.MightContain("foo"))

	fmt.Println()
	fmt.Println("4. Testing clear()")
	filter.Clear()
	fmt.Println("   Cleared the filter")
	fmt.Printf("   mightContain('apple')  = %v\n", filter.MightContain("apple"))
	fmt.Printf("   mightContain('banana') = %v\n", filter.MightContain("banana"))

	fmt.Println()
	fmt.Println("5. False positive demonstration")
	smallFilter, err := NewBloomFilterBuilder(10).
		FalsePositiveRate(0.1).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	smallFilter.Add("cat")
	smallFilter.Add("dog")
	smallFilter.Add("bird")
	fmt.Println("   Small filter (expected: 10, FP rate: 10%)")
	fmt.Println("   Added: cat, dog, bird")

	falsePositives := 0
	testWords := []string{"fish", "lion", "bear", "wolf", "deer",
		"frog", "hawk", "duck", "goat", "seal"}
	for _, word := range testWords {
		if smallFilter.MightContain(word) {
			falsePositives++
			fmt.Printf("   False positive: '%s'\n", word)
		}
	}
	fmt.Printf("False positives: %d/%d\n", falsePositives, len(testWords))

	fmt.Println()
	fmt.Println("=== Demo Complete ===")
}
