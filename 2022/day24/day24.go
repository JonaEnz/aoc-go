package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Blizzard struct {
	x, y      int
	direction rune
}

type State struct {
	time, x, y int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	blizzards := make([]Blizzard, 0)
	start_y, exit_y := 0, 0
	len_x, len_y := 0, 0
	x := 0
	for scanner.Scan() {
		if strings.Count(scanner.Text(), "#") > 2 {
			if x == 0 {
				start_y = strings.Index(scanner.Text(), ".")
			} else {
				exit_y = strings.Index(scanner.Text(), ".")
				len_x = x + 1
				len_y = len(scanner.Text())
			}
			x++
			continue
		}

		for y, c := range scanner.Text() {
			if c != '.' && c != '#' {
				blizzards = append(blizzards, Blizzard{x, y, c})
			}
		}
		x++
	}

	possiblePos := make([][][]bool, 1)
	possiblePos[0] = make([][]bool, len_x)
	for i := range possiblePos[0] {
		possiblePos[0][i] = make([]bool, len_y) // Init with false
	}
	possiblePos[0][0][start_y] = true
	stage := 0

	for true {
		possiblePos = append(possiblePos, make([][]bool, len_x))
		for i := range possiblePos[len(possiblePos)-1] {
			possiblePos[len(possiblePos)-1][i] = make([]bool, len_y) // Init with false
		}
		// Next blizzard positions
		blizzards = *NextStep(&blizzards, len_x, len_y)
		for x := 0; x < len_x; x++ {
			for y := 0; y < len_y; y++ {
				if possiblePos[len(possiblePos)-2][x][y] {
					state := State{len(possiblePos) - 1, x, y}
					for _, next_state := range state.NextStates(&blizzards, len_x, len_y) {
						possiblePos[len(possiblePos)-1][next_state.x][next_state.y] = true
					}
				}
			}
		}
		if stage%2 == 0 && possiblePos[len(possiblePos)-1][len_x-2][exit_y] {
			stage++
			possiblePos = append(possiblePos, make([][]bool, len_x))
			blizzards = *NextStep(&blizzards, len_x, len_y)
			for x := 0; x < len_x; x++ {
				possiblePos[len(possiblePos)-1][x] = make([]bool, len_y) // Init with false
			}
			possiblePos[len(possiblePos)-1][len_x-1][exit_y] = true
			fmt.Printf("Part %d: %d\n", (stage/2)+1, len(possiblePos)-1)
		} else if stage%2 == 1 && possiblePos[len(possiblePos)-1][1][start_y] {
			stage++
			possiblePos = append(possiblePos, make([][]bool, len_x))
			blizzards = *NextStep(&blizzards, len_x, len_y)
			for x := 0; x < len_x; x++ {
				possiblePos[len(possiblePos)-1][x] = make([]bool, len_y) // Init with false
			}
			possiblePos[len(possiblePos)-1][0][start_y] = true
		}
		if stage == 3 {
			break
		}
	}

}

func NextStep(blizzards *[]Blizzard, len_x, len_y int) *[]Blizzard {
	new_blizzards := make([]Blizzard, 0)
	for _, blizzard := range *blizzards {
		x, y := NextPosition(&blizzard, len_x, len_y)
		new_blizzards = append(new_blizzards, Blizzard{x, y, blizzard.direction})
	}
	return &new_blizzards
}

func NextPosition(blizzard *Blizzard, len_x, len_y int) (int, int) {
	switch blizzard.direction {
	case '>':
		if blizzard.y+1 >= len_y-1 {
			return blizzard.x, 1
		}
		return blizzard.x, blizzard.y + 1
	case '<':
		if blizzard.y-1 < 1 {
			return blizzard.x, len_y - 2
		}
		return blizzard.x, blizzard.y - 1
	case 'v':
		if blizzard.x+1 >= len_x-1 {
			return 1, blizzard.y
		}
		return blizzard.x + 1, blizzard.y
	case '^':
		if blizzard.x-1 < 1 {
			return len_x - 2, blizzard.y
		}
		return blizzard.x - 1, blizzard.y
	}
	return blizzard.x, blizzard.y
}

func (state *State) NextStates(blizzards *[]Blizzard, len_x, len_y int) []State {
	states := make([]State, 0)
	// Left
	if state.y-1 >= 1 && state.x > 0 && state.x < len_x-1 && !CheckBlizzards(blizzards, state.x, state.y-1) {
		states = append(states, State{state.time + 1, state.x, state.y - 1})
	}
	// Right
	if state.y+1 < len_y-1 && state.x > 0 && state.x < len_x-1 && !CheckBlizzards(blizzards, state.x, state.y+1) {
		states = append(states, State{state.time + 1, state.x, state.y + 1})
	}
	// Up
	if state.x-1 >= 1 && !CheckBlizzards(blizzards, state.x-1, state.y) {
		states = append(states, State{state.time + 1, state.x - 1, state.y})
	}
	// Down
	if (state.x+1 < len_x-1) && !CheckBlizzards(blizzards, state.x+1, state.y) {
		states = append(states, State{state.time + 1, state.x + 1, state.y})
	}
	// Stay
	if !CheckBlizzards(blizzards, state.x, state.y) {
		states = append(states, State{state.time + 1, state.x, state.y})
	}
	return states
}

func CheckBlizzards(blizzards *[]Blizzard, x, y int) bool {
	for _, blizzard := range *blizzards {
		if blizzard.x == x && blizzard.y == y {
			return true
		}
	}
	return false
}
