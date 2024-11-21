// A service layer for server components
package services

import (
	"log"
	"strconv"
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
		"load": strconv.FormatInt(load, 10) + " %",
	}, nil
}


func RamDetailInfo() (map[string]interface{}, error){
	// Get and return detail ram information
	ram, err := ReadMemoryStats()
	if err != nil{
		log.Fatal("Could not get Ram info")
	}

	return map[string]interface{}{
		"ram": ram,
	}, nil
}

