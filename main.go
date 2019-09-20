package main

import (
	// "gitee.com/martind/golcud/rabbit"
	"github.com/Martinsuper/golcud/rabbit"
)
func main() {
	rabbit.SendMessage("hello world! today")
}