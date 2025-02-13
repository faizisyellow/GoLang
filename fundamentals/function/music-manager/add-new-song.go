package main

import (
	"fmt"
	"strconv"

	"exampl.com/music-manager/utils"
)

func addNewSong() {
	var title, artist, duration string

	fmt.Print("Title: ")
	fmt.Scan(&title)

	fmt.Print("Artist: ")
	fmt.Scan(&artist)

	fmt.Print("Duration: ")
	fmt.Scan(&duration)

	numduration, err := strconv.Atoi(duration)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}

	newSong, err := New(title, artist, numduration)
	if err != nil {
		fmt.Printf("Error in line 55: %v", err.Error())
		return
	}

	songs, err := utils.ReadFile[[]Song]("songs.json")
	if err != nil {
		fmt.Println("error get song: ", err.Error())
		return
	}

	var songExist bool
	for _, v := range *songs {
		if v.Title == newSong.Title {
			songExist = true
			fmt.Println("Song already exists!")
			break
		}
	}

	if !songExist {
		*songs = append(*songs, *newSong)

		utils.WriteDataToJson("songs.json", *songs)
	}

}
