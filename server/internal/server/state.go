package server

import (
	"github.com/looplab/fsm"
)

func createFSM() *fsm.FSM {

	brain := fsm.NewFSM("stop",
		fsm.Events{
			{Name: "stop", Src: []string{"right"}, Dst: "stop"},
			{Name: "right", Src: []string{"stop"}, Dst: "right"},
		},
		fsm.Callbacks{},
	)

	return brain
}
