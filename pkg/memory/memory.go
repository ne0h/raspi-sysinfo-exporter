// Package memory returns information about memory usage.
package memory

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ne0h/raspi-sysinfo-exporter/internal/fileutil"
	"github.com/ne0h/raspi-sysinfo-exporter/pkg/types"
)

// Get returns information about the memory.
func Get() (*types.Memory, error) {
	input, err := fileutil.Get("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	var (
		total     = uint64(0)
		free      = uint64(0)
		available = uint64(0)
	)
	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Fields(line)
		switch tokens[0] {
		case "MemTotal:":
			t, err := strconv.ParseUint(tokens[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse '%s': %v", tokens[1], err)
			}
			total = t
		case "MemFree:":
			f, err := strconv.ParseUint(tokens[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse '%s': %v", tokens[1], err)
			}
			free = f
		case "MemAvailable:":
			a, err := strconv.ParseUint(tokens[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse '%s': %v", tokens[1], err)
			}
			available = a
		}
	}
	return &types.Memory{
		Total:     total,
		Free:      free,
		Available: available,
	}, nil
}
