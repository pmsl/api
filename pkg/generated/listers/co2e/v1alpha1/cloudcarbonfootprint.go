// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gocrane/api/co2e/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloudCarbonFootprintLister helps list CloudCarbonFootprints.
// All objects returned here must be treated as read-only.
type CloudCarbonFootprintLister interface {
	// List lists all CloudCarbonFootprints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudCarbonFootprint, err error)
	// Get retrieves the CloudCarbonFootprint from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CloudCarbonFootprint, error)
	CloudCarbonFootprintListerExpansion
}

// cloudCarbonFootprintLister implements the CloudCarbonFootprintLister interface.
type cloudCarbonFootprintLister struct {
	indexer cache.Indexer
}

// NewCloudCarbonFootprintLister returns a new CloudCarbonFootprintLister.
func NewCloudCarbonFootprintLister(indexer cache.Indexer) CloudCarbonFootprintLister {
	return &cloudCarbonFootprintLister{indexer: indexer}
}

// List lists all CloudCarbonFootprints in the indexer.
func (s *cloudCarbonFootprintLister) List(selector labels.Selector) (ret []*v1alpha1.CloudCarbonFootprint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudCarbonFootprint))
	})
	return ret, err
}

// Get retrieves the CloudCarbonFootprint from the index for a given name.
func (s *cloudCarbonFootprintLister) Get(name string) (*v1alpha1.CloudCarbonFootprint, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudcarbonfootprint"), name)
	}
	return obj.(*v1alpha1.CloudCarbonFootprint), nil
}