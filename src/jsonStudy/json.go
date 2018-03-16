package jsonStudy

import (
	"encoding/json"
	"fmt"
)

//转化为json串
func MarshalJson(v Student) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := string(b)
	fmt.Println("转化后的JSON串：", result)
}

//解析JSON串
func UnMarshalJSON(str string, s *Student) {
	json.Unmarshal([]byte(str), s)
}


