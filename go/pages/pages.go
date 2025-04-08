package pages

import (
	"embed"
)

//go:generate ./convert.sh

//go:embed *.md
var MD embed.FS

//go:embed *.html
var HTML embed.FS
