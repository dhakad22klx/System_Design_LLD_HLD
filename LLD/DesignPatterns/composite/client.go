package composite

/*
The Composite Design Pattern is a structural pattern that lets you treat individual objects
and compositions of objects uniformly.

It allows you to build tree-like structures (e.g., file systems, UI hierarchies, organizational charts)
where clients can work with both single elements and groups of elements using the same interface.

It’s particularly useful in situations where:
- You need to represent part-whole hierarchies.
- You want to perform operations on both leaf nodes and composite nodes in a consistent way.
- You want to avoid writing special-case logic to distinguish between "single" and "grouped" objects.

Composite is a structural design pattern that lets you compose objects into tree structures
and then work with these structures as if they were individual objects.

Composite became a pretty popular solution for the most problems that require building a tree structure.
Composite’s great feature is the ability to run methods recursively
over the whole tree structure and sum up the results.
*/

import "fmt"

func TestCompositePattern() {
	// Create root folder
	root := NewFolder("Root")

	// Create subfolders
	documents := NewFolder("Documents")
	downloads := NewFolder("Downloads")
	pictures := NewFolder("Pictures")

	// Create files
	report := NewFile("report.txt", 1024)
	image := NewFile("image.jpg", 2048)
	video := NewFile("video.mp4", 4096)

	// Build the structure
	root.Add(documents)
	root.Add(downloads)
	root.Add(pictures)

	documents.Add(report)
	downloads.Add(video)
	pictures.Add(image)

	// Print the structure
	fmt.Println("File System Structure:")
	root.Print("")

	// Print total size
	fmt.Printf("\nTotal size: %d bytes\n", root.GetSize())

	// Delete a folder
	fmt.Println("\nDeleting Documents folder:")
	root.Remove(documents)
	documents.Delete()

	// Print the structure again
	fmt.Println("\nUpdated File System Structure:")
	root.Print("")
}
