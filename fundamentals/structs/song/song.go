package main

import (
	artiststruct "example.com/struct-song/artist_struct"
	songstruct "example.com/struct-song/song_struct"
)

func main() {

	// Creat struct.
	song := songstruct.New("vortex", 5.24) // it is return pointer (0x17361263...)

	song.PrintTitleSong()
	song.UpdateSong("come down soon", 3.15)
	song.PrintTitleSong()

	song.DeleteSong()
	song.PrintTitleSong()

	artist := artiststruct.New("lizzy mcalpine", "united state", song)

	artist.PrintArtist()
	artist.UpdateArtist("elizabeth mcalpine", "philadelpia")
	artist.PrintArtist()
}
