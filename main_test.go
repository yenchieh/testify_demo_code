package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testNumber struct {
	Number int
	Expect int
}

var tests = []testNumber{
	{
		Number: 5,
		Expect: 15,
	},
	{
		Number: 15,
		Expect: 120,
	},
	{
		Number: 30,
		Expect: 465,
	},
}

func TestAddItUp(t *testing.T) {
	for _, num := range tests {
		count := addItUp(num.Number)
		if count != num.Expect {
			t.Errorf("Input number %d, Expected %d, but got %d", num.Number, num.Expect, count)
		}
	}

	for _, num := range tests {
		count := addItUp(num.Number)

		// It will continue the test if it fails
		assert.NotEqual(t, count, num.Expect, fmt.Sprintf("Input number %d, Expected %d, but got %d", num.Number, num.Expect, count))

		// It will stop the test if it fails
		require.NotEqual(t, count, num.Expect, fmt.Sprintf("Input number %d, Expected %d, but got %d", num.Number, num.Expect, count))
	}
}

func TestLooping(t *testing.T) {
	number := looping()
	if number != 4999950001 {
		t.Errorf("Error number: %d", number)
	}
}

func BenchmarkLooping(b *testing.B) {
	_ = looping()
	//fmt.Printf("\nNumber: %d\n", number)
}

func BenchmarkAddItUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, num := range tests {
			count := addItUp(num.Number)
			fmt.Sprintf("%d", count)
		}
	}
}
