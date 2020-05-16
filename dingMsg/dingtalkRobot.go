package dingMsg
import (
	"fmt"
	"net/http"
	"strings"
)

func SendDingMsg(msg string) {
	//请求地址模板
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=0c09ee0834fedc9c0cd25cc7ccadac240bb202f13b28a69048664409079b3cf7"
	content := `{
					"msgtype": "text",
					"text": {
								"content": "`+msg+`"
							},
					"at": {
						"isAtAll": true
					}
				}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		fmt.Println("handle error")
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}
}
