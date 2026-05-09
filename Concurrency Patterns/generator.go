package main

import (
	"context"
	"fmt"
	"time"
)

/*
  - Generator Pattern is used to generate a sequence of values which is used to produce some output.
  - In our example, we have a generator function that simply returns a channel from which we can read the values.
  - This works on the fact that sends and receives block until both the sender and receiver are ready.
    This property allowed us to wait until the next value is requested.
*/
func main() {
	// 1. Create a context that we can cancel manually
	ctx, cancel := context.WithCancel(context.Background())

	// 2. Ensure cancel is called when main finishes to clean up the generator
	defer cancel()

	// 3. Start the generator
	ch := generator(ctx)

	fmt.Println("Starting to consume values...")

	for i := 0; i < 5; i++ {
		value := <-ch
		fmt.Printf("Consumed Value: %d\n", value)

		// Simulating some work being done with the value
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Main is done. Cleaning up background workers...")
	// After this point, 'defer cancel()' runs, and the generator stops.
}

// generator returns a receive-only channel
func generator(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		defer fmt.Println("Generator goroutine has exited safely.")
		defer close(ch) // Close channel when the loop ends

		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				// 4. If main is done or calls cancel(), we exit the loop
				return
			case ch <- i:
				// 5. Normal operation: send the value into the channel
				// This still blocks until main is ready to read it!
			}
		}
	}()

	return ch
}
