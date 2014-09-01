package main

import(
	"fmt"
	"strings"
	"strconv"
	"os"
	"log"
	"bufio"
)

func GrabLines(args string) []string {
	var lines []string
	file, err := os.Open(args)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	lines = append(lines, scanner.Text())
	}
	return lines
}

func Reverse(line []string) []string {
	var reversed []string
	for i := len(line)-1; i >= 0; i-- {
		reversed = append(reversed, line[i])
	}
	return reversed
}

func main() {
	argsWithoutProgram := os.Args[1]; if argsWithoutProgram == "" {
		fmt.Println("Need a single filename argument to run this program")
	}
	lines := GrabLines(argsWithoutProgram)
	cases, _ := strconv.Atoi(lines[0])
	for i := 1; i <= cases; i++{
		line := strings.Split(lines[i], " ")
		reversed := strings.Join(Reverse(line), " ")
		fmt.Printf("Case #%d: ", i)
		fmt.Println(reversed)
	}
}
