package controller

import (
	"mongosteen/api"
	"mongosteen/config/queries"
	"mongosteen/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagController struct {
}

func (ctrl *TagController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.POST("/tags", ctrl.Create)
}

// CreateTag godoc
//
//	@Summary	创建标签
//	@Accept		json
//	@Produce	json
//
//	@Security	Bearer
//
//	@Param		name		body		string						true	"标签"	SchemaExample(通勤)
//	@Param		sign		body		string						true	"符号"	SchemaExample(🚌)
//	@Param		kind		body		queries.Kind			true	"类型"
//
//	@Success	200			{object}	api.CreateTagResponse	数据
//	@Failure	422			{string}	string					数据
//	@Router		/api/v1/tags [post]
func (ctrl *TagController) Create(c *gin.Context) {
	var body api.CreateTagRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(422, "参数错误")
		return
	}

	me, _ := c.Get("me")
	user, _ := me.(queries.User)
	q := database.NewQuery()
	tag, err := q.CreateTag(c, queries.CreateTagParams{
		UserID: user.ID,
		Name:   body.Name,
		Kind:   body.Kind,
		Sign:   body.Sign,
	})
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, api.CreateTagResponse{Resource: tag})
}

func (ctrl *TagController) Destroy(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *TagController) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *TagController) Get(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *TagController) GetPaged(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
