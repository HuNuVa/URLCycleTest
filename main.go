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

	for _,v := range conf {

		a := point.Newpoint(v.Name,v.Url)
		newSp = append(newSp, *a)
	}

	defer newSp.JsonOut()
	fmt.Println("\n检查开始===================================")
	fmt.Println("以下为各页面减少连接---------------")
	newSp.SliContrast(oldSp)

	fmt.Println("以下为各页面新增连接---------------")
	oldSp.SliContrast(newSp)

}
