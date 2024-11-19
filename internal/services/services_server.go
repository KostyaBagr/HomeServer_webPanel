package services

// A service layer for server functionality

import (
	"log"
	"os/exec"

	"github.com/dhamith93/systats"
)


func ServerConfiguration() (interface{}, error) {
	// Get Server configuration.

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

func RebootServer()  (string, error){
	// Reboot 
	cmd := exec.Command("sudo", "/sbin/reboot")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error rebooting computer:", err)
	}
	return "Ok", nil
}

func PowerOffserver() (string, error){
	// Poweroff 
	cmd := exec.Command("sudo", "/sbin/shutdown", "-h", "now")
	err := cmd.Run()
	if err != nil {
			log.Fatal("Error shutting down computer:", err)
	}
	return "Ok", nil
}
