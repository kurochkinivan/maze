package terminal

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/entities"
)

func parseToPoint(coordinates string) (entities.Point, error) {
	coordinates = strings.ReplaceAll(coordinates, " ", "")
	coords := strings.Split(coordinates, ",")

	if len(coords) != 2 {
		return entities.Point{}, fmt.Errorf("invalid amount of coordinates, expected 2, got %d", len(coords))
	}

	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return entities.Point{}, fmt.Errorf("failed to parce first coordinate %q: %w", coords[0], err)
	}

	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return entities.Point{}, fmt.Errorf("failed to parce second coordinate %q: %w", coords[1], err)
	}

	return entities.NewPoint(y, x), nil
}
