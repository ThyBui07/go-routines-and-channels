package main

import (
	"fmt"
	"sync"
	"time"
)

func f1() {
	for i := 0; i < 3; i++ {
		fmt.Println("func 1 running at:" + time.Now().Format("15:05:05\n"))
	}
}

func f2() {
	for i := 0; i < 4; i++ {
		fmt.Println("func 2 running at:" + time.Now().Format("15:05:05\n"))
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		f1()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		f2()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("done")

	//output
	// 	func 2 running at:00:45:45

	// func 2 running at:00:45:45

	// func 2 running at:00:45:45

	// func 2 running at:00:45:45

	// func 1 running at:00:45:45

	// func 1 running at:00:45:45

	// func 1 running at:00:45:45

	// run both f1 and f2 at the same time
}
