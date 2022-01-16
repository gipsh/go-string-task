package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	MATCH_STRING = "^\\w+$"
	MATCH_NUMBER = "^\\d+$"
	SPLIT_TOKEN  = "(\\w+-\\d+)"
)

// testValidity
// estimated: 1h
// used: 45min
func testValidity(input string) bool {

	tokens := strings.Split(input, "-")
	mString := regexp.MustCompile(MATCH_STRING)
	mNumber := regexp.MustCompile(MATCH_NUMBER)
	var m bool

	for i, v := range tokens {
		if i%2 == 0 {
			m = mString.MatchString(v)
		} else {
			m = mNumber.MatchString(v)
		}

		if !m {
			return false
		}
	}

	return true

}

// averageNumber
// estimated: 40min
// used: 30min
func averageNumber(input string) (ret float64, err error) {

	if !testValidity(input) {
		return 0, fmt.Errorf("invalid input")
	}

	re := regexp.MustCompile(SPLIT_TOKEN)

	split := re.FindAllString(input, -1)

	var sum float64

	for _, token := range split {
		p := strings.Split(token, "-")

		i, err := strconv.ParseInt(p[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("error")
		}

		sum = sum + float64(i)
	}

	ret = sum / float64(len(split))

	return ret, nil

}

// wholeStory
// estimated: 15min
// used: 15min
func wholeStory(input string) (string, error) {

	if !testValidity(input) {
		return "", fmt.Errorf("invalid input")
	}

	re := regexp.MustCompile(SPLIT_TOKEN)

	split := re.FindAllString(input, -1)

	var story string

	for i, token := range split {
		p := strings.Split(token, "-")
		if i == 0 {
			story = p[0]
		} else {
			story = story + " " + p[0]
		}
	}

	return story, nil

}

type StoryStats struct {
	Shortest           int
	Longest            int
	AverageLenght      float64
	AverageLenghtWords []string
}

//storyStats
// estimated: 1h
// used: 1h
func storyStats(input string) (*StoryStats, error) {

	if !testValidity(input) {
		return nil, fmt.Errorf("invalid input")
	}

	var (
		s, l     int     = math.MaxInt64, 0
		sum      float64 = 0
		words    []string
		lens     []int
		retwords []string
	)

	re := regexp.MustCompile(SPLIT_TOKEN)

	split := re.FindAllString(input, -1)

	for _, token := range split {
		p := strings.Split(token, "-")
		wl := len(p[0])
		if wl < s {
			s = wl
		}
		if wl > l {
			l = wl
		}
		sum = sum + float64(wl)
		words = append(words, p[0])
		lens = append(lens, wl)
	}

	avg := sum / float64(len(split))

	floor := math.Floor(avg)
	ceil := math.Ceil(avg)

	for i, w := range words {
		if lens[i] == int(floor) || lens[i] == int(ceil) {
			retwords = append(retwords, w)
		}
	}

	stats := StoryStats{
		Shortest:           s,
		Longest:            l,
		AverageLenght:      avg,
		AverageLenghtWords: retwords,
	}

	return &stats, nil

}

//generate
// estimated: 1h
// used: 40min
func generate(valid bool) string {

	rand.Seed(time.Now().UnixNano())

	var output string
	var token string
	tokens := rand.Intn(30)

	for i := tokens; i != 0; i-- {

		token = generateRandomToken(valid, rand.Intn(5)+1, rand.Intn(2)+1)

		if i == tokens {
			output = token
		} else {
			output = output + "-" + token
		}
	}

	// make sure its invalid
	if !valid {
		output = output + "-"
	}

	return output

}

func generateRandomToken(valid bool, sizeWord int, sizeNumber int) string {

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	var numbersRunes = []rune("0123456789")

	if !valid {
		allRunes := append(letterRunes, numbersRunes...)
		letterRunes = allRunes
		numbersRunes = allRunes
	}

	rand.Seed(time.Now().UnixNano())
	w := make([]rune, sizeWord)
	n := make([]rune, sizeNumber)
	for i := range w {
		w[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	for i := range n {
		n[i] = numbersRunes[rand.Intn(len(numbersRunes))]
	}

	return string(w) + "-" + string(n)

}
