package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/gole-dev/gole/pkg/app"
	"github.com/gole-dev/gole/pkg/log"
)

// Ping ping
// @Summary ping
// @Description ping
// @Tags system
// @Accept  json
// @Produce  json
// @Router /ping [get]
func Ping(c *gin.Context) {
	log.Info("Get function called.")

	app.Success(c, gin.H{})
}
