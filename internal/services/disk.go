package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
)

type DiskStatus struct {
	All  string `json:"All"`
	Used string `json:"Used"`
	Free string `json:"Free"`
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
			
			gbAll := strconv.FormatUint(uint64(fs.Blocks * uint64(fs.Bsize)) / GB, 10) + " GB"
			gbFree := strconv.FormatUint(uint64(fs.Bfree * uint64(fs.Bsize)) / GB, 10) + " GB"

			usedSpace := uint64(fs.Blocks * uint64(fs.Bsize)) / GB - uint64(fs.Bfree * uint64(fs.Bsize)) / GB
			gbUsed := strconv.FormatUint(usedSpace, 10) + " GB"

			disk := DiskStatus{
				All:  gbAll,
				Free: gbFree,
				Used: gbUsed,
			}
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
