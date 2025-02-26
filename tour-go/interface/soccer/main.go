package main

import (
	"fmt"
	"math/rand"
)

type SoccerPlayer interface {
	KickBall()
	Bio() string
}

type Haaland struct {
	stamina int
	power   int
}

func (h Haaland) KickBall() {
	shot := h.stamina + h.power
	fmt.Println("Haaland kicking the ball", shot)
}
func (h Haaland) Bio() string {
	return "haaland"
}

type Dolberg struct {
	stamina int
	power   int
	goals   int
}

func (d Dolberg) KickBall() {
	shot := d.stamina + d.power
	fmt.Println("Dolberg kicking the ball", shot)
}
func (d Dolberg) Bio() string {
	return "Kasper Dolberg"
}

func main() {
	player1 := Haaland{
		stamina: rand.Intn(10),
		power:   rand.Intn(10),
	}

	player2 := Dolberg{
		stamina: rand.Intn(10),
		power:   rand.Intn(10),
		goals:   10,
	}

	players := make([]SoccerPlayer, 2)
	players[0] = player1
	players[1] = player2

	for _, v := range players {
		v.KickBall()
		name := v.Bio()
		fmt.Println(name)

	}

	var i SoccerPlayer = Haaland{
		stamina: rand.Intn(10),
		power:   rand.Intn(10),
	}
	v, ok := i.(Dolberg)
	fmt.Printf("VALUE %v, %v, %T\n", v, ok, v)

	switch v := i.(type) {
	case Dolberg:
		fmt.Printf("Interface implement type %T!\n", v)
	case Haaland:
		fmt.Printf("Interface implement type %T!\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	var any interface{}
	any = "bisa string"
	any = 12
	any = 0.12

	switch any.(type) {
	case string:
		fmt.Printf("Interface implement type %T!\n", any)
	default:
		fmt.Printf("I don't know about type %T!\n", any)
	}
}
