//go:build linux
// +build linux

package wallpaper

import (
	"os/exec"
	"strconv"
)

func setDconfMode(mode Mode, path string) error {
	m, err := getGNOMEMode(mode)
	if err != nil {
		return err
	}
	return setDconf(m, path)
}

func setDconfFile(file string, path string) error {
	return setDconf(file, path)
}

func setDconf(value string, path string) error {
	return exec.Command("dconf", "write", path,
		strconv.Quote(value)).Run()
}
