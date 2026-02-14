package main

import (
	"errors"
	"fmt"
	"strings"
)

func strtoBool(s string) (bool, error) {
	s = strings.ToLower(s)
	switch s {
	case "true", "yes", "1":

		return true, nil
	case "false", "no", "0":
		return false, nil
	default:
		return false, errors.New("error")

	}
}
func main() {

	b, _ := strtoBool("11")
	fmt.Println(b)
}
