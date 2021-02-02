// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/cluster/v1alpha1"
	"github.com/karmada-io/karmada/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ClusterV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClustersGetter
}

// ClusterV1alpha1Client is used to interact with features provided by the cluster.karmada.io group.
type ClusterV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ClusterV1alpha1Client) Clusters() ClusterInterface {
	return newClusters(c)
}

// NewForConfig creates a new ClusterV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ClusterV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ClusterV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ClusterV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ClusterV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ClusterV1alpha1Client {
	return &ClusterV1alpha1Client{c}
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
func (c *ClusterV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}