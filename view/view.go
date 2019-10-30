package view

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/vvidovic/test-go/view/in"
	"github.com/vvidovic/test-go/view/out"
)

func Init(eventChan chan string) {
	fmt.Println("view/Init()")

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	// setScreenSize(m)
	// draw(m)
	//
	// keyPressedLoop(m)
	out.Init(eventChan)
	in.Start(eventChan)
}
