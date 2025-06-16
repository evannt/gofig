# GoFig

A CLI tool that displays inputted text as ASCII characters
**GoFig** is a lightweight and fast CLI tool written in Go used to generate colorful ASCII art from the command line. Decorate your terminal banners, prompts, scripts, and README files!

## Features

- Render any text as ASCII art
- Optional terminal color support
- Supports .flf font files
- Optional width wrapping for pretty formatting

## Installation

First, ensure you have go installed, then you can install gofig globally with `go install`:

```bash
go install github.com/evannt/gofig/cmd/gofig@latest
```

Make sure $GOPATH/bin is in your $PATH to use gofig globally.

## Usage

```
gofig [OPTIONS]
```

### Text Options

| Option     | Description                         | Default         |
|------------|-------------------------------------|-----------------|
| `-t`       | Text to display in ASCII format     | `"GoFig(ure)"`  |
| `-f`       | Font used to display the text       | `"stforek"`     |
| `-c`       | Color for text rendering            | Terminal default|
| `-w`       | Maximum width for rendering text    | `80`            |

### Information Options

| Option     | Description              |
|------------|--------------------------|
| `-lf`      | List available fonts     |
| `-lc`      | List available colors    |
| `-help`    | Show help message        |

## Examples

### Display Ascii Art
```
gofig -t "Hello GoFig"
 _  _   ___   _     _      __       __    __    ___   _    __
| || | | __| | |   | |    /__\     / _]  /__\  | __| | |  / _]
| >< | | _|  | |_  | |_  | \/ |   | [/\ | \/ | | _|  | | | [/\
|_||_| |___| |___| |___|  \__/     \__/  \__/  |_|   |_|  \__/
```
```
gofig -t "Hello GoFig" -f big -c red
  _    _          _   _              _____           ______   _
 | |  | |        | | | |            / ____|         |  ____| (_)
 | |__| |   ___  | | | |   ___     | |  __    ___   | |__     _    __ _
 |  __  |  / _ \ | | | |  / _ \    | | |_ |  / _ \  |  __|   | |  / _` |
 | |  | | |  __/ | | | | | (_) |   | |__| | | (_) | | |      | | | (_| |
 |_|  |_|  \___| |_| |_|  \___/     \_____|  \___/  |_|      |_|  \__, |
                                                                   __/ |
                                                                  |___/
```
```
gofig -t "Welcome To GoFig(ure)" -f standard -w 120 -c blue
 __        __         _                                       _____
 \ \      / /   ___  | |   ___    ___    _ __ ___     ___    |_   _|   ___
  \ \ /\ / /   / _ \ | |  / __|  / _ \  | '_ ` _ \   / _ \     | |    / _ \
   \ V  V /   |  __/ | | | (__  | (_) | | | | | | | |  __/     | |   | (_) |
    \_/\_/     \___| |_|  \___|  \___/  |_| |_| |_|  \___|     |_|    \___/

   ____           _____   _            __                       __
  / ___|   ___   |  ___| (_)   __ _   / /  _   _   _ __    ___  \ \
 | |  _   / _ \  | |_    | |  / _` | | |  | | | | | '__|  / _ \  | |
 | |_| | | (_) | |  _|   | | | (_| | | |  | |_| | | |    |  __/  | |
  \____|  \___/  |_|     |_|  \__, | | |   \__,_| |_|     \___|  | |
                              |___/   \_\                       /_/
```

### List available fonts
```
gofig -lf
1row
3-d
3d diagonal
3d-ascii
3d
3d_diagonal
3x5
4max
5 line oblique
5lineoblique
acrobatic
alligator
alligator2
alligator3
alpha
alphabet
...
```

### List available colors
```
gofig -lc
cyan
white
black
red
green
yellow
blue
magenta
```

### Display Usage Information
```
gofig -help
```
