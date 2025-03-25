package students

import (
	"github.com/fabiokusaba/aula/api/controller"
	"github.com/fabiokusaba/aula/entities"
	"github.com/fabiokusaba/aula/entities/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Details(c *gin.Context) {
	var studentFound entities.Student

	idString := c.Param("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("invalid id"))
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

	c.JSON(http.StatusOK, studentFound)
}
