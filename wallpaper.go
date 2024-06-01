package wallpaper

import (
	"fmt"
	"os"
)

// Mode wallpaper fill strategies.
type Mode uint

const (
	ModeZoom Mode = iota
	ModeScale
	ModeCenter
	ModeTile
	ModeOriginal
)

// SetFromFile sets the desktop wallpaper from a file.
func SetFromFile(file string, mode Mode) error {
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("%s is a directory", file)
	}
	return setFromFile(file, mode)
}

// Modes returns all the modes.
func Modes() []Mode {
	return modes()
}
