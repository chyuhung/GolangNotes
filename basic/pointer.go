package main

import "fmt"

// 指针
func main() {
	helloPointer()
	fmt.Println()
	// 示例
	var num = 4
	fmt.Printf("&num:%v\n", &num)
	squareVal(num)
	squareAdd(&num)
}

func helloPointer() {
	// 创建一个数值型变量
	i, j := 24, 15237
	// 创建一个指针
	p := &i
	fmt.Println(&i)       // i内存地址
	fmt.Println(p)        // p值为i内存地址
	fmt.Println(*p)       // 24
	fmt.Printf("%T\n", p) // *int
	*p = 42
	fmt.Println(i) // 42

	// assign p to &j
	p = &j
	fmt.Println(&j) // j内存地址
	fmt.Println(p)  // 赋值后变成j内存地址
	fmt.Println(*p) //15237
}

func squareVal(v int) {
	v *= v
	fmt.Println(&v, v)
}

func squareAdd(p *int) {
	*p *= *p
	fmt.Println(p, *p)
}
