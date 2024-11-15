// A service layer for CPU calculations
package services

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"log"
	"path/filepath"

	"os"
	"os/exec"
	"strconv"
	"strings"
)


type Process struct {
    pid int
    cpu float64
}


func getCpuLoad() (int64, error){
	// Get cpu load
	var total float64
	cmd := exec.Command("ps", "aux")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    processes := make([]*Process, 0)
    for {
        line, err := out.ReadString('\n')
        if err!=nil {
            break;
        }
        tokens := strings.Split(line, " ")
        ft := make([]string, 0)
        for _, t := range(tokens) {
            if t!="" && t!="\t" {
                ft = append(ft, t)
            }
        }
        log.Println(len(ft), ft)
        pid, err := strconv.Atoi(ft[1])
        if err!=nil {
            continue
        }
        cpu, err := strconv.ParseFloat(ft[2], 64)
        if err!=nil {
            log.Fatal(err)
        }
        processes = append(processes, &Process{pid, cpu})
    }
    for _, p := range(processes) {
		total += p.cpu

    }
	return int64(total), nil
}


func getCpuTemp() (map[string]interface{}, error) {
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
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	
	return temps, nil
}