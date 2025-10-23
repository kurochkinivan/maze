package main

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func main() {
	generator := domain.NewPrimGenerator()

	m := domain.New(5, 3, generator)
	m.Display()
}
