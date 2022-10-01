package common

type ContextTypeEnum int

const (
	CTX_QUERIES ContextTypeEnum = iota

	CTX_PARAMS
)

type HeaderEnum string

const (
	HEADER_CONTENT_TYPE HeaderEnum = "Content-Type"
)
