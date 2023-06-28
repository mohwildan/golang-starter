package middelware

import (
	"golang-starter/helpers"

	"github.com/gin-gonic/gin"
)

type AppMiddelaware struct {
}

func (m *AppMiddelaware) Auth(c *gin.Context) {
	appLocals := helpers.NewAppLocalsService()
	go appLocals.SetData("data", "mantap")

	c.Next()
}
