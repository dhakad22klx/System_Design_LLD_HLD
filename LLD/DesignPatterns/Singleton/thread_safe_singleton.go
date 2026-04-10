package singleton

import "sync"

type ThreadSafeSingleton struct {
	//
}

var (
	thredSafeInstance *ThreadSafeSingleton
	mu                sync.Mutex
)

func GetThreadSafeInstance() *ThreadSafeSingleton {
	mu.Lock()
	defer mu.Unlock()

	if thredSafeInstance == nil {
		thredSafeInstance = &ThreadSafeSingleton{}
	}
	return thredSafeInstance
}

/*
"Performance Consideration"

This approach is correct but has a performance cost: every call to getInstance() acquires a lock,
even after the instance has been created.
Once the instance exists, there is no reason to synchronize. The next approach fixes this.
*/
