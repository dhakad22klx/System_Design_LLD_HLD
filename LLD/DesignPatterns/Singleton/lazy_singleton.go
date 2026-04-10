package singleton
// Lazy Initialization (Not Thread-Safe)
// This approach creates the singleton instance only when it is needed, 
//saving resources if the singleton is never used in the application.

import (
	"fmt"
	"sync"
)

// LazySingleton implements lazy initialization singleton pattern
type LazySingleton struct{}

var (
	instance *LazySingleton
	once     sync.Once
)

// GetInstance returns the singleton instance
func GetInstance() *LazySingleton {

	fmt.Println("Returning LazySingleton Instance")
	once.Do(func() {
		instance = &LazySingleton{}
	})
	return instance
}
