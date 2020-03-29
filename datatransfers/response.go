package datatransfers

import (
	"encoding/json"
	"path"
	"runtime"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type ResponseBuilder interface {
	SetData(data interface{}) *responseBuild
	SetError(code int, message string) *responseBuild
	AddAdditionalError(domain string, reason string, message string) *responseBuild
	ServeJSON()
}

type responseBuild struct {
	responseContext *context.Response
	response        apiResponse
}

type apiResponse struct {
	Data  interface{}       `json:"data"`
	Error *apiResponseError `json:"error,omitempty"`
}

type apiResponseError struct {
	Code     int                          `json:"code"`
	Message  string                       `json:"message"`
	CallInfo *callInfo                    `json:"callInfo,omitempty"`
	Errors   []apiResponseAdditionalError `json:"errors"`
}

type apiResponseAdditionalError struct {
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type callInfo struct {
	PackageName string `json:"packageName"`
	FileName    string `json:"fileName"`
	FuncName    string `json:"funcName"`
	Line        int    `json:"line"`
}

func NewResponseBuilder(responseContext *context.Response) *responseBuild {
	r := &responseBuild{responseContext, apiResponse{}}
	r.responseContext.Header().Set("Content-Type", "application/json")

	return r
}

func (r *responseBuild) SetData(data interface{}) *responseBuild {
	r.response.Data = data

	return r
}

func (r *responseBuild) SetError(code int, message string) *responseBuild {
	r.responseContext.WriteHeader(code)
	r.response.Error = NewErrorResponse(code, message)

	return r
}

func (r *responseBuild) AddAdditionalError(domain string, reason string, message string) *responseBuild {
	r.response.Error.Errors = append(r.response.Error.Errors, apiResponseAdditionalError{domain, reason, message})

	return r
}

func (r *responseBuild) ServeJSON() {
	responseBody, _ := json.Marshal(r.response)
	r.responseContext.Write(responseBody)
}

func NewErrorResponse(code int, message string) *apiResponseError {
	response := &apiResponseError{code, message, nil, make([]apiResponseAdditionalError, 0)}

	if beego.BConfig.RunMode == "dev" {
		response.CallInfo = retrieveCallInfo()
	}

	return response
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
