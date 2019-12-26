package http

import (
	"fmt"

	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/k8s"
)

type HttpServerDefinition struct {
	config C
}

func DefineHttpServer(config C) {
	AppendConnectionDefinable(&HttpServerDefinition{
		config: config,
	})
}

func (this *HttpServerDefinition) SetConnectionForLocal(tc TC) (TC, error) {
	localConfig := this.config["local"].(C)
	url := fmt.Sprintf("http://%s:%d", localConfig["host"].(string), localConfig["port"].(int))
	tc[this.config["name"].(string)] = &Client{Url: url}
	return tc, nil
}

func (this *HttpServerDefinition) SetConnectionForK8S(tc TC, namespace string) (TC, error) {
	k8s, _ := CreateK8S()
	k8sConfig := this.config["k8s"].(C)
	port, _ := k8s.Port(namespace, k8sConfig["svc"].(string))
	url := fmt.Sprintf("http://%s:%d", k8s.Host(), port)
	tc[this.config["name"].(string)] = &Client{Url: url}
	return tc, nil
}

