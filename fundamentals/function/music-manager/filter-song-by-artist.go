package main

import "fmt"

func filterSongByArtist() {
	var artistname string

	fmt.Print("Artist name: ")
	fmt.Scan(&artistname)

	if artistname == "" {
		fmt.Println("Artist name is required")
		return
	}

	songsByArtist, err := filter(func(song Song) bool {
		return artistname == song.Artist
	},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Show filtered songs: ")
	for i, v := range songsByArtist {
		if i == 0 {
			fmt.Println("--------------------")
		}
		fmt.Println("Title: ", v.Title)
		fmt.Println("Song by : ", v.Artist)
		fmt.Println("--------------------")
	}
}
