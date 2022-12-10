package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	waiting := make(map[int](int))
	image := make(map[int](int))
	clock := 0
	x_register := 1
	sum, r := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		if line[:4] == "addx" {
			value, _ := strconv.Atoi(line[5:])
			waiting[clock+2] = value
			x_register, r = simulate(waiting, &image, clock, x_register)
			sum += r
			clock++
		}

		x_register, r = simulate(waiting, &image, clock, x_register)
		sum += r

		clock++
	}
	fmt.Println(sum)
	print_image(image)
}

func simulate(m map[int](int), img *map[int]int, clock int, reg int) (int, int) {
	v, ok := m[clock]
	if ok {
		reg += v
	}

	if clock%40 == reg || clock%40 == reg-1 || clock%40 == reg+1 {
		(*img)[clock+1] = reg
	}

	if (clock-20)%40 == 0 {
		fmt.Printf("%d: %d -> %d\n", clock, reg, reg*clock)
		return reg, reg * clock
	}

	return reg, 0
}

func print_image(img map[int]int) {
	for i := 1; i <= 240; i++ {
		_, ok := img[i]
		if ok {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if i%40 == 0 {
			fmt.Println()
		}

	}
}
