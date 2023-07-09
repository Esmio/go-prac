package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

// CreateValidationCode godoc
// @Summary      用来邮箱发送验证码
// @Description  接受邮箱地址，发送验证码
// @Tags         Ping
// @Accept       json
// @Produce      json
// @Param        id   path      int  false  "Account ID"
// @Success      200
// @Failure      500
// @Router       /validation_codes [post]
func CreateValidationCode(c *gin.Context) {
	var body struct {
		Email string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(err)
		c.String(400, "参数错误")
		return
	}
	log.Println("-------------------")
	log.Println(body.Email)
	c.String(400, "xxx")
}
