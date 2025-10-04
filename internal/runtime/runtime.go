package runtime

import (
	"drago/internal/event"
	Drago_Recorder "drago/internal/recorder"
)

type Mode int

const (
	ModeNone Mode = iota
	ModeRecord
	ModeReplay
)

var currentMode Mode
var recorder Drago_Recorder.Recorder

func init() {
	recorder = *Drago_Recorder.NewRecorder()

}

func CurrentMode() Mode {
	return currentMode
}
func RecordEvent(e event.Event) {
	if currentMode == ModeRecord {

	}
}
