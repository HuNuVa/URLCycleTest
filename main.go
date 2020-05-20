package main

import (
	"URLCycleTest/dingMsg"
	_ "URLCycleTest/logout"
	"URLCycleTest/point"
	"fmt"
	_ "io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	f, _ := os.OpenFile("result.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND,0755)

	os.Stdout = f

	os.Stderr = f

	if !point.Exists("out.json") {
		_,err := os.Create("out.json")
		if err != nil {
			log.Print(err)
			return
		}
	}
	//输出时间
	fmt.Println("\n"+"执行时间:  ",time.Now().Format("2006-01-02 15:04:05"))

	//定义一个结构体切片,用来导入json文件中的记录
	oldSp := point.Newspoint()
	oldSp, _ = oldSp.JsonIn()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//读取配置文件
	conf := point.Newspoint()
	conf = conf.SliInit()
	newSp := point.Newspoint()
	//fmt.Println("conf元素个数:", + len(conf))
	//fmt.Println(conf)
	//定义一个管道,用来保存协程传递过来的 point,长度为conf.json中定义的对象个数
	c := make(chan *point.Point, len(conf))
	for _,v := range conf {

		go func() {
			c <- point.Newpoint(v.Name,v.Url)
		}()
		a := <- c
		newSp = append(newSp, *a)
	}

	defer newSp.JsonOut()


	//输出到result.txt
	fmt.Println("===============检查开始=================")
	fmt.Println("\n---------------以下为各页面减少连接--------------")
	snew := newSp.SliContrast(oldSp)
	fmt.Println("\n+++++++++++以下为各页面新增连接++++++++++")
	sold := oldSp.SliContrast(newSp)
	fmt.Println("\n===============检查结束=================")

	//将内容拼接位字符串,发送给钉钉
	strs := "执行时间:  " + time.Now().Format("2006-01-02 15:04:05")+
		"\n===============检查开始\n" +
		"-------以下为各页面减少连接\n"
	for _,v := range snew {
		strs = strs + "-" + v + "\n"
	}
	strs = strs + "+++++以下为各页面新增连接\n"
	for _,v := range sold {
		strs = strs + "+"  + v + "\n"
	}
	strs = strs + "===============检查结束"

	dingMsg.SendDingMsg(strs)



	/*
	//将要发送给钉钉的消息写入一个临时文件
	ftmp,err := os.Create("./.tmp")
	if err != nil {
		log.Println("创建临时文件失败(./.tmp)")
	} else {
		_, _ = ftmp.WriteString("===============检查开始\n" +
			"-------以下为各页面减少连接\n")
		for _,v := range snew {
			_, _ = ftmp.WriteString(v+"\n")
		}
		_, _ = ftmp.WriteString("+++++以下为各页面新增连接\n")
		for _,v := range sold {
			_, _ = ftmp.WriteString(v+"\n")
		}
		_, _ = ftmp.WriteString("===============检查结束")

		err = ftmp.Close()
		if err != nil {
			log.Println("关闭临时文件(./tmp出错:)",err)
		}

		//输出到钉钉
		bytes, err := ioutil.ReadFile("./.tmp")
		if err != nil {
			log.Println(err)
			return
		}
		dingMsg.SendDingMsg(string(bytes))

		err = os.Remove("./.tmp")
		if err !=nil {
			log.Println(err)
		}
	}

	 */


	f.Close()
}
