package main

import (
	"fmt"
	"unicode"
)

func main() {
	/*
		In the game of Scrabble, players create words to score points, and the number of points is the sum of the point values of each letter in the word.
		For example, if we wanted to score the word “CODE”, we would note that the ‘C’ is worth 3 points, the ‘O’ is worth 1 point, the ‘D’ is worth 2 points, and the ‘E’ is worth 1 point. Summing these, we get that “CODE” is worth 7 points.

		implement a program in Go that determines the winner of a short Scrabble-like game. Your program should prompt for input twice: once for “Player 1” to input their word and once for “Player 2” to input their word. Then, depending on which player scores the most points, your program should either print “Player 1 wins!”, “Player 2 wins!”, or “Tie!” (in the event the two players score equal points).
	*/

	var inputPlayer1 string
	var inputPlayer2 string

	input("Player 1", &inputPlayer1)
	input("Player 2", &inputPlayer2)

	player1 := computeWord(inputPlayer1)
	player2 := computeWord(inputPlayer2)

	if player1 > player2 {
		fmt.Println("Player 1 wons!")
	} else if player1 < player2 {
		fmt.Println("Player 2 wons!")
	} else {
		fmt.Println("Tie!")
	}
}

func input(message string, word *string) {
	fmt.Printf("%v: ", message)
	fmt.Scan(word)
}

func computeWord(word string) int64 {
	// POINTS EACH LETTER
	POINTS := []int64{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}
	var score int64

	for i, v := range word {
		if unicode.IsUpper(v) {
			score += POINTS[word[i]-'A']
		} else if unicode.IsLower(v) {
			score += POINTS[word[i]-'a']
		}
	}

	return score
}
