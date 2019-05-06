/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	v1alpha1 "k8s.io/autoscaler/hvpa-controller/pkg/apis/autoscaling/v1alpha1"
	"k8s.io/autoscaler/hvpa-controller/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type AutoscalingV1alpha1Interface interface {
	RESTClient() rest.Interface
	HvpasGetter
}

// AutoscalingV1alpha1Client is used to interact with features provided by the autoscaling.k8s.io group.
type AutoscalingV1alpha1Client struct {
	restClient rest.Interface
}

func (c *AutoscalingV1alpha1Client) Hvpas(namespace string) HvpaInterface {
	return newHvpas(c, namespace)
}

// NewForConfig creates a new AutoscalingV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*AutoscalingV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &AutoscalingV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new AutoscalingV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AutoscalingV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AutoscalingV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *AutoscalingV1alpha1Client {
	return &AutoscalingV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AutoscalingV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
