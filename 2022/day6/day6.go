package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()
	for scanner.Scan() {
		line := scanner.Text()
		// Process input
		fmt.Printf("Part1: %d\n", part1(line))
		fmt.Printf("Part2: %d\n", part2(line))
	}
}

func part1(line string) int {
	return part12(line, 4)
}

func part2(line string) int {
	return part12(line, 14)
}

func part12(line string, length int) int {
	// Find index where the next <length> characters are all different
	for i := 0; i < len(line)-length; i++ {
		// Check if the next <length> characters are all different
		char_map := make(map[byte]bool)
		for j := 0; j < length; j++ {
			_, ok := char_map[line[i+j]]
			if ok {
				break
			}
			char_map[line[i+j]] = true
			if j == length-1 {
				return i + length
			}
		}

	}
	return -1
}
