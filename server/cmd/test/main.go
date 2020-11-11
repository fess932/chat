package main

import (
	"log"

	"github.com/looplab/fsm"
)

type event string

type state string

func (s state) String() string {
	return string(s)
}

//const (
//	open  event = "open"
//	close event = "close"
//)
//
//const (
//	opened state = "opened"
//	closed state = "closed"
//)

func onOpen(e *fsm.Event) {
	log.Println("Открыто!")
}

func main() {
	open := fsm.EventDesc{
		Name: "open",
		Src:  []string{"closed"},
		Dst:  "open",
	}

	brain := fsm.NewFSM("closed",
		fsm.Events{
			open,
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_open": onOpen,
		},
	)

	log.Println(brain.Current())

	if err := brain.Event("open"); err != nil {
		log.Println(err)
	}

	log.Println(brain.Current())

	log.Println(brain.AvailableTransitions())

}
