package main

import (
	"fmt"
	"github.com/evannt/gofig/internal/fontparser"
)

func main() {
	fmt.Println("Hello GoFig(ure)")
	big, err := fontparser.ParseFontFile("big")
	if err != nil {
		return
	}
	fmt.Println(big.Chars[int('a')])
}
