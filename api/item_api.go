package api

import (
	"mongosteen/config/queries"
	"time"
)

type CreateItemRequest struct {
	Amount     int32        `json:"amount" binding:"required" example:"2333"`
	Kind       queries.Kind `json:"kind" binding:"required"`
	HappenedAt time.Time    `json:"happened_at" binding:"required"`
	TagIds     []int32      `json:"tag_ids" binding:"required"`
}

type CreateItemResponse struct {
	Resource queries.Item
}

type GetPagedItemsRequest struct {
	Page           int32     `josn:"page"`
	HappenedAfter  time.Time `json:"happened_after"`
	HappenedBefore time.Time `json:"happened_before"`
}

type GetPagesItemsResponse struct {
	Resources []queries.Item `json:"resources" `
	Pager     Pager          `json:"pager"`
}
