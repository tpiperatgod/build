// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BuildStrategyLister helps list BuildStrategies.
type BuildStrategyLister interface {
	// List lists all BuildStrategies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BuildStrategy, err error)
	// BuildStrategies returns an object that can list and get BuildStrategies.
	BuildStrategies(namespace string) BuildStrategyNamespaceLister
	BuildStrategyListerExpansion
}

// buildStrategyLister implements the BuildStrategyLister interface.
type buildStrategyLister struct {
	indexer cache.Indexer
}

// NewBuildStrategyLister returns a new BuildStrategyLister.
func NewBuildStrategyLister(indexer cache.Indexer) BuildStrategyLister {
	return &buildStrategyLister{indexer: indexer}
}

// List lists all BuildStrategies in the indexer.
func (s *buildStrategyLister) List(selector labels.Selector) (ret []*v1alpha1.BuildStrategy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BuildStrategy))
	})
	return ret, err
}

// BuildStrategies returns an object that can list and get BuildStrategies.
func (s *buildStrategyLister) BuildStrategies(namespace string) BuildStrategyNamespaceLister {
	return buildStrategyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BuildStrategyNamespaceLister helps list and get BuildStrategies.
type BuildStrategyNamespaceLister interface {
	// List lists all BuildStrategies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BuildStrategy, err error)
	// Get retrieves the BuildStrategy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BuildStrategy, error)
	BuildStrategyNamespaceListerExpansion
}

// buildStrategyNamespaceLister implements the BuildStrategyNamespaceLister
// interface.
type buildStrategyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BuildStrategies in the indexer for a given namespace.
func (s buildStrategyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BuildStrategy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BuildStrategy))
	})
	return ret, err
}

// Get retrieves the BuildStrategy from the indexer for a given namespace and name.
func (s buildStrategyNamespaceLister) Get(name string) (*v1alpha1.BuildStrategy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("buildstrategy"), name)
	}
	return obj.(*v1alpha1.BuildStrategy), nil
}
