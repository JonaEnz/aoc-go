package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z int
}

type Queue struct {
	items []Cube
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	cubes := make([]Cube, 0)
	for scanner.Scan() {
		xyz := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		cubes = append(cubes, Cube{x, y, z})
	}
	fmt.Printf("Part 1: %d\n", part1(cubes))

	edges := minMax(cubes)

	queue := NewQueue()
	queue.enqueue(Cube{0, 0, 0})
	visited := make(map[Cube]bool)
	visited[Cube{0, 0, 0}] = true
	for _, cube := range cubes {
		visited[cube] = true
	}
	for !queue.isEmpty() {
		nextQueue := NewQueue()
		for !queue.isEmpty() {
			current := queue.dequeue()
			for _, neighbor := range current.neighbors() {
				if neighbor.x >= edges[0] && neighbor.x <= edges[1] && neighbor.y >= edges[2] && neighbor.y <= edges[3] && neighbor.z >= edges[4] && neighbor.z <= edges[5] && !visited[neighbor] {
					visited[neighbor] = true
					nextQueue.enqueue(neighbor)
				}
			}
		}
		queue = nextQueue
	}
	innerCubes := make([]Cube, 0)
	for x := edges[0]; x <= edges[1]; x++ {
		for y := edges[2]; y <= edges[3]; y++ {
			for z := edges[4]; z <= edges[5]; z++ {
				if !visited[Cube{x, y, z}] {
					innerCubes = append(innerCubes, Cube{x, y, z})
				}
			}
		}
	}

	innerNeighbors := 0
	for i := 0; i < len(innerCubes); i++ {
		for j := 0; j < len(innerCubes); j++ {
			if i != j && innerCubes[i].isNeighbor(innerCubes[j]) {
				innerNeighbors += 1
			}
		}
	}

	fmt.Printf("Part 2: %d\n", part1(cubes)-6*len(innerCubes)+innerNeighbors)
}

func (c1 *Cube) isNeighbor(c2 Cube) bool {
	neighbors := [][]int{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}}
	for _, n := range neighbors {
		if c1.x+n[0] == c2.x && c1.y+n[1] == c2.y && c1.z+n[2] == c2.z {
			return true
		}
	}

	return false
}

func part1(cubes []Cube) int {
	sides := 6 * len(cubes)

	for i := 0; i < len(cubes); i++ {
		for j := 0; j < len(cubes); j++ {
			if i != j && cubes[i].isNeighbor(cubes[j]) {
				sides -= 1
			}
		}
	}
	return sides
}

func NewQueue() Queue {
	return Queue{items: make([]Cube, 0)}
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(c Cube) {
	q.items = append(q.items, c)
}

func (q *Queue) dequeue() Cube {
	if len(q.items) == 0 {
		panic("Queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (c *Cube) neighbors() []Cube {
	neighbors := make([]Cube, 0)
	neighbors = append(neighbors, Cube{c.x - 1, c.y, c.z})
	neighbors = append(neighbors, Cube{c.x + 1, c.y, c.z})
	neighbors = append(neighbors, Cube{c.x, c.y - 1, c.z})
	neighbors = append(neighbors, Cube{c.x, c.y + 1, c.z})
	neighbors = append(neighbors, Cube{c.x, c.y, c.z - 1})
	neighbors = append(neighbors, Cube{c.x, c.y, c.z + 1})
	return neighbors
}

func minMax(cubes []Cube) []int {
	minX, maxX, minY, maxY, minZ, maxZ := 0, 0, 0, 0, 0, 0
	for i := 0; i < len(cubes); i++ {
		if cubes[i].x < minX {
			minX = cubes[i].x
		}
		if cubes[i].x > maxX {
			maxX = cubes[i].x
		}
		if cubes[i].y < minY {
			minY = cubes[i].y
		}
		if cubes[i].y > maxY {
			maxY = cubes[i].y
		}
		if cubes[i].z < minZ {
			minZ = cubes[i].z
		}
		if cubes[i].z > maxZ {
			maxZ = cubes[i].z
		}
	}
	return []int{minX - 1, maxX + 1, minY - 1, maxY + 1, minZ - 1, maxZ + 1}
}
