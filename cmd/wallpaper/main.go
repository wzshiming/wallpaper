package main

import (
	"fmt"
	"os"

	"github.com/wzshiming/wallpaper"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: wallpaper <file>")
		return
	}
	err := wallpaper.SetFromFile(args[1])
	if err != nil {
		panic(err)
	}
}
