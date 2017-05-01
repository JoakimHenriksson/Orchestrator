package service

import (
	"orchestrator/logging"
	"orchestrator/service/state"
	"orchestrator/service/testservice"
	"testing"
	"time"

	log "github.com/JoakimHenriksson/logrus"
)

func init() {
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	log.AddHook(logging.NewCallerHook(log.AllLevels...))
}

func TestStateMachine(t *testing.T) {
	var ts = testservice.New("Kaka")
	var sm = New(ts)
	sm.SetExpectedState(state.Creating)
	time.Sleep(time.Second)
}
