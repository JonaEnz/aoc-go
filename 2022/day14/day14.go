package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	field := make([][]bool, 1000)
	for i := 0; i < len(field); i++ {
		field[i] = make([]bool, 1000)
	}

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	highest_y := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " -> ")
		for i := 0; i < len(split)-1; i++ {
			x_start, y_start := parseXY(split[i])
			x_end, y_end := parseXY(split[i+1])
			drawLine(field, x_start, y_start, x_end, y_end)
			if y_end > highest_y {
				highest_y = y_end
			}
			if y_start > highest_y {
				highest_y = y_start
			}
		}
	}

	// Clone field into field2
	field2 := make([][]bool, 1000)
	for i := 0; i < len(field2); i++ {
		field2[i] = make([]bool, 1000)
	}
	for x := 0; x < len(field); x++ {
		for y := 0; y < len(field[x]); y++ {
			field2[x][y] = field[x][y]
		}
	}
	// Add line
	drawLine(field2, 0, highest_y+2, 999, highest_y+2)

	//fmt.Println(field)
	counter, counter2 := 0, 0
	for dropSand(field, 500, 0) {
		counter++
	}
	for dropSand(field2, 500, 0) {
		counter2++
	}
	fmt.Println(counter, counter2)

}

func drawLine(field [][]bool, x_start, y_start, x_end, y_end int) {
	if x_start > x_end {
		x_start, x_end = x_end, x_start
	}
	if y_start > y_end {
		y_start, y_end = y_end, y_start
	}

	if x_start == x_end {
		// vertical line
		for y := y_start; y <= y_end; y++ {
			field[x_start][y] = true
		}
	} else {
		// horizontal line
		for x := x_start; x <= x_end; x++ {
			field[x][y_start] = true
		}
	}
}

func parseXY(s string) (int, int) {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return x, y
}

func dropSand(field [][]bool, x, y int) bool {
	if y >= 900 || x < 0 || x >= 1000 {
		return false
	}

	if field[500][0] {
		return false
	}

	// Test down
	if !field[x][y+1] {
		return dropSand(field, x, y+1)
	}
	// Test left
	if x > 0 && !field[x-1][y+1] {
		return dropSand(field, x-1, y+1)
	}
	// Test right
	if x < 999 && !field[x+1][y+1] {
		return dropSand(field, x+1, y+1)
	}
	// Place sand
	field[x][y] = true
	return true
}
