// A service layer for RAM calculations
package services

import (
	"bufio"

	"os"
	"strconv"
	"strings"
)



type Memory struct {
    MemTotal     string
    MemFree      string
    MemAvailable string
}


func parseLine(raw string) (key string, value string) {
	// Parse lines in file /proc/meminfo
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	totalVal := int64((toInt(keyValue[1]) + 1048576 - 1) / 1048576)
	totalValStr := strconv.FormatInt(totalVal, 10) + " Gb"
	return keyValue[0], totalValStr
}

func toInt(raw string) int {
	// covert str to int
    if raw == "" {
        return 0
    }
    res, err := strconv.Atoi(raw) 
    if err != nil {
        panic(err)
    }
    return res
}

func ReadMemoryStats() (Memory, error) {
    file, err := os.Open("/proc/meminfo")
    if err != nil {
        panic(err)
    }

    defer file.Close()
    bufio.NewScanner(file)
    scanner := bufio.NewScanner(file)
    res := Memory{}
    for scanner.Scan() {
        key, value := parseLine(scanner.Text())
        switch key {
        case "MemTotal":
            res.MemTotal = value 
        case "MemFree":
            res.MemFree = value 
        case "MemAvailable":
            res.MemAvailable = value 
        }
    }
    return res, nil
}