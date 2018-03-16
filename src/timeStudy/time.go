package timeStudy

import (
	"fmt"
	"time"
)

const format string = "2006-01-02 15:04:05"
//const format string = "0000-00-00 00:00:00"

//获取当前时间戳
func GetNowUnix() {
	fmt.Println("时间戳：", time.Now().Unix())
}

//获取当前时间
func GetNow() {
	fmt.Println("当前时间：", time.Now())
}

//格式化当前时间
func GetNowFormat() {
	fmt.Println(time.Now().Format(format))
}

//将时间格式转为为时间戳
func TimeParse() {
	t, err := time.Parse("2006-01-02 15:04:05", "2018-03-16 10:34:50")
	if err != nil {
		fmt.Println(err)
		return
	}
	var time int64 = t.Unix()
	fmt.Println("格式化后的时间戳：", time)
}

//时间
func TimeInfo() {
	time := time.Now()
	fmt.Println("年(year)：", time.Year())
	fmt.Println("年(yearDay)：", time.YearDay())
	fmt.Println("月：", time.Month().String())
	fmt.Println("日：", time.Day())
	fmt.Println("时：", time.Hour())
	fmt.Println("分：", time.Minute())
	fmt.Println("秒：", time.Second())
	fmt.Println("String：", time.String())
	fmt.Println("weekDay：", time.Weekday())

}


