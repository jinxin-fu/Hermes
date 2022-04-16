/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/04/12/19:11
 * @Description:
 */
package querier

import (
	"context"
	"encoding/json"
	"fmt"
	promonitv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/dynamic"
)

var gvr = schema.GroupVersionResource{
	Group:    "monitoring.coreos.com",
	Version:  "v1",
	Resource: "prometheusrules",
}

func ListPromRules(clien dynamic.Interface, namespace string) (*PromruleList, error) {
	list, err := clien.Resource(gvr).Namespace(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	data, err := list.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var prList PromruleList
	if err := json.Unmarshal(data, &prList); err != nil {
		return nil, err
	}

	return &prList, nil
}

func GetPromRule(client dynamic.Interface, namespace string, name string) (*promonitv1.PrometheusRule, error) {
	utd, err := client.Resource(gvr).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	data, err := utd.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var pr promonitv1.PrometheusRule
	if err := json.Unmarshal(data, &pr); err != nil {
		return nil, err
	}
	return &pr, nil

}

func CreatePromruleWithYaml(clien dynamic.Interface, namespace string, yamlData string) (*promonitv1.PrometheusRule, error) {
	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	obj := &unstructured.Unstructured{}
	if _, _, err := decoder.Decode([]byte(yamlData), nil, obj); err != nil {
		return nil, err

	}

	ruleSpec := promonitv1.PrometheusRuleSpec{
		Groups: []promonitv1.RuleGroup{
			{
				Name: "test-demo.rules",
				Rules: []promonitv1.Rule{
					{
						Record: ":node_memory_MemAvailable_bytes:sum",
						Expr:   intstr.FromString("sum(node_memory_MemAvailable_bytes{job=\"node-exporter\"}) by (cluster)"),
						Labels: map[string]string{
							"testlabel": "label1",
						},
					},
					{
						Record: ":container_cpu_usage_seconds_total:sum",
						Expr:   intstr.FromString("sum(rate(container_cpu_usage_seconds_total{image!=\"\"}[1m])) by (pod_name, namespace)"),
						Labels: map[string]string{
							"testlabel": "label2",
						},
					},
					{
						Alert: "alert",
						Expr:  intstr.FromString("count by (namespace,service) (\n  count_values by (namespace,service) (\"config_hash\", alertmanager_config_hash{job=\"prometheus-kube-prometheus-alertmanager\",namespace=\"hypermonitor\"})\n)\n!= 1"),
						Labels: map[string]string{
							"alertlabel": "alertlabel",
							"alertname":  "TestAlert",
						},
						Annotations: map[string]string{
							"alertName":       "TestAlert",
							"receiverAddress": "http://ip+port",
							"aggerateRules":   "(count by (namespace,service) (changes(process_start_time_seconds{job=\"prometheus-kube-prometheus-alertmanager\",namespace=\"hypermonitor\"}[10m]) > 4)/ count by (namespace,service) (up{job=\"prometheus-kube-prometheus-alertmanager\",namespace=\"hypermonitor\"}))>= 0.5",
							"returnValueFlah": "true",
						},
					},
				},
			},
		},
	}

	obj.Object["spec"] = interface{}(ruleSpec)
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[SubscriberAnnotationKey] = "http://192.168.1.51:9857"

	labels := obj.GetLabels()
	if labels == nil {
		labels = make(map[string]string)
	}
	labels["label-test"] = "test1111"
	obj.SetLabels(labels)
	obj.SetAnnotations(annotations)
	obj.SetName("subscriber.rule")

	utd, err := clien.Resource(gvr).Namespace(namespace).Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	data, err := utd.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var pr promonitv1.PrometheusRule
	if err := json.Unmarshal(data, &pr); err != nil {
		return nil, err
	}

	return &pr, nil
}

func UpdatePromruleWithYaml(client dynamic.Interface, namespace string, yamlData string) (*promonitv1.PrometheusRule, error) {
	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	obj := &unstructured.Unstructured{}
	if _, _, err := decoder.Decode([]byte(yamlData), nil, obj); err != nil {
		return nil, err
	}
	utd, err := client.Resource(gvr).Namespace(namespace).Get(context.Background(), obj.GetName(), metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	obj.SetResourceVersion(utd.GetResourceVersion())
	utd, err = client.Resource(gvr).Namespace(namespace).Update(context.Background(), obj, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	data, err := utd.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var pr promonitv1.PrometheusRule
	if err := json.Unmarshal(data, &pr); err != nil {
		return nil, err
	}
	return &pr, nil
}

func PatchPromrule(client dynamic.Interface, namespace, name string, pt types.PatchType, data []byte) error {
	_, err := client.Resource(gvr).Namespace(namespace).Patch(context.Background(), name, pt, data, metav1.PatchOptions{})

	return err
}

func deletePromrule(client dynamic.Interface, namespace string, name string) error {
	return client.Resource(gvr).Namespace(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}

func CreatPr(client dynamic.Interface, namespace string, yamlData string) error {
	pr, err := CreatePromruleWithYaml(client, namespace, yamlData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s %s %V\n", pr.Namespace, pr.Name, pr.Spec)
	return nil
}

func UpdateCr(client dynamic.Interface, namespace string, yamlData string) error {
	//	updateData := `
	//apiVersion: "stable.example.com/v1"
	//kind: CronTab
	//metadata:
	//  name: cron-2
	//spec:
	//  cronSpec: "* * * * */15"
	//  image: my-awesome-cron-image-2-update122
	//`
	pr, err := UpdatePromruleWithYaml(client, namespace, yamlData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s %s %V\n", pr.Namespace, pr.Name, pr.Spec)
	return nil
}
