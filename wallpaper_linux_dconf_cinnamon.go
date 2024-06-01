//go:build linux
// +build linux

package wallpaper

func setChinnamon(file string, mode Mode) error {
	err := setDconfMode(mode, "/org/cinnamon/desktop/background/picture-options")
	if err != nil {
		return err
	}
	return setDconfFile("file://"+file, "//org/cinnamon/desktop/background/picture-uri")
}
