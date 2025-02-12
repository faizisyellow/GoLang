package main

import (
	"fmt"
)

type Song struct {
	duration float64
	like     float64
	title    string
	trending bool
}

func New(d, l float64, t string, tr bool) Song {
	return Song{
		duration: d,
		like:     l,
		title:    t,
		trending: tr,
	}
}

/*
   TODO:
    1. create a slice of struct.
    2. filter the slice by implementing, function as parameter.
*/

func main() {
	var songByLizzyMcAlpine []Song

	songByLizzyMcAlpine = append(songByLizzyMcAlpine, []Song{
		{duration: 3.4, like: 5324.2, title: "spring into summer", trending: false},
		{duration: 5.4, like: 2524.2, title: "pushing it down and praying", trending: false},
		{duration: 3.8, like: 1024.2, title: "older", trending: false},
		{duration: 3.7, like: 100324.2, title: "ceilings", trending: false},
		{duration: 5.9, like: 60324.2, title: "vortex", trending: true},
		{duration: 3.6, like: 2024.2, title: "method acting", trending: false},
		{duration: 5.0, like: 324.2, title: "march", trending: false},
	}...)

	fmt.Println(songByLizzyMcAlpine)

	// passing a function as an argument
	songByLizzyMcAlpineTrending := filterSong(&songByLizzyMcAlpine, filterBytrending)

	fmt.Printf("TRENDING: %v\n", songByLizzyMcAlpineTrending)
}

// describe shape of the function parameter should be
func filterSong(m *[]Song, filterby func(song Song) bool) []Song {
	var filteredSong []Song

	for _, v := range *m {
		// calling the result of function parameter
		if filterby(v) {
			filteredSong = append(filteredSong, v)
		}
	}

	return filteredSong
}

func filterBytrending(song Song) bool {

	return song.trending
}
