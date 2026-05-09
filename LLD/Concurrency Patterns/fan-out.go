package main

/*
Fan-out patterns allow us to essentially split our single input channel into multiple output channels.
This is a useful pattern to distribute work items into multiple uniform actors.

In our example, we break the input channel into 4 different output channels.
For a dynamic number of outputs, we can merge outputs into a shared "aggregate" channel and use select.

Note: fan-out pattern is different from pub/sub.
*/

import "fmt"

func main() {
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}
	in := generateWork(work)

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)
	out4 := fanOut(in)

	for i := 0; i < len(work); {
		select {
		case value, ok := <-out1:
			if ok {
				fmt.Println("Output 1 got:", value)
				i++
			} // if we print without condition, it can output 0 for closed channels.
		case value, ok := <-out2:
			if ok {
				fmt.Println("Output 2 got:", value)
				i++
			}
		case value, ok := <-out3:
			if ok {
				fmt.Println("Output 3 got:", value)
				i++
			}
		case value, ok := <-out4:
			if ok {
				fmt.Println("Output 4 got:", value)
				i++
			}
		}
	}
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for data := range in {
			out <- data
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}
