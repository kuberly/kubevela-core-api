/*
Copyright 2023 The KubeVela Authors.

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
	v1beta1 "github.com/kuberly/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TraitDefinitionLister helps list TraitDefinitions.
// All objects returned here must be treated as read-only.
type TraitDefinitionLister interface {
	// List lists all TraitDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.TraitDefinition, err error)
	// TraitDefinitions returns an object that can list and get TraitDefinitions.
	TraitDefinitions(namespace string) TraitDefinitionNamespaceLister
	TraitDefinitionListerExpansion
}

// traitDefinitionLister implements the TraitDefinitionLister interface.
type traitDefinitionLister struct {
	indexer cache.Indexer
}

// NewTraitDefinitionLister returns a new TraitDefinitionLister.
func NewTraitDefinitionLister(indexer cache.Indexer) TraitDefinitionLister {
	return &traitDefinitionLister{indexer: indexer}
}

// List lists all TraitDefinitions in the indexer.
func (s *traitDefinitionLister) List(selector labels.Selector) (ret []*v1beta1.TraitDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.TraitDefinition))
	})
	return ret, err
}

// TraitDefinitions returns an object that can list and get TraitDefinitions.
func (s *traitDefinitionLister) TraitDefinitions(namespace string) TraitDefinitionNamespaceLister {
	return traitDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TraitDefinitionNamespaceLister helps list and get TraitDefinitions.
// All objects returned here must be treated as read-only.
type TraitDefinitionNamespaceLister interface {
	// List lists all TraitDefinitions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.TraitDefinition, err error)
	// Get retrieves the TraitDefinition from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.TraitDefinition, error)
	TraitDefinitionNamespaceListerExpansion
}

// traitDefinitionNamespaceLister implements the TraitDefinitionNamespaceLister
// interface.
type traitDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TraitDefinitions in the indexer for a given namespace.
func (s traitDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.TraitDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.TraitDefinition))
	})
	return ret, err
}

// Get retrieves the TraitDefinition from the indexer for a given namespace and name.
func (s traitDefinitionNamespaceLister) Get(name string) (*v1beta1.TraitDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("traitdefinition"), name)
	}
	return obj.(*v1beta1.TraitDefinition), nil
}
