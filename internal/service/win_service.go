package service

import "github.com/rinpostk/ezPanel/internal/model"

// WinServiceManager provides operations for managing Windows services.
// TODO: implement actual Windows service management via syscall or sc.exe.
type WinServiceManager struct{}

// NewWinServiceManager creates a new WinServiceManager.
func NewWinServiceManager() *WinServiceManager {
	return &WinServiceManager{}
}

// List returns all registered services.
func (m *WinServiceManager) List() ([]model.Service, error) {
	// TODO: implement
	return []model.Service{}, nil
}

// Start starts a service by name.
func (m *WinServiceManager) Start(name string) error {
	// TODO: implement
	return nil
}

// Stop stops a service by name.
func (m *WinServiceManager) Stop(name string) error {
	// TODO: implement
	return nil
}

// Status returns the status of a service by name.
func (m *WinServiceManager) Status(name string) (*model.Service, error) {
	// TODO: implement
	return &model.Service{}, nil
}
