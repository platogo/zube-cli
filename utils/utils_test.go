package utils

import (
	"fmt"
	"testing"
)

func TestTruncateString(t *testing.T) {
	tests := []struct {
		exampleString string
		exampleMaxLen int
		want          string
	}{
		{"platogos", 3, "pla"},
		{"slotpark", 10, "slotpark"},
		{"早く", 1, "早"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%d,%s", tt.exampleString, tt.exampleMaxLen, tt.want)

		t.Run(testname, func(t *testing.T) {
			res := TruncateString(tt.exampleString, tt.exampleMaxLen)
			if res != tt.want {
				t.Errorf("got %s expected %s", res, tt.want)
			}
		})
	}
}

func TestSnakeCaseToTitleCase(t *testing.T) {
	example := "in_progress"
	want := "In Progress"

	res := SnakeCaseToTitleCase(example)

	if res != want {
		t.Errorf("expected %s got %s", want, res)
	}
}
