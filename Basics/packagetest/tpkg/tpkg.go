package tpkg

import "fmt"

func TestPublic() {
	fmt.Println("公用方法测试");
	fmt.Println("私有方法调用");
	TestPrivate();
}

func TestPrivate() {
	fmt.Println("私有方法测试");
}