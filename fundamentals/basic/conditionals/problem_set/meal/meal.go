package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1. Input user name
	// 2. If user's input between in hour and minute return what should do
	// 3. Otherwise return nothing

	var times string

	fmt.Print("What time is it? ")
	fmt.Scan(&times)

	hours, minutes := getTime(times)

	showMealTime(hours, minutes)
}

func getTime(time string) (float64, float64) {
	splitTime := strings.Split(time, ":")

	hours, _ := strconv.ParseFloat(splitTime[0], 64)
	minutes, _ := strconv.ParseFloat(splitTime[1], 64)

	return hours, minutes
}

func showMealTime(hours, minutes float64) {
	if hours == 7 || hours == 8 && minutes <= 59 {
		if hours == 8 && minutes > 0 {
			return
		}
		fmt.Println("breakfast time!")
	} else if hours == 12 || hours == 13 && minutes <= 59 {
		if hours == 13 && minutes > 0 {
			return
		}
		fmt.Println("lunch time!")
	} else if hours == 18 || hours == 19 && minutes <= 59 {
		if hours == 19 && minutes > 0 {
			return
		}
		fmt.Println("dinner time!")
	} else {
		return
	}
}
