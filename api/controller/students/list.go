package students

import (
	"github.com/fabiokusaba/aula/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, entities.Students)
}
