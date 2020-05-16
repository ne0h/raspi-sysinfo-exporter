// Package load returns the system load.
package load

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ne0h/raspi-sysinfo-exporter/internal/fileutil"
	"github.com/ne0h/raspi-sysinfo-exporter/pkg/types"
)

// Get returns the system load.
func Get() (*types.Load, error) {
	input, err := fileutil.Get("/proc/loadavg")
	if err != nil {
		return nil, err
	}
	tokens := strings.Split(input, " ")

	load1, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s': %v", tokens[0], err)
	}
	load5, err := strconv.ParseFloat(tokens[1], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s': %v", tokens[1], err)
	}
	load15, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s': %v", tokens[2], err)
	}

	return &types.Load{
		Load1:  load1,
		Load5:  load5,
		Load15: load15,
	}, nil
}
