package flyweight

import "fmt"

/*
Flyweight is a structural design pattern that allows programs to support vast quantities of objects
by keeping their memory consumption low.

The pattern achieves it by sharing parts of object state between multiple objects.
In other words, the Flyweight saves RAM by caching the same data used by different objects.

The Flyweight Design Pattern is a structural pattern that focuses on efficiently sharing
common parts of object state across many objects to reduce memory usage and boost performance.


Two characteristics define the Flyweight pattern and set it apart from other structural patterns:

1. Intrinsic/extrinsic state separation. The pattern divides object state into two parts.
 Intrinsic state is shared and immutable, stored inside the flyweight.
 Extrinsic state is context-dependent and unique, stored outside the flyweight and passed in when needed.
 This separation is what makes sharing possible.
2. Factory-managed caching. A dedicated factory object controls the creation and reuse of flyweights.
 Clients never instantiate flyweights directly.
 They go through the factory, which maintains a cache and ensures
 that identical flyweights are shared rather than duplicated.

*/

func TestFlyweight() {
	// Create a text editor client
	editor := NewTextEditorClient()

	// Render some text with the same properties
	fmt.Println("Rendering text with same properties:")
	editor.RenderText("Hello", 0, 0, "Arial", 12, "black")
	fmt.Printf("Number of unique characters: %d\n\n", editor.GetUniqueCharacterCount())

	// Render the same text with different properties
	fmt.Println("Rendering text with different properties:")
	editor.RenderText("Hello", 0, 20, "Times New Roman", 14, "blue")
	fmt.Printf("Number of unique characters: %d\n\n", editor.GetUniqueCharacterCount())

	// Render text with mixed properties
	fmt.Println("Rendering text with mixed properties:")
	editor.RenderText("World", 0, 40, "Arial", 12, "red")
	fmt.Printf("Number of unique characters: %d\n", editor.GetUniqueCharacterCount())
}
