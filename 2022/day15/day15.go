package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Sensor struct {
	sensor_x int
	sensor_y int
	beacon_x int
	beacon_y int
	distance int
}

type Field struct {
	sensors map[string]Sensor
	beacons map[string]Sensor
}

func main() {
	field := Field{sensors: make(map[string]Sensor), beacons: make(map[string]Sensor)}
	sensorRegex := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matches := sensorRegex.FindStringSubmatch(line)
		if len(matches) == 5 {
			sensor_x, _ := strconv.Atoi(matches[1])
			sensor_y, _ := strconv.Atoi(matches[2])
			beacon_x, _ := strconv.Atoi(matches[3])
			beacon_y, _ := strconv.Atoi(matches[4])
			field.AddSensor(sensor_x, sensor_y, beacon_x, beacon_y)
		} else {
			panic("Invalid line: " + line)
		}
	}
	fmt.Printf("Part 1: %d\n", field.CheckRow(2000000))
	fmt.Printf("Part 2: %d\n", field.Part2(0, 4000000, 0, 4000000))
}

func (f *Field) AddSensor(sensor_x, sensor_y, beacon_x, beacon_y int) {
	sensor := Sensor{sensor_x, sensor_y, beacon_x, beacon_y, 0}
	sensor.distance = sensor.Distance()
	f.sensors[fmt.Sprintf("%d-%d", sensor_x, sensor_y)] = sensor
	f.beacons[fmt.Sprintf("%d-%d", beacon_x, beacon_y)] = sensor
}

func (s *Sensor) Distance() int {
	//Manhattan distance
	return Abs(s.sensor_x-s.beacon_x) + Abs(s.sensor_y-s.beacon_y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (f *Field) BeaconPossible(x, y int) bool {
	if _, ok := f.beacons[fmt.Sprintf("%d-%d", x, y)]; ok {
		return true
	}
	for _, sensor := range f.sensors {
		distance := Abs(sensor.sensor_x-x) + Abs(sensor.sensor_y-y)
		if distance <= sensor.distance {
			return false
		}
	}
	return true
}

func (f *Field) MinMaxY() (int, int) {
	min, max := 0, 0
	for _, sensor := range f.sensors {
		if sensor.sensor_y+sensor.distance > max {
			max = sensor.sensor_y + sensor.distance
		}
		if sensor.sensor_y-sensor.distance < min {
			min = sensor.sensor_y - sensor.distance
		}
	}
	return min, max
}

func (f *Field) CheckRow(y int) int {
	count := 0
	min, max := f.MinMaxY()
	for x := min; x < max; x++ {
		if !f.BeaconPossible(x, y) {
			count++
		}
	}
	return count
}

func (f *Field) Part2(x_min, x_max, y_min, y_max int) int {
	for _, sensor := range f.sensors {
		// Check around edges of manhattan distance
		x, y := -sensor.distance-1, 0
		for x <= sensor.distance+1 {
			check_x := sensor.sensor_x + x
			check_y := sensor.sensor_y + y
			if check_x >= x_min && check_x <= x_max && check_y >= y_min && check_y <= y_max {
				if f.BeaconPossible(check_x, check_y) {
					return check_x*4000000 + check_y
				}
			}
			check_y = sensor.sensor_y - y
			if check_x >= x_min && check_x <= x_max && check_y >= y_min && check_y <= y_max {
				if f.BeaconPossible(check_x, check_y) {
					return check_x*4000000 + check_y
				}
			}
			x++
			y = sensor.distance - Abs(x) + 1
		}
	}
	return 0
}
