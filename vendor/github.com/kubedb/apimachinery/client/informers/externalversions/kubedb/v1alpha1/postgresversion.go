/*
Copyright 2018 The KubeDB Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubedb_v1alpha1 "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	versioned "github.com/kubedb/apimachinery/client/clientset/versioned"
	internalinterfaces "github.com/kubedb/apimachinery/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubedb/apimachinery/client/listers/kubedb/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PostgresVersionInformer provides access to a shared informer and lister for
// PostgresVersions.
type PostgresVersionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PostgresVersionLister
}

type postgresVersionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPostgresVersionInformer constructs a new informer for PostgresVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPostgresVersionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPostgresVersionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPostgresVersionInformer constructs a new informer for PostgresVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPostgresVersionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubedbV1alpha1().PostgresVersions().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubedbV1alpha1().PostgresVersions().Watch(options)
			},
		},
		&kubedb_v1alpha1.PostgresVersion{},
		resyncPeriod,
		indexers,
	)
}

func (f *postgresVersionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPostgresVersionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *postgresVersionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubedb_v1alpha1.PostgresVersion{}, f.defaultInformer)
}

func (f *postgresVersionInformer) Lister() v1alpha1.PostgresVersionLister {
	return v1alpha1.NewPostgresVersionLister(f.Informer().GetIndexer())
}
