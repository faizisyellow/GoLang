package artiststruct

import (
	"fmt"

	songstruct "example.com/struct-song/song_struct"
)

type Artist struct {
	Name  string
	Based string
	Song  songstruct.Song
}

func New(n, b string, song *songstruct.Song) *Artist {

	artist := &Artist{
		Name:  n,
		Based: b,
		Song:  *song,
	}

	return artist

}

func (artis *Artist) PrintArtist() {

	fmt.Println(*artis)

}

func (artist *Artist) UpdateArtist(name, based string) {
	artist.Name = name
	artist.Based = based
}
