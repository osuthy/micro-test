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
	return Request{Json: i.ToJson()}
}

func (this Request) WithJson(i Object) Request {
	this.Json = i.ToJson()
	return this
}

func WithParam(name string, value string) Request {
	return Request{Params: []*Param{&Param{Name: name, Value: value}}}
}

func (this Request) WithParam(name string, value string) Request {
	this.Params = append(this.Params, &Param{Name: name, Value: value})
	return this
}


