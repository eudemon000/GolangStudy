package jsonStudy

import (
	"time"
	"fmt"
)

const timeFormat  = "2006-01-02 15:04:05"

type JsonTime time.Time

type Student struct {
	Id			int			`json:"id"`
	Name		string		`json:"name"`
	Class		string		`json:"class"`
	Brithday	JsonTime	`json:"brithday"`
}

//JSON时间格式化，给time.Time设置一个包装类型，然后实现MarshalJSON()方法
func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

//将JSON字符串转为结构体，此处对时间格式进行出来，要实现UnmarshalJSON函数
func (j *JsonTime) UnmarshalJSON(b []byte) error {
	t, err := time.ParseInLocation(`"` + timeFormat + `"`, string(b), time.Local)
	if err != nil {
		fmt.Println(err)
	}
	*j = JsonTime(t)
	return nil
}

//在处理时间格式的JSON字段的时候，字段会显示为时间戳形式，这里通过格式化时间戳，转化为指定的字符串形式显示
func (j JsonTime) String() string {
	return time.Time(j).Format(timeFormat)
}
