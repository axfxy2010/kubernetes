// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/template/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TemplateInstanceLister helps list TemplateInstances.
type TemplateInstanceLister interface {
	// List lists all TemplateInstances in the indexer.
	List(selector labels.Selector) (ret []*v1.TemplateInstance, err error)
	// TemplateInstances returns an object that can list and get TemplateInstances.
	TemplateInstances(namespace string) TemplateInstanceNamespaceLister
	TemplateInstanceListerExpansion
}

// templateInstanceLister implements the TemplateInstanceLister interface.
type templateInstanceLister struct {
	indexer cache.Indexer
}

// NewTemplateInstanceLister returns a new TemplateInstanceLister.
func NewTemplateInstanceLister(indexer cache.Indexer) TemplateInstanceLister {
	return &templateInstanceLister{indexer: indexer}
}

// List lists all TemplateInstances in the indexer.
func (s *templateInstanceLister) List(selector labels.Selector) (ret []*v1.TemplateInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TemplateInstance))
	})
	return ret, err
}

// TemplateInstances returns an object that can list and get TemplateInstances.
func (s *templateInstanceLister) TemplateInstances(namespace string) TemplateInstanceNamespaceLister {
	return templateInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TemplateInstanceNamespaceLister helps list and get TemplateInstances.
type TemplateInstanceNamespaceLister interface {
	// List lists all TemplateInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TemplateInstance, err error)
	// Get retrieves the TemplateInstance from the indexer for a given namespace and name.
	Get(name string) (*v1.TemplateInstance, error)
	TemplateInstanceNamespaceListerExpansion
}

// templateInstanceNamespaceLister implements the TemplateInstanceNamespaceLister
// interface.
type templateInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TemplateInstances in the indexer for a given namespace.
func (s templateInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1.TemplateInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TemplateInstance))
	})
	return ret, err
}

// Get retrieves the TemplateInstance from the indexer for a given namespace and name.
func (s templateInstanceNamespaceLister) Get(name string) (*v1.TemplateInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("templateinstance"), name)
	}
	return obj.(*v1.TemplateInstance), nil
}
