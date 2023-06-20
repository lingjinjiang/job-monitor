package spark

import (
	"job-monitor/pkg/api"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const GROUP_NAME = "sparkoperator.k8s.io"

const GROUP_VERSION = "v1beta2"

type SparkApplicationSpec struct{}

type SparkApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metatdata"`
	Spec              SparkApplicationSpec `json:"spec"`
}

type SparkApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metatdata"`
	Items           []SparkApplication `json:"items"`
}

func (in *SparkApplication) DeepCopyInto(out *SparkApplication) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = in.Spec
}

func (in *SparkApplication) DeepCopyObject() runtime.Object {
	out := SparkApplication{}
	in.DeepCopyInto(&out)
	return &out
}

func (in *SparkApplicationList) DeepCopyObject() runtime.Object {
	out := SparkApplicationList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		out.Items = make([]SparkApplication, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
	return &out
}

func (s *SparkApplication) Convert() api.Job {
	return api.Job{}
}
