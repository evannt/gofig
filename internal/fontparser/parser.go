package fontparser

import (
	"bufio"
	"fmt"
	"github.com/evannt/gofig/assets"
	"strconv"
	"strings"
)

const fontFilePath = "fonts"
const fileExtension = ".flf"

const asciiStartCode = 32
const asciiEndCode = 126

type Font struct {
	HardBlank      byte
	EndMark        string
	Height         int
	Baseline       int
	MaxLength      int
	OldLayout      int
	PrintDirection int
	Chars          map[rune]string
}

func ParseFontFile(fileName string) (font Font, err error) {
	file, err := assets.GetFontDir().Open(fontFilePath + "/" + fileName + fileExtension)
	if err != nil {
		fmt.Printf("Font Not Supported: %s\n", fileName)
		return font, err
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

	// Parse lines with characters
	replacer := strings.NewReplacer(string(font.HardBlank), " ")
	font.Chars = make(map[rune]string)
	for asciiCode := asciiStartCode; asciiCode <= asciiEndCode; asciiCode++ {
		var builder strings.Builder
		for range font.Height {
			scanner.Scan()
			line := replacer.Replace(scanner.Text())
			builder.WriteString(line)
		}
		asciiChar := builder.String()
		font.Chars[rune(asciiCode)] = asciiChar[:len(asciiChar)-1]
	}

	// Parse Deutsch characters
	deutschAscii := []int{196, 214, 220, 228, 246, 252, 223}
	for _, ascii := range deutschAscii {
		var builder strings.Builder
		for range font.Height {
			scanner.Scan()
			line := replacer.Replace(scanner.Text())
			builder.WriteString(line)
		}
		asciiChar := builder.String()
		font.Chars[rune(ascii)] = asciiChar[:len(asciiChar)-1]
	}

	// Parse code tagged characters
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var builder strings.Builder
		tagLine := strings.Split(line, "  ")
		var tagCode int
		if num, err := parseTagCode(tagLine[0]); !hasError(err) {
			tagCode = num
			if tagCode == -1 { // Illegal code
				continue
			}
			for range font.Height {
				scanner.Scan()
				line := replacer.Replace(scanner.Text())
				builder.WriteString(line)
			}
			asciiChar := builder.String()
			font.Chars[rune(tagCode)] = asciiChar[:len(asciiChar)-1]
		}
	}

	firstChar := font.Chars['!']
	font.EndMark = string(firstChar[len(firstChar)-1])

	if err := scanner.Err(); hasError(err) {
		return font, err
	}
	return font, nil
}

func GetFonts() []string {
	fonts := []string{}

	files, err := assets.GetFontDir().ReadDir(fontFilePath)
	if err != nil {
		fmt.Println(err)
		return fonts
	}

	for _, f := range files {
		fonts = append(fonts, strings.Split(f.Name(), ".")[0])
	}

	return fonts
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
