package service

import (
	"context"
	"orchestrator/service/state"
)

// Service object
type Service interface {
	Create(context.Context) error
	Start(context.Context) error
	Stop(context.Context) error
	Delete(context.Context) error

	ServiceName() string
	ServiceID() string
	state.States
}

// Exec object
type Exec interface {
	Create(context.Context) error
	Start(context.Context) error
	Stop(context.Context) error
	Delete(context.Context) error

	ServiceName() string
	ServiceID() string
	state.States
}
