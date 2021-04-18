package feed

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  []int
		in2  []int
		out  int
	}{
		{
			name: "1",
			in1:  []int{1, 5, 9},
			in2:  []int{2, 4, 5, 6, 7, 9, 10, 12},
			out:  2,
		},
		{
			name: "2",
			in1:  []int{1, 5, 9},
			in2:  []int{1, 2, 4, 5, 6, 7, 9, 10, 12},
			out:  3,
		},
		{
			name: "3",
			in1:  []int{1, 5, 9},
			in2:  []int{1, 5, 9},
			out:  3,
		},
		{
			name: "4",
			in1:  []int{0, 1},
			in2:  []int{1, 5, 9},
			out:  1,
		},
		{
			name: "5",
			in1:  []int{},
			in2:  []int{},
			out:  0,
		},
		{
			name: "6",
			in1:  []int{0, 1, 2, 4, 5},
			in2:  []int{},
			out:  0,
		},
		{
			name: "7",
			in1:  []int{1, 2, 4, 5, 6, 7, 9, 10, 12},
			in2:  []int{1, 5, 9},
			out:  3,
		},
		{
			name: "8",
			in1:  []int{},
			in2:  []int{1, 5, 9},
			out:  0,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := calculateScore(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestIsBlacklisted(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  []int
		in2  []int
		out  bool
	}{
		{
			name: "1",
			in1:  []int{1, 5, 9},
			in2:  []int{2, 4, 5, 6, 7, 9, 10, 12},
			out:  true,
		},
		{
			name: "2",
			in1:  []int{1, 5, 9},
			in2:  []int{1, 2, 4, 6, 7, 9, 10, 12},
			out:  true,
		},
		{
			name: "3",
			in1:  []int{1, 5, 9},
			in2:  []int{1, 2, 4, 6, 7, 10, 12},
			out:  true,
		},
		{
			name: "4",
			in1:  []int{},
			in2:  []int{2, 4, 6, 7, 10, 12},
			out:  false,
		},
		{
			name: "5",
			in1:  []int{2, 4, 6, 7, 10, 12},
			in2:  []int{},
			out:  false,
		},
		{
			name: "8",
			in1:  []int{1, 5, 9},
			in2:  []int{2, 4, 6, 7, 10, 12},
			out:  false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := isBlacklisted(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}
