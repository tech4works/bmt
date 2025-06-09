package bmt

import (
	"fmt"
	"github.com/tech4works/converter"
	"regexp"
)

var formatRegex = regexp.MustCompile(`%[^%s]`)

func Sprint(a ...any) string {
	return fmt.Sprint(parse(a...)...)
}

func Sprintln(a ...any) string {
	return fmt.Sprintln(parse(a...)...)
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(replaceFormatSpecifiers(format), parse(a...)...)
}

func replaceFormatSpecifiers(format string) string {
	return formatRegex.ReplaceAllString(format, "%s")
}

func parse(a ...any) []any {
	s := make([]any, len(a))
	for i, v := range a {
		vs, err := converter.ToStringWithErr(v)
		if err != nil {
			vs = fmt.Sprint(v)
		}
		s[i] = vs
	}
	return s
}
