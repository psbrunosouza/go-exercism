package parsinglogfiles

import (
	"fmt"
	"regexp"
)

// IsValidLine check if the text matches with any of [TRC] [DBG] [INF] [WRN] [ERR] [FTL] codes
func IsValidLine(text string) bool {
	regex, _ := regexp.Compile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
	return regex.MatchString(text)
}

// SplitLogLine split the text based in your dividers and return a list of sections
func SplitLogLine(text string) []string {
	regex, _ := regexp.Compile(`<> | <(\* | = | ~ | -)+>`)
	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (quantity int) {
	regex, _ := regexp.Compile(`"(?i).*password.*"`)

	for _, value := range lines {
		println(value)

		if regex.MatchString(value) {
			quantity++
		}
	}

	return
}

func RemoveEndOfLineText(text string) string {
	regex, _ := regexp.Compile(`end-of-line[0-9]+`)
	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	regex, _ := regexp.Compile(`(\bUser\b.*)([A-Z]+[a-z]+[0-9]+)`)

	for index, value := range lines {
		words := regex.FindStringSubmatch(value)
		exists := regex.MatchString(value)
		if exists {
			lines[index] = fmt.Sprintf("[USR] %s %s", words[2], lines[index])
		}
	}

	return lines
}
