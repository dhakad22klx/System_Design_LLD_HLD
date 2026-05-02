package decorator

/*
Decorator is a structural design pattern that lets you attach new behaviors to objects
by placing these objects inside special wrapper objects that contain the behaviors.


The Decorator Design Pattern is a structural pattern that lets you dynamically add
new behavior or responsibilities to objects without modifying their underlying code.


Decorator is a structural pattern that allows adding new behaviors to objects dynamically
by placing them inside special wrapper objects, called decorators.

Using decorators you can wrap objects countless number of times since
both target objects and decorators follow the same interface.
The resulting object will get a stacking behavior of all wrappers.
*/

import (
	"fmt"
)

func TestDecoratorPattern() {

	pizza := &VeggieMania{}

	//Add cheese topping - decorate with cheese
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
