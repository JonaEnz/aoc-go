package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Make 2d int array
	grid := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	line_nr := 0
	for scanner.Scan() {
		line := scanner.Text()

		grid_line := make([]int, len(line))

		for i := 0; i < len(line); i++ {
			nr, _ := strconv.Atoi(string(line[i]))
			grid_line[i] = nr
		}
		grid = append(grid, grid_line)
		line_nr++
	}
	sum := 0
	scene_record := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d", get_scenic_score(grid, i, j))
			if is_visible(grid, i, j) {
				sum++
			}
			if get_scenic_score(grid, i, j) > scene_record {
				scene_record = get_scenic_score(grid, i, j)
			}
		}
		fmt.Println()
	}
	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", scene_record)
}

func is_visible(grid [][]int, x int, y int) bool {

	return check_direction(grid, x, y, 1) || check_direction(grid, x, y, 2) || check_direction(grid, x, y, 3) || check_direction(grid, x, y, 4)

}

func check_direction(grid [][]int, x int, y int, direction int) bool {

	switch direction {
	case 1:
		// Check if there is a tree in the way from the top
		for i := 0; i < y; i++ {
			if grid[i][x] >= grid[y][x] {
				return false
			}
		}
		return true
	case 2:
		// Check if there is a tree in the way from the bottom
		for i := y + 1; i < len(grid); i++ {
			if grid[i][x] >= grid[y][x] {
				return false
			}
		}
		return true
	case 3:
		// Check if there is a tree in the way from the left
		for i := 0; i < x; i++ {
			if grid[y][i] >= grid[y][x] {
				return false
			}
		}
		return true
	case 4:
		// Check if there is a tree in the way from the right
		for i := x + 1; i < len(grid[y]); i++ {
			if grid[y][i] >= grid[y][x] {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func get_scenic_score(grid [][]int, x int, y int) int {
	scenic_score := 1

	if x == 0 || y == 0 || x == len(grid)-1 || y == len(grid[x])-1 {
		return 0
	}

	for i := x - 1; i >= 0; i-- {
		if i == 0 {
			scenic_score *= x
			break
		}
		if grid[i][y] >= grid[x][y] {
			scenic_score *= x - i
			break
		}
	}

	for i := x + 1; i <= len(grid); i++ {
		if i == len(grid) {
			scenic_score *= i - x - 1
			break
		}
		if grid[i][y] >= grid[x][y] {
			scenic_score *= i - x
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		if i == 0 {
			scenic_score *= y
			break
		}
		if grid[x][i] >= grid[x][y] {
			scenic_score *= y - i
			break
		}
	}

	for i := y + 1; i <= len(grid[x]); i++ {
		if i == len(grid[x]) {
			scenic_score *= i - y - 1
			break
		}
		if grid[x][i] >= grid[x][y] {
			scenic_score *= i - y
			break
		}
	}

	return scenic_score

}
