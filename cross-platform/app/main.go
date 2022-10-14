package main

import (
	"fmt"
)

func main() {
	fmt.Println("跨平台测试案例")
	dir := GetDefaultUserHomePath()
	fmt.Println("当前用户家目录:" + dir)
}
