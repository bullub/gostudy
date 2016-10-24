package main

import (
	"fmt"
)

func main() {
	//c := make(chan bool);
	//go func() {
	//	fmt.Println("Go Go Go!!!");
	//	c <- true;
	//	close(c);
	//}()
	////<-c;
	//fmt.Println("dddd");
	//for v := range c {
	//	fmt.Println(v);
	//}

	//runtime.GOMAXPROCS(runtime.NumCPU());
	//
	//c := make(chan bool, 10)
	//
	//for i:=0; i<10; i ++ {
	//	go Go(c, i)
	//}
	//
	//for i:=0; i<10; i ++ {
	//	<-c
	//}



	//runtime.GOMAXPROCS(runtime.NumCPU());
	//wg := sync.WaitGroup{}
	//wg.Add(10);
	//
	//for i := 0; i < 10; i ++ {
	//	go Go(&wg, i);
	//}
	//
	//wg.Wait();

	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)

	go func() {
		for {
			select {
			case v, ok := <- c1:
				if !ok {
					o <- true
					break;

				}
				fmt.Println("c1", v)
			case v, ok := <- c2:
				if !ok {
					o <- true
					break;
				}
				fmt.Println("c2", v)
			}
		}
	}()


	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"
	close(c1);
	close(c2);

	for i := 0; i < 2; i ++ {
		<-o
	}
}

//func Go(c chan bool, index int)  {
//
//	a:=1;
//	for i:=0; i<100000; i ++ {
//		a += i
//	}
//
//	fmt.Println(index, a)
//
//	//if index == 9 {
//		c <- true
//	//}
//}

//func Go(wg *sync.WaitGroup, index int) {
//	a:=1;
//	for i:=0; i<100000; i ++ {
//		a += i
//	}
//
//	fmt.Println(index, a)
//
//	wg.Done();
//}