package main

import (
	singleton "designpatterns/Singleton"
	abstractfactory "designpatterns/abstract-factory"
	adapter "designpatterns/adapter"
	builder "designpatterns/builder"
	decorator "designpatterns/decorator"
	factory "designpatterns/factory-method"
	"designpatterns/prototype"
	"fmt"
)

func main() {
	fmt.Println("\nCalling Dependency Injection")
	testDependencyInjection()
	fmt.Println("\nCalling Singleton Patterns")
	singleton.GetInstance()
	fmt.Println("\nCalling Builder Pattern")
	builder.TestBuilderPattern()
	fmt.Println("\nCalling Factory Method")
	factory.TestFactoryMethod()
	fmt.Println("\nCalling Abstract Factory Method")
	abstractfactory.TestAbstractFactory()
	fmt.Println("\nCalling Prototype Method")
	prototype.TestPrototype()
	fmt.Println("\nCalling Adapter Pattern")
	adapter.TestAdapterPattern()
	fmt.Println("\nCalling Decorator Design pattern")
	decorator.TestDecoratorPattern()
}
