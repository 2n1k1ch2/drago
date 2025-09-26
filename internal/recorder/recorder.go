package recorder

import (
	"drago/internal/event"
)

type Recorder struct {
}

func NewRecorder() *Recorder {
	return &Recorder{}
}

func (recorder *Recorder) Record(event event.Event) {

}
