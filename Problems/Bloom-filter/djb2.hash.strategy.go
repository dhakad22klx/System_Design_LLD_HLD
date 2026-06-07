package main

type DJB2HashStrategy struct{}

func (s *DJB2HashStrategy) Hash(element string, seed int, bitArraySize int) int {
	hash := int32(5381 + seed)

	for i := 0; i < len(element); i++ {
		hash = ((hash << 5) + hash) + int32(element[i]) // hash * 33 + c
	}

	result := int(hash) % bitArraySize
	if result < 0 {
		result = -result
	}
	return result
}
