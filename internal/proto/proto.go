package proto

import (
	"drago/internal/event"
	"google.golang.org/protobuf/proto"
)

func ConvertToProto(e event.Event) *proto.Event {
	pb := &proto.{
		Seq: e.Seq(),
		Id:  e.ID(),
	}

	switch ev := e.(type) {
	case *GoEvent:
		pb.EventType = &proto.Event_Go{
			Go: &proto.GoEvent{RoutineId: ev.RoutineID},
		}
	case *ChanEvent:
		pb.EventType = &proto.Event_Chan{
			Chan: &proto.ChanEvent{ChanId: ev.ChanID, Direction: ev.Direction},
		}
	}
	case *MutexEvent:
		pb.EventType = &proto.Event_Mutex{
		Mutex: &proto.MutexEvent{MutexId: ev.MutexID},
	}
	case *TimerEvent:
		pb.EventType = &proto.Event_Timer{
		Timer: &proto.TimerEvent{TimerId: ev.TimerID},
	}

	return pb
}
