package health

import "time"

type (
	CheckResponseDTO struct {
		Ts     time.Time      `json:"ts"`
		Pid    int            `json:"pid"`
		Uptime string         `json:"uptime"`
		Memory MemoryUsageDTO `json:"memory"`
		Status string         `json:"status"`
	}

	MemoryUsageDTO struct {
		Alloc      uint64 `json:"alloc"`
		TotalAlloc uint64 `json:"totalAlloc"`
		Sys        uint64 `json:"sys"`
		HeapAlloc  uint64 `json:"heapAlloc"`
		HeapSys    uint64 `json:"heapSys"`
		NumGC      uint32 `json:"numGC"`
	}
)
