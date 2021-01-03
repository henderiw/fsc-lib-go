/*
Copyright The Kubernetes Authors.

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

package v1

import (
	internalinterfaces "github.com/fsc-demo-wim/fsc-lib-go/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// KubeControllersConfigurations returns a KubeControllersConfigurationInformer.
	KubeControllersConfigurations() KubeControllersConfigurationInformer
	// NodeTopologies returns a NodeTopologyInformer.
	NodeTopologies() NodeTopologyInformer
	// WorkLoads returns a WorkLoadInformer.
	WorkLoads() WorkLoadInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// KubeControllersConfigurations returns a KubeControllersConfigurationInformer.
func (v *version) KubeControllersConfigurations() KubeControllersConfigurationInformer {
	return &kubeControllersConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NodeTopologies returns a NodeTopologyInformer.
func (v *version) NodeTopologies() NodeTopologyInformer {
	return &nodeTopologyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WorkLoads returns a WorkLoadInformer.
func (v *version) WorkLoads() WorkLoadInformer {
	return &workLoadInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
