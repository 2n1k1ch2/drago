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

func (g GoEvent) Type() string {
	return "GoEvent"
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
func (g ChanEvent) Type() string {
	return "ChanEvent"
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
func (g TimerEvent) Type() string {
	return "TimerEvent"
}

type MutexEvent struct {
	id     string
	seq    uint64
	Action string
}

func (m MutexEvent) ID() string {
	return m.id
}
func (g MutexEvent) Seq() uint64 {
	return g.seq
}

func (g MutexEvent) Type() string {
	return "MutexEvent"
}
