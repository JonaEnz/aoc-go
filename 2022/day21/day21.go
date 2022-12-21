package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Operator int

const (
	Plus     Operator = 0
	Minus    Operator = 1
	Multiply Operator = 2
	Divide   Operator = 3
	Number   Operator = 4
)

type Node struct {
	name  string
	val   int
	left  string
	right string
	op    Operator
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := make(map[string]Node)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		node := parseLine(scanner.Text())
		m[node.name] = node
	}
	fmt.Printf("Part 1: %d\n", GetNodeValue("root", &m))

	solution := 0
	start := "root"
	if contains(m["root"].left, "humn", &m) {
		start = m["root"].left
		solution = GetNodeValue(m["root"].right, &m)
	} else {
		start = m["root"].right
		solution = GetNodeValue(m["root"].left, &m)
	}

	containsMap := make(map[string]bool)
	for name := range m {
		if name == start {
			containsMap[name] = true
		}
		containsMap[name] = contains(name, "humn", &m)
	}

	fmt.Printf("Part 2: %d\n", RecSolutionSearch(start, "humn", solution, &m, &containsMap))
}

func RecSolutionSearch(name, searchName string, solution int, nodeMap *map[string]Node, containsMap *map[string]bool) int {
	// containsMap[name] = true if name contains searchName
	if name == searchName {
		return solution
	}
	if (*containsMap)[(*nodeMap)[name].right] && (*containsMap)[(*nodeMap)[name].left] {
		return -1
	}
	switch (*nodeMap)[name].op {
	case Plus:
		if !(*containsMap)[(*nodeMap)[name].left] {
			return RecSolutionSearch((*nodeMap)[name].right, searchName, solution-GetNodeValue((*nodeMap)[name].left, nodeMap), nodeMap, containsMap)
		} else {
			return RecSolutionSearch((*nodeMap)[name].left, searchName, solution-GetNodeValue((*nodeMap)[name].right, nodeMap), nodeMap, containsMap)
		}
	case Minus:
		if !(*containsMap)[(*nodeMap)[name].left] {
			return RecSolutionSearch((*nodeMap)[name].right, searchName, GetNodeValue((*nodeMap)[name].left, nodeMap)-solution, nodeMap, containsMap)
		} else {
			return RecSolutionSearch((*nodeMap)[name].left, searchName, GetNodeValue((*nodeMap)[name].right, nodeMap)+solution, nodeMap, containsMap)
		}
	case Multiply:
		if !(*containsMap)[(*nodeMap)[name].left] {
			return RecSolutionSearch((*nodeMap)[name].right, searchName, solution/GetNodeValue((*nodeMap)[name].left, nodeMap), nodeMap, containsMap)
		} else {
			return RecSolutionSearch((*nodeMap)[name].left, searchName, solution/GetNodeValue((*nodeMap)[name].right, nodeMap), nodeMap, containsMap)
		}
	case Divide:
		if !(*containsMap)[(*nodeMap)[name].left] {
			return RecSolutionSearch((*nodeMap)[name].right, searchName, GetNodeValue((*nodeMap)[name].left, nodeMap)/solution, nodeMap, containsMap)
		} else {
			return RecSolutionSearch((*nodeMap)[name].left, searchName, GetNodeValue((*nodeMap)[name].right, nodeMap)*solution, nodeMap, containsMap)
		}
	case Number:
		return -1
	}

	return -1
}

func GetNodeValue(name string, nodeMap *map[string]Node) int {
	if (*nodeMap)[name].op == Number {
		return (*nodeMap)[name].val
	}
	switch (*nodeMap)[name].op {
	case Plus:
		return GetNodeValue((*nodeMap)[name].left, nodeMap) + GetNodeValue((*nodeMap)[name].right, nodeMap)
	case Minus:
		return GetNodeValue((*nodeMap)[name].left, nodeMap) - GetNodeValue((*nodeMap)[name].right, nodeMap)
	case Multiply:
		return GetNodeValue((*nodeMap)[name].left, nodeMap) * GetNodeValue((*nodeMap)[name].right, nodeMap)
	case Divide:
		return GetNodeValue((*nodeMap)[name].left, nodeMap) / GetNodeValue((*nodeMap)[name].right, nodeMap)
	}
	panic("Unknown operator")
}

func contains(current, name string, nodeMap *map[string]Node) bool {
	if current == name {
		return true
	}

	if (*nodeMap)[current].op == Number {
		return false
	}

	if (*nodeMap)[current].left == name || (*nodeMap)[current].right == name {
		return true
	}
	return contains((*nodeMap)[current].left, name, nodeMap) || contains((*nodeMap)[current].right, name, nodeMap)
}

func parseLine(line string) Node {
	nodeReg := regexp.MustCompile(`(\w+): (\w+) ([+\/\-*]) (\w+)`)
	matches := nodeReg.FindStringSubmatch(line)
	if len(matches) == 0 {
		val, _ := strconv.Atoi(line[6:])
		return Node{
			line[0:4],
			val,
			"",
			"",
			Number,
		}
	}
	switch matches[3] {
	case "+":
		return Node{matches[1], -1, matches[2], matches[4], Plus}
	case "-":
		return Node{matches[1], -1, matches[2], matches[4], Minus}
	case "*":
		return Node{matches[1], -1, matches[2], matches[4], Multiply}
	case "/":
		return Node{matches[1], -1, matches[2], matches[4], Divide}
	}
	panic("Unknown operator")
}
