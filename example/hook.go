package main

import (
	"net/http"

	"github/jackyxiong/luna"
)

type Hook struct{}

func (h *Hook) AfterResponseHook(resp *http.Response) (err error) {
	fmt.Println(resp)
	return
}

func (h *Hook) BeforeRequestHook(req *http.Request) (err error) {
	fmt.Println(req)
	return nil
}

func main() {
	// ======= example 1
	reqOpt := luna.NewReqOptions()
	hook := new(Hook)
	reqOpt.setHook(hook, nil)
	reqOpt.setHook(nil, hook)
	reqOpt.setHook(hook, hook)
	resp, err := luna.Request("http://www.ipip.net/ip.html", "GET", reqOpt)
	fmt.Println(resp, err)

	// ======= example 2
	reqOpt := luna.NewReqOptions()
	reqOpt.Hook = luna.Hook{
		ReqHook:  new(Hook),
		RespHook: new(Hook),
	}
	resp, err := luna.Request("http://www.ipip.net/ip.html", "GET", reqOpt)
	fmt.Println(resp, err)

}
