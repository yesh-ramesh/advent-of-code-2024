package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	shouldDo := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`do\(\)|don't\(\)|mul\([0-9]+,[0-9]+\)`)
		matches := re.FindAllString(line, -1)

		for _, match := range matches {
			if match == "do()" {
				shouldDo = true
			} else if match == "don't()" {
				shouldDo = false
			} else {
				if shouldDo {
					match = strings.TrimPrefix(match, "mul(")
					match = strings.TrimSuffix(match, ")")
					split := strings.Split(match, ",")

					one, _ := strconv.Atoi(split[0])
					two, _ := strconv.Atoi(split[1])

					total += one * two
				}
			}
		}
	}
	fmt.Println(total)
}
