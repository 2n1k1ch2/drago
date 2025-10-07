package buffer

import (
	"drago/internal/event"
)

type EventBuffer interface {
	Insert(e event.Event) error // add event in buffer
	Flush(batchSize int) error  // flush buffer
	Len() uint32
	Close() error // end and free resource
	Out() <-chan []event.Event
}
