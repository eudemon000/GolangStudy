package stringStudy

import (
	"strings"
	"fmt"
)

//判断一串字符串中是否存在指定的subString子串
func StrContains(s, subStr string) {
	r := strings.Contains(s, subStr)
	fmt.Println("Contains：", r)
}

//判断一串字符串中是否存在指定子串中的任意一字符，只要子串中的任意一字符存在于字符串中就返回true
func StrContainsAny(s string, chars string) {
	r := strings.ContainsAny(s, chars)
	fmt.Println("ContainsAny：", r)
}

//判断一串字符串中是否包含指定的字符
func StrContainsRune(s string, r rune) {
	result := strings.ContainsRune(s, r)
	fmt.Println("ContainsRune：", result)
}

//判断字符串是否以prefix开头
func StrHasPrefix(s, prefix string) {
	r := strings.HasPrefix(s, prefix)
	fmt.Println("HasPrefix：", r)
}

//清除字符串前后的空格
func StrTrimSpace(s string) {
	r := strings.TrimSpace(s)
	fmt.Println("trimSpace：", r)
}

//删除字符串前后指定的子字符串（注：只能删除前后的字符串，如果删除后的位置有空格，空格不会被删除）
func StrTrim(s, cutset string) {
	r := strings.Trim(s, cutset)
	fmt.Println("trim：", r)
}

//计算字符串中子串出现的次数
func StrCount(s, substr string) {
	r := strings.Count(s, substr)
	fmt.Println("Count：", r)
}

//返回两个字符串字典比较的整数结果，如果a > b，返回1，如果a < b，返回-1，如果a == b，返回0
func StrCompare(a, b string) {
	r := strings.Compare(a, b)
	fmt.Println("Compare：", r)
}

//判断两个字符串是否相等
func StrEqualFold(s, t string) {
	r := strings.EqualFold(s, t)
	fmt.Println("EqualFold：", r)
}

//讲字符串按照空格分割成一个字符串数组
func StrFields(s string) {
	r := strings.Fields(s)
	for _, a := range r {
		fmt.Println("Fields：", a)
	}
}

//将字符串按照指定的字符分割成字符串数组
func StrFieldsFunc(s string) {
	result := strings.FieldsFunc(s, func(r rune) bool {
		if r == rune('a') {
			return true
			} else {
				return false
				}
				})
	for _, a := range result {
		fmt.Println("FieldsFunc：", a)
	}
}

//判断字符串是否是以子串结尾的
func StrHasSuffix(s, suffix string) {
	r := strings.HasSuffix(s, suffix)
	fmt.Println("HasSuffix：", r)
}

//获取字符串中子串第一次出现的索引值，如果不存在就返回-1
func StrIndex(s, substr string) {
	 r := strings.Index(s, substr)
	 fmt.Println("Index：", r)
}

//获取字符串中子串中任意一个字符第一次出现的索引值（注：子串中的字符不需要全部出现，。只要任意一个字符出现即可）
func StrIndexAny(s, chars string) {
	r := strings.IndexAny(s, chars)
	fmt.Println("IndexAny：", r)
}

//查找字符串中第一个字符出现的索引值
func StrIndexByte(s string, c byte) {
	r := strings.IndexByte(s, c)
	fmt.Println("IndexByte：", r)
}

//按照制指定的规则获取
func StrIndexFunc(s string) {
	result := strings.IndexFunc(s, func(r rune) bool {
		if r == rune('f') {
			return true
		} else {
			return false
		}

	})
	fmt.Println("IndexFunc：", result)
}

//查找字符串中指定字符第一次出现的位置
func StrIndexRune(s string, r rune) {
	result := strings.IndexRune(s, r)
	fmt.Println("IndexRune：", result)
}

//用sep字符串将字符串数组a连接成一个字符串
func StrJoin(a []string, sep string) {
	r := strings.Join(a, sep)
	fmt.Println("Join：", r)
}

//从字符串末尾获取指定子串第一次出现的索引
func StrLastIndex(s, substr string) {
	r := strings.LastIndex(s, substr)
	fmt.Println("LastIndex：", r)
}

//返回一个Map的副本，返回的rune值替换对应位置的字符
func StrMap(s string) {
	result := strings.Map(func(r rune) rune {
		if r == rune('a') {
			return rune('s')
		}
		return r
	}, s)
	fmt.Println("Map：", result)
}

//产生一个新的字符串，该字符串包括字符串s的计数副本
func StrRepeat(s string, count int) {
	r := strings.Repeat(s, count)
	fmt.Println("Repeat：", r)
}

//用新的字符串代替字符串中指定的子串，n表示替代几个，当n=-1的时候表示全部替代
func StrReplace(s, old, new string, n int) {
	r := strings.Replace(s, old, new, n)
	fmt.Println("Replace：", r)
}

//将英文的每个单词首字母大写，转化位Title的格式
func StrTitle(s string) {
	r := strings.Title(s)
	fmt.Println("Title：", r)
}

//字符串转化位小写形式
func StrToLower(s string) {
	r := strings.ToLower(s)
	fmt.Println("ToLower：", r)
}

//字符串转化为大写
func StrToUpper(s string) {
	r := strings.ToUpper(s)
	fmt.Println("ToUpper：", r)
}

//按照指定的字符串将字符串分割位一个字符串数组
func StrSplit(s, sep string) {
	r := strings.Split(s, sep)
	for _, a := range r {
		fmt.Println("Split：", a)
	}
}

//按照指定的字符串分割字符串为一个字符串数组，指定的字符串后的位置
func StrSplitAfter(s, sep string) {
	r := strings.SplitAfter(s, sep)
	for _, a := range r {
		fmt.Println("SplitAfter：", a)
	}
}
