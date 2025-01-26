package main

import "fmt"

// Animal can move because has Move method.
type Animal interface {
	Move()
}

// Zebra can Move() because declared Move method.
type Zebra struct {
	iam string
}

func (z Zebra) Move() {
	fmt.Println("Zebra says: \"Animal\" can only move.")
}

/*
Fish can Move() because implement Animal interface.
Fish can Swim() because has Swim method.
*/
type Fish interface {
	Animal
	Swim()
}

// ClownFish can Move and Swim because has method Move
type ClownFish struct {
	iam string
}

func (cf ClownFish) Move() {
	fmt.Println("ClownFish says: \"Animal\" can only move. I can move too.")
}

func (cf ClownFish) Swim() {
	fmt.Println("Clownfish says: \"Fish\" can move and swim too.")
}

func main() {
	pico := Zebra{iam: "Zebra Animal"}
	AnimalSayer(pico)

	nemo := ClownFish{iam: "Clown Fish"}
	FishSayer(nemo)
}

func AnimalSayer(a Animal) {
	a.Move()
}

func FishSayer(f Fish) {
	f.Move()
	f.Swim()
}
