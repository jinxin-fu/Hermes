/**
 * Created with IntelliJ goland.
 * @Auther: jinxin
 * @Date: 2022/04/12/19:10
 * @Description:
 */
package querier

import (
	promonitv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SubscriberAnnotationKey = "subscriber/endpointed"
)

var PrometheusRuleTemplate string = `
apiVersion: monitoring.coreos.com/v1
kind: PrometheusRule
metadata:
  annotations:
    meta.helm.sh/release-name: prometheus
    meta.helm.sh/release-namespace: default
    prometheus-operator-validated: "true"
  labels:
    app: kube-prometheus-stack
    app.kubernetes.io/instance: prometheus
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/part-of: kube-prometheus-stack
    app.kubernetes.io/version: 0.1.0
    chart: kube-prometheus-stack-0.1.0
    heritage: Helm
    release: prometheus
  name: subscriber2.rules
  namespace: hypermonitor
spec:
  groups:
`

type PromruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []promonitv1.PrometheusRule `json:"items"`
}
