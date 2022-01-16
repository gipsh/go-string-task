package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// estimated: 10min
// userd: 10min
func Test_testValidity(t *testing.T) {

	input := "aasdasd-12312312-adasdasd-123123123"
	assert.True(t, testValidity(input))

	input = "aasdasd-12312312-adasdasd-123123123-"
	assert.False(t, testValidity(input))

	input = "-aasdasd-12312312-adasdasd-123123123-"
	assert.False(t, testValidity(input))

	input = "12312312-adasdasd-123123123-asdasdasd"
	assert.False(t, testValidity(input))

	input = "d-12-adsd-13"
	assert.True(t, testValidity(input))

	input = ""
	assert.False(t, testValidity(input))

	input = "12-jj-12-adsd-13"
	assert.False(t, testValidity(input))

}

// estimated: 5min
// userd: 5min
func Test_averagNumber(t *testing.T) {

	avg, err := averageNumber("aa-10-bb-30")

	assert.Nil(t, err)
	assert.Equal(t, avg, float64(20))

	avg, err = averageNumber("aa-40-bb-60-cc-24-dd-25")

	assert.Nil(t, err)
	assert.Equal(t, avg, float64(37.25))

	_, err = averageNumber("40-bb-60-cc-24-dd-25")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid input")

}

// estimated: 5min
// userd: 5min
func Test_wholeStory(t *testing.T) {

	story, err := wholeStory("hello-10-world-30")

	assert.Nil(t, err)
	assert.Equal(t, story, "hello world")

	story, err = wholeStory("aa-40-bb-60-cc-24-dd-25")

	assert.Nil(t, err)
	assert.Equal(t, story, "aa bb cc dd")

}

// estimated: 10min
// userd: 10min
func Test_storyStats(t *testing.T) {

	stats, err := storyStats("hello-10-world-30")

	assert.Nil(t, err)
	assert.Equal(t, stats.AverageLenght, float64(5))
	assert.Equal(t, stats.Shortest, int(5))
	assert.Equal(t, stats.Longest, int(5))
	assert.Equal(t, stats.AverageLenghtWords, []string{"hello", "world"})

	stats, err = storyStats("hello-22-whats-23-your-33-name-43-again-433-please-43-tell-311-mee-11")

	assert.Nil(t, err)
	assert.Equal(t, stats.AverageLenght, float64(4.5))
	assert.Equal(t, stats.Shortest, int(3))
	assert.Equal(t, stats.Longest, int(6))
	assert.Equal(t, stats.AverageLenghtWords, []string{"hello", "whats", "your", "name", "again", "tell"})

}
