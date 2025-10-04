package buffer

import (
	"drago/internal/event"
	"sync"
)

var (
	BATCH_SIZE = 256
	CH_SIZE    = 64
)

type ChanBuffer struct {
	mtx   sync.Mutex
	batch []event.Event
	ch    chan []event.Event
}

func NewChanBuffer() *ChanBuffer {
	return &ChanBuffer{
		ch:    make(chan []event.Event, CH_SIZE),
		batch: make([]event.Event, BATCH_SIZE),
	}
}
func (c *ChanBuffer) Insert(e event.Event) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.batch = append(c.batch, e)
	if len(c.batch) == BATCH_SIZE {
		err := c.Flush(BATCH_SIZE)
		if err != nil {
			return err
		}
	}

	return nil
}
func (c *ChanBuffer) Flush(batchSize int) error {
	c.ch <- c.batch
	c.batch = make([]event.Event, batchSize)
	return nil
}
func (c *ChanBuffer) Close() error {
	close(c.ch)
	return nil
}
func (c *ChanBuffer) Len() uint32 {
	return uint32(len(c.batch))
}
