//go:build !windows && !darwin && !linux
// +build !windows,!darwin,!linux

package wallpaper

import (
	"fmt"
)

func setFromFile(file string, mode Mode) error {
	return fmt.Errorf("invalid wallpaper mode: %s", mode)
}

func modes() []Mode {
	return []Mode{}
}
