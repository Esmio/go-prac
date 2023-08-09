package api

import "mongosteen/config/queries"

type GetMeResponse struct {
	Resource queries.User
}
