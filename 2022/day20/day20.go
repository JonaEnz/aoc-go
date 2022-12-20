package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Field struct {
	value int
	index int
	prev  *Field
	next  *Field
}

func (f *Field) Print() {
	current := f.next
	fmt.Print(f.value, " ")
	for current != f {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func (f *Field) Clone() *Field {
	current := f.next
	root := &Field{f.value, f.index, nil, nil}
	newCurrent := root
	for current != f {
		newCurrent.next = &Field{current.value, current.index, newCurrent, nil}
		newCurrent = newCurrent.next
		current = current.next
	}
	newCurrent.next = root
	root.prev = newCurrent
	return root
}

func (field *Field) mix(len int) *Field {
	for i := 0; i < len; i++ {
		for field.index != i {
			field = field.next
		}
		field.Rotate(field.value, len)
	}
	return field
}

func (f *Field) Part1() int {
	result := 0
	// Find value 0
	for f.value != 0 {
		f = f.next
	}
	c := f
	for j := 0; j < 3; j++ {
		for i := 0; i < 1000; i++ {
			c = c.next
		}
		result += c.value
	}
	return result
}

func (f *Field) Rotate(steps int, len int) *Field {
	steps = steps%len + steps/len // A full rotation is len-1 steps

	if steps == 0 {
		return f
	}
	if steps > 0 {
		// Rotate right
		f.next.prev = f.prev
		f.prev.next = f.next
		f.prev = f.next
		f.next = f.next.next
		f.prev.next = f
		f.next.prev = f
		return f.Rotate(steps-1, len)
	}
	// Rotate left
	f.next.prev = f.prev
	f.prev.next = f.next
	f.next = f.prev
	f.prev = f.prev.prev
	f.prev.next = f
	f.next.prev = f
	return f.Rotate(steps+1, len)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	root := &Field{0, 0, nil, nil}
	current := root
	j := 0
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		current.next = &Field{i, j, current, nil}
		current = current.next
		j++
	}
	current.next = root.next
	root.next.prev = current
	part1 := current.next
	part2 := current.next.Clone()

	part1.mix(j)
	fmt.Println(part1.Part1())
	part2.value *= 811589153
	for f := part2.next; f != part2; f = f.next {
		f.value *= 811589153
	}
	//part2.Print()
	for i := 0; i < 10; i++ {
		part2.mix(j)
		//part2.Print()
	}
	fmt.Println(part2.Part1())
}
