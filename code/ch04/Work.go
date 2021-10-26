package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	rand.Seed(time.Now().UnixNano()) // 隨機種子
	year := rand.Intn(2020) + 1 // 隨機年份
	leep := year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) // 判斷閏平年
	month := rand.Intn(12) + 1 // 隨機月份
	daysInMonth := 31 // 除了 2 4 6 9 11 以外都是 31 天 

	switch month {
	case 2:
		if leep {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1 // 隨機日期
	fmt.Println(era, year, month, day)
}