package fontparser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fontFilePath = "assets/fonts/"
const fileExtension = ".flf"

type FontDetails struct {
	hardBlank      byte
	height         int
	baseline       int
	maxLength      int
	oldLayout      int
	printDirection int
	fullLayout     int
	codetagCount   int
	chars          map[rune]string
}

func ParseFontFile(fileName string) (fontDetails FontDetails, e error) {
	fmt.Printf("Loading font \"%s\"\n", fileName)
	file, err := os.Open(fontFilePath + fileName + fileExtension)
	if err != nil {
		fmt.Println(err)
		return fontDetails, e
	}
	defer file.Close()

	lineNumber := 0
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	header := strings.Split(scanner.Text(), " ")
	commentLines, err := strconv.Atoi(header[5])
	fmt.Println(header)
	lineNumber++

	// Skip past comment lines
	for lineNumber <= commentLines {
		scanner.Scan()
		lineNumber++
	}
	fmt.Println("Successfully Skipped Comments")

	// Parse lines with characters

	// Parse Deutsch characters

	// Parse code tagged characters

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return fontDetails, e
	}
	return fontDetails, nil
}
