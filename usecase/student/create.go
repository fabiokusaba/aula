package student

import "github.com/fabiokusaba/aula/entities"

func CreateStudent(fullName string, age int) (student *entities.Student, err error) {
	student = entities.NewStudent(fullName, age)

	entities.Students = append(entities.Students, *student)

	return student, err
}
