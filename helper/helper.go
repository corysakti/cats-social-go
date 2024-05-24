package helper

import "strings"

func ConvertArrToString(urls []string) string {
	return strings.Join(urls, "|")
}
