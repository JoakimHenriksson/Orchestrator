package state

import (
	"fmt"
	"sync"
)

// ServiceState is the state a service can have
type ServiceState int

const (
	// New is valid but before Create has been called
	New ServiceState = iota

	// Creating is the state after Create has been called
	Creating

	// Created is the state the service comes to after create has finished
	Created

	// Starting is the state a service get just after Start has been called
	Starting

	// Started is the state reached when Start is finished
	Started

	// Stopping is the state directly after stop has been called.
	Stopping

	// Stopped is the state the service reach after stop is finished
	Stopped

	// Deleting is the state reached when Delete has been called.
	Deleting

	// Deleted is the state reached after Delete has finished
	Deleted

	// Invalid is an invalid state
	Invalid ServiceState = -1
)

// String returns the text representation of a state
func String(state ServiceState) string {
	switch state {
	case Invalid:
		return "Invalid"
	case New:
		return "New"
	case Creating:
		return "Creating"
	case Created:
		return "Created"
	case Starting:
		return "Starting"
	case Started:
		return "Started"
	case Stopping:
		return "Stopping"
	case Stopped:
		return "Stopped"
	case Deleting:
		return "Deleting"
	case Deleted:
		return "Deleted"
	default:
		panic(fmt.Sprintf("Unknown state: %d", state))
	}
}

// States for a service
type States interface {
	GetState() ServiceState
	SetState(ServiceState)
	GetExpectedState() ServiceState
	SetExpectedState(ServiceState)
}

// State keeps a service's state
type State struct {
	sync.RWMutex
	ExpectedState ServiceState
	CurrentState  ServiceState
}

// GetState gets the current state
func (state *State) GetState() ServiceState {
	state.RLock()
	defer state.RUnlock()
	return state.CurrentState
}

// SetState sets the current state
func (state *State) SetState(newState ServiceState) {
	state.Lock()
	defer state.Unlock()
	state.CurrentState = newState
}

// GetExpectedState gets the expected State
func (state *State) GetExpectedState() ServiceState {
	state.RLock()
	defer state.RUnlock()
	return state.ExpectedState
}

//SetExpectedState sets the expected State
func (state *State) SetExpectedState(newExpectedState ServiceState) {
	state.Lock()
	defer state.Unlock()
	state.ExpectedState = newExpectedState
}
