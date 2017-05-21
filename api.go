package luna

import (
	// "fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 封装的请求选项，用来构造http.Request
type ReqOptions struct {
	Params    map[string]string // get
	Data      map[string]string // post
	Headers   map[string]string
	Json      map[string]interface{} // post json data
	Timeout   int64                  // millsecond
	BasicAuth BasicAuth
	Hook      Hook
	Files     []File
	proxy     func(*http.Request) (*url.URL, error)
	proxyURL  string
}

func NewReqOptions() *ReqOptions {
	// headers := map[string]string{}
	return &ReqOptions{
		Data:    nil,
		Headers: nil,
		Timeout: DefaultTimeout,
	}
}

func newBody(reqOpt *ReqOptions) (body io.Reader, contentType string, err error) {
	if reqOpt.Data == nil && reqOpt.Files == nil && reqOpt.Json == nil {
		return nil, "", nil
	}
	if reqOpt.Files != nil {
		return newMultipartBody(reqOpt)
	}
	if reqOpt.Json != nil {
		return newJsonBody(reqOpt)
	}
	data := url.Values{}
	for k, v := range reqOpt.Data {
		data.Set(k, v)
	}
	contentType = "application/x-www-form-urlencoded"
	return strings.NewReader(data.Encode()), contentType, nil
}

func createRequest(reqOpt *ReqOptions, method string, url string) (req *http.Request, err error) {
	//构造body
	body, contentType, err := newBody(reqOpt)
	if err != nil {
		return nil, nil
	}

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	mergeHeaders(req, reqOpt, contentType)
	if reqOpt.BasicAuth.User != "" {
		req.SetBasicAuth(reqOpt.BasicAuth.User, reqOpt.BasicAuth.Password)
	}
	return
}

// 发送请求
func Request(url string, method string, reqOpt *ReqOptions) (resp *Response, err error) {
	req, err := createRequest(reqOpt, method, url)
	if err != nil {
		return nil, err
	}
	applyBeforeHooks(req, reqOpt)
	client := new(http.Client)
	// set timeout
	if reqOpt.Timeout != 0 {
		client.Timeout = time.Duration(reqOpt.Timeout) * time.Millisecond
	}
	// set proxy
	if reqOpt.proxy != nil {
		client.Transport = &http.Transport{Proxy: reqOpt.proxy}
	}

	if OriginResp, err := client.Do(req); err != nil
	if err != nil {
		return nil, err
	}
	applyAfterHooks(OriginResp, reqOpt)
	resp = &Response{OriginResp, nil}
	return
}

func Get(url string, reqOpt *ReqOptions) (resp *Response, err error) {
	return Request(url, "GET", reqOpt)
}

func Post(url string, reqOpt *ReqOptions) (resp *Response, err error) {
	return Request(url, "Post", reqOpt)
}
