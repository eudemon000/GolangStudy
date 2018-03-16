package fileStudy

import (
	"os"
	"fmt"
	"syscall"
	"io/ioutil"
	"bufio"
	"io"
)

//创建文件(如果当前已经存在该文件，文件会被替换为新文件)，
func CreateFile(filePath string) {
	//创建一个新文件并以读写方式打开，权限位"0666"，如果文件存在则会清空
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("创建文件失败，失败原因：", err)
		return
	}
	fmt.Println("创建文件成功", file.Name())
	defer file.Close()
}

//删除文件
func RemoveFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("删除失败，失败原因：", err)
		return
	}
	fmt.Println("删除成功")
}

//删除目录
func RemoveAll(filePath string) {
	err := os.RemoveAll(filePath)
	if err != nil {
		fmt.Println("删除失败，失败原因：", err)
		return
	}
	fmt.Println("删除成功")

}

//写入文件
func WriteFile(filePath, content string) {
	//查看当前的工作目录路径 得到测试文件的绝对路径
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
	//path := dir + "/" + filePath
	//该方式为最简单的文件打开方式，为只读模式
	//file, err := os.Open(filePath)
	/*
	其中第一个参数为文件路径，第二个参数为打开模式，第三个参数为文件权限
	打开模式分为:
		O_RDONLY：只读模式(read-only)
		O_WRONLY：只写模式(write-only)
		O_RDWR：读写模式(read-write)
		O_APPEND：追加模式(append)
		O_CREATE：文件不存在就创建(create a new file if none exists.)
		O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
		O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
		O_TRUNC：打开并清空文件
	文件权限（*nix权限位）：只有在创建文件时才需要，不需要创建文件可以设置为0
	 */
	file, err := os.OpenFile(filePath, os.O_CREATE | os.O_RDWR, syscall.S_IRWXU | syscall.S_IRWXG)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败，失败原因：", err)
		return
	}
	b := []byte(content)
	n, err := file.Write(b)
	if err != nil {
		fmt.Println("写入失败，失败原因：", err)
		return
	}
	fmt.Println(n)

	/*n, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)*/
}

//读取文件，这里可以控制你读取的字节数4
func ReadFile(filePath string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败，失败原因：", err)
		return
	}
	buffer := make([]byte, 5)
	file.Read(buffer)
	fmt.Println(string(buffer))

	/*n1, err := file.Read(buffer)
	fmt.Println(n1, string(buffer))*/
}

//通过ioutil包读取文件
func ReadFileIOUtil(filePath string) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("打开文件错误，错误原因：", err)
		return
	}
	fmt.Println(string(b))
}

//从指定的偏移量开始读取文件
//whence参数使用io包中的SeekCurrent，SeekEnd，SeekStart常量
func ReadFileSeek(filePath string, offSet int64, whence int) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件错误，错误原因：", err)
		return
	}
	file.Seek(offSet, whence)
	buffer := make([]byte, 5)
	file.Read(buffer)
	fmt.Println("Seek 执行结果：", string(buffer))
}

func  ReadFileBuffer(filePath string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件错误，错误原因：", err)
		return
	}
	reader := bufio.NewReader(file)
	//buffer := make([]byte, 5)
	//reader.Read(buffer)
	//fmt.Println("ReadFileBuffer：", string(buffer))
	b, err := reader.Peek(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ReadFileBuffer：", string(b))

}

//写文件，粗粒度
func WriteFileIOUtil(filePath, content string) {
	b := []byte(content)
	err := ioutil.WriteFile(filePath, b, syscall.S_IRWXG | syscall.S_IRWXU)
	if err != nil {
		fmt.Println(err)
	}
}

func WriteFileA(filePath, content string) {
	file, err := os.Create("aaa")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	b := []byte(content)
	n2, err := file.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n2)
}

func WriteFileB(filePath, content string) {
	file, err := os.Create("aaa")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	n2, err := file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n2)
	file.Sync()
}

func WriteFileC(filePath, content string) {
	file, err := os.Create("aaa")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(file)
	n2, err := w.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n2)
	w.Flush()
}

//查看文件信息，filePath为目录
func ReadInfo(filePath string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfos, err := file.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range fileInfos {
		fmt.Println(item.Name())
		fmt.Println(item.IsDir())
		fmt.Println(item.Mode())
		fmt.Println(item.ModTime())
		fmt.Println(item.Size())
		fmt.Println(item.Sys())
	}
	/*for i := 0; i < len(fileInfos); i++ {
		name := fileInfos[i].Name()
		fmt.Println(name)
	}*/
}

//向文件末尾追加内容，利用Seek方法偏移的文件末尾，然后通过WriteAt来进行内容的写入
func AppendFile(filePath, content string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteAt([]byte(content), n)
	if err != nil {
		fmt.Println(err)
	}

}

