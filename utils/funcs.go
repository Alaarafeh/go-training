package utils

import "github.com/google/uuid"

type NewCOOL struct {
	id string
}

func foo() {}
func Blah() string {
	return "some"
}

func NewCool() NewCOOL {
	id := uuid.NewString()
	return NewCOOL{id: id}
}
