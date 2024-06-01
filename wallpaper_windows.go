package wallpaper

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724947.aspx
const (
	spiSetDeskWallpaper = 0x0014

	uiParam = 0x0000

	spifUpdateINIFile = 0x01
	spifSendChange    = 0x02
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
)

func setFromFile(file string, mode string) error {
	err := setWindowsMode(mode)
	if err != nil {
		return err
	}

	return setWindowsFile(file)
}

func setWindowsFile(file string) error {
	filenameUTF16, err := syscall.UTF16PtrFromString(file)
	if err != nil {
		return err
	}

	systemParametersInfo.Call(
		uintptr(spiSetDeskWallpaper),
		uintptr(uiParam),
		uintptr(unsafe.Pointer(filenameUTF16)),
		uintptr(spifUpdateINIFile|spifSendChange),
	)
	return nil
}

func setWindowsMode(mode string) error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, "Control Panel\\Desktop", registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	err = key.SetStringValue("TileWallpaper", "0")
	if err != nil {
		return err
	}

	m, err := getWindowsMode(mode)
	if err != nil {
		return err
	}

	return key.SetStringValue("WallpaperStyle", m)
}

func getWindowsMode(mode Mode) (string, error) {
	switch mode {
	case ModeCenter:
		return "0", nil
	case ModeOriginal:
		return "6", nil
	case ModeZoom:
		return "22", nil
	case ModeScale:
		return "2", nil
	default:
		return "", fmt.Errorf("invalid wallpaper mode: %s", mode)
	}
}

func modes() []Mode {
	return []Mode{
		ModeZoom,
		ModeCenter,
		ModeOriginal,
		ModeScale,
	}
}
