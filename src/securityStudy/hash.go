package securityStudy

import (
	"crypto/md5"
	"fmt"
	"crypto/sha1"
	"encoding/hex"
	"os"
	"io/ioutil"
	"bytes"
	"io"
)

func Md5() {
	testString := "Hello world"
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(testString))
	result := Md5Inst.Sum([]byte(""))
	//fmt.Printf("%x\n\n", result)
	str := hex.EncodeToString(result)
	fmt.Println(str)
}

func Sha1() {
	testString := "Hello world"
	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(testString))
	result := sha1Inst.Sum([]byte(""))
	//fmt.Printf("%x\n\n", result)
	str := hex.EncodeToString(result)
	fmt.Println(str)
}

func FileMd5() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range fileInfos {
		if !f.IsDir() {
			b := bytes.Buffer{}
			b.WriteString(path)
			b.WriteString("/")
			b.WriteString(f.Name())
			tempPath := b.String()
			file, err := os.Open(tempPath)
			if err != nil {
				fmt.Println(err)
				return
			}
			md5h := md5.New()
			io.Copy(md5h, file)
			md5Result := md5h.Sum([]byte(""))

			sha1H := sha1.New()
			io.Copy(sha1H, file)
			sha1Result := sha1H.Sum([]byte(""))
			fmt.Println(tempPath)
			fmt.Println("Md5===>", hex.EncodeToString(md5Result))
			fmt.Println("sha1===>", hex.EncodeToString(sha1Result))
			fmt.Println("")
		}
	}


}