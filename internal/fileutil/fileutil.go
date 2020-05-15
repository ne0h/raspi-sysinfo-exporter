// Package fileutil provides utility functions to gather data from the filesystem.
package fileutil

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Get fetches the content of thee given file.
func Get(filename string) (string, error) {
	resp, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return "", err
	}
	return strings.Trim(string(resp), "\u0000\n"), nil
}
