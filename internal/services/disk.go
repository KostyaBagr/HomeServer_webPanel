package services

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

type DiskStatus struct {
	All  uint64 `json:"All"`
	Used uint64 `json:"Used"`
	Free uint64 `json:"Free"`
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func DiskUsage() (map[string]DiskStatus, error) {
	// Find all disks in system. count free, total ans used space
	diskStats := make(map[string]DiskStatus)

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		path := "/mnt/" + strings.ToLower(string(drive))
		_, err := os.Stat(path)
		if err == nil {
			fs := syscall.Statfs_t{}
			err = syscall.Statfs(path, &fs)
			if err != nil {
				log.Printf("Could not get disk info for %s: %v", path, err)
				continue
			}

			disk := DiskStatus{
				All:  uint64(fs.Blocks*uint64(fs.Bsize)) / GB,
				Free: uint64(fs.Bfree*uint64(fs.Bsize)) / GB, 
			}
			disk.Used = disk.All - disk.Free
			name := strings.Split(path, "/")
			diskLetter := name[len(name)-1] 
			diskStats[diskLetter] = disk
		}
	}

	if len(diskStats) == 0 {
		return nil, fmt.Errorf("no accessible drives found")
	}

	return diskStats, nil
}
