package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var structure = map[int][]rune{
	0: []rune{'<', '>', '[', ']', '{', '}', '~', '€', '\\', '^'},
	1: []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'},
	2: []rune{'a', 'b', 'c', 'A', 'B', 'C', ' ', '=', '@', '|'},
	3: []rune{'d', 'e', 'f', 'D', 'E', 'F', 'ä', 'ö', 'ü', 'ß'},
	4: []rune{'g', 'h', 'i', 'G', 'H', 'I', 'Ä', 'Ö', 'Ü', 'ẞ'},
	5: []rune{'j', 'k', 'l', 'J', 'K', 'L', ',', ';', '.', ':'},
	6: []rune{'m', 'n', 'o', 'M', 'N', 'O', '!', '?', '/', '*'},
	7: []rune{'p', 'q', 'r', 's', 'P', 'Q', 'R', 'S', '(', ')'},
	8: []rune{'t', 'u', 'v', 'T', 'U', 'V', '"', '%', '&', '#'},
	9: []rune{'w', 'x', 'y', 'z', 'W', 'X', 'Y', 'Z', '+', '-'},
}

func encode(s string, w int) string {
	res := ""
	for _, r := range s {
		for k, v := range structure {
			for i, c := range v {
				if r == c {
					res += fmt.Sprintf("%d%d", k, i)
				}
			}
		}
	}

	if w > 0 {
		var sb strings.Builder
		for i, r := range res {
			if i > 0 && i%w == 0 {
				sb.WriteString("\r\n")
			}
			sb.WriteRune(r)
		}
		res = sb.String()
	}

	return res
}

func decode(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "")
	res := ""
	for i := 0; i < len(s); i += 2 {
		k, _ := strconv.Atoi(string(s[i]))
		v, _ := strconv.Atoi(string(s[i+1]))
		res += string(structure[k][v])
	}
	return res
}

func usage() {
	fmt.Printf("Usage: %s [-d] [-w line width]\n", os.Args[0])
	fmt.Println("  -d\tdecode")
	fmt.Println("  -w\tset the line width for encoding")
}

func main() {
	decodePtr := flag.Bool("d", false, "decode")
	widthPtr := flag.Int("w", 0, "line width")
	flag.Parse()

	if len(flag.Args()) != 0 {
		usage()
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if *decodePtr {
			fmt.Println(decode(line))
		} else {
			fmt.Println(encode(line, *widthPtr))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
