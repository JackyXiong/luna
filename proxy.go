package luna

import (
	"net/http"
	"net/url"
	// "golang.org/x/net/proxy"
)

type Socks5Auth struct {
	username string
	password string
}

// set proxy by proxy server url with port
func (reqOpt *ReqOptions) SetHttpProxy(proxyURL string) (err error) {
	fixedURL, err := url.Parse(proxyURL)
	if err != nil {
		return
	}
	reqOpt.proxyURL = proxyURL
	reqOpt.proxy = http.ProxyURL(fixedURL)
	return
}

// func (reqOpt *ReqOptions) SetScoksProxy(proxyURL string, socks5Auth Socks5Auth) (err error) {
// var auth proxy.Auth
// if socks5Auth == nil {
// auth = nil
// } else {
// auth = proxy.Auth{User: socks5Auth.username, Password: socks5Auth.password}
// }

// dialer, err := proxy.SOCKS5("tcp", proxyURL, auth, proxy.Direct)
// if err != nil {
// return
// }
// reqOpt
// }

// return proxy url
func (reqOpt *ReqOptions) GetProxyURL() (url string) {
	return reqOpt.proxyURL
}
