package logger

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	Red     = color.New(color.FgRed).SprintFunc()
	Green   = color.New(color.FgGreen).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Blue    = color.New(color.FgBlue).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	Cyan    = color.New(color.FgCyan).SprintFunc()
	White   = color.New(color.FgWhite).SprintFunc()
	Black   = color.New(color.FgBlack).SprintFunc()
	DeadCol = color.New(color.FgHiBlack).SprintFunc()
	Bold    = color.New(color.Bold).SprintFunc()
)

func Info(format string, args ...interface{}) {
	prefix := Blue("[INFO]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, msg)
}

func Vulnerable(format string, args ...interface{}) {
	prefix := Green("[VULNERABLE]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, Bold(msg))
}

func NotVulnerable(format string, args ...interface{}) {
	prefix := Yellow("[NOT VULNERABLE]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, msg)
}

func Error(format string, args ...interface{}) {
	prefix := Red("[ERROR]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, msg)
}

func Debug(format string, args ...interface{}) {
	prefix := Magenta("[DEBUG]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, msg)
}

func Dead(format string, args ...interface{}) {
	prefix := DeadCol("[DEAD]")
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, msg)
}

func PrintBanner() {
	banner := `
    ______                                  
   / ____/____  ____ ____ ____ ____  ____ 
  / /    / __ \/ ___/ ___/ ___/ __ \/ __ \
 / /___ / /_/ / /  (__  ) /_ / / /_/ / / / 
 \____/ \____/_/  /____/\____\__,_/_/ /_/ 
                                           
    Corscan - CORS Vulnerability Scanner
	`
	color.Cyan(banner)
}
