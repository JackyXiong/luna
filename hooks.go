package luna

import (
	"net/http"
)

// hook 请求和响应
// BeforeRequestHook 在请求发送前，对请求hook
// AfterRequestHook 获取响应后对响应hook
type Hook interface {
	BeforeRequestHook(req *http.Request) (err error)
	AfterRequestHook(resp *http.Response) (err error)
}

func applyBeforeHooks(req *http.Request, reqOpt *ReqOptions) error {
	if reqOpt.Hook != nil {
		err := reqOpt.Hook.BeforeRequestHook(req)
		return err
	}
	return nil
}

func applyAfterHooks(resp *http.Response, reqOpt *ReqOptions) error {
	if reqOpt.Hook != nil {
		err := reqOpt.Hook.AfterRequestHook(resp)
		return err
	}
	return nil
}
