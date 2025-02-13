package main

import (
	"fmt"

	"exampl.com/music-manager/utils"
)

func removeSong() {
	var song string
	var newUpdateSong []Song

	fmt.Print("Song name: ")
	fmt.Scan(&song)

	songs, err := utils.ReadFile[[]Song]("songs.json")
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	var songFound bool
	for _, v := range *songs {
		if v.Title == song {
			songFound = true
		} else {
			newUpdateSong = append(newUpdateSong, v)
		}
	}

	if !songFound {
		fmt.Println("Song not found")
		return
	}

	utils.WriteDataToJson("songs.json", newUpdateSong)
	fmt.Println("Song successfuly deleted")
}
