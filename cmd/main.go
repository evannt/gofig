package main

import (
	"fmt"
	"github.com/evannt/gofig/internal/fontparser"
)

func main() {
	fmt.Println("Hello GoFig(ure)")
	fontparser.ParseFontFile("big")
}
