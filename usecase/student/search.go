package student

import (
	"errors"
	"github.com/fabiokusaba/aula/entities"
	"github.com/fabiokusaba/aula/entities/shared"
	"github.com/google/uuid"
)

func SearchStudentByID(id uuid.UUID) (student entities.Student, err error) {
	for _, studentElement := range entities.Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("student not found")
	}

	return student, err
}
