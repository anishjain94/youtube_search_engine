package common

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"youtube_search_engine/util"

	"github.com/gorilla/mux"
)

func HandleHTTPPost[InputDtoType any, OutputDtoType any](serviceFunc func(ctx *context.Context, dto *InputDtoType) *OutputDtoType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pCtx := &ctx

		queries := r.URL.Query()
		*pCtx = context.WithValue(*pCtx, CTX_QUERIES, &queries)

		params := mux.Vars(r)
		*pCtx = context.WithValue(*pCtx, CTX_PARAMS, &params)

		var dto InputDtoType

		_ = json.NewDecoder(r.Body).Decode(&dto)
		// ValidateDto(dto)

		response := serviceFunc(pCtx, &dto)

		w.Header().Set(string(HEADER_CONTENT_TYPE), "application/json")
		msg := MSG_SUCCESS
		json.NewEncoder(w).Encode(SuccessDto{
			Meta: AckDto{
				Success: true,
				Message: &msg,
			},
			Data: response,
		})
	}
}

func HandleHTTPGet[OutputDtoType any](serviceFunc func(ctx *context.Context) *OutputDtoType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		queries := r.URL.Query()
		ctx = context.WithValue(ctx, CTX_QUERIES, &queries)

		params := mux.Vars(r)
		ctx = context.WithValue(ctx, CTX_PARAMS, &params)

		response := serviceFunc(&ctx)

		// paginationParams, isPaginated := getPaginationData(&ctx)

		w.Header().Set(string(HEADER_CONTENT_TYPE), "application/json")
		msg := MSG_SUCCESS
		json.NewEncoder(w).Encode(SuccessDto{
			Meta: AckDto{
				Success: true,
				Message: &msg,
				// PaginationParams: paginationParams,
				// IsPaginated:      isPaginated,
			},
			Data: response,
		})
	}
}

// TODO: Add pagination data
func getPaginationData(ctx *context.Context) (*PaginationParams, bool) {
	var paginationParams *PaginationParams
	isPaginated := false

	queryParamsCtx := (*ctx).Value(CTX_QUERIES)
	if queryParamsCtx != nil {
		queryParams := queryParamsCtx.(*url.Values)

		paginationParams = &PaginationParams{}
		var page int
		var err error

		pageString := queryParams.Get(string("PARAM_PAGE"))
		pageSizeString := queryParams.Get(string("PARAM_PAGE_SIZE"))
		if pageString == "" || pageSizeString == "" {
			return nil, false
		}

		if pageString != "" {
			isPaginated = true
			page, err = strconv.Atoi(pageString)
			util.AssertError(err, "util.PARAM_ERROR", http.StatusBadRequest, "MSG_INCORRECT_PARAMS")
			paginationParams.Page = page
		}

		var pageSize int
		if pageSizeString != "" {
			isPaginated = true
			pageSize, err = strconv.Atoi(pageSizeString)
			util.AssertError(err, "util.PARAM_ERROR", http.StatusBadRequest, "MSG_INCORRECT_PARAMS")
			paginationParams.PageSize = pageSize
		}

		paginationParamsCtx := (*ctx).Value("CTX_PAGINATION_PARAMS")
		if paginationParamsCtx != nil {
			paginationParamsAsserted := paginationParamsCtx.(*PaginationParams)
			paginationParams.TotalCount = paginationParamsAsserted.TotalCount
		}
	}
	return paginationParams, isPaginated
}
