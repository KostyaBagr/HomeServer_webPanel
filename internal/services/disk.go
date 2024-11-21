package services

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
)
type DiskInfo struct {
	Device      string
	Mountpoint  string
	Fstype      string
	FreeSpace   uint64
	TotalSpace  uint64
	UsedSpace   uint64
	UsedPercent float64
}

type DisksInfo struct {
	Disks []DiskInfo
}

func (d *DisksInfo) Summary() string {
	var totalFree, totalUsed, totalSpace uint64
	for _, disk := range d.Disks {
		if disk.Fstype == "squashfs" {
			continue
		}
		totalFree += disk.FreeSpace
		totalUsed += disk.UsedSpace
		totalSpace += disk.TotalSpace
	}

	return fmt.Sprintf("Суммарное свободное место: %d GB\nСуммарный объем: %d GB\nСуммарно использовано: %d GB\n",
		totalFree/1024/1024/1024,
		totalSpace/1024/1024/1024,
		totalUsed/1024/1024/1024,
	)
}


func DiskUsage() (DisksInfo, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return DisksInfo{}, fmt.Errorf("ошибка при получении разделов: %w", err)
	}

	var disksInfo DisksInfo
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Printf("Ошибка для %s: %v\n", partition.Mountpoint, err)
			continue
		}

		diskInfo := DiskInfo{
			Device:      partition.Device,
			Mountpoint:  partition.Mountpoint,
			Fstype:      partition.Fstype,
			FreeSpace:   usage.Free,
			TotalSpace:  usage.Total,
			UsedSpace:   usage.Used,
			UsedPercent: usage.UsedPercent,
		}
		disksInfo.Disks = append(disksInfo.Disks, diskInfo)
	}
	return disksInfo, nil
}
