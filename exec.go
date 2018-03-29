package main

import (
	"Study/src/timeStudy"
	"Study/src/jsonStudy"
	"time"
	"fmt"
	"Study/src/xmlStudy"
	"Study/src/stringStudy"
	"Study/src/regexpStudy"
	"Study/src/channelStudy"
	"Study/src/httpStudy"
	"Study/src/securityStudy"
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
	stringStudy.StrContains("I am your father", "m y")
	stringStudy.StrContains("", "")
	stringStudy.StrContainsAny("I am your father", "a1")
	stringStudy.StrContainsRune("I am your father", 'I')
	stringStudy.StrHasPrefix("I am your father", "I")
	stringStudy.StrTrimSpace("   I am your father   ")
	stringStudy.StrTrim("I am your father", "father")
	stringStudy.StrCount("I am your fahter", "a")
	stringStudy.StrCompare("I am your father", "I a")
	stringStudy.StrEqualFold("I am your father", "I am your father")
	stringStudy.StrFields("I am your father")
	stringStudy.StrFieldsFunc("I am your father")
	stringStudy.StrHasSuffix("I am your father", "father")
	stringStudy.StrIndex("I am your father", "r1")
	stringStudy.StrIndexAny("I am your father", "er")
	stringStudy.StrIndexByte("I am your father", 'a')
	stringStudy.StrIndexFunc("I am your father")
	stringStudy.StrIndexRune("I am your father", 'a')
	a := []string{"I", "am", "your", "father"}
	stringStudy.StrJoin(a, " ")
	stringStudy.StrLastIndex("I am your father", "r")
	stringStudy.StrMap("I am your father")
	stringStudy.StrRepeat("I am your father", 3)
	stringStudy.StrReplace("I am your fatheraaaaaaaaaaaaaaaaaaaaaaaa", "a", "b", 0)
	stringStudy.StrTitle("I am your father")
	stringStudy.StrToLower("I am your father")
	stringStudy.StrToUpper("I am your father")
	stringStudy.StrSplit("I am your father", " ")
	stringStudy.StrSplitAfter("I am your father", "a")
	regexpStudy.RegeMatchString("a", "I am your father")
	channelStudy.ChannelTest()
	//channelStudy.SelectTest()
	//channelStudy.SelectTest2()
	//channelStudy.TimeOut()
	//A()
	//socketStudy.TcpStudy()
	//httpStudy.HttpGet()
	httpStudy.HttpHead()
	securityStudy.Md5()
	securityStudy.Sha1()
	securityStudy.FileMd5()
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

func A() {
	aaa := make(chan *channelStudy.PipeData)
	pip := new(channelStudy.PipeData)
	pip.Value = 1
	pip.Next = make(chan int)
	//pip.Next <- 2
	pip.Handler = func(a int) int {
		fmt.Println("传递过来的值是：", a)
		return 0
	}
	aaa <- pip
	channelStudy.Handler(aaa)

}
