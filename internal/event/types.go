package event

const (
	GO_START     string = "go_start"
	GO_STOP      string = "go_stop"
	CHANNEL_SEND string = "channel_send"
	CHANNEL_RECV string = "channel_recv"
	TIME_SLEEP   string = "time_sleep"
	TIME_NOW     string = "time_now"
	TIME_AFTER   string = "time_after"
)

type Event interface {
	ID() string
	Seq() uint64
	Type() string
}
type isEvent_EventType interface {
	isEvent_EventType()
}
type Event_Go struct{ Go *GoEvent }
type Event_Chan struct{ Chan *ChanEvent }
type Event_Timer struct{ Timer *TimerEvent }
type Event_Mutex struct{ Mutex *MutexEvent }
