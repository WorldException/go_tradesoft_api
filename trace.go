package go_tradesoft_api

import (
	"github.com/go-resty/resty/v2"
)

type TraceInfoEvent struct {
	TraceInfo  resty.TraceInfo
	Url        string
	Method     string
	BodySize   int64
	StatusCode int
}
