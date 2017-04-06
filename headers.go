package luna

import (
	"net/http"
)

var defaultHeaders = map[string]string{
	"Connection":      "keep-alive",
	"User-Agents":     "luna",
	"Accept":          "*/*",
	"Accept-Encoding": "gzip, deflate",
}


func mergeHeaders(req *http.Request, reqOpt *ReqOptions) {
	for k, v := range  defaultHeaders {
		req.Header.Set(k, v)
	}

	for k, v := range reqOpt.Headers {
		req.Header.Set(k, v)
	}
}