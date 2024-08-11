package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dict := make(map[int]bool)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numbers := strings.Split(input, " ")
	for _, num := range numbers {
		tmp, _ := strconv.Atoi(num)
		dict[tmp] = true
	}
	fmt.Println(len(dict))
}
