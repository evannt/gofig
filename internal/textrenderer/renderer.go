package textrenderer

import (
	"fmt"
	"github.com/evannt/gofig/internal/fontparser"
	"strings"
)

const colorReset = "\033[0m"

var colorMap = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}

func RenderText(font string, text string, cols int, color string) {
	fontDetails, err := fontparser.ParseFontFile(font)
	if err != nil {
		return
	}

	textColor := getColor(color)

	if cols < fontDetails.MaxLength {
		fmt.Printf("Specify a width of at least %s%d%s for font %s%s%s", textColor, fontDetails.MaxLength, colorReset, textColor, font, colorReset)
		return
	}

	if fontDetails.PrintDirection == 1 {
		runes := []rune(text)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		text = string(runes)
	}

	var replacementChar string
	if r, exists := fontDetails.Chars[0]; exists {
		replacementChar = r
	} else {
		replacementChar = ""
	}

	var builder strings.Builder
	for _, c := range text {
		if _, exists := fontDetails.Chars[c]; !exists {
			builder.WriteString(replacementChar)
		} else {
			builder.WriteRune(c)
		}
	}

	filteredText := builder.String()
	if filteredText == "" {
		return
	}

	asciiChars := make(map[rune][]string)
	for _, ascii := range filteredText {
		asciiChars[ascii] = strings.Split(fontDetails.Chars[ascii], fontDetails.EndMark)
	}

	fittedText := fitText(filteredText, cols, fontDetails)
	for _, line := range fittedText {
		for r := range fontDetails.Height {
			for _, ascii := range line {
				fmt.Printf("%s%s%s", textColor, asciiChars[ascii][r], colorReset)
			}
			fmt.Println()
		}
	}
}

func DisplaySupportedFonts() {
	fonts := fontparser.GetFonts()
	for _, f := range fonts {
		fmt.Println(f)
	}
}

func DisplaySupportedColors() {
	for c := range colorMap {
		fmt.Println(c)
	}
}

func fitText(text string, cols int, font fontparser.Font) []string {
	fittedText := []string{}
	words := strings.SplitSeq(text, " ")

	colCharCount := 0
	var builder strings.Builder

	for w := range words {
		wordWidth := 0
		for _, c := range w {
			wordWidth += strings.Index(font.Chars[c], font.EndMark)
		}

		spaceWidth := 0
		if builder.Len() > 0 {
			spaceWidth = strings.Index(font.Chars[' '], font.EndMark)
		}

		if colCharCount+spaceWidth+wordWidth <= cols {
			if builder.Len() > 0 {
				builder.WriteRune(' ')
				colCharCount += spaceWidth
			}
			builder.WriteString(w)
			colCharCount += wordWidth
		} else {
			if builder.Len() > 0 {
				fittedText = append(fittedText, builder.String())
				builder.Reset()
				colCharCount = 0
			}

			if wordWidth > cols {
				for _, c := range w {
					charWidth := strings.Index(font.Chars[c], font.EndMark)
					if colCharCount+charWidth > cols && builder.Len() > 0 {
						fittedText = append(fittedText, builder.String())
						colCharCount = 0
						builder.Reset()
					}
					builder.WriteRune(c)
					colCharCount += charWidth
				}
			} else {
				builder.WriteString(w)
				colCharCount = wordWidth
			}
		}
	}
	if colCharCount != 0 {
		fittedText = append(fittedText, builder.String())
	}
	return fittedText
}

func getColor(name string) string {
	if code, exists := colorMap[name]; exists {
		return code
	}
	return ""
}
