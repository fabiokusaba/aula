package infra

import (
	"github.com/fabiokusaba/aula/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Heart(c *gin.Context) {
	c.JSON(http.StatusOK, controller.NewResponseMessage("OK"))
}
