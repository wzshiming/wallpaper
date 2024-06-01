//go:build linux
// +build linux

package wallpaper

func setMeta(file string, mode Mode) error {
	err := setDconfMode(mode, "/org/mate/desktop/background/picture-options")
	if err != nil {
		return err
	}
	return setDconfFile(file, "/org/mate/desktop/background/picture-filename")
}
