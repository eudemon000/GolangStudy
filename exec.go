package main

import (
	"Study/src/timeStudy"
	"Study/src/jsonStudy"
	"time"
	"fmt"
	"Study/src/xmlStudy"
)

func main() {
	//fileStudy.CreateFile("aaa")
	//fileStudy.RemoveFile("abc.desktop")
	//fileStudy.RemoveAll("aaa")
	//fileStudy.WriteFile("aaa", "aaa")
	/*fileStudy.ReadFile("aaa")
	fileStudy.ReadFileIOUtil("aaa")
	fileStudy.ReadFileSeek("aaa", 3, io.SeekStart)
	fileStudy.ReadFileBuffer("aaa")
	fileStudy.WriteFileIOUtil("aaa", "你们不要搞事情")
	fileStudy.WriteFileA("aaa", "大谷歌")
	fileStudy.WriteFileB("aaa", "大谷歌")
	fileStudy.WriteFileC("aaa", "大Android")
	fileStudy.ReadInfo("test")
	fileStudy.AppendFile("aaa", "456")*/
	timeStudy.GetNowUnix()
	timeStudy.GetNow()
	timeStudy.GetNowFormat()
	timeStudy.TimeParse()
	timeStudy.TimeInfo()
	jsonTest()
	unMarshalTest()
	XMLMarshal()
}

func jsonTest() {
	s := jsonStudy.Student{}
	s.Name = "张三"
	s.Class = "三年级一班"
	s.Id = 1
	s.Brithday = jsonStudy.JsonTime(time.Now())
	jsonStudy.MarshalJson(s)
}

func unMarshalTest() {
	s := new(jsonStudy.Student)
	fmt.Println(s)
	str := `{"id":1,"name":"张三","class":"三年级一班","brithday":"2018-03-16 15:14:26"}`
	jsonStudy.UnMarshalJSON(str, s)
	fmt.Println(s)
}

func XMLMarshal() {
	t := new(xmlStudy.Teacher)
	t.Name = "李四"
	t.Age = 28
	t.School = "李四中学"
	t.Brithday = xmlStudy.XmlTime(time.Now())
	xmlStudy.XMLMarshal(t)
}
