package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func reverse(s string) (string, error) {
	rs := []rune(s)
	if len(rs) <= 1 {
		return "", errors.New("not reverse of string the actual lenght.")
	}
	last := len(rs) - 1
	for pos := 0; pos <= last/2; pos++ {
		rs[pos], rs[last-pos] = rs[last-pos], rs[pos]
	}
	return string(rs), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	s, err := reverse(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
