package httpStudy

import (
	"net/http"
	"fmt"
	"io"
	"os"
)

func HttpGet() {
	resp, err := http.Get("http://www.qq.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func HttpHead() {
	resp, err := http.Head("http://gaoqing.la/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}