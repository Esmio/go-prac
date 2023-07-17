package controller

import (
	"crypto/rand"
	"log"
	"mongosteen/config/queries"
	"mongosteen/internal/database"
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
	// 伪随机 math/rand
	// 真随机 crypto/rand
	len := 4
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("[rand.Read fail]", err)
		c.String(500, "发送失败")
		return
	}
	digits := make([]byte, len)
	for i := range b {
		digits[i] = b[i]%10 + 48
	}
	str := string(digits)
	q := database.NewQuery()
	vc, err := q.CreateValidationCode(c, queries.CreateValidationCodeParams{
		Email: body.Email,
		Code:  str,
	})
	if err != nil {
		// TODO
		c.Status(400)
		return
	}
	if err := email.SendValidationCode(vc.Email, vc.Code); err != nil {
		log.Println("[SendValidationCode fail]", err)
		c.String(500, "发送失败")
		return
	}
	log.Println(body.Email)
	c.String(200, "xxx")
}
