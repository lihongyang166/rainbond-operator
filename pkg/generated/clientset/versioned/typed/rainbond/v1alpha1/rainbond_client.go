// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/GLYASAI/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	"github.com/GLYASAI/rainbond-operator/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type RainbondV1alpha1Interface interface {
	RESTClient() rest.Interface
	RainbondClustersGetter
	RainbondPackagesGetter
	RbdComponentsGetter
}

// RainbondV1alpha1Client is used to interact with features provided by the rainbond.io group.
type RainbondV1alpha1Client struct {
	restClient rest.Interface
}

func (c *RainbondV1alpha1Client) RainbondClusters(namespace string) RainbondClusterInterface {
	return newRainbondClusters(c, namespace)
}

func (c *RainbondV1alpha1Client) RainbondPackages(namespace string) RainbondPackageInterface {
	return newRainbondPackages(c, namespace)
}

func (c *RainbondV1alpha1Client) RbdComponents(namespace string) RbdComponentInterface {
	return newRbdComponents(c, namespace)
}

// NewForConfig creates a new RainbondV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*RainbondV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &RainbondV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new RainbondV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *RainbondV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new RainbondV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *RainbondV1alpha1Client {
	return &RainbondV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *RainbondV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
