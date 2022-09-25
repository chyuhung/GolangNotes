package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0, 5)
	a = append(a, 1)
	b := append(a, 2)
	c := append(a, 3)
	fmt.Printf("v=%v || p=%p\n", a, &a) // v=[1] || p=0xc00000c060
	fmt.Printf("v=%v || p=%p\n", b, &b) // v=[1 3] || p=0xc00000c080
	fmt.Printf("v=%v || p=%p\n", c, &c) // v=[1 3] || p=0xc00000c0a0
}

// Why? Slice is a struct
//type slice struct {
//	array unsafe.Pointer
//	len   int
//	cap   int
//}

// the function append
//func append(slice []Type, elems ...Type) []Type

//因为 Go 语言内置函数 append 参数是值传递，所以 append 函数在追加新元素到切片时，append 会生成一个新切片，并且将原切片的值拷贝到新切片。
//在 Part 02 示例代码中，我们三次使用 append 参数追加新元素到切片 a 的操作，接收返回值的变量都不同。
//第二次操作时，因为 append 生成一个新切片，将原切片 a 的值拷贝到新切片，并且将新元素在原切片a[len(a)] 长度的位置开始追加，使用变量 b 接收 append 返回值 [1 2]，所以变量 b 的值是 [1 2]。
//第三次操作时，同样 append 生成一个新切片，将原切片 a 的值拷贝到新切片，并且将新元素在原切片a[len(a)] 长度的位置开始追加，使用变量 c 接收 append 返回值 [1 3]，所以变量 c 的值是 [1 3]。
//但是，因为三个切片的底层数组相同，Go 内置函数 append 会在原切片长度的位置开始追加新元素，所以第三次操作时，把第二次操作时得到的变量 b 的最后一个元素覆盖了。
