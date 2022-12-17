package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
)

type Game struct {
	blocks           *[][][]bool
	field            [][]bool
	jet_pattern      []bool
	block_x, block_y int
	highest_point    int
	block            int
	step             int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)
	input := make([]bool, 0)
	for scanner.Scan() {
		if scanner.Text() == "<" {
			input = append(input, true)
		} else {
			input = append(input, false)
		}
	}
	field := make([][]bool, 4)
	for i := range field {
		field[i] = make([]bool, 7)
	}

	blocks := make([][][]bool, 5)
	blocks[0] = [][]bool{{true, true, true, true}}
	blocks[1] = [][]bool{{false, true, false}, {true, true, true}, {false, true, false}}
	blocks[2] = [][]bool{{false, false, true}, {false, false, true}, {true, true, true}}
	blocks[3] = [][]bool{{true}, {true}, {true}, {true}}
	blocks[4] = [][]bool{{true, true}, {true, true}}
	game := Game{
		blocks:        &blocks,
		field:         field,
		jet_pattern:   input,
		block_x:       0,
		block_y:       0,
		highest_point: -1,
		block:         -1,
		step:          0,
	}
	game.NextBlock()
	part2hash := game.Hash()
	old_hash := game.Hash()
	fallenBlocks := 0
	diff_hashmap := make(map[uint64]int)
	next_hashmap := make(map[uint64]uint64)
	for fallenBlocks < 2022 {
		old_highest_point := game.highest_point
		if game.NextState() {
			fallenBlocks++
			gameHash := game.Hash()

			hashed_diff, ok := diff_hashmap[gameHash]
			hashed_next_hash, ok2 := next_hashmap[old_hash]
			if ok2 && hashed_next_hash != gameHash {
				panic("Hash function failed")
			}

			next_hashmap[old_hash] = gameHash
			old_hash = gameHash
			diff := game.highest_point - old_highest_point
			if ok && diff != hashed_diff {
				panic("Hash function failed")
			} else {
				diff_hashmap[gameHash] = diff
			}
		}
	}
	// Find loop size
	seen_at := make(map[uint64]int)
	seen_score := make(map[uint64]int)
	seen_hash := next_hashmap[part2hash]
	seen, loop_size, loop_diff := false, 0, 0
	for !seen {
		_, seen = seen_at[seen_hash]
		if seen {
			break
		}
		seen_at[seen_hash] = loop_size
		seen_score[seen_hash] = loop_diff
		loop_diff += diff_hashmap[seen_hash]
		loop_size++
		seen_hash = next_hashmap[seen_hash]
	}
	loop_size -= seen_at[seen_hash]
	loop_diff -= seen_score[seen_hash]

	high := -1
	high += loop_diff * (1000000000000 / loop_size)
	for i := 0; i < (1000000000000 % loop_size); i++ {
		new_hash := next_hashmap[part2hash]
		high += diff_hashmap[new_hash]
		part2hash = new_hash
	}

	fmt.Printf("Part 1: %d\n", game.highest_point+1)
	fmt.Printf("Part 2: %d\n", high+1)
}

func (game *Game) NextState() bool {
	//Apply jet burst (left/right)
	if game.jet_pattern[game.step%(len(game.jet_pattern))] {
		//left
		if game.CanMove(0, -1) {
			game.block_y--
		}
	} else {
		//right
		if game.CanMove(0, 1) {
			game.block_y++
		}
	}

	//Move down
	if game.CanMove(-1, 0) {
		game.block_x--
	} else {
		//Place block, move on to next
		for i := 0; i < len((*game.blocks)[game.block]); i++ {
			for j := 0; j < len((*game.blocks)[game.block][0]); j++ {
				if (*game.blocks)[game.block][i][j] {
					game.field[game.block_x-i][game.block_y+j] = true
					if game.block_x-i > game.highest_point {
						game.highest_point = game.block_x - i
					}
				}
			}
		}
		game.NextBlock()
		game.step++
		return true
	}
	game.step++
	return false
}

func (game *Game) CanMove(x, y int) bool {
	for i := 0; i < len((*game.blocks)[game.block]); i++ {
		for j := 0; j < len((*game.blocks)[game.block][0]); j++ {
			if game.block_x+x-i < 0 || game.block_x+x-i >= len(game.field) || game.block_y+y+j < 0 || game.block_y+y+j >= len(game.field[0]) {
				return false
			}
			if ((*game.blocks)[game.block][i][j]) && game.field[game.block_x+x-i][game.block_y+y+j] {
				return false
			}
		}
	}
	return true
}

func (game *Game) NextBlock() {
	game.block = (game.block + 1) % (len(*game.blocks))
	for len(game.field) < game.highest_point+len((*game.blocks)[game.block])+4 {
		game.field = append(game.field, make([]bool, 7))
	}
	for len(game.field) > game.highest_point+len((*game.blocks)[game.block])+4 {
		game.field = game.field[:len(game.field)-1]
	}
	game.block_x = len(game.field) - 1
	game.block_y = 2
}

func (game *Game) Draw() {
	for i := len(game.field) - 1; i >= 0 && i > len(game.field)-21; i-- {
		for j := 0; j < len(game.field[0]); j++ {
			if game.field[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (game *Game) Hash() uint64 {
	sha := sha256.New()
	h := ""
	for i := len(game.field) - 1; i > len(game.field)-21; i-- {
		if i <= 0 {
			h += "......."
			continue
		}
		for j := 0; j < len(game.field[0]); j++ {
			if game.field[i][j] {
				h += "#"
			} else {
				h += "."
			}
		}
	}
	h += fmt.Sprintf("%d%d", game.step%len(game.jet_pattern), game.block)
	sha.Write([]byte(h))
	res := binary.BigEndian.Uint64(sha.Sum(nil))
	return res
}
