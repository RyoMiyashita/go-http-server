package httprequest

import (
	"strings"
)

func getRequestString(buffer []byte) (string) {
	str := string(buffer)
	str = strings.TrimRight(str, "\n")
	return str
}
