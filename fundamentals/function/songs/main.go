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
    2. filter the slice by implementing,
       function that do filter by argm function
	   and return new slice.
*/

func main() {
	var songByLizzyMcAlpine []Song

	songByLizzyMcAlpine = append(songByLizzyMcAlpine, []Song{
		{duration: 3.4, like: 5324.2, title: "spring into summer", trending: true},
		{duration: 5.4, like: 2524.2, title: "pushing it down and praying", trending: true},
		{duration: 3.8, like: 1024.2, title: "older", trending: false},
		{duration: 3.7, like: 100324.2, title: "ceilings", trending: true},
		{duration: 5.9, like: 60324.2, title: "vortex", trending: true},
		{duration: 3.6, like: 2024.2, title: "method acting", trending: false},
		{duration: 5.0, like: 324.2, title: "march", trending: false},
	}...)

	// filterSongByTrending := filterSong[Song](&songByLizzyMcAlpine)
	// for i, v := range songByLizzyMcAlpine {

	// 	if v.trending {
	// 		songByLizzyMcAlpine = append(songByLizzyMcAlpine, songByLizzyMcAlpine...)
	// 	} else {
	// 		slices.Delete()
	// 	}
	// }
	// fmt.Println(songByLizzyMcAlpine)
	songByLizzyMcAlpineTrending := filterSong(&songByLizzyMcAlpine)
	fmt.Printf("TRENDING: %v\n", songByLizzyMcAlpineTrending)
}

// should passing function to filter song.
func filterSong(m *[]Song) []Song {
	var filteredSong []Song

	for _, v := range *m {
		if v.trending {
			filteredSong = append(filteredSong, v)
		}
	}

	return filteredSong
}

// func getTrending(m *[]Song) Song {
// 	var tredingSong []Song
// 	for i, v := range *m {

// 	}

// 	return tredingSong
// }
