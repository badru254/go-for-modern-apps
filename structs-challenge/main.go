package main

import (
	"fmt"
	"log"
)

func main() {

	s, err := NewSquare(1, 1, 10)

	if err != nil {
		log.Fatalln(err.Error())
	}
	s.Move(2, 3)

	fmt.Printf("%+v\n", s)

	fmt.Println(s.Area())

}

type Square struct {
	PositionX int
	PositionY int
	Length    int
}

func NewSquare(x int, y int, length int) (*Square, error) {

	if length <= 0 {
		return nil, fmt.Errorf("length should a  number above 0")
	}
	square := Square{x, y, length}

	return &square, nil
}

func (s Square) Area() int {
	return s.Length * s.Length
}

func (s *Square) Move(dx int, dy int) {
	s.PositionX += dx
	s.PositionY += dy
}
