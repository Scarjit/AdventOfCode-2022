package day01

import (
	"os"
	"testing"
)

const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1Example(t *testing.T) {
	var day01 Day01
	if err := day01.prepare(input); err != nil {
		t.Errorf("prepare() error = %v", err)
		return
	}

	t.Logf("Solution: %d", day01.run1())
}

func TestPart1Full(t *testing.T) {
	// Load inputs/1.txt

	inputString, err := os.ReadFile("../../inputs/1.txt")
	if err != nil {
		t.Errorf("os.ReadFile() error = %v", err)
		return
	}

	var day01 Day01
	err = day01.prepare(string(inputString))
	if err != nil {
		t.Errorf("prepare() error = %v", err)
		return
	}

	t.Logf("Solution: %d", day01.run1())
}

func TestPart2Example(t *testing.T) {
	var day01 Day01
	if err := day01.prepare(input); err != nil {
		t.Errorf("prepare() error = %v", err)
		return
	}

	t.Logf("Solution: %d", day01.run2())
}

func TestPart2Full(t *testing.T) {
	// Load inputs/1.txt

	inputString, err := os.ReadFile("../../inputs/1.txt")
	if err != nil {
		t.Errorf("os.ReadFile() error = %v", err)
		return
	}

	var day01 Day01
	err = day01.prepare(string(inputString))
	if err != nil {
		t.Errorf("prepare() error = %v", err)
		return
	}

	t.Logf("Solution: %d", day01.run2())
}
