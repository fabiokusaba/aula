package students

import (
	"github.com/fabiokusaba/aula/api/controller"
	student_usecase "github.com/fabiokusaba/aula/usecase/student"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var input Input
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := student_usecase.CreateStudent(input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, student)
}
