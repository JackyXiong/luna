package luna

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type BasicAuth struct {
	User     string
	Password string
}

// 封装的请求选项，用来构造http.Request
type ReqOptions struct {
	Params    map[string]string // GET
	Data      map[string]string // post
	Headers   map[string]string
	Timeout   int
	BasicAuth BasicAuth
	Hook      Hook
	// File 支持buf，file path 等
}

func NewReqOptions() *ReqOptions {
	// headers := map[string]string{}
	return &ReqOptions{
		Data:    nil,
		Headers: nil,
		Timeout: DefaultTimeout,
	}
}

func newBody(reqOpt *ReqOptions) (body io.Reader, err error) {
	if reqOpt.Data == nil {
		return nil, nil
	}
	data := url.Values{}
	for k, v := range reqOpt.Data {
		data.Set(k, v)
	}
	return strings.NewReader(data.Encode()), nil
}

func CreateRequest(reqOpt *ReqOptions, method string, url string) (req *http.Request, err error) {
	//构造body
	body, err := newBody(reqOpt)

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if reqOpt.Headers != nil {
		mergeHeaders(req, reqOpt)
	}
	if reqOpt.BasicAuth.User != "" {
		req.SetBasicAuth(reqOpt.BasicAuth.User, reqOpt.BasicAuth.Password)
	}
	// hook
	return
}

func Request(url string, method string, reqOpt *ReqOptions) (resp *http.Response, err error) {
	// 解析url
	// 构造request
	req, err := CreateRequest(reqOpt, method, url)
	if err != nil {
		return nil, err
	}
	applyBeforeHooks(req, reqOpt)
	client := new(http.Client)
	resp, err = client.Do(req)
	applyAfterHooks(resp, reqOpt)
	return
}

func Get(url string, reqOpt *ReqOptions) (resp *http.Response, err error) {
	return Request(url, "GET", reqOpt)
}

func Post(url string, reqOpt *ReqOptions) (resp *http.Response, err error) {
	return Request(url, "Post", reqOpt)
}
