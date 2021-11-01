// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gocrane-io/api/pkg/generated/clientset/versioned/typed/prediction/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePredictionV1alpha1 struct {
	*testing.Fake
}

func (c *FakePredictionV1alpha1) NodePredictions(namespace string) v1alpha1.NodePredictionInterface {
	return &FakeNodePredictions{c, namespace}
}

func (c *FakePredictionV1alpha1) PodGroupPredictions(namespace string) v1alpha1.PodGroupPredictionInterface {
	return &FakePodGroupPredictions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePredictionV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}