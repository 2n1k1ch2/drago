package recorder

import (
	"drago/internal/buffer"
	"drago/internal/event"
	"drago/internal/writer"
	"log"
)

type Recorder struct {
	buf    buffer.EventBuffer
	writer *writer.Writer
	ch     chan<- []event.Event
}

func NewRecorder() *Recorder {
	return &Recorder{
		buf: buffer.NewChanBuffer(),
	}
}
func (r *Recorder) Record(event event.Event) {
	er := r.buf.Insert(event)
	if er != nil {
		log.Print(er)
	}
}

func (r *Recorder) Start() {
	for batch := range r.buf.Out() {
		if err := r.writer.Write(batch); err != nil {
			log.Printf("write error: %v", err)
		}
	}
}
