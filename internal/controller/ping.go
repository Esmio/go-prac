package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      测试 API 是否可以正常工作
// @Description  测试 API 是否可以正常工作
// @Tags         Ping
// @Accept       json
// @Produce      json
// @Param        id   path      int  false  "Account ID"
// @Success      200
// @Failure      500
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")

}
