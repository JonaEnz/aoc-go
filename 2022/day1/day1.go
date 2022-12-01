package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, parse(line))
	}

	sum := 0
	top_three := []int{0, 0, 0}

	for _, line := range lines {
		if line == -1 {
			fmt.Printf("%d\n", sum)
			// Update the top three values
			for i := 0; i < 3; i++ {
				if sum > top_three[i] {
					top_three[i] = sum
					sort.Ints(top_three)
					break
				}
			}

			sum = 0
			continue
		}
		sum += line
	}

	fmt.Printf("Part 1: %d\n", top_three[0])
	fmt.Println("Part 2: ", top_three)
	fmt.Printf("Part 2: %d\n", top_three[2]+top_three[1]+top_three[0])
}

func parse(line string) int {
	if len(line) == 0 {
		return -1
	}
	a, _ := strconv.Atoi(line)
	return a
}
