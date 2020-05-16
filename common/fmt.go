package common

import "fmt"
type point struct {
	x, y int
}


/*
remeber: fmt 各种常见的输出格式
*/

func fmtPrintf(){
	p := point{1, 2}
	fmt.Printf("%v\n", p) // {1 2}
	//如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	fmt.Printf("%+v\n", p) // {x:1 y:2}
	//%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
	//需要打印值的类型，使用 %T。
	fmt.Printf("%T\n", p) // main.point
}