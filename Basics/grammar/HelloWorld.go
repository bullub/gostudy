package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i ++ {
		fmt.Printf("%d", i);
	}

	fmt.Println();
	var i int;

	i = 0;

	for {
		if i > 10 {
			break;
		}
		i ++;
		fmt.Println(i);
	}

	var iArr [3]int;

	iArr = [3]int{
		'a', 'b', 'c',
	}

	for k, v := range iArr {
		fmt.Println(k, "=", v)
	}

	j := 0;

	L:

	if j < 10 {
		fmt.Println("j=", j);
		j ++;
		goto L;
	}

	fmt.Printf("%d\n", 10)
}
