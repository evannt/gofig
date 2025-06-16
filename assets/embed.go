package assets

import (
	"embed"
)

//go:embed fonts/*.flf
var fontFiles embed.FS

func GetFontDir() embed.FS {
	return fontFiles
}
