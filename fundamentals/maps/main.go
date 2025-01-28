package main

import (
	"fmt"
	"maps"
)

type Song struct {
	title    string
	duration float64
}

// alias type for short type map.
type SongMap map[string]Song

func New(t string, a float64) Song {
	return Song{title: t, duration: a}
}

func main() {

	songs := SongMap{
		"lizzy mcalpine": New("March", 3.4),
		"liang lawrance": New("out Loud", 3.0),
		"grace anger":    New("Good Stuff", 3.0),
	}

	maps.DeleteFunc(songs, func(k string, v Song) bool {
		// delete all values expect
		return k != "lizzy mcalpine"
	})

	// CLONE
	songs2 := maps.Clone(songs)
	songs2["warren hue"] = New("runaway w me", 3.3)
	fmt.Printf("CLONE: %v\n", songs2)
	fmt.Println(songs)

	i, v := songs["lizzy mcalpines"]
	fmt.Println(i, "-", v)
}
