package main

import (
	"bufio"
	"fmt"
	"os"
)

type Elf struct {
	x, y           int
	intent         int
	next_x, next_y int
}

func (e *Elf) String() string {
	return fmt.Sprintf("%d,%d", e.x, e.y)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	elfMap := make(map[string]Elf)
	i := 0
	for scanner.Scan() {
		for j, c := range scanner.Text() {
			if c == '#' {
				e := Elf{i, j, 0, 0, 0}
				elfMap[e.String()] = e
			}
		}
		i++
	}

	intent := 0
	//printMap(&elfMap)
	//fmt.Println("------")
	for step := 0; step < 9999; step++ {
		anyMove := simulationStep(&elfMap, intent)
		//elfMap = elfMap2
		// printMap(&elfMap)
		// fmt.Println(len(elfMap), intent)
		// fmt.Println("------", step)
		if step == 9 {
			minX, minY, maxX, maxY := getMinMaxXY(&elfMap)
			fmt.Printf("Part 1: %d\n", (maxX-minX+1)*(maxY-minY+1)-len(elfMap))
		}
		intent = (intent + 1) % 4
		if !anyMove {
			fmt.Printf("Part 2: %d\n", step+1)
			return
		}
	}
	panic("No solution found")

}

func getMinMaxXY(elfMap *map[string]Elf) (int, int, int, int) {
	minX, minY, maxX, maxY := 99999, 99999, -99999, -99999
	for _, elf := range *elfMap {
		if elf.x < minX {
			minX = elf.x
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}
	return minX, minY, maxX, maxY
}

func simulationStep(elfMap *map[string]Elf, start int) bool {
	nextMap := make(map[string]int)
	for e, elf := range *elfMap {
		elf.intent = -1
		if adjacent(elfMap, elf) == 0 {
			(*elfMap)[e] = elf
			continue
		}
		for i := start; i < start+4; i++ {
			check, next_x, next_y := checkIntent(elfMap, (*elfMap)[e], i%4)
			if check {
				_, ok := nextMap[fmt.Sprintf("%d,%d", next_x, next_y)]
				if !ok {
					nextMap[fmt.Sprintf("%d,%d", next_x, next_y)] = 0
				}
				nextMap[fmt.Sprintf("%d,%d", next_x, next_y)]++
				elf.intent = i
				elf.next_x = next_x
				elf.next_y = next_y
				break
			}
		}
		(*elfMap)[e] = elf
	}
	anyMove := false

	for e, elf := range *elfMap {
		if elf.intent == -1 {
			continue
		}
		if nextMap[fmt.Sprintf("%d,%d", elf.next_x, elf.next_y)] == 1 {
			anyMove = true
			elf.x = elf.next_x
			elf.y = elf.next_y
			if _, ok := (*elfMap)[elf.String()]; !ok {
				(*elfMap)[elf.String()] = elf
				delete((*elfMap), e)
			}

		}
	}
	return anyMove
}

func adjacent(elfMap *map[string]Elf, elf Elf) int {
	count := 0
	for i := elf.x - 1; i <= elf.x+1; i++ {
		for j := elf.y - 1; j <= elf.y+1; j++ {
			if i == elf.x && j == elf.y {
				continue
			}
			if _, ok := (*elfMap)[fmt.Sprintf("%d,%d", i, j)]; ok {
				count++
			}
		}
	}
	return count
}

func checkIntent(elfMap *map[string]Elf, elf Elf, intent int) (bool, int, int) {
	switch intent {
	case 0: // North
		for i := elf.y - 1; i <= elf.y+1; i++ {
			if _, ok := (*elfMap)[fmt.Sprintf("%d,%d", elf.x-1, i)]; ok {
				return false, elf.x, elf.y
			}
		}
		return true, elf.x - 1, elf.y
	case 1: // South
		for i := elf.y - 1; i <= elf.y+1; i++ {
			if _, ok := (*elfMap)[fmt.Sprintf("%d,%d", elf.x+1, i)]; ok {
				return false, elf.x, elf.y
			}
		}
		return true, elf.x + 1, elf.y

	case 3: // East
		for i := elf.x - 1; i <= elf.x+1; i++ {
			if _, ok := (*elfMap)[fmt.Sprintf("%d,%d", i, elf.y+1)]; ok {
				return false, elf.x, elf.y
			}
		}
		return true, elf.x, elf.y + 1

	case 2: // West
		for i := elf.x - 1; i <= elf.x+1; i++ {
			if _, ok := (*elfMap)[fmt.Sprintf("%d,%d", i, elf.y-1)]; ok {
				return false, elf.x, elf.y
			}
		}
		return true, elf.x, elf.y - 1
	}
	panic("Should not happen")
}
