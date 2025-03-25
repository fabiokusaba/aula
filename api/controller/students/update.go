package students

import (
	"github.com/fabiokusaba/aula/api/controller"
	"github.com/fabiokusaba/aula/entities"
	"github.com/fabiokusaba/aula/entities/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(c *gin.Context) {
	var input Input
	var studentFound entities.Student
	var newStudents []entities.Student

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	idString := c.Param("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	for _, studentElement := range entities.Students {
		if studentElement.ID == id {
			studentFound = studentElement
		}
	}

	if studentFound.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("student not found"))
		return
	}

	studentFound.FullName = input.FullName
	studentFound.Age = input.Age

	for _, studentElement := range entities.Students {
		if studentElement.ID == studentFound.ID {
			newStudents = append(newStudents, studentFound)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	entities.Students = newStudents
	c.JSON(http.StatusOK, studentFound)
}
