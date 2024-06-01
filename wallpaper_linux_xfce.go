//go:build linux
// +build linux

package wallpaper

import (
	"fmt"
	"os/exec"
)

func setXFCE(file string, mode Mode) error {
	err := setXFCEMode(mode)
	if err != nil {
		return err
	}

	return setXFCEFile(file)
}

func setXFCEMode(mode Mode) error {
	m, err := getXFCEMode(mode)
	if err != nil {
		return err
	}
	styles, err := getXFCEProps("image-style")
	if err != nil {
		return err
	}
	for _, style := range styles {
		err = exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--property", style, "--set", m).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func setXFCEFile(file string) error {
	desktops, err := getXFCEProps("last-image")
	if err != nil {
		return err
	}
	for _, desktop := range desktops {
		err = exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--property", desktop, "--set", file).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func getXFCEProps(key string) ([]string, error) {
	output, err := exec.Command("xfconf-query", "--channel", "xfce4-desktop", "--list").Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.Trim(string(output), "\n"), "\n")
	var desktops []string

	for _, line := range lines {
		if path.Base(line) == key {
			desktops = append(desktops, line)
		}
	}
	return desktops, nil
}

func getXFCEMode(mode Mode) (string, error) {
	switch mode {
	case ModeCenter:
		return "1", nil
	case ModeZoom:
		return "4", nil
	case ModeOriginal:
		return "5", nil
	case ModeScale:
		return "3", nil
	case ModeTile:
		return "2", nil
	default:
		return "", fmt.Errorf("invalid wallpaper mode: %s", mode)
	}
}
