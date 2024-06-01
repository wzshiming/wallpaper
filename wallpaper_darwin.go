package wallpaper

import (
	"fmt"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func setFromFile(file string, mode Mode) error {
	err := getDarwinMode(mode)
	if err != nil {
		return err
	}

	workspace := appkit.Workspace_SharedWorkspace()
	screen := appkit.Screen_MainScreen()

	if ok := workspace.SetDesktopImageURLForScreenOptionsError(
		foundation.URL_FileURLWithPath(file),
		screen,
		map[appkit.WorkspaceDesktopImageOptionKey]objc.IObject{},
		foundation.IError(nil),
	); !ok {
		return fmt.Errorf("could not set desktop image: %s", file)
	}

	return nil
}

func modes() []Mode {
	return []Mode{
		ModeZoom,
	}
}

func getDarwinMode(mode Mode) error {
	switch mode {
	case ModeZoom:
		return nil
	default:
		return fmt.Errorf("invalid wallpaper mode: %s", mode)
	}
}
