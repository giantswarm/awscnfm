package color

import "fmt"

var (
	Info    = Print("\033[1;34m%s\033[0m")
	Notice  = Print("\033[1;36m%s\033[0m")
	Warning = Print("\033[1;33m%s\033[0m")
	Error   = Print("\033[1;31m%s\033[0m")
	Debug   = Print("\033[0;36m%s\033[0m")

	Infof    = Printf("\033[1;34m%s\033[0m")
	Noticef  = Printf("\033[1;36m%s\033[0m")
	Warningf = Printf("\033[1;33m%s\033[0m")
	Errorf   = Printf("\033[1;31m%s\033[0m")
	Debugf   = Printf("\033[0;36m%s\033[0m")
)

func Print(colorString string) func(...interface{}) string {
	colorSprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return colorSprint
}

func Printf(colorString string) func(string, ...interface{}) string {
	colorSprintf := func(format string, args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprintf(format, args...))
	}
	return colorSprintf
}
