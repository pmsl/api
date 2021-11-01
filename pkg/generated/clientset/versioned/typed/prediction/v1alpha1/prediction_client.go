// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/gocrane-io/api/pkg/generated/clientset/versioned/scheme"
	v1alpha1 "github.com/gocrane-io/api/prediction/v1alpha1"
	rest "k8s.io/client-go/rest"
)

type PredictionV1alpha1Interface interface {
	RESTClient() rest.Interface
	NodePredictionsGetter
	PodGroupPredictionsGetter
}

// PredictionV1alpha1Client is used to interact with features provided by the prediction.crane.io group.
type PredictionV1alpha1Client struct {
	restClient rest.Interface
}

func (c *PredictionV1alpha1Client) NodePredictions(namespace string) NodePredictionInterface {
	return newNodePredictions(c, namespace)
}

func (c *PredictionV1alpha1Client) PodGroupPredictions(namespace string) PodGroupPredictionInterface {
	return newPodGroupPredictions(c, namespace)
}

// NewForConfig creates a new PredictionV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*PredictionV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &PredictionV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new PredictionV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *PredictionV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new PredictionV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *PredictionV1alpha1Client {
	return &PredictionV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *PredictionV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}