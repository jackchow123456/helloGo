package main

import (
	"fmt"
	"math"
	"strings"
	"unsafe"
	"json"
)
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth","", 
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
	c := &Class{"101",make([]*Student,0,200)}
	for i := 0; i < 10; i++ {
		stu := &Student{
			fmt.Sprintf("stu_%02d",i), 
			i,
		}
		c.students = append(c.students, stu)
	}
	fmt.Sprintf("%#v\n",c)
	data, err := json.Marshal(c)
	if !err != nil {
		fmt.Println("json Marshal failed.")
		return
	}
	fmt.Sprintf("json:%s",data)
}

// 学生
type Student struct{
	name string
	id int
}

// 班级
type Class struct{
	title string
	students []*Student
}

// Animal 动物
type Animal struct {
	name string
}

// 动物的方法，走
func (animal *Animal) move(){
	fmt.Printf("%s会走\n",animal.name)
}

// Dog 狗
type Dog struct{
	feet int
	*Animal
}

// 狗的方法，叫
func (dog *Dog) wang(){
	fmt.Printf("%s会汪汪汪\n",dog.name)
}


//Address 地址结构体
type Address struct {
	Province string
	City string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name string
	Age int
	Address
	Email
}


// 结构体的匿名字段
type person1 struct {
	string
	int
}

type MyInt int
func (m MyInt)sayHello(){
	fmt.Println("Hellow,我是一个INt")
}

type person struct {
	name string
	city string
	age  int8
}
func newPerson(name, city string, age int8) *person {
	return &person{
		name : name,
		city : city,
		age : age,
	}
}

//Dream person做梦的方法
func (p person)dream(){
	fmt.Printf("%s的梦想就是学好go语言\n",p.name)
}

// 设置年龄
func (p *person)setAge(age int8){
	p.age = age
}


// 计算每个人所得的金币
func dispatchCoin() (int) {
	for _, name:= range users {
		coin := countCoin(name)
		distribution[name] = coin;
		fmt.Printf("%s 的金币为：%d \n",name, coin)
	}
	var sum = 0;
	for _, v := range distribution{
		sum+= v;
	}
	return coins - sum;
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
