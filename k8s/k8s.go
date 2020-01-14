package k8s

import (
	"net/url"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
)

type K8S struct {
	config *rest.Config
}

func (this *K8S) Port(namespace, nodePortName string) (int32, error) {
	client, err := kubernetes.NewForConfig(this.config)
	if err != nil {
		return 0, err
	}
	s, err := client.CoreV1().Services(namespace).Get(nodePortName, v1.GetOptions{})
	if err != nil {
		return 0, err
	}
	return s.Spec.Ports[0].NodePort, nil
}

func (this *K8S) Host() string {
	u, _ := url.Parse(this.config.Host)
	return u.Hostname()
}

func CreateK8S() (*K8S, error){
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}
	return &K8S{config: kubeConfig}, nil
}

