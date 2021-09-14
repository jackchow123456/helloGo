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

	findTheSame()
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

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

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
