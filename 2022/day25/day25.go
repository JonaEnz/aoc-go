package main

import (
	"bufio"
	"fmt"
	"strconv"

	"math"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	converterMap := make(map[rune]int)
	converterMap['1'] = 1
	converterMap['2'] = 2
	converterMap['0'] = 0
	converterMap['-'] = -1
	converterMap['='] = -2

	scanner := bufio.NewScanner(f)
	n := 0
	for scanner.Scan() {
		j := 0
		for i := len(scanner.Text()) - 1; i >= 0; i-- {
			n += int(math.Pow(5, float64(j))) * converterMap[rune(scanner.Text()[i])]
			j++
		}
	}

	s := ""
	for n > 0 {
		v := ((n + 2) % 5) - 2
		switch v {
		case -2:
			s = "=" + s
		case -1:
			s = "-" + s
		default:
			s = strconv.Itoa(v) + s
		}
		n -= v
		n /= 5
	}
	fmt.Println(s)

}
