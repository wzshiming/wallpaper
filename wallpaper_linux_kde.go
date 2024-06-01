//go:build linux
// +build linux

package wallpaper

import (
	"fmt"
	"os/exec"
	"strconv"
)

func setKDE(file string, mode Mode) error {
	err := setKDEMode(mode)
	if err != nil {
		return err
	}

	return setKDEFile(file)
}

func setKDEFile(file string) error {
	return evalKDE(`
for (const desktop of desktops()) {
	desktop.currentConfigGroup = ["Wallpaper", "org.kde.image", "General"]
	desktop.writeConfig("Image", ` + strconv.Quote("file://"+file) + `)
}
`)
}

func setKDEMode(mode Mode) error {
	m, err := getKDEMode(mode)
	if err != nil {
		return err
	}
	return evalKDE(`
for (const desktop of desktops()) {
	desktop.currentConfigGroup = ["Wallpaper", "org.kde.image", "General"]
	desktop.writeConfig("FillMode", ` + m + `)
}
`)
}

func evalKDE(script string) error {
	return exec.Command("qdbus", "org.kde.plasmashell", "/PlasmaShell", "org.kde.PlasmaShell.evaluateScript", script).Run()
}

func getKDEMode(mode Mode) (string, error) {
	switch mode {
	case ModeCenter:
		return "6", nil
	case ModeZoom:
		return "1", nil
	case ModeOriginal:
		return "2", nil
	case ModeScale:
		return "0", nil
	case ModeTile:
		return "3", nil
	default:
		return "", fmt.Errorf("invalid wallpaper mode: %s", mode)
	}
}
