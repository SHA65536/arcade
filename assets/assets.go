package assets

import (
	_ "embed"
)

//go:embed gameboy.txt
var Gameboy string

//go:embed mines_easy.txt
var EasyBoard string

//go:embed mines_medium.txt
var MediumBoard string

//go:embed mines_hard.txt
var HardBoard string
