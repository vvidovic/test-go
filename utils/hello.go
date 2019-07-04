package utils

import (
	"strings"
)

func Hello(name string) string {
	if strings.ContainsAny(name, "ABCDEFGHIJKLMNOPRQSTUVWXYZ") {
		return "Hello big " + name + "."
	} else {
		return "Hello " + name + "."
	}
}
