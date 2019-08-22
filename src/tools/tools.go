package tools

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var inputReader = bufio.NewReader(os.Stdin)
//方便读入数据
func ReadInt() int {
	input, _ := inputReader.ReadString('\n')
	FilterStr(&input)
	readInt, _ := strconv.Atoi(input)
	return readInt
}

func ReadString() string {
	input, _ := inputReader.ReadString('\n')
	FilterStr(&input)
	readString := input
	return readString
}
//按行入读
func ReadFileLine(filePath string) ([]string) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result []string
	for {
		line, err := buf.ReadString('\n')
		FilterStr(&line)
		//fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				return result
			}
			return nil
		}
		result = append(result, line)
	}
	return result
}
//自定义分割
func Split(s rune) bool {
	if s == '-' {
		return true
	} else {
		return false
	}
}
//把文件信息读入map
func ReadFileToMap(filePath string) map[string]string {
	fmt.Println("Loading    "+filePath)
	m := make(map[string]string)
	var buf []string = ReadFileLine(filePath)
	for _, v := range buf {
		FilterStr(&v)
		var tmp []string = strings.FieldsFunc(v, Split)
		m[tmp[0]] = tmp[1]
	}
	return m
}
//把map信息写入文件
func WriteMapToFile(m map[string]string, filePath string) {
	f, _ := os.Create(filePath)
	defer f.Close()
	buf := bufio.NewWriter(f)
	for k, v := range m {
		line := fmt.Sprintf("%s-%s", k, v)
		fmt.Fprintln(buf, line)
	}
	buf.Flush()
}
//判断输入路径是否有效
func PathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
//延时函数，x是秒
func MySleep(second int) {
	time.Sleep(time.Duration(second)*time.Second)
}
//实现按任意键继续
func MyPause() {
	fmt.Print("按任意键继续...")
	var tmpInput string
	fmt.Scanf("%s", tmpInput)
}
//实现清屏
func ClrScreen () {
	system := runtime.GOOS
	if  system == "linux"{
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if system == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
//读入一行后，过滤掉末尾的\r\n
func FilterStr(str *string) {
	(*str) = strings.Replace((*str),"\r", "", -1)
	(*str) = strings.Replace((*str),"\n", "", -1)
}
