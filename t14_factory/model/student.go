package model

import "fmt"

type student struct {
	name  string
	score float64
}

func (s *student) String() string {
	return fmt.Sprintf("{name: %v, score: %v}\n", s.name, s.score)
}
func NewStudent(n string, s float64) *student {
	return &student{
		name:  n,
		score: s,
	}
}

func (s *student) GetName() string {
	return s.name
}

func (s *student) GetScore() float64 {
	return s.score
}
