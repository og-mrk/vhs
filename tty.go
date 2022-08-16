package dolly

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// TTYOptions is the set of options to pass to `ttyd`.
type TTYOptions struct {
	Port       int
	FontFamily string
	FontSize   int
	LineHeight float64
	Debug      bool
	Theme      Theme
}

// DefaultTTYOptions is the set of options to pass to `ttyd` by default.
var DefaultTTYOptions = TTYOptions{
	Port:       randomPort(),
	FontFamily: "SF Mono",
	FontSize:   22,
	LineHeight: 1.2,
	Debug:      false,
	Theme:      DefaultTheme,
}

// Theme is a terminal theme for xterm
type Theme struct {
	Background    string `json:"background"`
	Foreground    string `json:"foreground"`
	Black         string `json:"black"`
	BrightBlack   string `json:"brightBlack"`
	Red           string `json:"red"`
	BrightRed     string `json:"brightRed"`
	Green         string `json:"green"`
	BrightGreen   string `json:"brightGreen"`
	Yellow        string `json:"yellow"`
	BrightYellow  string `json:"brightYellow"`
	Blue          string `json:"blue"`
	BrightBlue    string `json:"brightBlue"`
	Magenta       string `json:"magenta"`
	BrightMagenta string `json:"brightMagenta"`
	Cyan          string `json:"cyan"`
	BrightCyan    string `json:"brightCyan"`
	White         string `json:"white"`
	BrightWhite   string `json:"brightWhite"`
}

// DefaultTheme is the default theme to use for recording demos and
// screenshots. Taken from https://github.com/meowgorithm/dotfiles.
var DefaultTheme = Theme{
	Background:    "#171717",
	Foreground:    "#dddddd",
	Black:         "#000000",
	BrightBlack:   "#4d4d4d",
	Red:           "#c73b1d",
	BrightRed:     "#e82100",
	Green:         "#00a800",
	BrightGreen:   "#00db00",
	Yellow:        "#acaf15",
	BrightYellow:  "#e5e900",
	Blue:          "#3854FC",
	BrightBlue:    "#566BF9",
	Magenta:       "#d533ce",
	BrightMagenta: "#e83ae9",
	Cyan:          "#2cbac9",
	BrightCyan:    "#00e6e7",
	White:         "#bfbfbf",
	BrightWhite:   "#e6e6e6",
}

// StartTTY starts the ttyd process on the given port and options.
func StartTTY(opts TTYOptions) *exec.Cmd {
	theme, _ := json.Marshal(opts.Theme)

	cmd := exec.Command(
		"ttyd", fmt.Sprintf("--port=%d", opts.Port),
		"-t", fmt.Sprintf("fontFamily='%s'", opts.FontFamily),
		"-t", fmt.Sprintf("fontSize=%d", opts.FontSize),
		"-t", fmt.Sprintf("lineHeight=%f", opts.LineHeight),
		"-t", fmt.Sprintf("theme=%s", theme),
		"-t", "customGlyphs=true",
		"zsh", "-d", "-f",
	)

	if opts.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd
}