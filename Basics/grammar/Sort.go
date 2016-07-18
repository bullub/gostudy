package main

import "fmt"

//冒泡排序
func bubbleSort(args ...int64) (queue []int64) {

	argsLen := len(args);
	queue = make([]int64, argsLen);
	copy(queue, args[:]);

	for i := 0; i < argsLen; i ++ {
		for j := 0; j < argsLen; j ++ {
			if(queue[i] < queue[j]) {
				queue[i], queue[j] = queue[j], queue[i]
			}
		}

	}
	return ;
}

//快速排序
func quickSort(args ...int64) (queue []int64) {

	right := len(args)
	queue = make([]int64, right)
	copy(queue, args[:])

	if(right < 2) {
		return ;
	}

	right--
	left := 0
	position := right / 2


	for left < right {
		//从右向左找,如果右边的比当前位置的大,则略过
		for left < right && queue[position] < queue[right] {
			right --
		}

		//走到这里,如果right还是比left大,则表示找到了比positon位置的数小的数了
		if left < right {
			//找到了比position位置的数大的数了,将它们交换
			queue[position], queue[right] = queue[right], queue[position]
			//设置当前位置为被交换的数字的位置
			position = right
		}

		//从左向右找,如果左边的比当前位置的小,略过
		for left < right && queue[position] >= queue[left] {
			left ++
		}

		//如果执行到这里,left还是比right小,则表示left位置的数一定比position处的数要大,则调换它们的位置
		if(left < right) {
			//找到了比position位置的数小的数了,将它们交换
			queue[position], queue[left] = queue[left], queue[position]
			//设置当前位置为被交换的数字的位置
			position = left
		}
	}

	lQ := quickSort(queue[:position]...)
	rQ := quickSort(queue[position + 1:]...)

	queue = append(lQ, queue[position])
	queue = append(queue, rQ...)

	return ;

}
//插入排序
func insertSort(args ...int64) (queue []int64) {



	length := len(args)

	queue = make([]int64, length)

	copy(queue, args[:])

	for i := 1; i < length; i ++ {

		for j := i - 1; j >= 0 && queue[j] > queue[j + 1]; j -- {
			queue[j], queue[j + 1] = queue[j + 1], queue[j];
		}

	}
	return

}


func main() {
	int64Arr := []int64{
		1,6,3,9,6,7,3,4,9,1,6,3,9,6,7,3,4,9,1,6,3,9,6,7,3,4,9,1,6,3,9,6,7,3,4,9,
	};
	fmt.Println(int64Arr, "排序结果: ", insertSort(int64Arr...));
	fmt.Println(len(int64Arr), "排序结果: ", len(insertSort(int64Arr...)));
}
