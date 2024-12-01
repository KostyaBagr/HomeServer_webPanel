// A service layer for CPU calculations
package services
// #include <time.h>
import "C"
import (
	"fmt"
	"time"
	"os"

	"bufio"
	"strconv"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/KostyaBagr/HomeServer_webPanel/pkg/settings"
)



func calcSingleCoreUsage(curr, prev linuxproc.CPUStat) float32 {

	PrevIdle := prev.Idle + prev.IOWait
	Idle := curr.Idle + curr.IOWait

	PrevNonIdle := prev.User + prev.Nice + prev.System + prev.IRQ + prev.SoftIRQ + prev.Steal
	NonIdle := curr.User + curr.Nice + curr.System + curr.IRQ + curr.SoftIRQ + curr.Steal

	PrevTotal := PrevIdle + PrevNonIdle
	Total := Idle + NonIdle
	totald := Total - PrevTotal
	idled := Idle - PrevIdle

	if totald == 0 {
			return 0.0
	}

	CPU_Percentage := (float32(totald) - float32(idled)) / float32(totald) * 100.0

	return CPU_Percentage
}

func getCPUStats() (linuxproc.CPUStat, error) {
	stats, err := linuxproc.ReadStat("/proc/stat")
	if err != nil{
		return linuxproc.CPUStat{}, err
	}
	return stats.CPUStatAll, nil
	
}

func GetCPUInfo() (float32, error){
	prevStat, err := getCPUStats()
	if err != nil {
			fmt.Println("Ошибка чтения CPUStat:", err)
			return 0.0, err
	}
	time.Sleep(1 * time.Second)
	currStat, err := getCPUStats()
	if err != nil {
			fmt.Println("Ошибка чтения CPUStat:", err)
			return 0.0, err
	}

	load := calcSingleCoreUsage(currStat, prevStat)
	return load, nil

}


func GetCPUTemp() (int, error) {
	// Get and return CPU temp
	// TODO: check temp file on linux machine
	path := settings.AppSetting.TempFilePath

	file, err := os.Open(path)
    if err != nil {
        return 0, err
    }
    defer func() {
        if err = file.Close(); err != nil {
            return
        }
    }()

    scanner := bufio.NewScanner(file)
	var res int
    for scanner.Scan() {
        output, _ := strconv.Atoi(scanner.Text())
		res = output / 1000

    }
	return res, nil
}