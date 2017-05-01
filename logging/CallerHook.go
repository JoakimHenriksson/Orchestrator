package logging

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	log "github.com/JoakimHenriksson/logrus"
)

// CallerHook logrushook for callstack
type CallerHook struct {
	levels []log.Level
}

// NewCallerHook returns a new callerhook
func NewCallerHook(levels ...log.Level) (hook CallerHook) {
	hook.levels = levels
	return
}

// Levels returns the loglevels that should fire hook.
func (hook CallerHook) Levels() []log.Level {
	if len(hook.levels) == 0 {
		return log.AllLevels
	}
	return hook.levels
}

// Fire the hook
func (hook CallerHook) Fire(entry *log.Entry) error {
	var frame runtime.Frame
	var more = true
	var pc = make([]uintptr, 20)
	runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc)
	for more {
		frame, more = frames.Next()
		if !strings.Contains(frame.File, "logrus") {
			entry.Data["Trace"] = frameToString(frame)
			break
		}
	}
	return nil
}

func frameToString(frame runtime.Frame) string {
	return fmt.Sprintf("%s:%d:%s",
		filepath.Base(frame.File),
		frame.Line,
		frame.Function[strings.LastIndex(frame.Function, ".")+1:])
}
