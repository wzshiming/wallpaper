package wallpaper

import (
	"fmt"
	"os"
	"strings"
)

type linuxDesktop uint

const (
	linuxDesktopUnkonwn linuxDesktop = iota
	linuxDesktopGnome
	linuxDesktopKDE
	linuxDesktopXFCE
	linuxDesktopLXDE
	linuxDesktopCinnamon
	linuxDesktopMATE
	linuxDesktopDeepin
)

func getDesktop(desktop string) (linuxDesktop, error) {
	switch {
	case desktop == "Unity",
		desktop == "Pantheon",
		strings.Contains(desktop, "GNOME"):
		return linuxDesktopGnome, nil
	case desktop == "KDE":
		return linuxDesktopKDE, nil
	case desktop == "XFCE":
		return linuxDesktopXFCE, nil
	case desktop == "LXDE",
		desktop == "LXQT":
		return linuxDesktopLXDE, nil
	case desktop == "X-Cinnamon":
		return linuxDesktopCinnamon, nil
	case desktop == "MATE":
		return linuxDesktopMATE, nil
	case desktop == "Deepin":
		return linuxDesktopDeepin, nil
	default:
		return linuxDesktopUnkonwn, fmt.Errorf("unknown desktop: %v", desktop)
	}
}

func setFromFile(file string, mode Mode) error {
	var desktop = os.Getenv("XDG_CURRENT_DESKTOP")
	d, err := getDesktop(desktop)
	if err != nil {
		return err
	}
	switch d {
	case linuxDesktopGnome:
		return setGnome(file, mode)
	case linuxDesktopKDE:
		return setKDE(file, mode)
	case linuxDesktopXFCE:
		return setXFCE(file, mode)
	case linuxDesktopLXDE:
		return setLXDE(file, mode)
	case linuxDesktopCinnamon:
		return setChinnamon(file, mode)
	case linuxDesktopMATE:
		return setMeta(file, mode)
	case linuxDesktopDeepin:
		return setDeepin(file, mode)
	default:
		return fmt.Errorf("unknown desktop: %v", d)
	}
}

func modes() []Mode {
	return []Mode{
		ModeZoom,
		ModeCenter,
		ModeTile,
		ModeOriginal,
		ModeScale,
	}
}
