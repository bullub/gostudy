package main

import "fmt"

func _map(f func(int) int, series []int) []int {
	series1 := make([]int, len(series))

	//copy(series1, series)

	for index, value := range series {
		series1[index] = f(value)
	}
	return series1
}

func _mapStr(f func(string) string, series []string) []string {
	series1 := make([]string, len(series))

	//copy(series1, series)

	for index, value := range series {
		series1[index] = f(value)
	}
	return series1
}

func main() {

	series := []int{
		1,2,3,4,
	}

	series1 := []string{
		"hh", "mm",
	}

	fmt.Println("数列:", series, "执行map操作后的结果=", _map(func (v int) int{
		return v * 2;
	}, series))

	fmt.Println("数列:", series1, "执行map操作后的结果=", _mapStr(func (v string) string{
		return v+"heh";
	}, series1))

}
