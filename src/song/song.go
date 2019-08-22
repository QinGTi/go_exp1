package song

import (
	"awesomeProject/src/tools"
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

type Song struct {
	style 	string
	name 	string
	time 	int
	size 	int
	singer 	string
	album	string
	lrc		[]string
	buf		[]string
}
//根据路径读取歌曲信息
func (s *Song) ReadFile(filePath string) {
	//获取文件后缀
	s.style 	= path.Ext(filePath)
	//读取文件内容
	s.buf		= tools.ReadFileLine(filePath)
	//拆分文件，设置具体属性值
	s.name 		= s.buf[0]
	s.time, _	= strconv.Atoi(s.buf[1])
	s.size, _	= strconv.Atoi(s.buf[2])
	s.singer	= s.buf[3]
	s.album		= s.buf[4]
	s.lrc		= s.buf[5:]
}
func (s *Song) WriteFile(filePath string) {
	s.buf[0] = s.name
	s.buf[1] = strconv.Itoa(s.time)
	s.buf[2] = strconv.Itoa(s.size)
	s.buf[3] = s.singer
	s.buf[4] = s.album
	f, _ := os.Create(filePath)
	defer f.Close()
	buf := bufio.NewWriter(f)
	for _, v := range s.buf {
		fmt.Fprintln(buf, v)
	}
	buf.Flush()
}
//播放歌词
func (s *Song) ShowLrc(suffix string) {
	var endAdd string
	for i := 0; i < 5; i++ {
		endAdd += suffix
	}
	for i, v := range s.lrc {
		tools.MySleep(1)
		fmt.Printf("%d.%s%s\n", i, v, endAdd)
	}
}
//一些无聊的函数
func (s *Song) SetName() {
	fmt.Print("输入新名称:")
	rName := tools.ReadString()
	s.name = rName
}

func (s *Song) SetTime() {
	fmt.Print("输入新时长:")
	rTime := tools.ReadInt()
	s.time = rTime
}

func (s *Song) SetSize() {
	fmt.Print("输入新大小:")
	rSize := tools.ReadInt()
	s.time = rSize
}

func (s *Song) SetSinger() {
	fmt.Print("输入新歌手:")
	rSinger := tools.ReadString()
	s.singer = rSinger
}

func (s *Song) SetAlbum() {
	fmt.Print("输入新专辑:")
	rAlbum := tools.ReadString()
	s.album = rAlbum
}

func (s *Song) GetStyle() string {
	return s.style
}

func (s *Song) GetName() string {
	return s.name
}