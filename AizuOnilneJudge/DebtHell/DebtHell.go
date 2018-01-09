package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		debt := float64(100000)
		for i:=1; i<=n; i++ {
			debt = debt*1.05
			debt = math.Ceil(debt*0.001) * 1000
		}		
		fmt.Println(int(debt))
	}
}