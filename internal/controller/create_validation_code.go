package controller

import (
	"log"
	"mongosteen/internal/email"

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
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(err)
		c.String(400, "参数错误")
		return
	}
	if err := email.SendValidationCode(body.Email, "123456"); err != nil {
		log.Println("[SendValidationCode fail]", err)
		c.String(500, "发送失败")
		return
	}
	log.Println(body.Email)
	c.String(200, "xxx")
}
