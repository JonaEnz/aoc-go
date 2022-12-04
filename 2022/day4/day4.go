package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type AssignmentPair struct {
	low_a, high_a, low_b, high_b int
}

func main() {
	assignment_regex := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	partA, partB := 0, 0
	for scanner.Scan() {
		assignment := scanner.Text()
		match := assignment_regex.FindStringSubmatch(assignment)
		low_a, _ := strconv.Atoi(match[1])
		high_a, _ := strconv.Atoi(match[2])
		low_b, _ := strconv.Atoi(match[3])
		high_b, _ := strconv.Atoi(match[4])
		assignment_pair := AssignmentPair{low_a, high_a, low_b, high_b}
		if assignment_pair.fully_contained() {
			partA++
		}
		if assignment_pair.overlap() {
			partB++
		}
	}
	fmt.Printf("Part 1: %d\n", partA)
	fmt.Printf("Part 2: %d\n", partB)
}

func (a AssignmentPair) fully_contained() bool {
	return (a.low_a >= a.low_b && a.high_a <= a.high_b) || (a.low_b >= a.low_a && a.high_b <= a.high_a)
}

func (a AssignmentPair) overlap() bool {
	return (a.low_a <= a.low_b && a.high_a >= a.low_b) || (a.low_b <= a.low_a && a.high_b >= a.low_a)
}
