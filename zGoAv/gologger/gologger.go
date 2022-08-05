package gologger

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"os"
	"strings"
	"sync"
)

type Level int

const (
	Null Level = iota
	Fatal
	Silent
	Label
	Misc
	Error
	Warning
	Info
	Debug
	Verbose
)

var (
	UseColors = true

	MaxLevel = Info

	labels = map[Level]string{
		Warning: "Warning",
		Error:   "Error",
		Label:   "WRN",
		Fatal:   "Fatal",
		Debug:   "DEBUG",
		Info:    "INFO",
	}
	mutex  = &sync.Mutex{}
	output = colorable.NewColorableStdout()
)

var stringBuilderpool = &sync.Pool{New: func() interface{} {
	return new(strings.Builder)
}}

func wrap(label string, level Level) string {
	if !UseColors {
		return label
	}
	switch level {
	case Silent:
		return label
	case Info, Verbose:
		return aurora.Blue(label).String()
	case Fatal:
		return aurora.Bold(aurora.Red(label)).String()
	case Error:
		return aurora.Red(label).String()
	case Debug:
		return aurora.Magenta(label).String()
	case Warning, Label:
		return aurora.Yellow(label).String()
	default:
		return label

	}
}
func getLabel(level Level, label string, sb *strings.Builder) {
	switch level {
	case Silent, Misc:
		return
	case Error, Fatal, Info, Warning, Debug, Label:
		sb.WriteString("[")
		sb.WriteString(wrap(labels[level], level))
		sb.WriteString("]")
		sb.WriteString(" ")
		return
	case Verbose:
		sb.WriteString("[")
		sb.WriteString(wrap(label, level))
		sb.WriteString("]")
		sb.WriteString(" ")
		return
	default:
		return
	}
}
func log(level Level, lable string, format string, args ...interface{}) {
	if level == Null {
		return
	}
	if level <= MaxLevel {
		sb := stringBuilderpool.Get().(*strings.Builder)
		getLabel(level, lable, sb)
		message := fmt.Sprintf(format, args...)
		sb.WriteString(message)
		mutex.Unlock()
		sb.Reset()
		stringBuilderpool.Put(sb)
	}
}
func Infof(format string, args ...interface{}) {
	log(Info, "", format, args...)
}
func Warningf(format string, args ...interface{}) {
	log(Warning, "", format, args...)
}
func Errorf(format string, args ...interface{}) {
	log(Error, "", format, args)
}
func Debuf(format string, args ...interface{}) {
	log(Debug, "", format, args)
}
func Verbosef(format string, args ...interface{}) {
	log(Verbose, "", format, args)
}
func Silentf(format string, args ...interface{}) {
	log(Silent, "", format, args)
}
func Fatalf(format string, args ...interface{}) {
	log(Fatal, "", format, args)
	os.Exit(1)
}
func Printf(format string, args ...interface{}) {
	log(Misc, "", format, args)
}
func Labelf(format string, args ...interface{}) {
	log(Label, "", format, args)
}
