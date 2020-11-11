package wsClient

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

//import "log"
//
//type State int
//
//const (
//	stop = iota
//	move
//)
//
//func (s State) Run() {
//	switch s {
//	case stop:
//		log.Println(stop)
//	default:
//		log.Println("no actions!")
//	}
//
//}
//
//type StackFSM struct {
//	stack []State
//}
//
//func (s *StackFSM) update() {
//	var currentStateFunction = s.getCurrentState()
//
//	if currentStateFunction != nil {
//		currentStateFunction.Run()
//	}
//}
//
//func (s *StackFSM) popState() State {
//	x, a := s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
//	s.stack = a
//	return x
//}
//
//func (s *StackFSM) pushState(state State) {
//	if s.getCurrentState() != state {
//		s.stack = append(s.stack, state)
//	}
//}
//
//func (s *StackFSM) getCurrentState() State {
//	if len(s.stack) > 0 {
//		return s.stack[len(s.stack)-1]
//	}
//	return nil
//}
