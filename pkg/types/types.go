// Package types encapsulates the types.
package types

// Load holds the system load.
type Load struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

// Memory holds information about the memory.
type Memory struct {
	Total     uint64 `json:"total"`
	Free      uint64 `json:"free"`
	Available uint64 `json:"available"`
}

// Platform holds information about the platform.
type Platform struct {
	Model  string `json:"model"`
	Serial string `json:"serial"`
}

// SysInfo holds information about the system.
type SysInfo struct {
	Platform    *Platform `json:"platform"`
	Load        *Load     `json:"load"`
	Memory      *Memory   `json:"memory"`
	Temperature float64   `json:"temperature"`
}
