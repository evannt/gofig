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

const asciiStartCode = 32
const asciiEndCode = 126

type Font struct {
	HardBlank      byte
	Height         int
	Baseline       int
	MaxLength      int
	OldLayout      int
	PrintDirection int
	Chars          map[int]string
}

func ParseFontFile(fileName string) (font Font, e error) {
	fmt.Printf("Loading font \"%s\"\n", fileName)
	file, err := os.Open(fontFilePath + fileName + fileExtension)
	if hasError(err) {
		return font, e
	}
	defer file.Close()

	lineNumber := 0
	scanner := bufio.NewScanner(file)

	// Read flf header information
	scanner.Scan()
	header := strings.Split(scanner.Text(), " ")
	lineNumber++
	font.HardBlank = header[0][5]

	headerValues := parseHeader(header[1:])
	font.Height = headerValues[0]
	font.Baseline = headerValues[1]
	font.MaxLength = headerValues[2]
	font.OldLayout = headerValues[3]
	commentLines := headerValues[4]

	if len(headerValues) > 5 {
		font.PrintDirection = headerValues[5]
	} else {
		font.PrintDirection = 0
	}

	// Skip past comment lines
	for lineNumber <= commentLines {
		scanner.Scan()
		lineNumber++
	}

	font.Chars = make(map[int]string)
	// Parse lines with characters
	for asciiCode := asciiStartCode; asciiCode <= asciiEndCode; asciiCode++ {
		var builder strings.Builder
		for range font.Height {
			scanner.Scan()
			line := scanner.Text()
			builder.WriteString(line)

		}
		font.Chars[asciiCode] = builder.String()
	}

	// Parse Deutsch characters
	deutschAscii := []int{196, 214, 220, 228, 246, 252, 223}
	for _, ascii := range deutschAscii {
		var builder strings.Builder
		for range font.Height {
			scanner.Scan()
			line := scanner.Text()
			builder.WriteString(line)
		}
		font.Chars[ascii] = builder.String()
	}

	// Parse code tagged characters
	for scanner.Scan() {
		var builder strings.Builder
		tagLine := strings.Split(scanner.Text(), "  ")
		var tagCode int
		if num, err := parseTagCode(tagLine[0]); !hasError(err) {
			tagCode = num
			if tagCode == -1 { // Illegal code
				continue
			}
			for range font.Height {
				scanner.Scan()
				line := scanner.Text()
				builder.WriteString(line)
			}
			font.Chars[tagCode] = builder.String()
		}
	}

	if err := scanner.Err(); hasError(err) {
		return font, e
	}
	return font, nil
}

func parseHeader(header []string) []int {
	headerValues := make([]int, len(header))
	for i, s := range header {
		if num, err := strconv.Atoi(s); !hasError(err) {
			headerValues[i] = num
		}
	}
	return headerValues
}

func parseTagCode(value string) (int, error) {
	base := 10
	negative := false
	parseValue := value

	if value[0] == '-' {
		negative = true
		value = value[1:]
	}
	if value[0] == '0' {
		if len(value) >= 3 && (value[1] == 'x' || value[1] == 'X') {
			// hex
			base = 16
			parseValue = value[2:]
		} else {
			// octal
			base = 8
			parseValue = value[1:]
		}
	}
	if negative {
		parseValue = "-" + parseValue
	}
	num, err := strconv.ParseInt(parseValue, base, 0)
	return int(num), err
}

func hasError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
