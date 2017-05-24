package luna

func Get(url string, reqOpt *ReqOptions) (resp *Response, err error) {
	return Request(url, "GET", reqOpt)
}

func Post(url string, reqOpt *ReqOptions) (resp *Response, err error) {
	return Request(url, "Post", reqOpt)
}
