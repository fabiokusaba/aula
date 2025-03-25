package students

import (
	"github.com/fabiokusaba/aula/api/controller"
	"github.com/fabiokusaba/aula/entities/shared"
	student_usecase "github.com/fabiokusaba/aula/usecase/student"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Delete(c *gin.Context) {
	idString := c.Param("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("invalid id"))
		return
	}

	if err = student_usecase.DeleteByID(id); err != nil {
		c.JSON(http.StatusNotFound, controller.NewResponseMessageError("student not found"))
	}

	c.JSON(http.StatusOK, controller.NewResponseMessage("student deleted successfully"))
}
