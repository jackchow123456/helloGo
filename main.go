package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth", "",
	}
	distribution = make(map[string]int, len(users))
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

	/** Go语言基础之数组 **/
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

	/** Go语言基础之切片 **/
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

	/** Go语言基础之map **/
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

	/** Go语言基础之函数 **/
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

	// 练习题 - 分金币
	/*
		你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
		分配规则如下：
		a. 名字中每包含1个'e'或'E'分1枚金币
		b. 名字中每包含1个'i'或'I'分2枚金币
		c. 名字中每包含1个'o'或'O'分3枚金币
		d: 名字中每包含1个'u'或'U'分4枚金币
		写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
		程序结构如下，请实现 ‘dispatchCoin’ 函数
	*/

	// left := dispatchCoin()
	// fmt.Println("剩下：", left)

	/** Go语言基础之指针 **/

	// 指针地址和指针类型
	// a := 10
	// b := &a
	// fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	// fmt.Printf("b:%d, b:%p type:%T\n", b, b, b) // b:0xc00001a078 type:*int
	// fmt.Println(&b)                    // 0xc00000e018

	// 指针取值
	// a := 10
	// b := &a // 取变量a的地址，将指针保存到b中
	// fmt.Printf("type of b:%T\n", b)
	// c := *b // 指针取值（根据指针去内存取值）
	// fmt.Printf("type of c:%T\n", c)
	// fmt.Printf("value of c:%v\n", c)
	// 总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

	// new 和 make
	// var a *int
	// *a = 100
	// fmt.Println(*a)

	// var b map[string]int
	// b["沙河娜扎"] = 100
	// fmt.Println(b)

	// 执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
	// 分配内存空间
	// var a *int
	// a = new(int)
	// *a = 10
	// fmt.Println(*a)

	// 而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。
	// 分配内存
	// var b map[string]int
	// b = make(map[string]int, 10)
	// b["沙河娜扎"] = 100
	// fmt.Println(b)

	// new与make的区别
	// 1. 二者都是用来做内存分配的
	// 2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	// 3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

	/** Go语言基础之结构体 **/
	// 结构体实例化
	// 基本实例化
	type person struct {
		name string
		city string
		age  int8
	}

	// var p1 person

	// p1.name = "沙河娜扎"
	// p1.city = "北京"
	// p1.age = 18
	// fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
	// fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}

	// 匿名结构体
	// var user struct {Name string; Age int}
	// user.Name = "xhz"
	// user.Age = 18
	// fmt.Printf("%#v\n",user)

	// 创建指针类型结构体
	// var p2 = new(person)
	// fmt.Printf("%T\n",p2)
	// fmt.Printf("p2=%#v\n",p2)
	// p2.name = "小伙子"
	// p2.age = 18
	// p2.city = "上海"
	// fmt.Printf("p2=%#v\n",p2)

	// 取结构体的地址实例化
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	// p3 := &person{}
	// fmt.Printf("%T\n",p3)
	// fmt.Printf("p3=%#v\n",p3)
	// p3.name = "七米"
	// p3.age = 30
	// p3.city = "成都"
	// fmt.Printf("p3=%#v\n",p3)

	// 结构体初始化
	// 没有初始化的结构体，其成员变量都是对应其类型的零值。
	// var p4 person
	// fmt.Printf("p4=%#v\n",p4)

	// 使用键值对初始化
	// p5 := person{
	// 	name:"小王子",
	// 	city:"北京",
	// 	age:18,
	// }
	// fmt.Printf("p5=%#v\n",p5)

	// 结构体内存布局
	// 结构体占用一块连续的内存。
	// type test struct {
	// 	a int8
	// 	b int8
	// 	c int8
	// 	d int8
	// }
	// n := test{
	// 	1, 2, 3, 4,
	// }
	// fmt.Printf("n.a %p\n", &n.a)
	// fmt.Printf("n.b %p\n", &n.b)
	// fmt.Printf("n.c %p\n", &n.c)
	// fmt.Printf("n.d %p\n", &n.d)
	// 输出
	// 	n.a 0xc0000a0060
	// n.b 0xc0000a0061
	// n.c 0xc0000a0062
	// n.d 0xc0000a0063

	// 空结构体 - 空结构体是不占用空间的。
	var v struct{}
	fmt.Println(unsafe.Sizeof(v))

	// 面试题 - 请问下面代码的执行结果是什么？
	// type student struct {
	// 	name string
	// 	age  int
	// }

	// m := make(map[string]*student)
	// stus := []student{
	// 	{name: "小王子", age: 18},
	// 	{name: "娜扎", age: 23},
	// 	{name: "大王八", age: 9000},
	// }

	// for _, stu := range stus {
	// 	m[stu.name] = &stu
	// }
	// for k, v := range m {
	// 	fmt.Println(k, "=>", v.name)
	// }

	// 构造函数 newPerson
	// Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。

	// 调用构造函数
	// p9:= newPerson("张三", "沙河", 90)
	// fmt.Printf("%#v\n",p9)

	// 方法和接收者
	// p1 := newPerson("周周", "沙河", 18)
	// p1.dream()

	// 指针类型的接收者
	// p2 := newPerson("周周", "广州", 28)
	// fmt.Printf("周周的年龄是：%d\n",p2.age)
	// p2.setAge(18)
	// fmt.Printf("周周的年龄是：%d\n",p2.age)

	// 任意类型添加方法
	// 在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。
	// var m1 MyInt
	// m1.sayHello()
	// m1 = 100
	// fmt.Printf("%#v,%T\n",m1,m1)

	// 结构体的匿名字段
	// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
	// p := person1{"周周",18}
	// fmt.Printf("%#v\n",p)
	// 注意：这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。

	// 嵌套结构体
	// 一个结构体中可以嵌套包含另一个结构体或结构体指针，就像下面的示例代码那样。
	// address := Address{"广东省","广州市"}
	// user := User{"周周",18,address}
	// fmt.Printf("%#v\n",user)

	// 嵌套匿名字段
	// // 上面user结构体中嵌套的Address结构体也可以采用匿名字段的方式，例如：
	// user := User{"周周",18,Address{}}
	// fmt.Printf("%#v\n", user)
	// user.Address.Province = "广东"
	// fmt.Printf("%#v\n", user)
	// user.City = "广州"
	// fmt.Printf("%#v\n", user)
	// 当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。

	// 嵌套结构体的字段名冲突
	// 嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。
	// var user User
	// user.Name = "周周"
	// user.Age = 18
	// // user.CreateTime = "2021-12-02 14:47:52" // ambiguous selector user.CreateTim
	// user.Address.CreateTime = "2021-12-02 14:47:52"
	// fmt.Printf("%#v\n", user)

	// 结构体的“继承”
	// Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
	//  dog := &Dog{4,&Animal{"乐乐"}}
	//  dog.move()
	//  dog.wang()

	// 结构体字段的可见性
	// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

	// 结构体与JSON序列化
	// c := &Class{
	// 	Title:    "101",
	// 	Students: make([]*Student, 0, 200),
	// }
	// for i := 0; i < 10; i++ {
	// 	stu := &Student{
	// 		Name:   fmt.Sprintf("stu%02d", i),
	// 		Gender: "男",
	// 		ID:     i,
	// 	}
	// 	c.Students = append(c.Students, stu)
	// }

	// //JSON序列化：结构体-->JSON格式的字符串
	// data, err := json.Marshal(c)
	// if err != nil {
	// 	fmt.Println("json marshal failed")
	// 	return
	// }
	// fmt.Printf("json:%s\n", data)

	// str := data
	// c1 := &Class{}
	// //JSON反序列化：JSON格式的字符串-->结构体
	// err = json.Unmarshal([]byte(str), c1)
	// if err != nil {
	// 	fmt.Println("json unmarshal failed!")
	// 	return
	// }
	// fmt.Printf("%#v\n", c1)

	// 结构体标签（Tag）
	// Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：`key1:"value1" key2:"value2"`
	// 结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。
	// 例如我们为Student结构体的每个字段定义json序列化时使用的Tag：

	// s1 := Student{
	// 	ID:     1,
	// 	Gender: "男性",
	// 	Name:   "小明",
	// }
	// data, err := json.Marshal(s1)
	// if err != nil {
	// 	fmt.Println("json marshal failed.")
	// 	return
	// }
	// fmt.Printf("json str:%s\n", data)

	// 结构体和方法补充知识点
	// 因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。我们来看下面的例子：
	// p := Person{name: "周周", age: 18}
	// data := []string{"吃饭", "睡觉"}
	// p.setDreams(data)
	// data[1] = "不睡觉"
	// fmt.Println(p.dreams)

	// 练习题
	// 使用“面向对象”的思维方式编写一个学生信息管理系统
	// 1. 学生有id、姓名、年龄、分数等信息
	// 2. 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能

	/** Go语言基础之接口 **/
	// 接口的定义

	// 实现接口的条件

	// 接口类型变量
	// var x sayer // 声明一个sayer的变量 x
	// c := cat{}  // 实例化 一个 cat
	// d := dog{}  // 实例化 一个 dog
	// x = c       // 把 cat 实例直接赋值给 x
	// x.say()     // 喵喵喵
	// x = d
	// x.say()

	// Tips： 观察下面的代码，体味此处_的妙用
	// type IRouter interface{}
	// type RouterGroup struct{}
	// var _ IRouter = &RouterGroup{} // 确保RouterGroup实现了接口IRouter

	// 接收者和指针接收者实现接口的区别
	// 不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。

	// 面试题 - 首先请观察下面的这段代码，然后请回答这段代码能不能通过编译？
	// 不能，
	// var peo Peoplet = Studentt{}
	// think := "bitch"
	// fmt.Println(peo.Speak(think))

	// 一个类型实现多个接口
	// var x Sayer
	// var y Mover
	// var a = dogg{name:"旺旺"}
	// x = a
	// y = a
	// x.say()
	// y.move()

	// 多个类型实现同一接口
	// var x Mover
	// d := dogg{name:"旺财"}
	// x = d
	// x.move()
	// c := catt{branch:"咖啡猫"}
	// x = c
	// x.move()
	// 并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
	// var x WashingMachine
	// d := Dryer{}
	// h := haier{d}
	// x = h
	// x.wash()
	// x.dry()

	// 接口嵌套
	// 接口与接口间可以通过嵌套创造出新的接口。
	// 嵌套得到的接口的使用与普通接口一样，这里我们让cat实现animal接口：
	// var x animal
	// c := catt{branch:"加菲猫"}
	// x = c
	// x.say()
	// x.move()

	// 空接口
	// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	// 空接口类型的变量可以存储任意类型的变量。
	// var x interface{}
	// s := "Hello 沙河"
	// x = s
	// fmt.Printf("type:%T value:%v\n", x, x)
	// i := 100
	// x = i
	// fmt.Printf("type:%T value:%v\n", x, x)
	// b := true
	// x = b
	// fmt.Printf("type:%T value:%v\n", x, x)

	// 空接口的应用
	// 空接口作为函数的参数
	// 空接口作为map的值
	// var studentInfo = make(map[string]interface{})
	// studentInfo["name"] = "沙河娜扎"
	// studentInfo["age"] = 18
	// studentInfo["married"] = false
	// fmt.Println(studentInfo)

	// 类型断言
	// xx = "Hello 沙河"
	// vv, ok := xx.(string)
	// if ok {
	// 	fmt.Println(vv)
	// } else {
	// 	fmt.Println("类型断言失败")
	// }

	/** Go语言基础之反射 **/
	// reflect包，，任何接口值都由是一个具体类型和具体类型的值两部分组成的(我们在上一篇接口的博客中有介绍相关概念)。
	// 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，
	// 并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。
	// 在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息
	// var aa int32 = 100
	// reflectType(aa)

	// type name和type kind
	// type myInt int64
	// var a *float32 // 指针
	// var b myInt    // 自定义类型
	// var c rune     // 类型别名
	// reflectType(a) // type: kind:ptr
	// reflectType(b) // type:myInt kind:int64
	// reflectType(c) // type:int32 kind:int32

	// var d = person{
	// 	name: "沙河小王子",
	// 	age:  18,
	// }
	// var e = book{title: "《跟小王子学Go语言》"}
	// reflectType(d) // type:person kind:struct
	// reflectType(e) // type:book kind:struct

	// ValueOf
	// reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。

	// var a float32 = 3.14
	// var b int64 = 100
	// reflectValue(a) // type is float32, value is 3.140000
	// reflectValue(b) // type is int64, value is 100
	// // 将int类型的原始值转换为reflect.Value类型
	// c := reflect.ValueOf(10)
	// fmt.Printf("type c :%T\n", c) // type c :reflect.Value

	// 通过反射设置变量的值
	// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
	// var a int64 = 100
	// // reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	// reflectSetValue2(&a)
	// fmt.Println(a)

	// isNil()和isValid()
	// IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
	// IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。

	// // *int类型空指针
	// var a *int
	// fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// // nil值
	// fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// // 实例化一个匿名结构体
	// b := struct{}{}
	// // 尝试从结构体中查找"abc"字段
	// fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// // 尝试从结构体中查找"abc"方法
	// fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// // map
	// c := map[string]int{}
	// // 尝试从map中查找一个不存在的键
	// fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())

	// 结构体反射

	// stu1 := student{
	// 	Name:  "小王子",
	// 	Score: 90,
	// }

	// t := reflect.TypeOf(stu1)
	// fmt.Println(t.Name(), t.Kind()) // student struct
	// // 通过for循环遍历结构体的所有字段信息
	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)
	// 	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	// }

	// // 通过字段名获取指定结构体字段信息
	// if scoreField, ok := t.FieldByName("Score"); ok {
	// 	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	// }

	// printMethod(stu1)

	/** Go语言中的并发编程 **/
	// 	goroutine
	// go hello()
	// fmt.Println("main goroutine done!")
	// time.Sleep(time.Second)

	// 启动多个goroutine
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1) // 启动一个goroutine就登记+1
	// 	go hello(i)
	// }
	// wg.Wait() // 等待所有登记的goroutine都结束

	// goroutine与线程

	// channel
	// ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	// ch <- 10
	// fmt.Println("发送成功")

	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// // 开启goroutine将0~100的数发送到ch1中
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		ch1 <- i
	// 	}
	// 	close(ch1)
	// }()
	// // 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	// go func() {
	// 	for {
	// 		i, ok := <-ch1 // 通道关闭后再取值ok=false
	// 		if !ok {
	// 			break
	// 		}
	// 		ch2 <- i * i
	// 	}
	// 	close(ch2)
	// }()
	// // 在主goroutine中从ch2中接收值打印
	// for i := range ch2 { // 通道关闭后会退出for range循环
	// 	fmt.Println(i)
	// }

	// worker pool（goroutine池）
	// 在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨。
	// 一个简易的work pool示例代码如下：

	// jobs := make(chan int, 100)
	// results := make(chan int, 100)
	// // 开启3个goroutine
	// for w := 1; w <= 3; w++ {
	// 	go worker(w, jobs, results)
	// }
	// // 5个任务
	// for j := 1; j <= 5; j++ {
	// 	jobs <- j
	// }
	// close(jobs)
	// // 输出结果
	// for a := 1; a <= 5; a++ {
	// 	<-results
	// }

	// select多路复用

	// 在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。你也许会写出如下代码使用遍历的方式来实现：

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

// func hello() {
// 	fmt.Println("Hello Goroutine!")
// }

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

type book struct{ title string }

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	// fmt.Printf("type:%v\n", v)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

var xx interface{}

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

type animal interface {
	Mover
	Sayer
}

type WashingMachine interface {
	wash()
	dry()
}

type Dryer struct{}

func (d Dryer) dry() {
	fmt.Println("脱水烘干")
}

type haier struct {
	Dryer
}

func (h haier) wash() {
	fmt.Println("洗衣服")
}

type catt struct {
	branch string
}

type dogg struct {
	name string
}

func (d dogg) say() {
	fmt.Printf("%s叫了：汪汪汪\n", d.name)
}

func (d dogg) move() {
	fmt.Printf("%s走了\n", d.name)
}

func (c catt) move() {
	fmt.Printf("%s速度70迈\n", c.branch)
}

func (c catt) say() {
	fmt.Printf("%s叫了：猫猫猫\n", c.branch)
}

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

type Peoplet interface {
	Speak(string) string
}

type Studentt struct{}

func (stu *Studentt) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

// 接口的定义
type sayer interface {
	say()
}

// 实现接口的条件
type dog struct{}

type cat struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}
func (c cat) say() {
	fmt.Println("喵喵喵")
}

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) setDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

// 学生
type Student struct {
	ID     int `json:"id"`
	Gender string
	Name   string
}

// 班级
type Class struct {
	Title    string
	Students []*Student
}

// Animal 动物
type Animal struct {
	name string
}

// 动物的方法，走
func (animal *Animal) move() {
	fmt.Printf("%s会走\n", animal.name)
}

// Dog 狗
type Dog struct {
	feet int
	*Animal
}

// 狗的方法，叫
func (dog *Dog) wang() {
	fmt.Printf("%s会汪汪汪\n", dog.name)
}

//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name string
	Age  int
	Address
	Email
}

// 结构体的匿名字段
type person1 struct {
	string
	int
}

type MyInt int

func (m MyInt) sayHello() {
	fmt.Println("Hellow,我是一个INt")
}

type person struct {
	name string
	city string
	age  int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//Dream person做梦的方法
func (p person) dream() {
	fmt.Printf("%s的梦想就是学好go语言\n", p.name)
}

// 设置年龄
func (p *person) setAge(age int8) {
	p.age = age
}

// 计算每个人所得的金币
func dispatchCoin() int {
	for _, name := range users {
		coin := countCoin(name)
		distribution[name] = coin
		fmt.Printf("%s 的金币为：%d \n", name, coin)
	}
	var sum = 0
	for _, v := range distribution {
		sum += v
	}
	return coins - sum
}

// 计算名字中分配的金币数
func countCoin(name string) (result int) {

	result = 0
	name = strings.ToUpper(name)
	bs := []byte(name)
	for _, value := range bs {
		if value == 'E' {
			result++
		} else if value == 'I' {
			result += 2
		} else if value == 'O' {
			result += 3
		} else if value == 'U' {
			result += 4
		}
	}
	return
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
