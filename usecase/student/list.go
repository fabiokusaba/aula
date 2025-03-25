package student

import "github.com/fabiokusaba/aula/entities"

func GetAllStudents() (students []entities.Student, err error) {
	students = entities.Students

	return students, err
}
