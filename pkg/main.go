package main

import (
	"job-monitor/pkg/message"
	"job-monitor/pkg/spark"
	"job-monitor/pkg/storage"
	"net/http"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/etc/kubernetes/kubectl.kubeconfig")
	if err != nil {
		panic(err)
	}
	clientset, err := spark.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	queue := message.NewQueue("local")
	_, controller := spark.NewSparkApplicationInformer(clientset, queue)
	spark.AddToScheme(scheme.Scheme)
	storage.NewStorage("local", queue)
	go queue.Run()
	go controller.Run(wait.NeverStop)
	http.ListenAndServe(":8080", nil)
}
