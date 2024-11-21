package services

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

type DiskSummary struct {
	TotalFreeSpace uint64 `json:"total_free_space"` 
	TotalUsedSpace uint64 `json:"total_used_space"` 
	TotalSpace     uint64 `json:"total_space"`      
}

func DiskUsageSummary() (DiskSummary, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return DiskSummary{}, fmt.Errorf("Error: %w", err)
	}

	var totalFree, totalUsed, totalSpace uint64
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Printf("Error for %s: %v\n", partition.Mountpoint, err)
			continue
		}

		if usage.Fstype == "squashfs" || usage.Total == 0 {
			continue
		}

		totalFree += usage.Free
		totalUsed += usage.Used
		totalSpace += usage.Total
	}

	return DiskSummary{
		TotalFreeSpace: totalFree / (1024 * 1024 * 1024),
		TotalUsedSpace: totalUsed / (1024 * 1024 * 1024),
		TotalSpace:     totalSpace / (1024 * 1024 * 1024),
	}, nil
}
