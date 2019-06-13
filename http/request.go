package http

import(
	. "github.com/ShoichiroKitano/micro_test/json"
	"fmt"
	"strings"
)

type Request struct {
	Json []byte
	Params []*Param
}

type Param struct {
	Name string
	Value string
}

func (this Param) ToParam() string {
	return fmt.Sprintf("%s=%s", this.Name, this.Value)
}

func (this Request) ToQueryParam() string {
	if len(this.Params)>0 {
		var builder strings.Builder
		for _, param := range this.Params {
			builder.WriteString(param.ToParam())
			builder.WriteString("&")
		}
		return builder.String()[:len(builder.String()) - 1]
	}
	return ""
}

func WithJson(i Object) Request {
	request := Request{Json: i.ToJson()}
	return request
}

func (this Request) WithJson(i Object) Request {
	this.Json = i.ToJson()
	return this
}

func WithParam(name string, value string) Request {
	param := &Param{Name: name, Value: value}
	request := Request{Params: []*Param{param}}
	return request
}

func (this Request) WithParam(name string, value string) Request {
	param := &Param{Name: name, Value: value}
	this.Params = append(this.Params, param)
	return this
}


