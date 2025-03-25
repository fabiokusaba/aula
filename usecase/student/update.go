package student

import (
	"errors"
	"github.com/fabiokusaba/aula/entities"
	"github.com/fabiokusaba/aula/entities/shared"
	"github.com/google/uuid"
)

func UpdateStudent(id uuid.UUID, fullName string, age int) (student entities.Student, err error) {
	var newStudents []entities.Student

	for _, studentElement := range entities.Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("student does not exist")
	}

	student.FullName = fullName
	student.Age = age

	for _, studentElement := range entities.Students {
		if studentElement.ID == student.ID {
			newStudents = append(newStudents, student)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	entities.Students = newStudents
	return student, err
}
