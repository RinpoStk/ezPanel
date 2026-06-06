package service

import (
	"runtime"
	"time"

	"github.com/rinpostk/ezPanel/internal/model"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// MonitorCollector collects system monitoring data.
type MonitorCollector struct{}

// NewMonitorCollector creates a new MonitorCollector.
func NewMonitorCollector() *MonitorCollector {
	return &MonitorCollector{}
}

// Collect gathers current system metrics.
func (c *MonitorCollector) Collect() (*model.MonitorData, error) {
	data := &model.MonitorData{
		Timestamp: time.Now(),
	}

	// CPU usage
	if percentages, err := cpu.Percent(0, false); err == nil && len(percentages) > 0 {
		data.CPU = percentages[0]
	}

	// Memory usage
	if vm, err := mem.VirtualMemory(); err == nil {
		data.Memory = vm.UsedPercent
	}

	// Disk usage
	if d, err := disk.Usage("/"); err == nil {
		data.Disk = d.UsedPercent
	}
	if runtime.GOOS == "windows" {
		if d, err := disk.Usage("C:"); err == nil {
			data.Disk = d.UsedPercent
		}
	}

	// Network stats
	if counters, err := net.IOCounters(false); err == nil && len(counters) > 0 {
		data.Network = model.NetStats{
			BytesSent: counters[0].BytesSent,
			BytesRecv: counters[0].BytesRecv,
		}
	}

	return data, nil
}
