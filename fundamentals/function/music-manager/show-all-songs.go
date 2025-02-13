package main

import (
	"fmt"

	"exampl.com/music-manager/utils"
)

func showAllSong() {
	songs, err := utils.ReadFile[[]Song]("songs.json")

	if err != nil {
		fmt.Println("error get songs: ", err.Error())
		return
	}

	fmt.Println("Show all songs: ")
	for i, v := range *songs {
		if i == 0 {
			fmt.Println("--------------------")
		}
		fmt.Println("Title: ", v.Title)
		fmt.Println("Song by : ", v.Artist)
		fmt.Println("--------------------")
	}
}
