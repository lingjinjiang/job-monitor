package spark

import (
	"fmt"
	"job-monitor/pkg/message"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnowntype)
	AddToScheme   = schemeBuilder.AddToScheme
)

func addKnowntype(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		schema.GroupVersion{Group: GROUP_NAME, Version: GROUP_VERSION},
		&SparkApplication{},
		&SparkApplicationList{})
	metav1.AddToGroupVersion(scheme,
		schema.GroupVersion{Group: GROUP_NAME, Version: GROUP_VERSION})
	return nil
}

func NewSparkApplicationInformer(clientset SparkApplicationV1Beta2Interface, queue message.Queue) (cache.Store, cache.Controller) {
	return cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
				return clientset.SparkApplications("").List(opts)
			},
			WatchFunc: func(opts metav1.ListOptions) (watch.Interface, error) {
				return clientset.SparkApplications("").Watch(opts)
			},
		},
		&SparkApplication{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    addSparkApplication,
			UpdateFunc: updateSparkApplication,
			DeleteFunc: deleteSparkApplication,
		},
	)
}

func addSparkApplication(obj interface{}) {
	app := obj.(*SparkApplication)
	fmt.Println("add", app.Namespace, app.Name)
}

func updateSparkApplication(oldObj interface{}, newObj interface{}) {
	app := newObj.(*SparkApplication)
	fmt.Println("update", app.Namespace, app.Name)
}

func deleteSparkApplication(obj interface{}) {
	app := obj.(*SparkApplication)
	fmt.Println("delete", app.Namespace, app.Name)
}
