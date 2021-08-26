package service

import (
	"strings"
	"text/scanner"
)

/*SanityzeCommand provides a fast way to extract the command from a string*/
func SanitizeCommand(text string) string {
	var stringScan scanner.Scanner
	stringScan.Init(strings.NewReader(text))
	_ = stringScan.Scan()
	command := stringScan.TokenText()
	return command
}
