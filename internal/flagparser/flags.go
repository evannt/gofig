package flagparser

import (
	"flag"
	"fmt"
	"github.com/evannt/gofig/internal/textrenderer"
	"os"
)

const commandName = "gofig"

const gofigAsciiArt = ` _____  _____  _____  ___  _____  ___ __ __  _____  _____ ___
/   __\/  _  \/   __\/___\/   __\/  //  |  \/  _  \/   __\\  \
|  |_ ||  |  ||   __||   ||  |_ || | |  |  ||  _  <|   __| | |
\_____/\_____/\__/   \___/\_____/\__\\_____/\__|\_/\_____//__/
`

type Config struct {
	Text  string
	Font  string
	Cols  int
	Color string
}

func ParseFlags() (*Config, bool) {
	flag.Usage = gofigUsage

	help := flag.Bool("help", false, "Show usage information.")
	listFonts := flag.Bool("lf", false, "List available fonts.")
	listColors := flag.Bool("lc", false, "List available colors")
	text := flag.String("t", "GoFig(ure)", "Text to display in ascii format.")
	font := flag.String("f", "stforek", "Font used to display the provided text.")
	cols := flag.Int("w", 80, "The maximum width size when rendering text.")
	color := flag.String("c", "", "The color in which text will be rendered. Providing a non supported color will cause the terminal default color to be used.")

	flag.Parse()
	config := &Config{*text, *font, *cols, *color}
	displayText := !*help && !*listFonts && !*listColors

	if *help {
		flag.Usage()
	}

	if *listFonts && *listColors {
		fmt.Println("Specify -lf or -lc independently.")
	}

	if *listFonts {
		textrenderer.DisplaySupportedFonts()
	}

	if *listColors {
		textrenderer.DisplaySupportedColors()
	}
	return config, displayText
}

func gofigUsage() {
	fmt.Fprintf(os.Stderr, "%s\n", gofigAsciiArt)
	fmt.Fprintf(os.Stderr, "ASCII Art Text Generator\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", commandName)

	fmt.Fprintf(os.Stderr, "Text Options:\n")
	fmt.Fprintf(os.Stderr, "  -t string             Text to display in ASCII format (default: \"GoFig(ure)\")\n")
	fmt.Fprintf(os.Stderr, "  -f string             Font used to display the text (default: \"stforek\")\n")
	fmt.Fprintf(os.Stderr, "  -c string             Color for text rendering (default: terminal default)\n")
	fmt.Fprintf(os.Stderr, "  -w int                Maximum width when rendering text (default: 80)\n")

	fmt.Fprintf(os.Stderr, "\nInformation:\n")
	fmt.Fprintf(os.Stderr, "  -lf                   List available fonts\n")
	fmt.Fprintf(os.Stderr, "  -lc                   List available colors\n")
	fmt.Fprintf(os.Stderr, "  -help                 Show this help message\n")

	fmt.Fprintf(os.Stderr, "\nExamples:\n")
	fmt.Fprintf(os.Stderr, "  %s -t \"Hello World\"\n", commandName)
	fmt.Fprintf(os.Stderr, "  %s -t \"Welcome\" -f big -c red\n", commandName)
	fmt.Fprintf(os.Stderr, "  %s -t \"\" -w 120 -c blue\n", commandName)
	fmt.Fprintf(os.Stderr, "  %s -lf\n", commandName)
	fmt.Fprintf(os.Stderr, "  %s -lc\n", commandName)

	fmt.Fprintf(os.Stderr, "\nNotes:\n")
	fmt.Fprintf(os.Stderr, "  • Use -lf to see all available fonts before choosing one\n")
	fmt.Fprintf(os.Stderr, "  • Use -lc to see all supported colors\n")
	fmt.Fprintf(os.Stderr, "  • Invalid colors will fallback to terminal default\n")
	fmt.Fprintf(os.Stderr, "  • Width affects text wrapping and formatting\n")
}
