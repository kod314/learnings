package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := []rune(scanner.Text())
		for i:=len(str)-1; i>=0; i-- {
			fmt.Printf(string(str[i]))
		}
		fmt.Println()
	}
}