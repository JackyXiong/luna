package luna

import (
	"net/http"
	"net/url"
)

func (reqOpt *ReqOptions) SetProxy(proxyURL string) (err error) {
	fixedURL, err := url.Parse(proxyURL)
	if err != nil {
		return
	}
	reqOpt.proxyURL = proxyURL
	reqOpt.proxy = http.ProxyURL(fixedURL)
	return
}

func (reqOpt *ReqOptions) GetProxy() (url string) {
	return reqOpt.proxyURL
}
