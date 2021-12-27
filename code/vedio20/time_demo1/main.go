package main

import (
	"fmt"
	"time"
)

// 时间操作

func main() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	// year := now.Year()     //年
	// month := now.Month()   //月
	// day := now.Day()       //日
	// hour := now.Hour()     //小时
	// minute := now.Minute() //分钟
	// second := now.Second() //秒
	// fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// timestamp1 := now.Unix()     //时间戳
	// timestamp2 := now.UnixNano() //纳秒时间戳
	// fmt.Printf("current timestamp1:%v\n", timestamp1)
	// fmt.Printf("current timestamp2:%v\n", timestamp2)

	// // 将时间戳转换为具体格式
	// t := time.Unix(1639481696, 0)
	// fmt.Println(t)

	// // 时间间隔
	// n := 5
	// time.Sleep(time.Duration(n) * time.Second)
	// fmt.Println("over")

	// now + 1 hour
	// fmt.Println(now)
	// t2 := now.Add(time.Hour)
	// fmt.Println(t2)
	// // sub
	// fmt.Println(t2.Sub(now))

	// 定时器
	// for tmp := range time.Tick(time.Microsecond * 2) {
	// 	fmt.Println(tmp)
	// }

	// 时间格式化
	// ret1 := now.Format("2006-01-02")
	// fmt.Println(ret1)

	// ret2 := now.Format("2006-01-02 15:04:05")
	// fmt.Println(ret2)

	// 根据时区解析一个字符串格式的时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	timeStr := "2020/01/01 15:00:00"

	timeObj, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("parse timeStr failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)

	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println("parse timeStr failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj2)
}
