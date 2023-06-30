package middelware

import (
	"github.com/gin-gonic/gin"
)

type AppMiddelaware struct {
}

func (m *AppMiddelaware) Auth(c *gin.Context) {
	c.Next()
}
