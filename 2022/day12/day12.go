package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyzES"
	hierarchy := make(map[string]int)
	for i := 0; i < len(alphabet); i++ {
		hierarchy[string(alphabet[i])] = i
	}

	sp := make([][]int, 0)
	sp_2 := make([][]int, 0)

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	field := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, make([]int, 0))
		sp = append(sp, make([]int, 0))
		sp_2 = append(sp_2, make([]int, 0))
		for i := 0; i < len(line); i++ {
			field[len(field)-1] = append(field[len(field)-1], hierarchy[string(line[i])])
			if string(line[i]) == "S" {
				sp[len(sp)-1] = append(sp[len(sp)-1], 0)
			} else {
				sp[len(sp)-1] = append(sp[len(sp)-1], -1)
			}
			if string(line[i]) == "a" {
				sp_2[len(sp_2)-1] = append(sp_2[len(sp_2)-1], 0)
			} else {
				sp_2[len(sp_2)-1] = append(sp_2[len(sp_2)-1], -1)
			}
		}
	}

	for i := 0; i < 10000; i++ {
		a := next(field, sp, i)
		b := next(field, sp_2, i)
		if a != 0 {
			fmt.Printf("Part 1: %d\n", a)
		}
		if b != 0 {
			fmt.Printf("Part 2: %d\n", b)
		}
		// for j := 0; j < len(sp); j++ {
		// 	for k := 0; k < len(sp[0]); k++ {
		// 		if sp[j][k] != -1 {
		// 			fmt.Print(sp[j][k])
		// 		} else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()
	}
}

func next(field [][]int, sp [][]int, round int) int {
	for i := 0; i < len(sp); i++ {
		for j := 0; j < len(sp[0]); j++ {
			if sp[i][j] == round {
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						if (math.Abs(float64(k))+math.Abs(float64(l)) < 2) && i+k >= 0 && j+l >= 0 && i+k < len(sp) && j+l < len(sp[0]) {
							if is_valid(i, j, i+k, j+l, field) && sp[i+k][j+l] == -1 {
								sp[i+k][j+l] = round + 1
								if field[i+k][j+l] == 26 {
									return round + 1
								}
							}
						}

					}
				}
			}
		}
	}
	return 0
}

func is_valid(x_now, y_now, x_next, y_next int, field [][]int) bool {
	if x_next < 0 || y_next < 0 || x_next >= len(field) || y_next >= len(field[0]) {
		return false
	}
	if x_now == x_next && y_now == y_next {
		return false
	}
	return field[x_next][y_next]-field[x_now][y_now] <= 1
}
