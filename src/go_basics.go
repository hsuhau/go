/*
1.包
每个 Go 程序都是由包构成的。

程序从 main 包开始运行。

本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。

按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。

*注意：* 此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。 （要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed。 练习场中的时间为常量，因此你需要用其它的值作为种子数。）
*/
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*2.导入
此代码用圆括号组合了导入，这是“分组”形式的导入语句。

当然你也可以编写多个导入语句，例如：

import "fmt"
import "math"
不过使用分组导入语句是更好的形式。*/

func add1(x int, y int) int {
	return x + y
}
func add2(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func spilt(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("that's12 not cool!")

	/*3.导出名
	在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，Pizza 就是个已导出名，Pi 也同样，它导出自 math 包。

	pizza 和 pi 并未以大写字母开头，所以它们是未导出的。

	在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。

	执行代码，观察错误输出。

	然后将 math.pi 改名为 math.Pi 再试着执行一次。*/
	fmt.Println(math.Pi)

	/*4.函数
	函数可以没有参数或接受多个参数。

	在本例中，add 接受两个 int 类型的参数。

	注意类型在变量名 之后。

	（参考 这篇关于 Go 语法声明的文章了解这种类型声明形式出现的原因。）*/
	fmt.Println(add1(1, 2))

	/*5.函数（续）
	当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。

	在本例中，

	x13 int, y13 int
	被缩写为

	x13, y13 int*/
	fmt.Println(add2(1, 2))

	/*
		6.多值返回
		函数可以返回任意数量的返回值。
		swap 函数返回了两个字符串。
	*/
	fmt.Println(swap("hello", "world"))

	/*
		7.命名返回值
		Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

		返回值的名称应当具有一定的意义，它可以作为文档使用。

		没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。

		直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
	*/
	fmt.Println(spilt(17))

	/*8.变量
	var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。

	就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。*/
	var i int
	fmt.Println(i, c, python, java)

	/*9.变量的初始化
	变量声明可以包含初始值，每个变量对应一个。

	如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。*/
	var c9, python9, java9 = true, false, "no!"
	fmt.Println(i, j, c9, python9, java9)

	/*10.短变量声明
	在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。

	函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。*/
	var i10, j10 int = 1, 2
	k10 := 3
	c10, python10, java10 := true, false, "cool!"
	fmt.Println(i10, j10, k10, c10, python10, java10)

	/*11.基本类型
	Go 的基本类型有

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // uint8 的别名

	rune // int32 的别名
	    // 表示一个 Unicode 码点

	float32 float64

	complex64 complex128
	本例展示了几种类型的变量。 同导入语句一样，变量声明也可以“分组”成一个语法块。

	int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
	*/
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	/*12.零值
	没有明确初始值的变量声明会被赋予它们的 零值。

	零值是：

	数值类型为 0，
	布尔类型为 false，
	字符串为 ""（空字符串）。*/
	var i12 int
	var f12 float64
	var b12 bool
	var s12 string
	fmt.Printf("%v %v %v %q\n", i12, f12, b12, s12)

	/*13.类型转换
	表达式 T(v) 将值 v 转换为类型 T。

	一些关于数值的转换：

	var i int = 42
	var f12 float64 = float64(i)
	var u uint = uint(f12)
	或者，更加简单的形式：

	i := 42
	f12 := float64(i)
	u := uint(f12)
	与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。试着移除例子中 float64 或 uint 的转换看看会发生什么。*/
	var x13, y13 int = 3, 4
	var f13 float64 = math.Sqrt(float64(x13*x13 + y13*y13))
	var z13 uint = uint(f13)
	fmt.Println(x13, y13, z13)

	/*14.类型推导
	在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。

	当右值声明了类型时，新变量的类型与其相同：

	var i int
	j := i // j 也是一个 int
	不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度：

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	尝试修改示例代码中 v 的初始值，并观察它是如何影响类型的。*/
	v14 := 32
	fmt.Printf("v14 is the type %T\n", v14)

	/*15.常量
	常量的声明与变量类似，只不过是使用 const 关键字。

	常量可以是字符、字符串、布尔值或数值。

	常量不能用 := 语法声明。*/
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules ?", Truth)

	/*16.数值常量
	数值常量是高精度的 值。

	一个未指定类型的常量由上下文来决定其类型。

	再尝试一下输出 needInt(Big) 吧。

	（int 类型最大可以存储一个 64 位的整数，有时会更小。）

	（int 可以存放最大64位的整数，根据平台不同有时会更少。）*/
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
