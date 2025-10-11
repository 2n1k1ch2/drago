package proto

import (
	"drago/internal/event"
	EventPB "drago/internal/proto/proto" // корректное имя пакета protobuf
)

func ConvertToProto(e event.Event) *EventPB.Event {
	pb := &EventPB.Event{}

	switch ev := e.(type) {
	case event.GoEvent:
		pb.EventType = &EventPB.Event_Go{
			Go: &EventPB.GoEvent{
				RoutineId: ev.ID(),
				Seq:       ev.Seq(),
			},
		}
	case event.ChanEvent:
		pb.EventType = &EventPB.Event_Chan{
			Chan: &EventPB.ChanEvent{
				ChanId:    ev.ID(),
				Seq:       ev.Seq(),
				Direction: ev.Action,
				Payload:   ev.Value,
			},
		}
	case event.MutexEvent:
		pb.EventType = &EventPB.Event_Mutex{
			Mutex: &EventPB.MutexEvent{
				MutexId: ev.ID(),
				Op:      ev.Action,
			},
		}
	case event.TimerEvent:
		pb.EventType = &EventPB.Event_Timer{
			Timer: &EventPB.TimerEvent{
				TimerId:      ev.ID(),
				Seq:          ev.Seq(),
				Action:       ev.Action,
				Duration:     ev.Duration,
				ReturnedTime: ev.ReturnedTime,
			},
		}
	default:

	}

	return pb
}
