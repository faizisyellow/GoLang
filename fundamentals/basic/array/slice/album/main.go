package main

import "fmt"

type Song struct {
	title    string
	duration float64
	artist   string
}

func New(duration float64, title, artist string) Song {
	newSong := Song{title: title, duration: duration, artist: artist}

	return newSong
}

func (song Song) DisplaySong() {
	fmt.Printf("TITLE: %v\nDURATION: %v\nARTIST: %v\n", song.title, song.duration, song.artist)
}

func main() {
	var albums []Song

	// var title string
	// var duration float64
	// var artist string

	// for i := 0; i < 3; i++ {
	// 	fmt.Printf("title: ")
	// 	fmt.Scan(&title)

	// 	fmt.Printf("duration: ")
	// 	fmt.Scan(&duration)

	// 	fmt.Printf("artist: ")
	// 	fmt.Scan(&artist)
	// 	fmt.Println()

	// 	albums = append(albums, New(duration, title, artist))
	// }

	songs := []Song{
		New(3.4, "older", "lizzy-mcalpine"),
		New(3.4, "vortex", "lizzy-mcalpine"),
		New(3.4, "march", "lizzy-mcalpine"),
	}

	albums = append(albums, songs...)

	fmt.Printf("MY ALBUMS: %v\n", albums)
	fmt.Printf("len=%d cap=%d %v\n", len(albums), cap(albums), albums)

	songByLizzyMclpine := make([]string, 2)

	for i := range albums {
		if albums[i].artist == "lizzy-mcalpine" {
			// update the first element
			if len(songByLizzyMclpine) > i {
				songByLizzyMclpine[i] = albums[i].title
			} else {
				songByLizzyMclpine = append(songByLizzyMclpine, albums[i].title)
			}

		}
	}

	fmt.Println(songByLizzyMclpine)

}
