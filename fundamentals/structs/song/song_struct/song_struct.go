package songstruct

import "fmt"

// Song type
type Song struct {
	title    string
	duration float64
}

// constructor to creating new struct.
func New(t string, d float64) *Song {
	newSong := &Song{title: t, duration: d}

	return newSong
}

// Method here

// pointer receivers functions (method in struct)
func (song *Song) PrintTitleSong() {
	fmt.Println(song.title)
}

func (song *Song) UpdateSong(t string, d float64) *Song {
	song.title = t
	song.duration = d

	return song
}

func (song *Song) DeleteSong() {
	song.title = ""
	song.duration = 0
}
