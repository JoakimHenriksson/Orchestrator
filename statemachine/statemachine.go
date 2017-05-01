package service

import (
	"fmt"
	"orchestrator/service"
	"orchestrator/service/state"

	log "github.com/JoakimHenriksson/logrus"
)

var logger *log.Logger

func init() {
	logger = log.New()
}

// StateMachine handles states
type StateMachine struct {
	newStateChan chan state.ServiceState
	stateChan    chan state.ServiceState

	LOG *log.Entry

	service service.Service
}

// New creates a new statemachine
func New(service service.Service) (sm *StateMachine) {
	sm = &StateMachine{
		newStateChan: make(chan state.ServiceState, 1),
		stateChan:    make(chan state.ServiceState, 1),
		service:      service,
	}

	sm.updateLogger()
	go sm.newState()

	return
}

// SetExpectedState sets the expectedstate
func (sm *StateMachine) SetExpectedState(state state.ServiceState) {
	sm.newStateChan <- state
}

func (sm *StateMachine) newState() {
	for newState := range sm.newStateChan {
		fmt.Println("NewState")
		var err error
		log := sm.LOG.WithField("NewState", state.String(newState))
		sm.service.SetExpectedState(newState)
		log.WithError(err).Trace("New State")
		log.WithError(err).Debug("New State2")
		sm.updateLogger()
	}
}

func (sm *StateMachine) updateLogger() {
	sm.LOG = log.WithFields(log.Fields{
		"ServiceName":   sm.service.ServiceName(),
		"ServiceID":     sm.service.ServiceID(),
		"State":         state.String(sm.service.GetState()),
		"ExpectedState": state.String(sm.service.GetExpectedState()),
	})
}
