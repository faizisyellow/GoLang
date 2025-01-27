package main

import "fmt"

func ScoresOut() []int {
	scores := []int{}
	var input int
	for i := 0; i < 3; i++ {
		fmt.Scan(&input)
		scores = append(scores, input)
	}
	return scores
}
