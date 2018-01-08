package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func getInputAsInt(in string) (int, int) {
	inArr := strings.Split(in, " ")
	// Atoi(s string) (i int, err error) のerrは捨てる
	val1, _ := strconv.Atoi(inArr[0])
	val2, _ := strconv.Atoi(inArr[1])
	return val1, val2
}

func euclidean(val1, val2 int) (int) {
	surp := 1
	for  {
		surp = val1 % val2
		if surp != 0 {
			val1 = val2
			val2 = surp
		} else {
			break
		}
	}
	return val2
}

func calcLCM(val1, val2, gcd int) (int) {
	return val1/gcd*val2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		val1, val2 := getInputAsInt(scanner.Text())
		gcd := euclidean(val1, val2)
		lcm := calcLCM(val1, val2, gcd)
		fmt.Println(gcd, lcm)
	}
}

