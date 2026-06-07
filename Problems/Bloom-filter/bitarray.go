package main

type BitArray struct {
	bits []bool
	size int
}

func NewBitArray(size int) *BitArray {
	return &BitArray{
		bits: make([]bool, size),
		size: size,
	}
}

func (ba *BitArray) Set(position int) {
	ba.bits[position] = true
}

func (ba *BitArray) Get(position int) bool {
	return ba.bits[position]
}

func (ba *BitArray) Clear() {
	for i := range ba.bits {
		ba.bits[i] = false
	}
}

func (ba *BitArray) Size() int {
	return ba.size
}
