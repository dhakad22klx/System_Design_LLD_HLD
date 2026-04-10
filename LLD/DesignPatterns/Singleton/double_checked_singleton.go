package singleton

import "sync"

// DoubleCheckedSingleton implements double-checked locking singleton pattern
type DoubleCheckedSingleton struct{}

var (
	doubleCheckedInstance *DoubleCheckedSingleton
	dcMutex               sync.Mutex
)

// Returning Single Instance
func GetDoubleCheckedInstance() *DoubleCheckedSingleton {
	if doubleCheckedInstance == nil {
		dcMutex.Lock()
		defer dcMutex.Unlock()

		if doubleCheckedInstance == nil {
			doubleCheckedInstance = &DoubleCheckedSingleton{}
		}
	}
	return doubleCheckedInstance
}

/*
Double-checked locking reduces the performance overhead by only synchronizing during the first object creation.
After the instance exists, threads skip the lock entirely.

It can drastically reduce performance overhead, especially when the singleton is accessed frequently.
*/
