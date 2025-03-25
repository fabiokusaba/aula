package entities

import (
	"github.com/fabiokusaba/aula/entities/shared"
	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

func NewStudent(fullName string, age int) *Student {
	return &Student{
		ID:       shared.GetUuid(),
		FullName: fullName,
		Age:      age,
	}
}

var Students = []Student{
	{ID: shared.GetUuid(), FullName: "Joao", Age: 18},
	{ID: shared.GetUuid(), FullName: "Gabriel", Age: 19},
}
