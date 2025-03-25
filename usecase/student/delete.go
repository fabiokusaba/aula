package student

import (
	"github.com/fabiokusaba/aula/entities"
	"github.com/google/uuid"
)

func DeleteByID(id uuid.UUID) (err error) {
	var newStudents []entities.Student

	for _, studentElement := range entities.Students {
		if studentElement.ID != id {
			newStudents = append(newStudents, studentElement)
		}
	}

	entities.Students = newStudents
	return err
}
