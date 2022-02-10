package chapter

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Printf("hello %d\n", i)
}

func GoHello() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

func Channel() {
	ch := make(chan int, 1)
	ch <- 1
	<-ch
	ch <- 2
	x := <-ch
	fmt.Println(x)
	close(ch)
}

func ForRange() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
	}()
	time.Sleep(1000)
	go func() {
		defer close(ch2)
		for {
			// 当值取完ok为false
			i, ok := <-ch1
			fmt.Printf("ch1 i=%#v\n", i)
			if !ok {
				fmt.Printf("ch1 closed i=%#v\n", i)
				if i == 0 {
					break
				}
			}
			ch2 <- i
		}
	}()

	for j := range ch2 {
		fmt.Printf("ch2 j=%#v\n", j)
	}
}

func Select() {
	ch1 := make(chan int, 1)

	go func() {
		for i := 0; ; i++ {
			if rand.Int()%2 == 0 {
				ch1 <- i
			}
			fmt.Printf("noting %d\n", i)
		}

	}()
	for i := 0; i < 100; i++ {
		select {
		case x := <-ch1:
			fmt.Printf("x=%#v\n", x)
		case y := <-ch1:
			fmt.Printf("y=%#v\n", y)
		}
	}
}
