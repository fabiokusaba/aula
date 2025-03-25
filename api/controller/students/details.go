package students

import (
	"github.com/fabiokusaba/aula/api/controller"
	"github.com/fabiokusaba/aula/entities"
	"github.com/fabiokusaba/aula/entities/shared"
	student_usecase "github.com/fabiokusaba/aula/usecase/student"
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

	studentFound, err = student_usecase.SearchStudentByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, studentFound)
}
