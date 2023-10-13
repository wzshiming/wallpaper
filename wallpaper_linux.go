package wallpaper

import (
	"github.com/elias-gill/wallpaper"
)

func setFromFile(file string) error {
	return wallpaper.SetFromFile(file)
}
