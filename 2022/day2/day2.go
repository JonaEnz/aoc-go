package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sumA, sumB := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		a, b := parse(line)
		sumA += a
		sumB += b
	}
	fmt.Printf("Part 1: %d\n", sumA)
	fmt.Printf("Part 2: %d\n", sumB)

}

func parse(line string) (int, int) {
	input := strings.Fields(line)
	shape := 0
	outcomeMapA := map[string]int{}
	outcomeMapB := map[string]int{}
	own_score := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	win_lose_draw := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	switch input[0] {
	case "A": // Rock
		shape = 1
		outcomeMapA["X"] = 3 // Rock draws with Rock
		outcomeMapA["Y"] = 6 // Rock beats Scissors
		outcomeMapA["Z"] = 0 // Rock loses to Paper

		outcomeMapB["X"] = 3 //Lose -> Scissors
		outcomeMapB["Y"] = 1 //Draw -> Rock
		outcomeMapB["Z"] = 2 //Win -> Paper
		break
	case "B": // Paper
		shape = 2
		outcomeMapA["X"] = 0 // Paper loses to Rock
		outcomeMapA["Y"] = 3 // Paper draws with Paper
		outcomeMapA["Z"] = 6 // Paper beats Scissors

		outcomeMapB["X"] = 1 //Lose -> Rock
		outcomeMapB["Y"] = 2 //Draw -> Paper
		outcomeMapB["Z"] = 3 //Win -> Scissors
		break
	case "C": // Scissors
		shape = 3
		outcomeMapA["X"] = 6 // Scissors beats Rock
		outcomeMapA["Y"] = 0 // Scissors loses to Paper
		outcomeMapA["Z"] = 3 // Scissors draws with Scissors

		outcomeMapB["X"] = 2 //Lose -> Paper
		outcomeMapB["Y"] = 3 //Draw -> Scissors
		outcomeMapB["Z"] = 1 //Win -> Rock
		break
	default:
		panic("Invalid shape")
	}
	outcome := outcomeMapA[input[1]]
	_ = shape
	return scoreA(own_score[input[1]], outcome), scoreB(win_lose_draw[input[1]], outcomeMapB[input[1]])

}

func scoreA(own_score int, outcome int) int {
	return own_score + outcome
}

func scoreB(wld int, outcome int) int {
	return wld + outcome
}
