package students

import (
	"github.com/fabiokusaba/aula/api/controller"
	"github.com/fabiokusaba/aula/entities"
	"github.com/fabiokusaba/aula/entities/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Delete(c *gin.Context) {
	var newStudents []entities.Student

	idString := c.Param("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("invalid id"))
		return
	}

	for _, studentElement := range entities.Students {
		if studentElement.ID != id {
			newStudents = append(newStudents, studentElement)
		}
	}

	entities.Students = newStudents
	c.JSON(http.StatusOK, controller.NewResponseMessage("student deleted successfully"))
}
