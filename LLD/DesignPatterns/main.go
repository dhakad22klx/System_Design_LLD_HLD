package main

import (
	singleton "designpatterns/Singleton"
	abstractfactory "designpatterns/abstract-factory"
	adapter "designpatterns/adapter"
	"designpatterns/bridge"
	builder "designpatterns/builder"
	chainofresponsibility "designpatterns/chain-of-responsibility"
	commandpattern "designpatterns/command-pattern"
	"designpatterns/composite"
	decorator "designpatterns/decorator"
	factory "designpatterns/factory-method"
	"designpatterns/flyweight"
	"designpatterns/iterator"
	"designpatterns/observer"
	"designpatterns/prototype"
	"designpatterns/proxy"
	"designpatterns/state"
	"designpatterns/strategy"
	templatemethod "designpatterns/template-method"
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
	fmt.Println("\nCalling Iterator Pattern")
	iterator.TestIteratorPattern()
	fmt.Println("\nCalling Observer Pattern")
	observer.TestObserverPattern()
	fmt.Println("\nCalling Command Pattern")
	commandpattern.TestCommandPattern()
	fmt.Println("\nCalling State Pattern")
	state.TestStatePattern()
	fmt.Println("\nCalling Template Pattern")
	templatemethod.TestTemplateMethod()
	fmt.Println("\nCalling Chain of Responsibility")
	chainofresponsibility.TestChainOfResponsibility()
}
