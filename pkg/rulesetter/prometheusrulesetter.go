/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/04/12/19:14
 * @Description:
 */
package querier

import (
	"fmt"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

const kubeconfigPath = "C:\\Users\\Lenovo\\.kube"

var clientGlobal dynamic.Interface

type RuleSetter struct {
	client dynamic.Interface
}

func tes() {
	var rs RuleSetter
	err := rs.NewRuleSetter(kubeconfigPath, "config")
	if err != nil {
		panic(err)
	}
	rs.CreatePrometheusRule("hypermonitor", PrometheusRuleTemplate)

}

func (r *RuleSetter) NewRuleSetter(path, name string) error {
	kubeconfig := filepath.Join(path, "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	r.client = client
	return nil
}

func (r *RuleSetter) CreatePrometheusRule(namespace string, yamlData string) error {
	if r.client == nil {
		return fmt.Errorf("rulesetter need to init first\n")
	}
	err := CreatPr(r.client, namespace, yamlData)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *RuleSetter) DeletePrometheusRule(namespace string, ruleName string) error {
	if r.client == nil {
		return fmt.Errorf("rulesetter need to init first\n")
	}
	err := deletePromrule(r.client, namespace, ruleName)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *RuleSetter) UpdatePrometheusRule(namespace string, ruleName string) error {
	if r.client == nil {
		return fmt.Errorf("rulesetter need to init first\n")
	}
	err := UpdateCr(r.client, namespace, ruleName)
	if err != nil {
		panic(err)
	}
	return nil
}
