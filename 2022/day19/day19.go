package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
)

type Blueprint struct {
	id           int
	recipe       [][3]int
	maxResources [4]int
}

type State struct {
	blueprint *Blueprint
	resources [4]int
	robots    [4]int
	time      int
}

func (s *State) Copy() State {
	resources := [4]int{}
	robots := [4]int{}
	copy(resources[:], s.resources[:])
	copy(robots[:], s.robots[:])
	return State{s.blueprint, resources, robots, s.time}
}

func (s *State) CanBuild(res int) bool {
	for i := 0; i < 3; i++ {
		if s.resources[i] < s.blueprint.recipe[res][i] {
			return false
		}
	}
	return true
}

func (s *State) RobotWork() {
	for i := 0; i < 4; i++ {
		s.resources[i] += s.robots[i]
	}
	s.time++
}

func (s *State) Build(res int) {
	if !s.CanBuild(res) {
		panic("Cannot build")
	}
	for i := 0; i < 3; i++ {
		s.resources[i] -= s.blueprint.recipe[res][i]
	}
	s.robots[res]++
}

func part1(blueprint *Blueprint) int {

	max := 0
	for j := 0; j < 4; j++ {
		state := State{
			blueprint: blueprint,
			resources: [4]int{0, 0, 0, 0},
			robots:    [4]int{1, 0, 0, 0},
			time:      0,
		}
		d := dfs(&state, j, 24, 0)
		if d > max {
			max = d
		}
	}
	return max
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	buffer := make([]string, 0)
	blueprints := make([]Blueprint, 0)
	for scanner.Scan() {
		if scanner.Text() != "" {
			buffer = append(buffer, scanner.Text())
		}
		if len(buffer) == 5 {
			blueprints = append(blueprints, parseBlueprint(buffer))
			buffer = make([]string, 0)
		}
	}
	totalQuality := 0
	var wg sync.WaitGroup
	wg.Add(len(blueprints))
	results := make([]int, len(blueprints))
	for i := 0; i < len(blueprints); i++ {
		go func(i int) {
			results[i] = part1(&blueprints[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < len(blueprints); i++ {
		totalQuality += results[i] * blueprints[i].id
	}

	fmt.Printf("Part 1: %d\n", totalQuality)

	// Part 2
	part2 := 1
	results = make([]int, 3)
	var wg2 sync.WaitGroup
	wg2.Add(3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			state := State{
				blueprint: &blueprints[i],
				resources: [4]int{0, 0, 0, 0},
				robots:    [4]int{1, 0, 0, 0},
				time:      0,
			}
			max := 0
			for j := 0; j < 4; j++ {
				d := dfs(&state, j, 32, 0)
				if d > max {
					max = d
				}
			}
			results[i] = max
			wg2.Done()
		}(i)
	}
	wg2.Wait()
	for i := 0; i < 3; i++ {
		part2 *= results[i]
	}

	fmt.Printf("Part 2: %d\n", part2)
}

func parseBlueprint(lines []string) Blueprint {
	robotReg := regexp.MustCompile(`^\s*Each \S+ robot costs (\d+) ore(?: and (\d+) clay)?(?: and (\d+) obsidian)?.*$`)

	id, _ := strconv.Atoi(lines[0][10 : len(lines[0])-1])
	recipe := make([][3]int, 0)

	ores := robotReg.FindStringSubmatch(lines[1])
	ore := [3]int{}
	ore[0], _ = strconv.Atoi(ores[1])
	ore[1], _ = strconv.Atoi(ores[2])
	ore[2], _ = strconv.Atoi(ores[3])
	recipe = append(recipe, ore)

	clay := [3]int{}
	clays := robotReg.FindStringSubmatch(lines[2])
	clay[0], _ = strconv.Atoi(clays[1])
	clay[1], _ = strconv.Atoi(clays[2])
	clay[2], _ = strconv.Atoi(clays[3])
	recipe = append(recipe, clay)

	obsidian := [3]int{}
	obsidians := robotReg.FindStringSubmatch(lines[3])
	obsidian[0], _ = strconv.Atoi(obsidians[1])
	obsidian[1], _ = strconv.Atoi(obsidians[2])
	obsidian[2], _ = strconv.Atoi(obsidians[3])
	recipe = append(recipe, obsidian)

	geode := [3]int{}
	geodes := robotReg.FindStringSubmatch(lines[4])
	geode[0], _ = strconv.Atoi(geodes[1])
	geode[1], _ = strconv.Atoi(geodes[2])
	geode[2], _ = strconv.Atoi(geodes[3])
	recipe = append(recipe, geode)

	maxResources := [4]int{}
	for i := 0; i < 3; i++ {
		for _, r := range recipe {
			if r[i] > maxResources[i] {
				maxResources[i] += r[i]
			}
		}
	}
	maxResources[3] = 9999
	return Blueprint{
		id,
		recipe,
		maxResources,
	}
}

func best_case(state *State, maxTime int) int {
	// Assume a new geode robot is built every turn
	return quad_sum(maxTime-state.time) + state.robots[3]*(maxTime-state.time+1) + state.resources[3]
}

func dfs(state *State, goal_type, maxTime, record int) int {
	if state.resources[3] > record {
		record = state.resources[3]
	}
	if goal_type == 0 && state.robots[0] >= state.blueprint.maxResources[0] {
		return record
	}
	if goal_type == 1 && state.robots[1] >= state.blueprint.maxResources[1] {
		return record
	}
	if goal_type == 2 && state.robots[2] >= state.blueprint.maxResources[2] {
		return record
	}
	if goal_type == 3 && state.robots[2] == 0 {
		return record
	}
	if best_case(state, maxTime) <= record {
		return record
	}

	for state.time < maxTime {
		if state.CanBuild(goal_type) {
			//newState := state
			state.RobotWork()
			state.Build(goal_type)
			for i := 0; i < 4; i++ {
				newState := state.Copy()
				res := dfs(&newState, i, maxTime, record)
				if res > record {
					record = res
				}
			}
			return record
		}
		state.RobotWork()
		if state.resources[3] > record {
			record = state.resources[3]
		}
	}
	return record
}

func quad_sum(n int) int {
	return n * (n - 1) / 2
}
