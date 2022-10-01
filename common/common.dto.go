package common

import "youtube_search_engine/util"

type ErrorDto struct {
	Meta AckDto      `json:"meta"`
	Data interface{} `json:"data"`
}

type SuccessDto struct {
	Meta AckDto      `json:"meta"`
	Data interface{} `json:"data"`
}

type AckDto struct {
	RequestId        string             `json:"requestId,omitempty"`
	Success          bool               `json:"success"`
	Code             util.ErrorCodeEnum `json:"code,omitempty"`
	Message          *string            `json:"message"`
	IsPaginated      bool               `json:"isPaginated"`
	PaginationParams *PaginationParams  `json:"paginationParams,omitempty"`
}

type PaginationParams struct {
	TotalCount int64 `json:"totalCount"`
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
}
