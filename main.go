package main

import (
	"github.com/antchfx/jsonquery"
	"github.com/vvidovic/test-go/utils"
	"strings"
)

func main() {
	println("Message:", utils.Hello("Biggy"))
	println("Message:", utils.Hello("smally"))
	println("")
	jsonmsg := "{\"msg\":\"" + utils.Hello("JSON message") + "\"}"
	doc, err := jsonquery.Parse(strings.NewReader(jsonmsg))
	if err != nil {
		println("err: ", err)
	}
	println("jsonmsg:", jsonmsg)
	msg := jsonquery.Find(doc, "/msg")[0].InnerText()
	println("msg:", msg)
}
