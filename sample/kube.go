package main

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/k0kubun/pp"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}
	for _, namespace := range []string{"name1", "name2"} {
		for _, serviceName := range []string{"mysql"} {
			s, _ := client.CoreV1().Services(namespace).Get(serviceName, meta_v1.GetOptions{})
			pp.Println(s.Spec.Ports[0].NodePort)
		}
	}
}

func newClient() (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(kubeConfig)
}
