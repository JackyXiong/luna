package luna

import (
	"net/http"
)

const defaultUserAgent = "go-luna"

var defaultHeaders = map[string]string{
	"Connection":      "keep-alive",
	"User-Agent":      defaultUserAgent,
	"Accept":          "*/*",
	"Accept-Encoding": "gzip, deflate",
}

func mergeHeaders(req *http.Request, reqOpt *ReqOptions, contentType string) {
	for k, v := range defaultHeaders {
		req.Header.Set(k, v)
	}

	for k, v := range reqOpt.Headers {
		req.Header.Set(k, v)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
}
