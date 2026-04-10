package main

import (
	singleton "designpatterns/Singleton"
	"fmt"
)

func main() {
	fmt.Println("Calling Dependency Injection")
	testDependencyInjection()
	fmt.Println("Calling Singleton Patterns")
	singleton.GetInstance()
}
