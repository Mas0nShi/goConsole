package console

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var colorFmt Color

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

func getFormatTimeStr() string {
	t := time.Now()
	nanoT := strconv.FormatInt(t.UnixNano(), 10)
	return t.Format("2006-01-02 15:04:05") + "." + nanoT[len(nanoT)-9:len(nanoT)-6]
}

func getStack() string {
	stack := string(debug.Stack())
	stackSlices := strings.Split(stack, "\n")

	//_start := strings.Index(stackSlices[5], ".") + 1
	//_last := strings.LastIndex(stackSlices[5], "()")
	//funcName := stackSlices[5][_start:_last]

	document := strings.Split(stackSlices[8], "/")
	indexOnRunning := strings.Split(document[len(document)-1], " +")[0]

	//fmt.Println(funcName, indexOnRunning)

	return indexOnRunning
}

func Log(args ...interface{}) {
	pJoin := strings.Join
	stackstr := getStack()
	fmt.Print(pJoin(colorFmt.bold, getFormatTimeStr()+"  "+pJoin(colorFmt.red, "|")+"   LOG "+pJoin(colorFmt.red, "|")+"  ["+stackstr+"]"+"   "+"\uF135 "+"    "))
	fmt.Print(args...)
	fmt.Print("\n")
}

func Debug(args ...interface{}) {
	pJoin := strings.Join
	stackstr := getStack()
	fmt.Print(pJoin(colorFmt.bold, getFormatTimeStr()+"  "+pJoin(colorFmt.red, "|")+pJoin(colorFmt.blue, " DEBUG ")+pJoin(colorFmt.red, "|")+"  ["+stackstr+"]"+"   "+"\uF188 "+"    "))
	fmt.Print(args...)
	fmt.Print("\n")
}

func Info(args ...interface{}) {
	pJoin := strings.Join
	stackstr := getStack()
	fmt.Print(pJoin(colorFmt.bold, getFormatTimeStr()+"  "+pJoin(colorFmt.red, "|")+pJoin(colorFmt.cyan, "  INFO ")+pJoin(colorFmt.red, "|")+"  ["+stackstr+"]"+"   "+"\uF1D9 "+"    "))
	fmt.Print(args...)
	fmt.Print("\n")
}

func Warn(args ...interface{}) {
	pJoin := strings.Join
	stackstr := getStack()
	fmt.Print(pJoin(colorFmt.bold, getFormatTimeStr()+"  "+pJoin(colorFmt.red, "|")+pJoin(colorFmt.yellow, "  WARN ")+pJoin(colorFmt.red, "|")+"  ["+stackstr+"]"+"   "+"\uF21E "+"    "))
	fmt.Print(args...)
	fmt.Print("\n")
}

func Error(args ...interface{}) {
	pJoin := strings.Join
	stackstr := getStack()
	fmt.Print(pJoin(colorFmt.bold, getFormatTimeStr()+"  "+pJoin(colorFmt.red, "|")+pJoin(colorFmt.red, " ERROR ")+pJoin(colorFmt.red, "|")+"  ["+stackstr+"]"+"   "+"\uF127 "+"    "))
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
