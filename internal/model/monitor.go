package model

import "time"

// MonitorData holds system monitoring metrics.
type MonitorData struct {
	CPU       float64   `json:"cpu"`
	Memory    float64   `json:"memory"`
	Disk      float64   `json:"disk"`
	Network   NetStats  `json:"network"`
	Timestamp time.Time `json:"timestamp"`
}

// NetStats holds network statistics.
type NetStats struct {
	BytesSent uint64 `json:"bytes_sent"`
	BytesRecv uint64 `json:"bytes_recv"`
}
