package students

import (
	"github.com/fabiokusaba/aula/api/controller"
	student_usecase "github.com/fabiokusaba/aula/usecase/student"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	students, err := student_usecase.GetAllStudents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, students)
}
