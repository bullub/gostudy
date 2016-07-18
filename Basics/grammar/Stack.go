package main

import (
	"fmt"
	"strconv"
)

type stack struct{
	length 	int
	data	[10]int
}


func (s *stack) push(data int) {
	s.data[s.length] = data;
	s.length ++;
}

func (s *stack) pop() int {
	if(s.length < 1) {
		return 0;
	}
	s.length --;

	return s.data[s.length]
}

func (s *stack) String() (str string) {

	for index := 0; index < s.length; index ++ {
		str += "[" + strconv.Itoa(index) + ":" + strconv.Itoa(s.data[index]) + "] "
	}

	return
}

func main() {

	s := new(stack);

	s.push(10);
	fmt.Println("stack = ", s);

	s.push(100);

	fmt.Println("stack = ", s);

	s.pop();
	fmt.Println("stack = ", s);

}
