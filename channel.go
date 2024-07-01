package gokit

import "sync"

// FanOut is a function that will fan out the input channel to multiple workers
func FanOut[T any](input <-chan T, workerNum int, worker func(v T)) {
	for i := 0; i < workerNum; i++ {
		go func() {
			for v := range input {
				worker(v)
			}
		}()
	}
}

// FanIn is a function that will fan in multiple input channels to a single output channel
func FanIn[T any](input ...<-chan T) <-chan T {
	output := make(chan T)

	go func() {
		var wg sync.WaitGroup
		defer close(output)

		for _, v := range input {
			wg.Add(1)
			go func(c <-chan T) {
				defer wg.Done()

				for v := range c {
					output <- v
				}
			}(v)
		}

		wg.Done()
	}()

	return output
}
