package main

import (
	"fmt"
	"sync"
)

func main() {
	//f()
	//f2()
	f2()
}

func f() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 1 {
			fmt.Println(i)
			ch <- i
			<-ch
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 1 {
			<-ch
			fmt.Println(i)
			ch <- i
		}
	}()

	wg.Wait()
	close(ch)
}

func f2() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
			ch <- struct{}{}
			<-ch
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-ch
			fmt.Println(i)
			ch <- struct{}{}
		}
	}()

	wg.Wait()
	close(ch)
}

func g(goroutineCount, maxNumber int) {
	taskChan := make(chan int, goroutineCount)
	var wg sync.WaitGroup
	wg.Add(goroutineCount)
	for i := 0; i < goroutineCount; i += 1 {
		go func(id int) {
			defer wg.Done()
			for num := range taskChan {
				fmt.Println(num)
			}
		}(i)
	}

	for i := 1; i <= maxNumber; i += 1 {
		taskChan <- i
	}
	close(taskChan)
	wg.Wait()
}
