//go:build linux
// +build linux

package wallpaper

func setDeepin(file string, mode Mode) error {
	err := setDconfMode(mode, "/com/deepin/wrap/gnome/desktop/background/picture-options")
	if err != nil {
		return err
	}
	return setDconfFile("file://"+file, "/com/deepin/wrap/gnome/desktop/background/picture-uri")
}
