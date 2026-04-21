package main

import (
	singleton "designpatterns/Singleton"
	abstractfactory "designpatterns/abstract-factory"
	builder "designpatterns/builder"
	factory "designpatterns/factory-method"
	"designpatterns/prototype"
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
	fmt.Println("Calling Abstract Factory Method")
	abstractfactory.TestAbstractFactory()
	fmt.Println("Calling Prototype Method")
	prototype.TestPrototype()
}
