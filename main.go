package main

import (
	_ "URLCycleTest/logout"
	"URLCycleTest/point"
	"fmt"
	"log"
	"os"
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


	fmt.Println("\n\n\n================================检查开始===================================")
	fmt.Println("\n---------------------以下为各页面减少连接---------------")
	newSp.SliContrast(oldSp)
	fmt.Println()
	fmt.Println("\n+++++++++++++++++++++以下为各页面新增连接+++++++++++++++")
	oldSp.SliContrast(newSp)
	fmt.Println("\n================================检查结束===================================")

}
