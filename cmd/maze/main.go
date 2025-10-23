package main

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/generator/prim"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

func main() {
	generator := prim.New()

	m := maze.New(10, 10, generator)
	m.Display()
}
