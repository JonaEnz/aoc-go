package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Split into two equal parts
		compartment_a := []rune(line[:len(line)/2])
		compartment_b := []rune(line[len(line)/2:])
		intersect := intersection(compartment_a, compartment_b)
		intersect_scanner := bufio.NewScanner(strings.NewReader(string(intersect)))
		intersect_scanner.Split(bufio.ScanRunes)
		for intersect_scanner.Scan() {
			sum += score(intersect_scanner.Text())
		}
	}
	fmt.Printf("Part 1: %d\n", sum)
	file.Close()

	file, err = os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)
	sum = 0
	rucksacks := make([]string, 0)
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
		if len(rucksacks) == 3 {
			inter_12 := intersection([]rune(rucksacks[0]), []rune(rucksacks[1]))
			inter_123 := intersection(inter_12, []rune(rucksacks[2]))
			sum += score(string(inter_123[0]))
			rucksacks = make([]string, 0)
		}
	}
	fmt.Printf("Part 2: %d\n", sum)

}

func intersection(a, b []rune) []rune {
	m := make(map[rune]bool)
	for _, v := range a {
		m[v] = true
	}
	var r []rune
	for _, v := range b {
		if m[v] && !contains(r, v) {
			r = append(r, v)
		}
	}
	return r
}

func contains(a []rune, b rune) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func score(s string) int {
	a := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, v := range a {
		if s == string(v) {
			return i + 1
		}
	}
	return -1
}
