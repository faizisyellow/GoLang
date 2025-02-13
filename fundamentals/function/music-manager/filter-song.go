package main

import (
	"errors"

	"exampl.com/music-manager/utils"
)

func filter(filterFunc func(song Song) bool) ([]Song, error) {
	var filteredSong []Song
	songs, err := utils.ReadFile[[]Song]("songs.json")
	if err != nil {
		return []Song{}, errors.New("songs not found")
	}

	for _, v := range *songs {
		if filterFunc(v) {
			filteredSong = append(filteredSong, v)
		}
	}

	return filteredSong, nil
}
