package main

import (
	"golcud/rabbit"
	// "golcud/crawler/getSoft"
	"fmt"
	"encoding/json"
	"flag"
)
type message struct {
	Title  string `json:"title"`
	Content   string  `json:"content"`
 }

func main() {
	var title string 
	var content string
	flag.StringVar(&title, "t","title","邮件title")
	flag.StringVar(&content, "c","content","邮件正文")
	flag.Parse()

	//转换成JSON字符串
	myMessage := message{
		Title: title,
		Content: content,
	}
	jsons, errs := json.Marshal(myMessage) //转换成JSON返回的是byte[]
	if errs != nil {
	   fmt.Println(errs.Error())
	}
	fmt.Println(string(jsons)) //byte[]转换成string 输出
	rabbit.SendMessage(string(jsons),"email.task.queue")
}