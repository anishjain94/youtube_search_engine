package common

import (
	"context"
	"net/url"
)

func GetQueryValueFromCtx(ctx *context.Context) *url.Values {

	query := (*ctx).Value(CTX_QUERIES).(*url.Values)

	return query
}
