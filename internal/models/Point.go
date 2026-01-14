package models

import (
	"errors"
	"snake-game/internal/input"
)

type Point struct {
	X, Y int
}

func (p Point) UpPoint() Point {
	return Point{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p Point) LeftPoint() Point {
	return Point{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p Point) DownPoint() Point {
	return Point{
		X: p.X,
		Y: p.Y - 1,
	}
}

func (p Point) RightPoint() Point {
	return Point{
		X: p.X + 1,
		Y: p.Y,
	}
}

func (p1 Point) OppositePoint(p2 Point) (Point, error) {
	dir, err := p1.GetDirectionIndex(p2)
	if err != nil {
		return p1, err
	}
	oppositeDir := (dir + 2) % 4
	return p1.ByIntDirectionPoint(oppositeDir), nil
}

func (p Point) IsCorrect(minX, maxX, minY, maxY int) bool {
	return p.X >= minX &&
		p.X < maxX &&
		p.Y >= minY &&
		p.Y < maxY
}

func (p Point) ByIntDirectionPoint(direction int) Point {
	switch direction % 4 {
	case 0:
		return p.UpPoint()
	case 1:
		return p.LeftPoint()
	case 2:
		return p.DownPoint()
	case 3:
		return p.RightPoint()
	}
	return p
}

// Возвращает себя в случае если направление NONE
func (p Point) ByDirectionPoint(direction input.Direction) Point {
	if direction == input.NONE {
		return p
	}
	return p.ByIntDirectionPoint(int(direction) - 1)
}

func (p1 Point) GetDirectionIndex(p2 Point) (int, error) {
	switch p2 {
	case p1.UpPoint():
		return 0, nil
	case p1.LeftPoint():
		return 1, nil
	case p1.DownPoint():
		return 2, nil
	case p1.RightPoint():
		return 3, nil
	}
	return 0, errors.New("Points not neardy")
}

func (p1 Point) IsEqual(p2 Point) bool {
	return (p1.X == p2.X) && (p1.Y == p2.Y)
}
