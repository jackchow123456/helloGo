package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello World!")

	// var name string = "提伯斯"

	// fmt.Println("Hello " + name + " !")

	// traversalString()

	// var zhengxing int = 10

	// fmt.Printf("type:%T,value:%d", zhengxing, zhengxing)
	// fmt.Println()

	// countString()

	// name := "aabcc"

	// for i := 0; i < len(name); i++ {
	// 	fmt.Println(name[i])
	// }

	// changeString()
	// sqrtDemo()

	// findTheSame()

	// ifDemo()

	// gotoDemo()

	// breakDemo()

	// 打印99乘法表
	// printTheMul()

	// 数组的demo
	// 数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。 基本语法：
	// var testArray [3]int
	// var numArray = [3]int{1, 2}
	// var cityArray = [3]string{"北京", "上海", "成都"}
	// fmt.Println(testArray)
	// fmt.Println(numArray)
	// fmt.Println(cityArray)

	// 二维数组的定义
	// aryDemo2()

	// 练习题 - 求数组[1, 3, 5, 7, 8]所有元素的和
	// sumArtNum()

	// 找出数组中和为指定值的两个元素的下标
	// findTheAryKey(8)

	// 切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	// 切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
	// 切片声明
	// var testa []string
	// fmt.Println(testa)

	// 切片的表达
	// var testb = [5]int{1, 2, 3, 4, 5}
	// var b = testb[1:3]
	// fmt.Printf("b:%v,len(b):%v,cap(b):%v", b, len(b), cap(b))

	// 使用make()函数构造切片
	// a := make([]int, 2, 10)
	// fmt.Println(a)
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	// append()方法为切片添加元素
	// var s []int
	// s = append(s, 1, 2, 3, 4)
	// fmt.Println(s)
	// s2 := []int{5, 6, 7}
	// s = append(s, s2...)
	// fmt.Println(s)

	// 从切片中删除元素
	// a := []int{1, 2, 3}
	// b := append(a[:1], a[2:]...)
	// fmt.Println(b)

	// 练习题 - 请写出下面代码的输出结果。
	// var a = make([]string, 5, 10)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a)

	// 请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序（附加题，自行查资料解答）
	// a := [...]int{3, 7, 8, 9, 1}
	// b := a[:]
	// sort.Ints(b)
	// fmt.Println(b)

	// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	// scoreMap := make(map[string]int, 8)
	// scoreMap["小明"] = 8
	// scoreMap["张三"] = 100
	// fmt.Println(scoreMap)

	// 判断某个键是否存在
	// sourceMap := make(map[string]int, 8)
	// sourceMap["小明"] = 8
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	// v, ok := sourceMap["小明"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("查无此人")
	// }

	// 使用delete()函数删除键值对
	// sourceMap := make(map[string]int, 8)
	// sourceMap["小明"] = 8
	// sourceMap["张三"] = 100
	// fmt.Println(sourceMap)
	// delete(sourceMap, "小明")
	// fmt.Println(sourceMap)
	// for k, v := range sourceMap {
	// 	fmt.Println(k, v)
	// }

	// 按照指定顺序遍历map
	// var scoureMap = make(map[string]int)
	// for i := 0; i < 100; i++ {
	// 	key := fmt.Sprintf("stu%02d", i)
	// 	value := rand.Intn(100) //生成0~99的随机整数
	// 	scoureMap[key] = value
	// }

	// // 新建切片
	// keys := make([]string, 0, 200)

	// for k := range scoureMap {
	// 	keys = append(keys, k)
	// }

	// // 排序
	// sort.Strings(keys)

	// fmt.Println(keys)
	// for _, v := range keys {
	// 	fmt.Println(v, scoureMap[v])
	// }

	// 元素为map类型的切片
	// var mapSlice = make([]map[string]string, 3)
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }
	// fmt.Println("after init")
	// mapSlice[0] = make(map[string]string, 10)
	// mapSlice[0]["name"] = "小周"
	// mapSlice[0]["age"] = "26岁"
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }

	// 值为切片类型的map
	// var sliceMap = make(map[string][]string, 3)
	// fmt.Println(sliceMap)
	// fmt.Println("after init")
	// key := "中国"
	// value, ok := sliceMap[key]
	// if !ok {
	// 	value = make([]string, 0, 2)
	// }
	// value = append(value, "北京", "上海")
	// sliceMap[key] = value
	// fmt.Println(sliceMap)

	// 练习题
	// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	// value := "how do you do"
	// var sliceMapDemo = make(map[string]int)
	// demo := strings.Split(value, " ")
	// // fmt.Printf("%T", demo)
	// for _, v := range demo {
	// 	sliceMapDemo[v]++
	// }
	// for k, v := range sliceMapDemo {
	// 	fmt.Println(k, v)
	// }
	// 	the_limit := len(value)
	// 	// fmt.Printf("the_limit:%v\n", the_limit)
	// 	var str string
	// 	var sliceMap = make(map[string]int)
	// forloop1:
	// 	for k, v := range value {
	// 		if v == 32 {
	// 			sliceMap[str]++
	// 			fmt.Printf("the_value:%v\n", str)
	// 			str = ""
	// 			continue forloop1
	// 		} else {
	// 			str = fmt.Sprintf("%v%v", str, v)
	// 		}
	// 		// fmt.Printf("key:%v,value:%v\n", k, v)
	// 		if k+1 == the_limit {
	// 			sliceMap[str]++
	// 			fmt.Printf("the_value:%v\n", str)
	// 		}
	// 	}
	// 	fmt.Println(sliceMap)

	// 观察下面代码，写出最终的打印结果。
	// type Map map[string][]int
	// m := make(Map)
	// s := []int{1, 2}
	// s = append(s, 3)
	// fmt.Printf("%+v\n", s)
	// m["q1mi"] = s
	// s = append(s[:1], s[2:]...)
	// fmt.Printf("%+v\n", s)
	// fmt.Printf("%+v\n", m["q1mi"])

	// 函数的调用
	// sayHello()
	// ret := intSum(10, 20)
	// fmt.Println(ret)

	// 可变参数
	// ret1 := intSum2()
	// ret2 := intSum2(10)
	// ret3 := intSum2(10, 20)
	// ret4 := intSum2(10, 20, 30)
	// fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60

	// 返回值
	// 多返回值

	// 函数类型与变量
	// type calculation func(int, int) int

	// var c calculation
	// c = add
	// fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	// fmt.Println(c(1, 2))            // 像调用add一样调用c

	// f := add                        // 将函数add赋值给变量f1
	// fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	// fmt.Println(f(10, 20))          // 像调用add一样调用f

	// 高阶函数
	// 函数作为参数
	// func add(x, y int) int {
	// 	return x + y
	// }
	// func calc(x, y int, op func(int, int) int) int {
	// 	return op(x, y)
	// }
	// func main() {
	// 	ret2 := calc(10, 20, add)
	// 	fmt.Println(ret2) //30
	// }

	// 函数作为返回值
	// func do(s string) (func(int, int) int, error) {
	// 	switch s {
	// 	case "+":
	// 		return add, nil
	// 	case "-":
	// 		return sub, nil
	// 	default:
	// 		err := errors.New("无法识别的操作符")
	// 		return nil, err
	// 	}
	// }

	// 匿名函数和闭包
	// 函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：
	// 匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

	// 将匿名函数保存到变量
	// add := func(x, y int) {
	// 	fmt.Println(x + y)
	// }
	// add(10, 20) // 通过变量调用匿名函数

	// //自执行函数：匿名函数定义完加()直接执行
	// func(x, y int) {
	// 	fmt.Println(x + y)
	// }(10, 20)

	// 闭包
	// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。 首先我们来看一个例子：
	// func adder() func(int) int {
	// 	var x int
	// 	return func(y int) int {
	// 		x += y
	// 		return x
	// 	}
	// }
	// func main() {
	// 	var f = adder()
	// 	fmt.Println(f(10)) //10
	// 	fmt.Println(f(20)) //30
	// 	fmt.Println(f(30)) //60

	// 	f1 := adder()
	// 	fmt.Println(f1(40)) //40
	// 	fmt.Println(f1(50)) //90
	// }

	// defer语句
	// Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
	// fmt.Println("start")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("end")
	// 输出结果 start end 3 2 1

	// defer经典案例
	// fmt.Println(f1())
	// fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(f4())

	// panic/recover
	// Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。 panic可以在任何地方引发，但recover只有在defer调用的函数中有效。 首先来看一个例子：
	// funcA()
	// funcB()
	// funcC()
}

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 变量作用域
func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 变量作用域
func testLocalVar2(x, y int) {
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
}

func intSum2(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func intSum(x int, y int) int {
	return x + y
}

func sayHello() {
	fmt.Println("Hello 沙河")
}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// 统计字数
func countString() {
	s := "hello沙河小王子"
	num := 0
	for _, r := range s {
		if len(string(r)) >= 3 {
			num++
		}
	}

	fmt.Printf("一共有%d个字", num)
}

// 字符更换
func changeString() {
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "红萝卜"
	runteS2 := []rune(s2)
	runteS2[0] = '白'
	fmt.Println(string(runteS2))
}

// 类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 找出只出现过一次的字符
func findTheSame() {
	var ary = [...]int{1, 2, 3, 4, 5, 4, 2, 1, 5, 2, 1}
	var ary2 = make([]int, len(ary))
	for _, v := range ary {
		ary2[v]++
	}

	for k, v := range ary2 {
		if v == 1 {
			fmt.Println(k, v)
		}
	}
}

// 流程控制 if
func ifDemo() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// goto Demo
func gotoDemo() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 && i == 2 {
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

// break Demo
func breakDemo() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 && i == 2 {
				break BREAKDEMO1
			}

			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

// 打印99乘法表
func printTheMul() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

// 二维数组的定义
func aryDemo2() {
	a := [3][2]string{
		{"北京", "上海"},
		{"北京", "上海"},
		{"北京", "上海"},
	}
	fmt.Println(a)
}

// 计算数组内元素之和
func sumArtNum() {
	sum := 0
	ary := [...]int{1, 3, 5, 7, 8}
	for _, v := range ary {
		sum += v
	}
	fmt.Println(sum)
}

// 找出数组中和为指定值的两个元素的下标
func findTheAryKey(x int) {
	ary := [...]int{1, 3, 5, 7}
	for i := 0; i < len(ary); i++ {
		for j := i + 1; j < len(ary); j++ {
			if ary[i]+ary[j] == x {
				fmt.Println(i, j)
			}
		}
	}
}
