package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type State struct {
	x, y   int
	facing int
}

func parseMove(line string) []int {
	var move []int
	b := []rune{}
	for _, c := range line {
		if c == 'L' || c == 'R' && len(b) > 0 {
			n, _ := strconv.Atoi(string(b))
			move = append(move, n)
		}
		if c == 'L' {
			move = append(move, -1)
			b = []rune{}
		} else if c == 'R' {
			move = append(move, -2)
			b = []rune{}
		} else {
			b = append(b, c)
		}
	}
	n, _ := strconv.Atoi(string(b))
	move = append(move, n)
	move = append(move, 0)

	return move
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	field := make([][]byte, 250)
	for i := range field {
		field[i] = make([]byte, 250)
	}
	j := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		for i, c := range scanner.Text() {
			if c == '#' {
				field[j][i] = 1
			} else if c == '.' {
				field[j][i] = 2
			} else {
				field[j][i] = 0
			}
		}
		j++
	}
	scanner.Scan()
	moveInput := parseMove(scanner.Text())
	x, y := 0, 0
	for field[y][x] != 2 {
		x++
	}
	state := &State{x, y, 0}
	moveMap := make(map[string](State))
	for y := 0; y < len(field); y++ {
		left := -2
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] != 0 && left == -2 {
				left = x - 1
			}
			if field[y][x] == 0 && left != -2 {
				moveMap[fmt.Sprintf("H%d,%d", y, left)] = State{x - 1, y, 2}
				moveMap[fmt.Sprintf("H%d,%d", y, x)] = State{left + 1, y, 0}
				break
			}
		}
	}

	for x := 0; x < len(field[0]); x++ {
		top := -2
		for y := 0; y < len(field); y++ {
			if field[y][x] != 0 && top == -2 {
				top = y - 1
			}
			if field[y][x] == 0 && top != -2 {
				moveMap[fmt.Sprintf("V%d,%d", top, x)] = State{x, y - 1, 3}
				moveMap[fmt.Sprintf("V%d,%d", y, x)] = State{x, top + 1, 1}
				break
			}
		}
	}

	for m := 0; m+1 < len(moveInput); m += 2 {
		state = move(moveInput[m], &field, state, &moveMap)
		state = turn(moveInput[m+1]+2, state)
	}
	fmt.Printf("Part 1: %d\n", 1000*(state.y+1)+(state.x+1)*4+state.facing)

	moveMap2 := getPart2Map()
	x = 0
	for field[0][x] != 2 {
		x++
	}
	state2 := &State{x, 0, 0}
	for m := 0; m+1 < len(moveInput); m += 2 {
		state2 = move(moveInput[m], &field, state2, moveMap2)
		state2 = turn(moveInput[m+1]+2, state2)
	}
	fmt.Printf("Part 2: %d\n", 1000*(state2.y+1)+(state2.x+1)*4+state2.facing)
}

func turn(dir int, state *State) *State {
	if dir == 1 {
		state.facing = (state.facing + 3) % 4 //L
	} else if dir == 0 {
		state.facing = (state.facing + 1) % 4 //R
	}
	return state
}

func move(len int, field *[][]byte, state *State, moveMap *map[string]State) *State {

	for i := 0; i < len; i++ {
		prev_x, prev_y, prev_facing := state.x, state.y, state.facing
		direction_x, direction_y, d := 0, 0, 'V'
		switch state.facing {
		case 0:
			direction_x = 1
			d = 'H'
		case 1:
			direction_y = 1
		case 2:
			direction_x = -1
			d = 'H'

		case 3:
			direction_y = -1
		}

		state.x += direction_x
		state.y += direction_y
		if state.y < 0 || state.x < 0 || (*field)[state.y][state.x] == 0 {
			newState, ok := (*moveMap)[fmt.Sprintf("%c%d,%d", d, state.y, state.x)]
			if !ok {
				panic("not found")
			}
			state = &newState
		}
		if (*field)[state.y][state.x] == 1 { // wall
			state.x = prev_x
			state.y = prev_y
			state.facing = prev_facing
			return state
		}
	}
	return state
}

func getPart2Map() *map[string]State {
	// hardcoded, only guaranteed to work for my input
	moveMap := make(map[string]State)
	for i := 0; i < 50; i++ {
		// Edge A
		moveMap[fmt.Sprintf("%c%d,%d", 'V', -1, 50+i)] = State{0, 150 + i, 0}
		moveMap[fmt.Sprintf("%c%d,%d", 'H', 150+i, -1)] = State{50 + i, 0, 1}
	}
	for i := 0; i < 50; i++ {
		// Edge B
		moveMap[fmt.Sprintf("%c%d,%d", 'H', i, 49)] = State{0, 149 - i, 0}
		moveMap[fmt.Sprintf("%c%d,%d", 'H', 100+i, -1)] = State{50, 49 - i, 0}
	}
	for i := 0; i < 50; i++ {
		// Edge C
		moveMap[fmt.Sprintf("%c%d,%d", 'H', 50+i, 49)] = State{i, 100, 1}
		moveMap[fmt.Sprintf("%c%d,%d", 'V', 99, i)] = State{50, 50 + i, 0}
	}
	for i := 0; i < 50; i++ {
		// Edge D
		moveMap[fmt.Sprintf("%c%d,%d", 'V', -1, 100+i)] = State{i, 199, 3}
		moveMap[fmt.Sprintf("%c%d,%d", 'V', 200, i)] = State{100 + i, 0, 1}
	}
	for i := 0; i < 50; i++ {
		// Edge E
		moveMap[fmt.Sprintf("%c%d,%d", 'H', i, 150)] = State{99, 149 - i, 2}
		moveMap[fmt.Sprintf("%c%d,%d", 'H', 100+i, 100)] = State{149, 49 - i, 2}
	}
	for i := 0; i < 50; i++ {
		// Edge F
		moveMap[fmt.Sprintf("%c%d,%d", 'V', 50, 100+i)] = State{99, 50 + i, 2}
		moveMap[fmt.Sprintf("%c%d,%d", 'H', 50+i, 100)] = State{100 + i, 49, 3}
	}
	for i := 0; i < 50; i++ {
		// Edge G
		moveMap[fmt.Sprintf("%c%d,%d", 'V', 150, 50+i)] = State{49, 150 + i, 2}
		moveMap[fmt.Sprintf("%c%d,%d", 'H', 150+i, 50)] = State{50 + i, 149, 3}
	}

	return &moveMap
}
