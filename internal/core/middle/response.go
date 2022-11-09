package middle

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

type BaseResponse struct {
	Code    string      `json:"code" dc:"code"`
	Type    string      `json:"type" dc:"type"`
	Message string      `json:"message" dc:"message"`
	Data    interface{} `json:"data" dc:"data"`
}

func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		msg         = "操作成功！"
		err         = r.GetError()
		res         = r.GetHandlerResponse()
		resultModel = BaseResponse{}
	)
	if err != nil {

		msg = err.Error()
		resultModel.Code = "error"
		resultModel.Type = "error"
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			resultModel.Code = "error"
		case http.StatusForbidden:
			resultModel.Code = "error"
		}
		resultModel.Type = "warning"
	} else {
		resultModel.Type = "success"
		resultModel.Code = "success"
	}

	resultModel.Message = msg
	resultModel.Data = res
	r.Response.WriteJson(resultModel)
}
