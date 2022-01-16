package main

import (
	"regexp"
	"strings"
)

const (
	MATCH_TOKEN = "(\\w+-\\d+)-*"
)

// testValidity
// estimated: 1h
// used: 45min
func testValidity(input string) bool {

	// border case not covered by regexp
	if strings.HasSuffix(input, "-") {
		return false
	}

	match, _ := regexp.MatchString(MATCH_TOKEN, input)

	return match

}
