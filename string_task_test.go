package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_testValidity(t *testing.T) {

	input := "aasdasd-12312312-adasdasd-123123123"
	assert.True(t, testValidity(input))

	input = "aasdasd-12312312-adasdasd-123123123-"
	assert.False(t, testValidity(input))

	input = "-aasdasd-12312312-adasdasd-123123123-"
	assert.False(t, testValidity(input))

	input = "12312312-adasdasd"
	assert.False(t, testValidity(input))

	input = "d-12-adsd-13"
	assert.True(t, testValidity(input))

	input = ""
	assert.False(t, testValidity(input))

}

func Test_averagNumber(t *testing.T) {

	avg, err := averageNumber("aa-10-bb-30")

	assert.Nil(t, err)
	assert.Equal(t, avg, float64(20))

	avg, err = averageNumber("aa-40-bb-60-cc-24-dd-25")

	assert.Nil(t, err)
	assert.Equal(t, avg, float64(37.25))

}
