package kata_test

import (
	"testing"

	"github.com/jaedle-kata/golang-testify-template/kata"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert.Equal(t, true, kata.Kata())
	assert.NotEqual(t, false, kata.Kata())
}
