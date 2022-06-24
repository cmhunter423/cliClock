package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	screen.Clear()

	for {

		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		ssec := now.Nanosecond() / 1e8

		//fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
			dot,
			digits[ssec],
		}

		for line := range clock[0] {
			// Print a line for each placeholder in clock
			for index, digit, space := range clock {
				next := clock[index][line][space]
				/* if digit == colon && sec%2 == 0 {
					next = "      "
				}
				//fmt.Print(next, time.Millisecond.Seconds())
				if digit == dot && sec%2 == 0 {
					next = "      "
				} */
				//same as 2 if statements above
				if (digit == colon || digit == dot) && sec%2 == 0 {
					next = "      "
				}
				fmt.Print(next, "      ")
			}
			fmt.Println()

		}
		fmt.Println()

		const splitSecond = time.Second / 10

		time.Sleep(splitSecond)

	}

}
