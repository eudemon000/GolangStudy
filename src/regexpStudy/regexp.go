package regexpStudy

import (
	"regexp"
	"fmt"
)

func RegeMatchString(pattern, s string) {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("MatchString：", matched)
}