package runtime

import (
	"github.com/google/uuid"
	"sync/atomic"
)

var seq_generator uint64 = 0

func NEXT() uint64 {
	return atomic.AddUint64(&seq_generator, 1)
}

func GET_ID() string {
	id := uuid.New().String()
	return id
}
