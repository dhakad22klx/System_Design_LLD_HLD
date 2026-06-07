package main

type MurmurHash3Strategy struct{}

func (s *MurmurHash3Strategy) Hash(element string, seed int, bitArraySize int) int {
	data := []byte(element)
	h := int32(seed)

	for _, b := range data {
		h ^= int32(b)
		h *= 0x5bd1e995
		h ^= h >> 13
	}

	// Finalization mix
	h ^= h >> 16
	h *= -2048144777 // 0x85ebca6b as int32
	h ^= h >> 13

	result := int(h) % bitArraySize
	if result < 0 {
		result = -result
	}
	return result
}
