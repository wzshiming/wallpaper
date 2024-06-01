//go:build linux
// +build linux

package wallpaper

import (
	"fmt"
	"os/exec"
	"strconv"
)

func setGnome(file string, mode Mode) error {
	m, err := getGNOMEMode(mode)
	if err != nil {
		return err
	}
	err = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-options", strconv.Quote(m)).Run()
	if err != nil {
		return err
	}
	return exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", strconv.Quote("file://"+file)).Run()
}

func getGNOMEMode(mode Mode) (string, error) {
	switch mode {
	case ModeCenter:
		return "centered", nil
	case ModeZoom:
		return "zoom", nil
	case ModeScale:
		return "scaled", nil
	case ModeTile:
		return "wallpaper", nil
	default:
		return "", fmt.Errorf("invalid wallpaper mode: %s", mode)
	}
}
