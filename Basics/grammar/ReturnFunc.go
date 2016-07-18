package main

import "fmt"

func plusTwo() func(int) int {
	return func(n int) int {
		return n + 2;
	}
}

func plusX(x int) func(int) int {
	return func(n int) int {
		return x + n
	}
}


func main() {

	plus2 := plusTwo();

	fmt.Println("增加2调用= ", plus2(2))

	plus5 := plusX(5)

	fmt.Println("增加5调用=", plus5(2))
}
