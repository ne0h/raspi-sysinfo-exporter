// Package platform gathers information about the platform.
package platform

import (
	"github.com/ne0h/raspi-sysinfo-exporter/internal/fileutil"
	"github.com/ne0h/raspi-sysinfo-exporter/pkg/types"
)

// Get returns information about the platform.
func Get() (*types.Platform, error) {
	model, err := fileutil.Get("/proc/device-tree/model")
	if err != nil {
		return nil, err
	}
	serial, err := fileutil.Get("/proc/device-tree/serial-number")
	if err != nil {
		return nil, err
	}

	return &types.Platform{Model: model, Serial: serial}, nil
}
