// helloword
package main

import (
	"fmt"
)

//全局变量声明，在未赋值时，应该指明变量的数据类型
//全局变量可以仅声明不使用
var v1 int
var v2 int

//以上两行可以精简如下:
// var v1, v2 int

//以下声明方式一般用于全局变量

var (
	g_var1 string
	g_var2 bool
)

// 全局常量声明
const c1 = "const1"

//const 用于枚举
const (
	MALE    = 1
	FEMALE  = 0
	UNKNOWN = -1
)

func init() {
	/* 如果存在init,会在main之前执行*/
	fmt.Println("main 之前...")
	g_var1 = "global var1"
	fmt.Println(g_var1)
}
func print() {
	// 变量声明时可以不指定类型，由编译器推断，但是只能在函数内部使用
	//即为：初始化声明方式 不声明数据类型 使用:= 直接赋值，由编译器推断数据类型，这种方式不能对同一个变量声明两次
	//但是可以对变量重新赋值, 局部变量不允许仅声明不使用
	a := 1
	b := 2
	b = 500
	fmt.Println("a+b=", a+b)
	// 交换两个同类型变量的值
	a, b = b, a
	fmt.Println("交换a,b值后a=", a)
}

func testCondition() {
	a := 2
	b := 3
	if a > b {
		fmt.Println(a, ">", b)
	} else if a == b { // 格式要求：else不能换行
		fmt.Println(a, "=", b)
	} else {
		fmt.Println(a, "<", b)
	}
}

//函数定义
func add(a int, b int) int {
	return a + b
}

//函数闭包，声明一个返回值为函数的函数（闭包函数）,闭包函数也可以有参数和返回值
func getSequence() func() int {
	i := 0
	//匿名函数
	return func() int {
		i += 1
		return i

	}
}
func testArray() {
	//声明数组需要指定长度
	var ints [10]int
	var i int
	for i = 0; i < 10; i++ {
		ints[i] = i
		fmt.Print(ints[i], " ")
	}
	fmt.Println()

}

// 定义结构体
type Person struct {
	name string
	age  int
}

func descPerson(person Person) {
	// 结构体的成员访问方式跟c语言类似
	fmt.Print("name: ", person.name)
	fmt.Println(" age: ", person.age)

}
func main() {
	var a = "a string"
	//获取变量a的地址
	fmt.Println(&a)
	// 声明指针变量
	var pa *string = &a
	fmt.Println("*pa=", *pa)
	fmt.Println("Hello World!")
	fmt.Println(a)
	print()
	// var a:="another string" :=符号标记的变量此前不能声明过，否则会报错
	fmt.Println(c1)
	testCondition()
	fmt.Println(add(1, 2))

	// 这里返回了一个匿名函数，它可以直接使用getSequence函数的内部变量
	nextNum := getSequence()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	testArray()
	descPerson(Person{"lty", 18})

}
