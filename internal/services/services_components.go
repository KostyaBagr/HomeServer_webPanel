// A service layer for server components
package services

import (
	"log"
)


func CpuDetailInfo() (map[string]interface{}, error) {
	// Get and return detail information about the CPU
	load, err := getCpuLoad()
	if err != nil{
		log.Fatal("Error during getting Cpu load")
	}
	temp, err := getCpuTemp()
	if err != nil{
		log.Fatal("Error during getting Cpu temp")
	}

	return map[string]interface{}{
		"temp": temp,
		"load": load,
	}, nil
}