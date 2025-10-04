package recorder

import (
	"drago/internal/buffer"
	"drago/internal/event"
	"log"
)

type Recorder struct {
	buf buffer.EventBuffer
}

func NewRecorder() *Recorder {
	return &Recorder{
		buf: buffer.NewChanBuffer(),
	}
}
func (r Recorder) Record(event event.Event) {
	er := r.buf.Insert(event)
	if er != nil {
		log.Print(er)
	}
}
