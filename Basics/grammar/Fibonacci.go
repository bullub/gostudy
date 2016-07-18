package main

import "fmt"

//生成斐波纳契数列
func Fibonacci(n uint64) (series []uint64) {

	series = make([]uint64, n)

	if(n == 1) {
		series[0] = uint64(1)
	}

	if(n >= 2) {
		series[0], series[1] = uint64(1),uint64(1)
	}

	for i := uint64(2); i < n; i ++ {
		series[i] = series[i - 1] + series[i - 2]
	}
	return
}


//func (s slice) String() (str string) {
//
//	for index, value := range *s {
//		str += "[" + strconv.Itoa(index) + ":" + strconv.Itoa(value) + "]"
//	}
//
//	return
//}

func main() {

	n := uint64(100)

	fiboSeries := Fibonacci(n)

	fmt.Println(n, "项斐波那契数列: ", fiboSeries)

}
