package testservice

import (
	"context"
	"orchestrator/service"
	"orchestrator/service/state"
)

// TestService type
type TestService struct {
	ID string

	state.States
}

// New creates a new TestService
func New(ID string) (service service.Service) {
	service = &TestService{
		ID:     ID,
		States: &state.State{},
	}

	return
}

// Create a testservice
func (service *TestService) Create(ctx context.Context) (err error) {
	return
}

// Start a test service
func (service *TestService) Start(ctx context.Context) (err error) {
	return
}

// Stop a test service
func (service *TestService) Stop(ctx context.Context) (err error) {
	return
}

// Delete a test service
func (service *TestService) Delete(ctx context.Context) (err error) {
	return
}

// ServiceName returns the name
func (service TestService) ServiceName() string {
	return "TestService"
}

// ServiceID returns ID
func (service TestService) ServiceID() string {
	return service.ID
}
