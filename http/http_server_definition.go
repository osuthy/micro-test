package http

import (
	"fmt"
	. "github.com/osuthy/micro-test"
)

type HttpServerDefinition struct {
	config C
}

func DefineHttpServer(config C) {
	AppendConnectionDefinable(&HttpServerDefinition{
		config: config,
	})
}

func (this *HttpServerDefinition) SetConnectionForK8S(tc TC, namespace string) TC {
	return tc
}

func (this *HttpServerDefinition) SetConnectionForLocal(tc TC) TC {
	localConfig := this.config["local"].(C)
	url := fmt.Sprintf("http://%s:%d", localConfig["host"].(string), localConfig["port"].(int))
	tc[this.config["name"].(string)] = &Client{url: url}
	return tc
}
