package students

import (
	"github.com/fabiokusaba/aula/api/controller"
	"github.com/fabiokusaba/aula/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var input Input
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student := entities.NewStudent(input.FullName, input.Age)
	entities.Students = append(entities.Students, *student)

	c.JSON(http.StatusCreated, student)
}
