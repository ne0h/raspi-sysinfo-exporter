package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ne0h/raspi-sysinfo-exporter/internal/fileutil"
	"github.com/ne0h/raspi-sysinfo-exporter/pkg/load"
	"github.com/ne0h/raspi-sysinfo-exporter/pkg/memory"
	"github.com/ne0h/raspi-sysinfo-exporter/pkg/platform"
	"github.com/ne0h/raspi-sysinfo-exporter/pkg/types"
)

func getTemperature() (float64, error) {
	temp, err := fileutil.Get("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}
	resp, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse '%s': %v", temp, err)
	}
	return resp / 1000, nil
}

func main() {

	// command-line arguments
	addr := flag.String("addr", ":8082", "Listen to <interface>:<port>, e.g. 'localhost:8082'")
	flag.Parse()

	// gather platform information. this is done only once when the exporter starts
	platform, err := platform.Get()
	if err != nil {
		log.Fatalf("Failed to get platform info: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temperature, err := getTemperature()
		if err != nil {
			log.Printf("[ERROR] Failed to get temperature: %v", err)
			return
		}
		load, err := load.Get()
		if err != nil {
			log.Printf("[ERROR] Failed to get system load: %v", err)
			return
		}
		memory, err := memory.Get()
		if err != nil {
			log.Printf("[ERROR] Failed to get memory usage: %v", err)
			return
		}

		resp := types.SysInfo{
			Platform:    platform,
			Load:        load,
			Memory:      memory,
			Temperature: temperature,
		}

		enc, err := json.MarshalIndent(resp, "", "\t")
		if err != nil {
			return
		}
		if _, err = w.Write(enc); err != nil {
			log.Printf("[ERROR] Failed to send response: %v", err)
		}
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
