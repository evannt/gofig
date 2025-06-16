package main

import (
	"github.com/evannt/gofig/internal/flagparser"
	"github.com/evannt/gofig/internal/textrenderer"
)

func main() {
	config, display := flagparser.ParseFlags()

	if display {
		textrenderer.RenderText(config.Font, config.Text, config.Cols, config.Color)
	}
}
