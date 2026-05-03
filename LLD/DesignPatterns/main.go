package main

import (
	singleton "designpatterns/Singleton"
	abstractfactory "designpatterns/abstract-factory"
	adapter "designpatterns/adapter"
	"designpatterns/bridge"
	builder "designpatterns/builder"
	"designpatterns/composite"
	decorator "designpatterns/decorator"
	factory "designpatterns/factory-method"
	"designpatterns/flyweight"
	"designpatterns/prototype"
	"designpatterns/proxy"
	"designpatterns/strategy"
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
	fmt.Println("\nCalling Composite Design Pattern")
	composite.TestCompositePattern()
	fmt.Println("\nCalling Proxy Pattern")
	proxy.TestProxyPattern()
	fmt.Println("\nCalling Bridge Pattern")
	bridge.TestBridgePattern()
	fmt.Println("\nCalling Flyweight Pattern")
	flyweight.TestFlyweight()
	fmt.Println("\nCalling Strategy")
	strategy.TestStrategyPattern()
}
