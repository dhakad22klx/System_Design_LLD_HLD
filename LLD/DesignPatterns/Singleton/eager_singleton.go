package singleton

// EagerSingleton implements eager initialization singleton pattern
type EagerSingleton struct{}

// The single instance, created immediately
var eagerInstance = &EagerSingleton{}

// Returns the singleton instance
func GetEagerInstance() *EagerSingleton {
	return eagerInstance
}