package day01

import (
	"sort"
	"strconv"
	"strings"
)

type Day01 struct {
	preparedSolution [][]int
}

func (s *Day01) prepare(in string) error {
	// Remove \r from the input
	in = strings.ReplaceAll(in, "\r", "")

	splits := strings.Split(in, "\n")
	s.preparedSolution = make([][]int, 0, len(splits))
	tempSplits := make([]int, 0, len(splits))
	for _, split := range splits {
		if len(split) == 0 {
			s.preparedSolution = append(s.preparedSolution, tempSplits)
			tempSplits = make([]int, 0, len(splits))
			continue
		}
		parseInt, err := strconv.ParseInt(split, 10, 64)
		if err != nil {
			return err
		}
		tempSplits = append(tempSplits, int(parseInt))
	}
	s.preparedSolution = append(s.preparedSolution, tempSplits)

	return nil
}

func (s *Day01) run1() int {
	calSums := make([]int, 0, len(s.preparedSolution))
	for _, cal := range s.preparedSolution {
		calN := 0
		for _, calX := range cal {
			calN += calX
		}
		calSums = append(calSums, calN)
	}
	// Sort calSums and output the first

	sort.SliceStable(
		calSums, func(i, j int) bool {
			return calSums[i] > calSums[j]
		})

	return calSums[0]
}

func (s *Day01) run2() int {
	calSums := make([]int, 0, len(s.preparedSolution))
	for _, cal := range s.preparedSolution {
		calN := 0
		for _, calX := range cal {
			calN += calX
		}
		calSums = append(calSums, calN)
	}
	// Sort calSums and output the first

	sort.SliceStable(
		calSums, func(i, j int) bool {
			return calSums[i] > calSums[j]
		})

	return calSums[0] + calSums[1] + calSums[2]
}
