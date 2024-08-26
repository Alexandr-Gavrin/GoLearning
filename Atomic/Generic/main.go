package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Numbers[T Num] []T

func sum[V Num](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func showAny[V any](elmts ...V) {
	for _, el := range elmts {
		fmt.Print(el, " ")
	}
	fmt.Println()

}

func contains[T comparable](elmts []T, elem T) bool {
	for _, el := range elmts {
		if el == elem {
			return true
		}
	}
	return false
}

func main() {
	nums := []int32{1, 2, 3, 4}
	fmt.Println(sum(nums))
	tmp := []interface{}{1, "sdaf", "3234", true, []int{1, 2, 3}}
	showAny(tmp)

}

func withoutMutex() {
	var (
		counter int64
		wg      sync.WaitGroup
	)
	wg.Add(1000)

	timer := time.Now()
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter, time.Since(timer).Seconds())

}

func withMutex() {
	var (
		counter int64
		mtx     sync.Mutex
		wg      sync.WaitGroup
	)
	timer := time.Now()
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			mtx.Lock()
			counter++
			mtx.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter, time.Since(timer).Seconds())

}
