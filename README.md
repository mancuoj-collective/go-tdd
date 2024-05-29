# Learn Go with Tests

- https://quii.gitbook.io/learn-go-with-tests
- https://tour.go-zh.org/list
- https://gobyexample.com

## 包、变量与函数

- 每个 Go 程序都由包构成
- 程序从 main 包开始运行
- 包名与导入路径最后一个元素一致

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("我最喜欢的数字是 ", rand.Intn(10))
}
```

- 可以「分组」导入也可以分开导入

```go
import "fmt"
import "math"
```

- 如果一个名字以大写字母开头，那么它就是导出的

```go
fmt.Println(math.Pi)
```

- 函数可接受零个或多个参数
- 多个参数类型相同时可以省略前面的类型定义
- 函数可以返回任意数量的返回值

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

- 返回值可被命名，它们会被视作定义在函数顶部的变量
- 没有参数的 `return` 语句会直接返回已命名的返回值，也就是「裸」返回值

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

- `var` 语句用于声明一系列变量，可以出现在包或函数的层级
- 如果提供了初始值，则类型可以省略，变量会从初始值中推断出类型
- 也可以「分组」声明

```go
var i, j int = 1, 2
var i, j = 1, 2

var (
	i int = 1
	j int = 2
)
```

- 隐式确定类型的 `var` 声明可以使用 `:=` 短赋值语句
- 不能在函数外使用，函数外的每个语句都必须以关键字开始

```go
func main() {
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
}
```

- Go 的基本类型可以互相显式转换，`T(v)` 表达式将值 `v` 转换为类型 `T`，比如 `float64(42)`
- `int`、`uint` 和 `uintptr` 类型在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名
rune // int32 的别名，表示一个 Unicode 码位

float32 float64

complex64 complex128
```

- 没有明确初始化的变量声明会被赋予对应类型的零值
  - 数值类型为 `0`
  - 布尔类型为 `false`
  - 字符串为 `""`，也就是空字符串

```go
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

- 常量的声明与变量类似，只不过使用 `const` 关键字
- 不能用 `:=` 语法声明

```go
const Pi = 3.14
```

## 流程控制语句

- Go 只有一种循环结构：`for` 循环，由三部分组成，它们用分号隔开
  - 初始化语句：在第一次迭代前执行
  - 条件表达式：在每次迭代前求值
  - 后置语句：在每次迭代的结尾执行
- 初始化语句通常为一句短变量声明，仅在 `for` 语句的作用域中可见
- 一旦条件表达式求值为 `false`，循环迭代就会终止

```go
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

- 初始化语句和后置语句是可选的
- 都没有时可以去掉分号变为 `while` 循环

```go
for sum < 1000 {
	sum += sum
}
```

- 省略循环条件会变成无限循环

```go
for {
}
```

- `if, else if, else` 语句的语法与其他语言类似
- 可以在条件表达式前执行一个简短语句，该语句声明的变量作用域仅在 `if` 之内

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```

- `switch` 语句会在每个 `case` 后自动添加 `break`，除非使用 `fallthrough` 关键字
- `switch` 语句的表达式可以是任意类型，不一定是整数
- `case` 语句可以是一个值列表，用逗号分隔，表示多个值可以匹配

```go
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("macOS.")
case "linux":
	fmt.Println("Linux.")
default:
	fmt.Printf("%s.\n", os)
}
```

- 无条件 `switch` 语句可以使一长串的 `if-else` 语句更加清晰

```go
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}
```

- `defer` 语句会将函数推迟到外层函数返回之后执行
- 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
- 推迟调用的函数调用会被压入一个栈中，当外层函数返回时，被推迟的调用会按照后进先出的顺序调用

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
// counting
// done
// 9
// 8
// 7
// ...
// 0
```

## 结构体、数组、切片和映射

- 指针保存了值的内存地址，类型 `*T` 是指向 `T` 类型值的指针，其零值为 `nil`
- `&` 操作符会生成一个指向其操作数的指针
- `*` 操作符表示指针指向的底层值
- 与 C 不同，Go 没有指针运算

```go
func main() {
	i := 42
	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值
}
```

- 一个结构体（struct）就是一组字段（field）
- 字段可通过点号 `.` 来访问，或者通过结构体指针来访问

```go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	p := &v
	p.X = 1e9 // 隐式解引用 (*p).X
	fmt.Println(v.X)
}
```

- 使用 `Name:` 语法可以仅列出部分字段，顺序无关

```go
var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予零值
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)
```

- 类型 `[n]T` 表示一个数组，它拥有 n 个类型为 `T` 的值
- 数组的长度是其类型的一部分，因此数组不能改变大小

```go
var a [10]int
primes := [6]int{2, 3, 5, 7, 11, 13}
```

- 切片为数组元素提供了动态大小的、灵活的视角，所以实践中，切片比数组更常用
- 类型 `[]T` 表示一个元素类型为 `T` 的切片
