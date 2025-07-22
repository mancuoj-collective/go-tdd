package di

import (
	"fmt"
	"io"
	"net/http"
)

// 没有自己创建 io.Writer（比如直接使用 os.Stdout 或其他具体实现）
// 而是通过参数接收一个外部提供的 io.Writer 实现。
// 控制反转 Inversion of Control
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// 就像此处，Greet 并不知道它在处理 HTTP 响应，只是将输出写入到传入的 writer
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

// 依赖注入的核心思想是将依赖的创建和管理交给调用者，而不是让被调用的函数或对象自己去创建依赖
// 类比现实生活
// 想象一个咖啡机（Greet 函数）需要水（io.Writer）来制作咖啡（输出结果）。
// 没有依赖注入：咖啡机自己去取水（比如直接用自来水 os.Stdout），这限制了它只能用一种水源。
// 使用依赖注入：你给咖啡机提供水（通过参数传入 io.Writer），可以是自来水、矿泉水、或蒸馏水（os.Stdout、bytes.Buffer、http.ResponseWriter），咖啡机只负责冲泡咖啡，不关心水的来源。
