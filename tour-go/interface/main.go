package main

import "fmt"

// Song struct and method's
type Song struct {
	title, artist string
	duration      float64
}

func (s Song) Display() {
	fmt.Printf("Displayed song: %v by %v. From DisplaySong method\n", s.title, s.artist)
}

// City struct and method's
type City struct {
	name, state string
}

func (c City) Display() {
	fmt.Printf("Displayed city: %v in state %v. From DisplayCity method\n", c.name, c.state)
}

// interfaces
type D interface {
	Display()
}

func main() {
	var song D
	song = Song{title: "vortex", artist: "lizzy mcalpine", duration: 4.3}

	var city D
	city = City{name: "los angeles", state: "california"}

	song.Display()
	describe(song)
	city.Display()
	describe(city)
}

func describe(i D) {
	fmt.Printf("(%v, %T)\n", i, i)
}
