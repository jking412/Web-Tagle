package helper

import "strings"

func ParseActiveWeekdays(activeWeekdays string) []string {
	return strings.Split(activeWeekdays, ",")
}
