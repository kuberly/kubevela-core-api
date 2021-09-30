/*
Copyright 2021 The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PolicyDefinitionLister helps list PolicyDefinitions.
// All objects returned here must be treated as read-only.
type PolicyDefinitionLister interface {
	// List lists all PolicyDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.PolicyDefinition, err error)
	// PolicyDefinitions returns an object that can list and get PolicyDefinitions.
	PolicyDefinitions(namespace string) PolicyDefinitionNamespaceLister
	PolicyDefinitionListerExpansion
}

// policyDefinitionLister implements the PolicyDefinitionLister interface.
type policyDefinitionLister struct {
	indexer cache.Indexer
}

// NewPolicyDefinitionLister returns a new PolicyDefinitionLister.
func NewPolicyDefinitionLister(indexer cache.Indexer) PolicyDefinitionLister {
	return &policyDefinitionLister{indexer: indexer}
}

// List lists all PolicyDefinitions in the indexer.
func (s *policyDefinitionLister) List(selector labels.Selector) (ret []*v1beta1.PolicyDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PolicyDefinition))
	})
	return ret, err
}

// PolicyDefinitions returns an object that can list and get PolicyDefinitions.
func (s *policyDefinitionLister) PolicyDefinitions(namespace string) PolicyDefinitionNamespaceLister {
	return policyDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PolicyDefinitionNamespaceLister helps list and get PolicyDefinitions.
// All objects returned here must be treated as read-only.
type PolicyDefinitionNamespaceLister interface {
	// List lists all PolicyDefinitions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.PolicyDefinition, err error)
	// Get retrieves the PolicyDefinition from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.PolicyDefinition, error)
	PolicyDefinitionNamespaceListerExpansion
}

// policyDefinitionNamespaceLister implements the PolicyDefinitionNamespaceLister
// interface.
type policyDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PolicyDefinitions in the indexer for a given namespace.
func (s policyDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.PolicyDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PolicyDefinition))
	})
	return ret, err
}

// Get retrieves the PolicyDefinition from the indexer for a given namespace and name.
func (s policyDefinitionNamespaceLister) Get(name string) (*v1beta1.PolicyDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("policydefinition"), name)
	}
	return obj.(*v1beta1.PolicyDefinition), nil
}