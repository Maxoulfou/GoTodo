package utilities

import (
	"strconv"
	"strings"
)

func GetInput(table []string) string {

	return strings.Join(table," ")
}

func GetIntInput(table []string) int {
	justString, _ := strconv.ParseInt(table[0], 10, 64)

	return int(justString)
}

