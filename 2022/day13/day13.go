package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type IntOrList struct {
	isInt bool
	val   int
	list  []IntOrList
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	index := 1
	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		first := scanner.Text()[1 : len(scanner.Text())-1]
		scanner.Scan()
		second := scanner.Text()[1 : len(scanner.Text())-1]
		scanner.Scan() // empty line, throw away
		//fmt.Printf("first: %s, second: %s\n", first, second)

		end := false
		for !end {
			comp := 0
			c := 0
			f, s := first, second
			for comp == 0 && f != "" && s != "" {
				left, f2 := getNext(f)
				right, s2 := getNext(s)
				comp = compare(left, right)
				fmt.Printf("Compared %v and %v, result: %d\n", left.String(), right.String(), comp)
				//fmt.Println(f2, s2)
				c++
				if c > 100 {
					break
				}
				f, s = f2, s2
			}
			if comp == 0 {
				if f == "" && s != "" {
					comp = -1
				} else if s == "" && f != "" {
					comp = 1
				}
			}
			if comp == -1 {
				//fmt.Println("Correct order. Left: " + first + ", Right: " + second + " (index: " + strconv.Itoa(index) + ")")
				sum += index
			}

			first, second = f, s
			end = f == "" || s == "" || comp != 0
		}
		//fmt.Println("first: " + first + ", second: " + second)
		index++
	}

	fmt.Println("Part 1: " + strconv.Itoa(sum))

	f.Close()
	f, err = os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	allPackets := []IntOrList{}

	div_packet := IntOrList{isInt: true, val: 2}
	div_packet = convertToList(div_packet)
	div_packet = IntOrList{isInt: false, list: []IntOrList{div_packet}}
	allPackets = append(allPackets, div_packet)

	div_packet2 := IntOrList{isInt: true, val: 6}
	div_packet2 = convertToList(div_packet2)
	div_packet2 = IntOrList{isInt: false, list: []IntOrList{div_packet2}}
	allPackets = append(allPackets, div_packet2)

	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			allPackets = append(allPackets, parseLine(line))
		}
	}
	bubbleSort(allPackets)

	index2, index6 := 0, 0
	for i, p := range allPackets {
		fmt.Println(p.String())
		if len(p.list) == 1 && len(p.list[0].list) == 1 && p.list[0].list[0].isInt && p.list[0].list[0].val == 2 {
			index2 = i + 1
		} else if len(p.list) == 1 && len(p.list[0].list) == 1 && p.list[0].list[0].isInt && p.list[0].list[0].val == 6 {
			index6 = i + 1
		}
	}
	fmt.Printf("Part 2: %d\n", index2*index6)

}

func parseLine(s string) IntOrList {
	result := IntOrList{}
	for len(s) > 0 {
		next, remaining := getNext(s)
		result.list = append(result.list, next)
		for len(remaining) > 0 && remaining[0] == ']' {
			remaining = remaining[1:]
		}
		s = remaining
	}

	return result
}

func getNext(s string) (IntOrList, string) {
	if s[0] == '[' {
		// list
		bracketDepth := 1
		for i := 1; i < len(s); i++ {
			if s[i] == '[' {
				bracketDepth++
			} else if s[i] == ']' {
				bracketDepth--
			}
			if bracketDepth == 0 && s[i] == ',' {
				return IntOrList{isInt: false, list: stringToList(s[1:i])}, s[i+1:]
			}
		}
		return IntOrList{isInt: false, list: stringToList(s[1:])}, ""

	} else {
		// int
		for i := 0; i < len(s); i++ {
			if s[i] == ',' || s[i] == ']' {
				j, _ := strconv.Atoi(s[:i])
				return IntOrList{isInt: true, val: j}, s[i+1:]
			}
		}
		j, _ := strconv.Atoi(s)
		return IntOrList{isInt: true, val: j}, ""
	}
}

func stringToList(s string) []IntOrList {
	result := []IntOrList{}
	for len(s) > 0 && s[0] != ']' {
		next, remaining := getNext(s)
		result = append(result, next)
		for len(remaining) > 0 && remaining[0] == ']' {
			remaining = remaining[1:]
		}
		s = remaining
	}
	return result
}

func convertToList(i IntOrList) IntOrList {
	if i.isInt {
		return IntOrList{isInt: false, list: []IntOrList{i}}
	} else {
		return i
	}
}

func compare(first, second IntOrList) int {
	//fmt.Printf("Comparing %v and %v\n", first.String(), second.String())
	if first.isInt && second.isInt {
		if first.val == second.val {
			return 0
		} else if first.val < second.val {
			return -1
		} else {
			return 1
		}
	} else if !first.isInt && !second.isInt {
		left, right := 0, 0
		for left < len(first.list) && right < len(second.list) {
			comp := compare(first.list[left], second.list[right])
			if comp != 0 {
				return comp
			}
			left++
			right++
		}
		if left == len(first.list) && right == len(second.list) {
			return 0
		} else if left == len(first.list) {
			return -1
		} else {
			return 1
		}
	} else {
		// one is int, one is list
		if first.isInt {
			first = convertToList(first)
		} else {
			second = convertToList(second)
		}
		return compare(first, second)
	}
}

func (iol IntOrList) String() string {
	if iol.isInt {
		return strconv.Itoa(iol.val)
	} else {
		result := "["
		for i, v := range iol.list {
			if i > 0 {
				result += ","
			}
			result += v.String()
		}
		result += "]"
		return result
	}
}

func bubbleSort(list []IntOrList) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			if compare(list[j], list[j+1]) == 1 {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
