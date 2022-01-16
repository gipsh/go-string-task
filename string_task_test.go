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

}
