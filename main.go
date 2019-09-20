package main

import (
	"golcud/rabbit"
	"golcud/crawler/getSoft"
	"fmt"
)
func main() {
	link := getSoft.GetGoLink()
	fmt.Println(link)
	// rabbit.SendMessage("hello world! today")
}