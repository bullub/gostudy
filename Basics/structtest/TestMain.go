package main

import (
	"./tstruct"
	"fmt"
)

func main() {
	var a *tstruct.A = new(tstruct.A);
	fmt.Printf("x=%d,y=%d", a.GetPosition().X, a.GetPosition().Y);
}
