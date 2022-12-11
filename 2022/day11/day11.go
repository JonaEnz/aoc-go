package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items         []int64
	operation_op  string
	operation_val int
	test_div      int
	true_target   int
	false_target  int
	inspect_count int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	monkeys := make(map[int]Monkey, 0)
	monkeys_2 := make(map[int]Monkey, 0)

	read_buffer := make([]string, 0)
	for scanner.Scan() {
		read_buffer = append(read_buffer, scanner.Text())
		if scanner.Text() == "" {
			m, i, err := parse_monkey(read_buffer)
			if err != nil {
				panic(err)
			}
			monkeys[i] = m
			monkeys_2[i] = m
			read_buffer = make([]string, 0)
		}
	}
	m, i, err := parse_monkey(read_buffer)
	if err != nil {
		panic(err)
	}
	monkeys[i] = m
	monkeys_2[i] = m

	fmt.Println()
	for round := 0; round < 20; round++ {
		for i := 0; i < len(monkeys); i++ {
			m := monkeys[i]
			for len(m.items) > 0 {
				new_worry, target := inspect(&m, 3)
				monkeys[i] = m
				t := monkeys[target]
				t.items = append(monkeys[target].items, new_worry)
				monkeys[target] = t
			}
			monkeys[i] = m
		}
	}
	fmt.Printf("Part 1: %d\n", part1(monkeys))

	var lcm int64 = 1
	for _, m := range monkeys_2 {
		lcm = lcm * int64(m.test_div)
	}

	for round := 0; round < 10000; round++ {
		for i := 0; i < len(monkeys_2); i++ {
			m := monkeys_2[i]
			for len(m.items) > 0 {
				new_worry, target := inspect(&m, lcm)
				monkeys_2[i] = m
				t := monkeys_2[target]
				t.items = append(monkeys_2[target].items, new_worry)
				monkeys_2[target] = t
			}
			monkeys_2[i] = m
		}
	}
	fmt.Printf("Part 2: %d\n", part1(monkeys_2))
}

func parse_monkey(lines []string) (Monkey, int, error) {
	m := Monkey{}
	monkey_regex := regexp.MustCompile(`^Monkey (\d+):$`)
	items_regex := regexp.MustCompile(`^\s+Starting items: (.*)$`)
	operation_regex := regexp.MustCompile(`^\s+Operation: new = old (.+)$`)
	test_div_regex := regexp.MustCompile(`^\s+Test: divisible by (\d+)$`)
	true_regex := regexp.MustCompile(`^\s+If true: throw to monkey (\d+)$`)
	false_regex := regexp.MustCompile(`^\s+If false: throw to monkey (\d+)$`)

	i, _ := strconv.Atoi(monkey_regex.FindStringSubmatch(lines[0])[1])
	m.items = make([]int64, 0)
	for _, item := range strings.Split(items_regex.FindStringSubmatch(lines[1])[1], ", ") {
		it, _ := strconv.Atoi(item)
		m.items = append(m.items, int64(it))
	}
	op := operation_regex.FindStringSubmatch(lines[2])[1]
	m.operation_op = string(op[0])
	val, err := strconv.Atoi(op[2:])
	if err != nil {
		m.operation_val = -1
	} else {
		m.operation_val = val
	}

	m.test_div, _ = strconv.Atoi(test_div_regex.FindStringSubmatch(lines[3])[1])
	m.true_target, _ = strconv.Atoi(true_regex.FindStringSubmatch(lines[4])[1])
	m.false_target, _ = strconv.Atoi(false_regex.FindStringSubmatch(lines[5])[1])
	m.inspect_count = 0

	return m, i, nil
}

func inspect(monkey *Monkey, div int64) (int64, int) {
	var new_worry int64 = 0

	switch monkey.operation_op {
	case "+":
		new_worry = monkey.items[0] + int64(monkey.operation_val)
		break
	case "*":
		if monkey.operation_val == -1 {
			new_worry = monkey.items[0] * monkey.items[0]
		} else {
			new_worry = monkey.items[0] * int64(monkey.operation_val)
		}
		break
	default:
		panic("Unknown operation")
	}

	monkey.items = monkey.items[1:]

	if div == 3 {
		new_worry = new_worry / div
	} else {
		new_worry = new_worry % div
	}

	monkey.inspect_count++

	if new_worry%int64(monkey.test_div) == 0 {
		return new_worry, monkey.true_target
	} else {
		return new_worry, monkey.false_target
	}
}

func part1(monkeys map[int]Monkey) int {
	activity_list := make([]int, 0)
	for i := 0; i < len(monkeys); i++ {
		activity_list = append(activity_list, monkeys[i].inspect_count)
	}
	sort.Ints(activity_list)
	return activity_list[len(activity_list)-1] * activity_list[len(activity_list)-2]
}
