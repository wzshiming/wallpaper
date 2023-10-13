package wallpaper

import (
	"github.com/getlantern/byteexec"
)

func setFromFile(file string) error {
	e, err := byteexec.New(wallpaperBytesExec, "wallpaper")
	if err != nil {
		return err
	}
	cmd := e.Command("set", file)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
