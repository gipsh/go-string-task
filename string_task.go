package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	MATCH_TOKEN = "(\\w+-\\d+)-*"
	SPLIT_TOKEN = "(\\w+-\\d+)"
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

// averageNumber
// estimated: 40min
// used: 30min
func averageNumber(input string) (int64, error) {

	if !testValidity(input) {
		return 0, fmt.Errorf("invalid input")
	}

	re := regexp.MustCompile(SPLIT_TOKEN)

	split := re.FindAllString(input, -1)

	var avg int64

	for _, token := range split {
		p := strings.Split(token, "-")

		i, err := strconv.ParseInt(p[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("error")
		}

		avg = avg + i
	}

	ret := avg / int64(len(split))

	return ret, nil

}
