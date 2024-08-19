package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	timer := time.After(time.Second * 10)

	select {
	case <-timer:
		fmt.Println("times up")
		return
	case sig := <-sigChan:
		fmt.Println("Stopped by: ", sig)
		return
	}
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}

func workerPool() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	wg := sync.WaitGroup{}
	numberToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numberToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			// if i == 500 {
			// 	cancel()
			// }
			numberToProcess <- i
		}
		close(numberToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	counter := 0
	for resVal := range processedNumbers {
		counter++
		fmt.Println(resVal)
	}
	fmt.Println(counter)
}

func makeReq(num int) <-chan string {
	respChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		respChan <- fmt.Sprintf("String is: %d", num)
	}()
	return respChan
}

func chanAsPromise() {
	firstResp := makeReq(1)
	secondResp := makeReq(2)

	fmt.Println(<-firstResp, <-secondResp)
}

func someThings() {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("First Task")
			time.Sleep(time.Second)
		}
	}()

	go func() {

		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Second task")
			err = fmt.Errorf("Any error")
			cancel()
		}
	}()

	go func() {

		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Third Task")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println(err)
}

func main() {
	someThings()
}
