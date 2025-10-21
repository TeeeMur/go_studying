package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func main() {
	var colors [5]func(format string, a ...interface{}) = [5]func(format string, a ...interface{}){
		color.Blue,
		color.Red,
		color.Green,
		color.Cyan,
		color.Yellow,
	}
	var input string
	fmt.Fscan(os.Stdin, &input)
	var num, _ = strconv.Atoi(input)

	for num < 12307 {
		if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num%9 == 0 {
			num *= 13
			num += 1
			continue
		} else {
			num += 2
			num *= 3
		}

		if num%13 == 0 {
			fmt.Fprintf(os.Stdout, "%d", &num)
			var res = strconv.Itoa(num)
			rand.New(rand.NewSource(time.Now().Unix()))
			colors[rand.Intn(len(colors))](res)
			break
		} else {
			num += 1
		}
	}
	fmt.Fprintln(os.Stdout, "service error")
}
