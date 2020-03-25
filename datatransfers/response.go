package datatransfers

import (
	"path"
	"runtime"
	"strings"

	"github.com/astaxie/beego"
)

type ApiResponse struct {
	Data  interface{}       `json:"data"`
	Error *ApiResponseError `json:"error,omitempty"`
}

type ApiResponseError struct {
	Code     int                          `json:"code"`
	Message  string                       `json:"message"`
	CallInfo *callInfo                    `json:"callInfo"`
	Errors   []ApiResponseAdditionalError `json:"errors"`
}

type ApiResponseAdditionalError struct {
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func NewSuccessResponse(data interface{}) *ApiResponse {
	return &ApiResponse{Data: data}
}

func NewErrorResponse(code int, message string) *ApiResponse {
	response := &ApiResponse{nil, &ApiResponseError{code, message, retrieveCallInfo(), make([]ApiResponseAdditionalError, 0)}}
	if beego.AppConfig.String("runmode") == "dev" {
		response.Error.CallInfo = retrieveCallInfo()
	}

	return response
}

func (r *ApiResponse) AddError(domain string, reason string, message string) {
	r.Error.Errors = append(r.Error.Errors, ApiResponseAdditionalError{domain, reason, message})
}

type callInfo struct {
	PackageName string `json:"packageName"`
	FileName    string `json:"fileName"`
	FuncName    string `json:"funcName"`
	Line        int    `json:"line"`
}

// source https://stackoverflow.com/a/25265493/11337921
func retrieveCallInfo() *callInfo {
	pc, file, line, _ := runtime.Caller(2)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &callInfo{packageName, fileName, funcName, line}
}
