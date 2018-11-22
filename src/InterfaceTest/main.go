// InterfaceTest project main.go
package main

import (
	"fmt"
)

//定义接口
type CommonInter interface {
	// 定义接口方法
	raiseError()
}

type Person struct {
	name string
	age  int
}

// 实现接口方法，前半部分不能出现方法名
func (person Person) raiseError() {
	fmt.Println("something wrong with person " + person.name)

}

func main() {
	fmt.Println("")
	person := Person{"lty", 17}

	person.raiseError()
}
