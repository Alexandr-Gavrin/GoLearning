package main

import "fmt"

func deferPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic happend: ", err)
		}
	}()
	fmt.Println("Some job")
	panic("russian win!")
	return
}

func main() {

	//deferPanic()

}
