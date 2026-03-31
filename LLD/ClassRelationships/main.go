package main

import (
	"fmt"
)


func main(){
	fmt.Println("Calling Association")
	testAssociation()
	fmt.Println("Calling Aggregation")
	testAggregation()
	fmt.Println("Calling Composition")
	testCompsition()
	fmt.Println("Calling Dependency")
	testDependency()
	fmt.Println("Calling Realization")
	testRealization()
}