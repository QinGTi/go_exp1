package player

import (
	"awesomeProject/src/song"
	"awesomeProject/src/tools"
	"fmt"
)

type Player struct {
	songList 	map[string]string //歌曲列表
	removeList  map[string]string //回收站列表
	suffixList  map[string]string //后缀名映射
}
//播放器初始化
func (p *Player) PlayerInit() {
    tools.ClrScreen()
	fmt.Println("init")
	p.songList 		= tools.ReadFileToMap("../save/songList.txt")
	p.removeList	= tools.ReadFileToMap("../save/removeList.txt")
	p.suffixList	= make(map[string]string)
	p.suffixList[".st1"] = "!"
	p.suffixList[".st2"] = "@"
	p.suffixList[".st3"] = "#"
	p.suffixList[".st4"] = "$"
	p.suffixList[".st5"] = "%"
    tools.MyPause()
}
//播放器菜单
func (p Player) PlayerMenu() {
	var choose int
	for {
		tools.ClrScreen()
		fmt.Println("1.管理歌曲")
		fmt.Println("2.播放歌曲")
		fmt.Println("3.回收站")
		fmt.Println("0.退出")
		fmt.Print("输入选择序号：")
		choose = tools.ReadInt()
		switch choose {
		case 0:
			p.PlayerQuit()
		case 1:
			p.PlayerCtrl()
		case 2:
			p.PlayerRun()
		case 3:
			p.PlayerRecycle()
		default:
			fmt.Println("输入不合法")
		}
        tools.MyPause()
		if choose == 0 {
			break;
		}
	}
}
//退出并保存数据
func (p *Player) PlayerQuit() {
	tools.ClrScreen()
	tools.WriteMapToFile(p.songList, "../save/songList.txt")
	tools.WriteMapToFile(p.removeList, "../save/removeList.txt")
	fmt.Println("信息保存完毕")
}
//管理歌曲
func (p *Player) PlayerCtrl() {
	for {
		tools.ClrScreen()
		fmt.Println("1.添加歌曲")
		fmt.Println("2.删除歌曲")
		fmt.Println("3.修改歌曲")
		fmt.Println("4.查找歌曲")
		fmt.Println("5.歌曲列表")
		fmt.Println("0.返回")
		fmt.Print("输入选择序号：")
		choose := tools.ReadInt()
		switch choose {
		case 0:
			return
		case 1:
			p.AddSong()
		case 2:
			p.DelSong()
		case 3:
			p.RevSong()
		case 4:
			p.FndSong()
		case 5:
			p.ShowSongList()
		default:
			fmt.Println("输入不合法")
		}
        tools.MyPause()
	}
}
//播放歌曲
func (p *Player) PlayerRun() {
	if p.ShowSongList() == false {
        return
    }
	fmt.Print("输入待播放歌曲名:")
	sName := tools.ReadString()
	_, ok := p.songList[sName]
	if ok == false {
		fmt.Println("歌名不存在")
		return
	}
	var tmpSong song.Song
	tmpSong.ReadFile(p.songList[sName])
	fmt.Println(tmpSong.GetStyle())
	tmpSong.ShowLrc(p.suffixList[tmpSong.GetStyle()])
}
//回收站
func (p *Player) PlayerRecycle() {
	if p.ShowRemoveList() == false {
        return
    }
	fmt.Print("是否要恢复歌曲[yes/other]:")
	if tools.ReadString() == "yes" {
		fmt.Print("输入待恢复歌曲名:")
		sName := tools.ReadString()
		_, ok := p.removeList[sName]
		if ok == false {
			fmt.Println("歌名不存在")
			return
		}
		p.songList[sName] = p.removeList[sName]
		delete(p.removeList, sName)
		fmt.Println("恢复成功")
	}
}
//管理歌曲的函数
func (p *Player) AddSong() {
	tools.ClrScreen()
	fmt.Print("输入歌曲路径:")
	songPath := tools.ReadString()
	if tools.PathExist(songPath) == false {
		fmt.Println("路径不存在")
		return
	}
	var tmpSong song.Song
	tmpSong.ReadFile(songPath)
	p.songList[tmpSong.GetName()] = songPath
	fmt.Println("添加成功")
}

func (p *Player) DelSong() {
	if p.ShowSongList() == false {
        return
    }
	fmt.Print("输入待删除歌曲名：")
	sName := tools.ReadString()
	_, ok := p.songList[sName]
	if ok == false {
		fmt.Println("歌名不存在")
		return
	}
	p.removeList[sName] = p.songList[sName]
	delete(p.songList, sName)
	fmt.Println("删除成功")
}

func (p *Player) RevSong() {
	if p.ShowSongList() == false {
        return
    }
	fmt.Print("输入待修改歌名:")
	sName := tools.ReadString()
	_, ok := p.songList[sName]
	if ok == false {
		fmt.Println("歌名不存在")
		return
	}
	var tmpSong song.Song
	filePath := p.songList[sName]
	tmpSong.ReadFile(filePath)
	for {
		tools.ClrScreen()
		fmt.Println("可修改属性如下")
		fmt.Println("1.名称")
		fmt.Println("2.时长")
		fmt.Println("3.大小")
		fmt.Println("4.歌手")
		fmt.Println("5.专辑")
		fmt.Println("0.退出")
		fmt.Print("输入选择序号：")
		choose := tools.ReadInt()
		switch choose {
		case 0:
		    return
		case 1:
			tmpSong.SetName()
			if _, ok := p.songList[tmpSong.GetName()]; !ok {
				p.songList[tmpSong.GetName()] = filePath
				delete(p.songList, sName)
			}
		case 2:
			tmpSong.SetTime()
		case 3:
			tmpSong.SetSize()
		case 4:
			tmpSong.SetSinger()
		case 5:
			tmpSong.SetAlbum()
		default:
			fmt.Println("输入有误")
			continue
		}
		tmpSong.WriteFile(filePath)
		fmt.Print("修改完毕\n是否继续修改[yes/other]")
		if tools.ReadString() != "yes" {
			break;
		}
	}
}
//通过歌名查找歌曲
func (p *Player) FndSong() {
	tools.ClrScreen()
	fmt.Print("输入待查找歌名:")
	sName := tools.ReadString()
	_, ok := p.songList[sName]
	if ok {
		fmt.Println("查找成功")
	} else {
		fmt.Println("查找失败")
	}
}
//一些小函数
func (p *Player) ShowSongList() bool {
	tools.ClrScreen()
    if len(p.songList) == 0 {
        fmt.Println("歌曲列表为空")
        return false
    }
	fmt.Println("歌曲列表如下")
	i := 0
	for x := range p.songList {
		i++
		fmt.Printf("%d.%s\n", i, x)
	}
    return true
}

func (p *Player) ShowRemoveList() bool {
	tools.ClrScreen()
    if len(p.removeList) == 0 {
        fmt.Println("回收站列表为空")
        return false
    }
	fmt.Println("回收站列表如下")
	i := 0
	for x := range p.removeList {
		i++
		fmt.Printf("%d.%s\n", i, x)
	}
    return true
}
