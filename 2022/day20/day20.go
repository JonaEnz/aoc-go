package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Field struct {
	value  int
	index  int
	prev   *Field
	next   *Field
	length *int
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
	root := &Field{f.value, f.index, nil, nil, f.length}
	newCurrent := root
	for current != f {
		newCurrent.next = &Field{current.value, current.index, newCurrent, nil, f.length}
		newCurrent = newCurrent.next
		current = current.next
	}
	newCurrent.next = root
	root.prev = newCurrent
	return root
}

func (field *Field) mix(decryptionKey int) *Field {
	for i := 0; i < *field.length; i++ {
		for field.index != i {
			field = field.next
		}
		field.Rotate(field.value * decryptionKey)
	}
	return field
}

func (f *Field) Solution(len int) int {
	result := 0
	// Find value 0
	for f.value != 0 {
		f = f.next
	}
	c := f
	for j := 0; j < 3; j++ {
		for i := 0; i < 1000%(len-1); i++ {
			c = c.next
		}
		result += c.value
	}
	return result
}

func (f *Field) Rotate(steps int) *Field {
	steps = steps % (*f.length - 1) // A full rotation is len-1 steps

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
		return f.Rotate(steps - 1)
	}
	// Rotate left
	f.next.prev = f.prev
	f.prev.next = f.next
	f.next = f.prev
	f.prev = f.prev.prev
	f.prev.next = f
	f.next.prev = f
	return f.Rotate(steps + 1)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	length := 0
	root := &Field{
		value:  0,
		index:  0,
		prev:   nil,
		next:   nil,
		length: &length,
	}
	current := root
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		current.next = &Field{
			value:  i,
			index:  length,
			prev:   current,
			next:   nil,
			length: &length,
		}
		current = current.next
		length++
	}
	current.next = root.next
	root.next.prev = current
	part1 := current.next
	part2 := current.next.Clone()

	part1.mix(1)
	fmt.Println(part1.Solution(length))

	for i := 0; i < 10; i++ {
		part2.mix(811589153)
	}
	fmt.Println(part2.Solution(length) * 811589153)
}
