package tpkg

import "fmt"

func TextPublic1() {
	fmt.Println("公用方法测试1");
	fmt.Println("私有方法调用1");
	TestPrivate();
}