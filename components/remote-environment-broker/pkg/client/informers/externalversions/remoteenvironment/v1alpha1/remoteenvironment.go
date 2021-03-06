// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	remoteenvironment_v1alpha1 "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/apis/remoteenvironment/v1alpha1"
	versioned "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/client/listers/remoteenvironment/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RemoteEnvironmentInformer provides access to a shared informer and lister for
// RemoteEnvironments.
type RemoteEnvironmentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RemoteEnvironmentLister
}

type remoteEnvironmentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRemoteEnvironmentInformer constructs a new informer for RemoteEnvironment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRemoteEnvironmentInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRemoteEnvironmentInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredRemoteEnvironmentInformer constructs a new informer for RemoteEnvironment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRemoteEnvironmentInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RemoteenvironmentV1alpha1().RemoteEnvironments().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RemoteenvironmentV1alpha1().RemoteEnvironments().Watch(options)
			},
		},
		&remoteenvironment_v1alpha1.RemoteEnvironment{},
		resyncPeriod,
		indexers,
	)
}

func (f *remoteEnvironmentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRemoteEnvironmentInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *remoteEnvironmentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&remoteenvironment_v1alpha1.RemoteEnvironment{}, f.defaultInformer)
}

func (f *remoteEnvironmentInformer) Lister() v1alpha1.RemoteEnvironmentLister {
	return v1alpha1.NewRemoteEnvironmentLister(f.Informer().GetIndexer())
}
