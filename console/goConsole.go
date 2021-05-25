package console

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type Color struct {
	bold          []string
	italic        []string
	underline     []string
	inverse       []string
	strikethrough []string
	white         []string
	grey          []string
	black         []string
	blue          []string
	cyan          []string
	green         []string
	magenta       []string
	red           []string
	yellow        []string
	whiteBG       []string
	greyBG        []string
	blackBG       []string
	blueBG        []string
	cyanBG        []string
	greenBG       []string
	magentaBG     []string
	redBG         []string
	yellowBG      []string
}

var colorFmt Color

func init() {
	var (
		bold          = []string{"\x1B[1m", "\x1B[22m"}
		italic        = []string{"\x1B[1m", "\x1B[22m"}
		underline     = []string{"\x1B[4m", "\x1B[24m"}
		inverse       = []string{"\x1B[7m", "\x1B[27m"}
		strikethrough = []string{"\x1B[9m", "\x1B[29m"}
		white         = []string{"\x1B[37m", "\x1B[39m"}
		grey          = []string{"\x1B[90m", "\x1B[39m"}
		black         = []string{"\x1B[30m", "\x1B[39m"}
		blue          = []string{"\x1B[34m", "\x1B[39m"}
		cyan          = []string{"\x1B[36m", "\x1B[39m"}
		green         = []string{"\x1B[32m", "\x1B[39m"}
		magenta       = []string{"\x1B[35m", "\x1B[39m"}
		red           = []string{"\x1B[31m", "\x1B[39m"}
		yellow        = []string{"\x1B[33m", "\x1B[39m"}

		whiteBG   = []string{"\x1B[47m", "\x1B[49m"}
		greyBG    = []string{"\x1B[49;5;8m", "\x1B[49m"}
		blackBG   = []string{"\x1B[40m", "\x1B[49m"}
		blueBG    = []string{"\x1B[44m", "\x1B[49m"}
		cyanBG    = []string{"\x1B[46m", "\x1B[49m"}
		greenBG   = []string{"\x1B[42m", "\x1B[49m"}
		magentaBG = []string{"\x1B[45m", "\x1B[49m"}
		redBG     = []string{"\x1B[41m", "\x1B[49m"}
		yellowBG  = []string{"\x1B[43m", "\x1B[49m"}
	)
	colorFmt = Color{
		bold,
		italic,
		underline,
		inverse,
		strikethrough,
		white,
		grey,
		black,
		blue,
		cyan,
		green,
		magenta,
		red,
		yellow,
		whiteBG,
		greyBG,
		blackBG,
		blueBG,
		cyanBG,
		greenBG,
		magentaBG,
		redBG,
		yellowBG,
	}
}

func stringAdd(sources string, dest ...string) string {
	var pack strings.Builder
	n := 0
	for i := 0; i < len(dest); i++ {
		n += len(dest[i])
	}
	pack.Grow(n)
	pack.WriteString(sources)
	for i := 0; i < len(dest); i++ {
		pack.WriteString(dest[i])
	}
	return pack.String()
}

func getFormatTimeStr() string {
	t := time.Now()
	nanoT := strconv.FormatInt(t.UnixNano(), 10)
	return t.Format("2006-01-02 15:04:05") + "." + nanoT[10:13]
}

func getStack() string {
	stack := string(debug.Stack())
	stackSlices := strings.Split(stack, "\n")
	document := strings.Split(stackSlices[8], "/")
	indexOnRunning := strings.Split(document[len(document)-1], " +")[0]
	return indexOnRunning
}

func Log(args ...interface{}) {
	stacker := getStack()
	fmt.Print(stringAdd(strings.Join(colorFmt.bold, getFormatTimeStr()), "  ", strings.Join(colorFmt.red, "|"), "   LOG ", strings.Join(colorFmt.red, "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF135     "))
	fmt.Print(args...)
	fmt.Print("\n")
}

func Debug(args ...interface{}) {
	stacker := getStack()
	fmt.Print(stringAdd(strings.Join(colorFmt.bold, getFormatTimeStr()), "  ", strings.Join(colorFmt.red, "|"), strings.Join(colorFmt.blue, " DEBUG "), strings.Join(colorFmt.red, "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF188     "))
	fmt.Print(args...)
	fmt.Print("\n")
}

func Info(args ...interface{}) {
	stacker := getStack()
	fmt.Print(stringAdd(strings.Join(colorFmt.bold, getFormatTimeStr()), "  ", strings.Join(colorFmt.red, "|"), strings.Join(colorFmt.cyan, "  INFO "), strings.Join(colorFmt.red, "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF1D9     "))
	fmt.Print(args...)
	fmt.Print("\n")
}

func Warn(args ...interface{}) {
	stacker := getStack()
	fmt.Print(stringAdd(strings.Join(colorFmt.bold, getFormatTimeStr()), "  ", strings.Join(colorFmt.red, "|"), strings.Join(colorFmt.yellow, "  WARN "), strings.Join(colorFmt.red, "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF21E     "))
	fmt.Print(args...)
	fmt.Print("\n")
}

func Error(args ...interface{}) {
	stacker := getStack()
	fmt.Print(stringAdd(strings.Join(colorFmt.bold, getFormatTimeStr()), "  ", strings.Join(colorFmt.red, "|"), strings.Join(colorFmt.red, " ERROR "), strings.Join(colorFmt.red, "|"), "  ["+stacker+"]", strings.Repeat(" ", 6-len(strings.Split(stacker, ":")[1]))+"\uF127     "))
	fmt.Print(args...)
	fmt.Print("\n")
}

func main() {
	Log("Hello World !")
	Info("Hello World !")
	Debug("Hello World !")
	Warn("Hello World !")
	Error("Hello World !")
}
