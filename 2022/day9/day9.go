package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Game struct {
	head_x, head_y int
	tail_x, tail_y int
	direction      string
	movement       int
}

type Game2 struct {
	parts_x, parts_y []int
	direction        string
	movement         int
}

func main() {
	//Read input
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	game := Game{0, 0, 0, 0, "R", 0}
	game_2 := Game2{make([]int, 10), make([]int, 10), "R", 0}
	visited, visited2 := make(map[string]bool), make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		move_string := line[0]
		move_length, _ := strconv.Atoi(line[2:])

		game.direction = string(move_string)
		game.movement = move_length

		game_2.direction = string(move_string)
		game_2.movement = move_length

		stop, stop2 := false, false
		for !stop {
			game, stop = update(game)
			//print_game(game)
			visited[strconv.Itoa(game.tail_x)+","+strconv.Itoa(game.tail_y)] = true
		}

		for !stop2 {
			game_2, stop2 = update2(game_2)
			visited2[strconv.Itoa(game_2.parts_x[9])+","+strconv.Itoa(game_2.parts_y[9])] = true
		}

	}

	fmt.Println(len(visited))
	fmt.Println(len(visited2))

}

func update(game Game) (Game, bool) {
	if game.movement == 0 {
		return game, true
	}
	switch game.direction {
	case "R":
		game.head_x++
	case "L":
		game.head_x--
	case "U":
		game.head_y++
	case "D":
		game.head_y--
	}
	game.movement--
	game = move_tail(game)
	return game, game.movement <= 0
}

func move_tail(game Game) Game {
	x_dist := int(math.Abs(float64(game.tail_x - game.head_x)))
	y_dist := int(math.Abs(float64(game.tail_y - game.head_y)))
	if x_dist > 1 || y_dist > 1 {
		if x_dist > 0 {
			if game.tail_x < game.head_x {
				game.tail_x++
			} else {
				game.tail_x--
			}
		}
		if y_dist > 0 {
			if game.tail_y < game.head_y {
				game.tail_y++
			} else {
				game.tail_y--
			}
		}
	}
	return game
}

func update2(game Game2) (Game2, bool) {
	if game.movement == 0 {
		return game, true
	}
	switch game.direction {
	case "R":
		game.parts_x[0]++
	case "L":
		game.parts_x[0]--
	case "U":
		game.parts_y[0]++
	case "D":
		game.parts_y[0]--
	}
	game.movement--
	for i := 1; i < len(game.parts_x); i++ {
		game = move_tail2(game, i)
	}
	return game, game.movement <= 0
}

func move_tail2(game Game2, tail_number int) Game2 {
	// Compare part tail_number to tail_number - 1, number 0 is the head
	x_dist := int(math.Abs(float64(game.parts_x[tail_number] - game.parts_x[tail_number-1])))
	y_dist := int(math.Abs(float64(game.parts_y[tail_number] - game.parts_y[tail_number-1])))
	if x_dist > 1 || y_dist > 1 {
		if x_dist > 0 {
			if game.parts_x[tail_number] < game.parts_x[tail_number-1] {
				game.parts_x[tail_number]++
			} else {
				game.parts_x[tail_number]--
			}
		}
		if y_dist > 0 {
			if game.parts_y[tail_number] < game.parts_y[tail_number-1] {
				game.parts_y[tail_number]++
			} else {
				game.parts_y[tail_number]--
			}
		}
	}
	return game
}
