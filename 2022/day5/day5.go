package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	crate_regex := regexp.MustCompile(`\s*\[(.)\]`)
	slot_regex := regexp.MustCompile(`^(?:\s*\d+\s*)+$`)
	move_regex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	// Read input
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	slot_map_a, slot_map_b := make(map[int][]string), make(map[int][]string)

	for scanner.Scan() {
		line := scanner.Text()
		crate_matches := crate_regex.FindAllStringSubmatch(line, -1)
		slot_matches := slot_regex.FindAllString(line, -1)
		move_matches := move_regex.FindStringSubmatch(line)

		// Process input
		if len(crate_matches) > 0 {
			for i := 0; i < len(line); i += 4 {
				slot := i/4 + 1
				crate := line[i+1]
				if crate != ' ' {
					slot_map_a[slot] = append(slot_map_a[slot], string(crate))
				}
			}
			continue
		}
		if len(slot_matches) > 0 {
			fmt.Printf("Slot matches: %v\n", slot_matches)
			// Reverse all slots
			for i, slot := range slot_map_a {
				slot_map_a[i] = reverse(slot)
			}
			// Clone slot map
			for i, slot := range slot_map_a {
				slot_map_b[i] = make([]string, len(slot))
				copy(slot_map_b[i], slot)
			}
			continue
		}
		if len(move_matches) > 0 {
			count, _ := strconv.Atoi(move_matches[1])
			from, _ := strconv.Atoi(move_matches[2])
			to, _ := strconv.Atoi(move_matches[3])
			move(slot_map_a, from, to, count, true)
			move(slot_map_b, from, to, count, false)
			continue
		}
		fmt.Printf("No matches: %s\n", line)
	}

	fmt.Printf("Message: %s\n", message(slot_map_a))
	fmt.Printf("Message: %s\n", message(slot_map_b))

}

func move(slot_map map[int][]string, from int, to int, count int, rev bool) {
	fmt.Printf("Moving %d from %d to %d\n", count, from, to)
	// Move count crates from from to to
	moved_crates := slot_map[from][len(slot_map[from])-count:]
	if rev {
		moved_crates = reverse(moved_crates)
	}
	slot_map[to] = append(slot_map[to], moved_crates...)
	slot_map[from] = slot_map[from][:len(slot_map[from])-count]
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func message(slot_map map[int][]string) string {
	message := ""
	for i := 1; i <= len(slot_map); i++ {
		message += slot_map[i][len(slot_map[i])-1]
	}
	return message
}
