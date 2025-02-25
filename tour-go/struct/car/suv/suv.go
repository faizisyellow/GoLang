package suv

import "fmt"

type Suv struct {
	Name  string
	Merch string
}

func New(name, merch string) *Suv {
	return &Suv{
		Name:  name,
		Merch: merch,
	}
}

func (s Suv) DisplaySuv() {
	fmt.Printf("Suv with name: %v\n", s.Name)
}
