// 定义包名
package main

// 引入系统依赖包 fmt（格式化IO）
import "fmt"

// 定义变量
var num = 2.5
var str1, str2, str3 = "str1", "str2", "str3"

// 主程序入口函数（如果有 init() 函数则会先执行该函数）
func main() {
	num2 := 3.3
	fmt.Print(num, num2, str1, str2, str3)
}