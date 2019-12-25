package db

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/db/infra"
)

type RDBDefinition struct {
	config C
}

func (this *RDBDefinition) SetConnectionForLocal(tc TC) TC {
	driver := this.config["driver"].(string)
	localConfig := this.config["local"].(C)
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		localConfig["user"].(string),
		localConfig["password"].(string),
		localConfig["host"].(string),
		localConfig["port"].(string),
		this.config["database"].(string),
	)
	c, _ := sql.Open(driver, source)
	tc[this.config["name"].(string)] = NewConnection(c, driver)
	return tc
}

func (this *RDBDefinition) SetConnectionForK8S(tc TC, namespace string) TC {
	driver := this.config["driver"].(string)
	kubeConfig, _ := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	client, _ := kubernetes.NewForConfig(kubeConfig)
	s, _ := client.CoreV1().Services(namespace).Get(this.config["name"].(string), meta_v1.GetOptions{})
	u, _ := url.Parse(kubeConfig.Host)
	k8sConfig := this.config["k8s"].(C)
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		k8sConfig["user"].(string),
		k8sConfig["password"].(string),
		u.Hostname(),
		s.Spec.Ports[0].NodePort,
		this.config["database"].(string),
	)
	c, _ := sql.Open(driver, source)
	tc[this.config["name"].(string)] = NewConnection(c, driver)
	return tc
}

func DefineRDB(config C) {
	AppendConnectionDefinable(&RDBDefinition{
		config: config,
	})
}
