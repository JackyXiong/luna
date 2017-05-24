package luna

import (
	"net/http"
)

type RequestHook interface {
	BeforeRequestHook(req *http.Request) (err error)
}

type ResponseHook interface {
	AfterResponseHook(resp *http.Response) (err error)
}

type Hook struct {
	ReqHook  RequestHook  // hook before send request
	RespHook ResponseHook // hook after get reesponse
}

func applyReqHook(req *http.Request, reqOpt *ReqOptions) error {
	if &reqOpt.Hook != nil && reqOpt.Hook.ReqHook != nil {
		return reqOpt.Hook.ReqHook.BeforeRequestHook(req)
	}
	return nil
}

func applyRespHook(resp *http.Response, reqOpt *ReqOptions) error {
	if &reqOpt.Hook != nil && reqOpt.Hook.RespHook != nil {
		return reqOpt.Hook.RespHook.AfterResponseHook(resp)
	}
	return nil

}

// convenient function to set hook
func (reqOpt *ReqOptions) SetHook(reqHook RequestHook, respHook ResponseHook) (err error) {
	reqOpt.Hook = Hook{
		ReqHook:  reqHook,
		RespHook: respHook,
	}
	return nil
}
