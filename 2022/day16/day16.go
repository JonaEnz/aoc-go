package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	name      string
	index     int
	flow_rate int
	connected []int
}

type State struct {
	open  []bool
	pos   []int
	time  int
	score int
	flow  int
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	valve_map := map[string]Valve{}
	neighbor_map := map[string][]string{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		valve, n := NewValve(line)
		neighbor_map[valve.name] = n
		valve.index = i
		valve_map[valve.name] = valve
		i++
	}

	valves := make([]Valve, len(valve_map))
	//i := 0
	for s := range valve_map {
		valves[valve_map[s].index] = valve_map[s]
	}
	start := 0
	for i := range valves {
		if valves[i].name == "AA" {
			start = i
		}
		for j := range neighbor_map[valves[i].name] {
			for k := range valves {
				if valves[k].name == neighbor_map[valves[i].name][j] {
					valves[i].connected = append(valves[i].connected, k)
				}
			}
		}
	}
	state := NewState(&valves, 1, 30, start)

	_, rec := state.Next(valves, getShortestPaths(valves), 0)
	fmt.Printf("Part 1: %d\n", rec)

}

func NewState(valves *[]Valve, pos, time, start int) State {
	pos_list := make([]int, pos)
	for i := range pos_list {
		pos_list[i] = start
	}

	return State{
		open:  make([]bool, len(*valves)),
		pos:   pos_list,
		time:  time,
		score: 0,
	}
}

func NewValve(line string) (Valve, []string) {
	regex := regexp.MustCompile(`^Valve (\S+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)
	matches := regex.FindStringSubmatch(line)
	if matches == nil {
		panic("Invalid input")
	}
	flow_rate, _ := strconv.Atoi(matches[2])
	return Valve{
		name:      matches[1],
		flow_rate: flow_rate,
	}, strings.Split(matches[3], ", ")
}

func getShortestPaths(valves []Valve) [][]int {
	// Use Floyd-Warshall algorithm
	dist := make([][]int, len(valves))
	for i := range dist {
		dist[i] = make([]int, len(valves))
		for j := range dist[i] {
			dist[i][j] = 9999999
		}
	}
	for i := range valves {
		for j := range valves[i].connected {
			dist[i][valves[i].connected[j]] = 1
		}
	}
	for k := range valves {
		for i := range valves {
			for j := range valves {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	for i := range dist {
		dist[i][i] = 0
	}

	return dist
}

func (state State) Copy() State {
	s := State{
		open:  make([]bool, len(state.open)),
		pos:   make([]int, len(state.pos)),
		time:  state.time,
		score: state.score,
		flow:  state.flow,
	}
	copy(s.open, state.open)
	copy(s.pos, state.pos)
	return s
}

func (state State) Next(valves []Valve, sp [][]int, record int) ([]State, int) {
	if state.score > record {
		record = state.score
	}

	if state.time <= 0 {
		return []State{}, state.score
	}

	states := []State{}

	nothing_to_open := true
	for i := range state.open {
		if !state.open[i] && valves[i].flow_rate > 0 && state.time >= sp[state.pos[0]][i]+1 {
			nothing_to_open = false
			s := state.Copy()
			dist := sp[state.pos[0]][i] + 1
			s.pos[0] = i
			s.open[i] = true
			s.time -= dist
			s.score += s.flow * dist
			s.flow += valves[i].flow_rate
			if s.score > record && s.time >= 0 {
				record = s.score
			}
			states = append(states, s)
		}
	}

	if nothing_to_open {
		state.score += state.flow * state.time
		return []State{}, state.score
	}

	// Get next states
	new_states := []State{}
	for len(states) > 0 {
		for i := range states {
			s, rec := states[i].Next(valves, sp, record)
			if rec > record {
				record = rec
			}
			new_states = append(new_states, s...)
		}
		states = make([]State, len(new_states))
		copy(states, new_states)
		new_states = []State{}
	}

	return new_states, record
}
