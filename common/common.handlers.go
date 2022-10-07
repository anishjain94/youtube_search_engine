package common

import (
	"context"
	"encoding/json"
	"net/http"

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

		response := serviceFunc(&ctx)

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
