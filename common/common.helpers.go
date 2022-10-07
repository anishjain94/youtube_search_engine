package common

import (
	"context"
	"net/url"
	"strings"
)

func GetQueryValueFromCtx(ctx *context.Context) *url.Values {

	query := (*ctx).Value(CTX_QUERIES).(*url.Values)

	return query
}

func ParseToTsQuery(str string) string {
	strs := strings.Split(str, " ")

	return strings.Join(strs, " & ")
}
