package main

import "fmt"
//求平均值
func average(n []float64) (avg float64) {
	sum := 0.0;
	l := len(n);
	for i := 0; i < l; i ++ {
		sum += n[i];
	}

	avg = sum / float64(l);
	return avg;
}

func main() {
	sli := []float64 {
		1, 2, 3, 4, 5,
	}
	fmt.Println(sli, "的平均值=", average(sli))
}
