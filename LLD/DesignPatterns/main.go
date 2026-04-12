package main

import (
	singleton "designpatterns/Singleton"
	builder "designpatterns/builder"
	"fmt"
)

func main() {
	fmt.Println("Calling Dependency Injection")
	testDependencyInjection()
	fmt.Println("Calling Singleton Patterns")
	singleton.GetInstance()
	fmt.Println("Calling Builder Pattern")
	builder.TestBuilderPattern()
}
