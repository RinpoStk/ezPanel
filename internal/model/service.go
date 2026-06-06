package model

// Service represents a system service.
type Service struct {
	Name   string  `json:"name"`
	Status string  `json:"status"`
	PID    int32   `json:"pid"`
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}
