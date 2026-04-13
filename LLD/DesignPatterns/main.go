package main

import (
	singleton "designpatterns/Singleton"
	builder "designpatterns/builder"
	factory "designpatterns/factory-method"
	"fmt"
)

func main() {
	fmt.Println("Calling Dependency Injection")
	testDependencyInjection()
	fmt.Println("Calling Singleton Patterns")
	singleton.GetInstance()
	fmt.Println("Calling Builder Pattern")
	builder.TestBuilderPattern()
	fmt.Println("Calling Factory Method")
	factory.TestFactoryMethod()
}
