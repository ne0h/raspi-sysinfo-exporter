// Package types encapsulates the types.
package types

// Platform holds information about the platform.
type Platform struct {
	Model  string `json:"model"`
	Serial string `json:"serial"`
}

// SysInfo holds information about the system.
type SysInfo struct {
	Platform    *Platform `json:"platform"`
	Temperature float64   `json:"temperature"`
}
