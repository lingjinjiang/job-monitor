package spark

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type SparkApplicationClientInterface interface {
	List(opts metav1.ListOptions) (*SparkApplicationList, error)
	Get(name string, opts metav1.GetOptions) (*SparkApplication, error)
	Create(*SparkApplication) (*SparkApplication, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type SparkApplicationClient struct {
	restClient rest.Interface
	namespace  string
}

func (c *SparkApplicationClient) List(opts metav1.ListOptions) (*SparkApplicationList, error) {
	result := SparkApplicationList{}
	err := c.restClient.Get().Namespace(c.namespace).
		Resource("sparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).Into(&result)
	return &result, err
}

func (c *SparkApplicationClient) Get(name string, opts metav1.GetOptions) (*SparkApplication, error) {
	result := SparkApplication{}
	err := c.restClient.Get().Namespace(c.namespace).
		Resource("sparkapplication").Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).Into(&result)
	return &result, err
}

func (c *SparkApplicationClient) Create(app *SparkApplication) (*SparkApplication, error) {
	result := SparkApplication{}
	err := c.restClient.Post().Namespace(c.namespace).
		Resource("sparkapplication").Body(app).
		Do(context.Background()).Into(&result)
	return &result, err
}

func (c *SparkApplicationClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.Get().Namespace(c.namespace).
		Resource("sparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}

type SparkApplicationV1Beta2Interface interface {
	SparkApplications(namespace string) SparkApplicationClientInterface
}

type SparkApplicationV1Beta2Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*SparkApplicationV1Beta2Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: GROUP_NAME, Version: GROUP_VERSION}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SparkApplicationV1Beta2Client{client}, err
}

func (c *SparkApplicationV1Beta2Client) SparkApplications(namespace string) SparkApplicationClientInterface {
	return &SparkApplicationClient{restClient: c.restClient, namespace: namespace}
}
