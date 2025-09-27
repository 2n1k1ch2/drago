package event

type GoEvent struct {
	id  string
	seq uint64
}

func (g GoEvent) ID() string {
	return g.id
}
func (g GoEvent) Seq() uint64 {
	return g.seq
}

type ChanEvent struct {
	id     string
	seq    uint64
	Action string
	Value  []byte
}

func (c ChanEvent) ID() string {
	return c.id
}
func (g ChanEvent) Seq() uint64 {
	return g.seq
}

type TimerEvent struct {
	id           string
	seq          uint64
	Action       string
	Duration     int64
	ReturnedTime int64
}

func (g TimerEvent) ID() string {
	return g.id
}
func (g TimerEvent) Seq() uint64 {
	return g.seq
}
