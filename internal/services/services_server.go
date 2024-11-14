package services
// A service layer for server functionality

import (
	"log"
	
	"github.com/dhamith93/systats"
)


func ServerConfiguration() (interface{}, error) {
	// Get CPU info from the service layer
	syStats := systats.New()

	cpu, err := syStats.GetCPU()
	if err != nil {
		log.Printf("Could not get the CPU info: %v", err)
		return nil, err
	}

	// Get memory info
	memory, err := syStats.GetMemory(systats.Megabyte)
	if err != nil {
		log.Printf("Could not get the memory info: %v", err)
		return nil, err
	}

	// Get network info
	networks, err := syStats.GetNetworks()
	if err != nil {
		log.Printf("Could not get the network info: %v", err)
		return nil, err
	}

	
	return map[string]interface{}{
		"cpu":       cpu,
		"memory":    memory,
		"networks":  networks,
	}, nil
}
