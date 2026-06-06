package builder

/*The Builder Design Pattern is a creational pattern that lets you construct complex objects step-by-step,
separating the construction logic from the final representation.

It’s particularly useful in situations where:

- An object has many optional fields, and most callers only need a subset.
- You want to avoid telescoping constructors or long parameter lists.
- The object must be assembled through multiple steps, possibly in a specific order.
----

Builder is a creational design pattern, which allows constructing complex objects step by step.

Unlike other creational patterns, Builder doesn’t require products to have a common interface.
That makes it possible to produce different products using the same construction process.
*/
import (
	"fmt"
)

func TestBuilderPattern() {
	// Create a GET request
	getRequest, err := NewHttpRequestBuilder("https://api.example.com/users")
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	getRequest.
		Method("GET").
		Header("Accept", "application/json").
		QueryParam("page", "1").
		QueryParam("limit", "10").
		Timeout(5000)

	request := getRequest.Build()
	fmt.Println("GET Request:", request)

	// Create a POST request
	postRequest, err := NewHttpRequestBuilder("https://api.example.com/users")
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	postRequest.
		Method("POST").
		Header("Content-Type", "application/json").
		Header("Authorization", "Bearer token123").
		Body(`{"name": "John Doe", "email": "john@example.com"}`).
		Timeout(10000)

	request = postRequest.Build()
	fmt.Println("\nPOST Request:", request)
}
