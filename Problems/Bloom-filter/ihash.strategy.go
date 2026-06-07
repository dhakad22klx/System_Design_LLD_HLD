package main

type IHashStrategy interface {
	Hash(element string, seed int, bitArraySize int) int
}
