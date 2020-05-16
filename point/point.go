package point

import (
	_ "URLCycleTest/logout"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Point struct {
	Name 		string
	Url			string
	Link 		[]string
}

//提供一个方法,提供point
func Newpoint(name string,u string) *Point {

	scode,err := getSourceCode(u)
	if err != nil {
		log.Println("源代码未成功获取")
		panic(err)
	}

	sl ,_ := getLink(u,scode)



	return &Point{ //创建实例

		Name:	name,
		Url:	u,
		Link:	sl,
	}
}

//提供一个方法,传入一个[]string,point对象循环用[]string中的参数循环对比是否有重复
func (p Point) DiffLink (s []string) []string {
	var v []string
	var linkStr	string
	for _,str := range p.Link {
		linkStr = linkStr + str + "\n"
	}


	for _,str := range s {
		if !strings.Contains(linkStr, str) {
			v = append(v, str)
		}
	}
	return v
}




type Slipoint []Point
//获取新point切片
func Newspoint() Slipoint {
	var pointSlice Slipoint
	return pointSlice
}

//定义方法,取得页面源代码
func  getSourceCode (url string) (string,error) {
	//生成client 参数为默认
	client := &http.Client{}

	//生成要访问的url
	url = "https://"+url

	//输出当前时间
	//log.Println(url+"  Test start------")

	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
		state := "请求创建失败"
		log.Println(state)
		return "",err
	}


	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		log.Println("服务器未成功响应: "+url)
		return "",err
	}

	//返回的状态码
	status := response.StatusCode

	if status > 400 {
		state := url+" : 状态码大于400,请检查url是否可用"
		log.Println(state)
		return "",err
	}

	//将源代码存入变量
	body := response.Body
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(body)
	scode := buf.String()

	return scode,nil
}

func getLink(u string,s string) ([]string,error) {

	//创建切片,获取所有匹配到的链接
	var link []string

	//创建正则表达式
	reg := regexp.MustCompile("(http://|https://|art/|col/)(.*?)(\"|html|htm)")
	if strings.Contains(s, ".html") {
		//fmt.Println("检查到html连接")
		dataSlice := reg.FindAll([]byte(s), -1)
		for _, v := range dataSlice {
			if !strings.Contains(string(v), "http") {
				//fmt.Println(u + "/" + string(v))
				link = append(link, u + "/" + string(v))
			} else {
				//fmt.Println(string(v))
				link = append(link, string(v))
			}

		}
	}

	return link,nil
}

//将连接信息导出为json,方便持久保存
func (a Slipoint) JsonOut() {

	b,err := json.Marshal(a)
	if err != nil {
		log.Println("json数据转化失败")
	}
	//fmt.Println(string(b))

	//创建json文件,保存旧数据
	fileName := "out.json"
	dstFile,err := os.Create(fileName)
	if err!=nil{
		log.Println("创建json文件失败")
		return
	}
	defer dstFile.Close()
	_, _ = dstFile.Write(b)
}

//从json中导入旧连接信息,与新连接做对比
func (a Slipoint) JsonIn() (Slipoint, error) {
	f, err := os.Open("out.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	conf, err := ioutil.ReadAll(f)
	if err !=nil {
		return nil,err
	}
	err = json.Unmarshal(conf, &a)
	if err !=nil {
		return nil,err
	}
	return a,nil
}

//从conf.json中提取配置文件
func (s Slipoint)SliInit() Slipoint {
	var Conf Slipoint
	f, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	conf, err := ioutil.ReadAll(f)
	if err !=nil {
		panic(err)
	}
	err = json.Unmarshal(conf, &Conf)
	if err !=nil {
		panic(err)
	}
	return Conf
}

//判断文件是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}