// A service layer for CPU calculations
package services
// #include <time.h>
import "C"
import (
	// "bytes"
	"fmt"

	"io/ioutil"
	"log"
	"path/filepath"

	"os"
	"time"
	"strconv"
	"strings"
	linuxproc "github.com/c9s/goprocinfo/linux"
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


func GetCpuTemp() (map[string]interface{}, error) {
	// Get and return CPU temp
	// TODO: check temp file on linux machine
	thermalDir := "/sys/class/thermal/"
	temps := make(map[string]interface{})

	err := filepath.Walk(thermalDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if strings.HasSuffix(path, "temp") {
			fmt.Printf("Reading temperature from %s\n", path)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalf("Could not read file %s", path)
				return err
			}

			tempStr := strings.TrimSpace(string(data))
			fmt.Printf("Raw temperature data: %s\n", tempStr)
			temp, err := strconv.Atoi(tempStr)
			if err != nil {
				log.Fatalf("Error converting temperature data %q: %v\n", tempStr, err)
				return err
			}
			tempCelsius := float64(temp) / 1000.0
			temps[path] = tempCelsius
			fmt.Println(temps)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	fmt.Println(temps)
	return temps, nil
}