package controller

import (
	"mongosteen/api"
	"mongosteen/config/queries"
	"mongosteen/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	PerPage int32
}

func (ctrl *TagController) RegisterRoutes(rg *gin.RouterGroup) {
	ctrl.PerPage = 20
	v1 := rg.Group("/v1")
	v1.GET("/tags", ctrl.GetPaged)
	v1.POST("/tags", ctrl.Create)
	v1.PATCH("/tags/:id", ctrl.Update)
	v1.DELETE("/tags/:id", ctrl.Destroy)
	v1.GET("/tags/:id", ctrl.Get)
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
//	@Param		kind		body		string						true	"类型"
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
	idString, has := c.Params.Get("id")
	if !has {
		c.String(422, "参数错误")
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(422, "参数错误")
		return
	}

	q := database.NewQuery()
	err = q.DeleteTag(c, int32(id))
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (ctrl *TagController) Update(c *gin.Context) {
	var body api.UpdateTagRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(422, "参数错误")
		return
	}
	idString, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(422, "参数错误")
		return
	}

	me, _ := c.Get("me")
	user, _ := me.(queries.User)
	q := database.NewQuery()
	tag, err := q.UpdateTag(c, queries.UpdateTagParams{
		ID:     int32(id),
		UserID: user.ID,
		Name:   body.Name,
		Kind:   body.Kind,
		Sign:   body.Sign,
	})
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, api.UpdateTagResponse{Resource: tag})
}

func (ctrl *TagController) Get(c *gin.Context) {
	me, _ := c.Get("me")
	user, _ := me.(queries.User)
	idString, has := c.Params.Get("id")
	if !has {
		c.String(422, "参数错误")
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(422, "参数错误")
		return
	}
	q := database.NewQuery()
	tag, err := q.FindTag(c, queries.FindTagParams{
		UserID: user.ID,
		ID:     int32(id),
	})

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, api.GetTagResponse{Resource: tag})
}

func (ctrl *TagController) GetPaged(c *gin.Context) {
	me, _ := c.Get("me")
	user, ok := me.(queries.User)
	if !ok {
		c.Status(http.StatusUnauthorized)
	}
	var params api.GetPagedTagsRequest
	pageStr, _ := c.Params.Get("page")
	if page, err := strconv.Atoi(pageStr); err != nil {
		params.Page = int32(page)
	}

	if params.Page == 0 {
		params.Page = 1
	}

	kind, _ := c.Params.Get("kind")

	if kind == "" {
		kind = "expenses"
	}

	params.Kind = kind

	q := database.NewQuery()
	tags, err := q.ListTags(c, queries.ListTagsParams{
		Offset: (params.Page - 1) * ctrl.PerPage,
		Limit:  ctrl.PerPage,
		Kind:   params.Kind,
		UserID: user.ID,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}
	count, err := q.CountItems(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}
	c.JSON(http.StatusOK, api.GetPagesTagsResponse{
		Resources: tags,
		Pager: api.Pager{
			Page:    params.Page,
			PerPage: ctrl.PerPage,
			Count:   count,
		},
	})
}
