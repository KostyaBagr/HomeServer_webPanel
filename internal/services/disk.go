package services

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

type DiskSummary struct {
	TotalFreeSpace string `json:"total_free_space"`
	TotalUsedSpace string `json:"total_used_space"` 
	TotalSpace     string `json:"total_space"`      
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
		TotalFreeSpace: fmt.Sprintf("%d GB", totalFree/(1024*1024*1024)),
		TotalUsedSpace: fmt.Sprintf("%d GB", totalUsed/(1024*1024*1024)),
		TotalSpace:     fmt.Sprintf("%d GB", totalSpace/(1024*1024*1024)),
	}, nil
}
