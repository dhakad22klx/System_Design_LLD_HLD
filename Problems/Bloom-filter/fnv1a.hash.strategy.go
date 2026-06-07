package main

const (
	fnvOffsetBasis int32 = -2128831035 // 0x811c9dc5 as int32
	fnvPrime       int32 = 16777619    // 0x01000193
)

type FNV1aHashStrategy struct{}

func (s *FNV1aHashStrategy) Hash(element string, seed int, bitArraySize int) int {
	hash := fnvOffsetBasis ^ int32(seed)

	for i := 0; i < len(element); i++ {
		hash ^= int32(element[i])
		hash *= fnvPrime
	}

	result := int(hash) % bitArraySize
	if result < 0 {
		result = -result
	}
	return result
}
