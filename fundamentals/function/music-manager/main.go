package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to music manager")
	option := chooseOptions()

	switch option {
	case 1:
		addNewSong()
	case 2:
		showAllSong()
	case 3:
		removeSong()
	case 4:
		filterSongByArtist()
	default:
		os.Exit(1)
	}
}
