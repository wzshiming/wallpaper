//go:build linux
// +build linux

package wallpaper

import (
	"fmt"
	"os/exec"
)

func setLXDE(file string, mode Mode) error {
	m, err := getLXDEMode(mode)
	if err != nil {
		return err
	}
	err = exec.Command("pcmanfm", "--wallpaper-mode",
		m).Run()
	if err != nil {
		return err
	}
	return exec.Command("pcmanfm", "-w", file).Run()
}

func getLXDEMode(mode Mode) (string, error) {
	switch mode {
	case ModeCenter:
		return "center", nil
	case ModeOriginal:
		return "center", nil
	case ModeScale:
		return "fit", nil
	case ModeZoom:
		return "fill", nil
	case ModeTile:
		return "tile", nil
	default:
		return "", fmt.Errorf("invalid wallpaper mode: %s", mode)
	}
}
