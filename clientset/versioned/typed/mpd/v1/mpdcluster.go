/* 
*Copyright (c) 2019-2021, Alibaba Group Holding Limited;
*Licensed under the Apache License, Version 2.0 (the "License");
*you may not use this file except in compliance with the License.
*You may obtain a copy of the License at

*   http://www.apache.org/licenses/LICENSE-2.0

*Unless required by applicable law or agreed to in writing, software
*distributed under the License is distributed on an "AS IS" BASIS,
*WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*See the License for the specific language governing permissions and
*limitations under the License.
 */


/*


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

package v1

import (
	"time"

	v1 "gitlab.alibaba-inc.com/polar-as/polar-mpd-controller/apis/mpd/v1"
	scheme "gitlab.alibaba-inc.com/polar-as/polar-mpd-controller/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MPDClustersGetter has a method to return a MPDClusterInterface.
// A group's client should implement this interface.
type MPDClustersGetter interface {
	MPDClusters(namespace string) MPDClusterInterface
}

// MPDClusterInterface has methods to work with MPDCluster resources.
type MPDClusterInterface interface {
	Create(*v1.MPDCluster) (*v1.MPDCluster, error)
	Update(*v1.MPDCluster) (*v1.MPDCluster, error)
	UpdateStatus(*v1.MPDCluster) (*v1.MPDCluster, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MPDCluster, error)
	List(opts metav1.ListOptions) (*v1.MPDClusterList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MPDCluster, err error)
	MPDClusterExpansion
}

// mPDClusters implements MPDClusterInterface
type mPDClusters struct {
	client rest.Interface
	ns     string
}

// newMPDClusters returns a MPDClusters
func newMPDClusters(c *MpdV1Client, namespace string) *mPDClusters {
	return &mPDClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mPDCluster, and returns the corresponding mPDCluster object, and an error if there is any.
func (c *mPDClusters) Get(name string, options metav1.GetOptions) (result *v1.MPDCluster, err error) {
	result = &v1.MPDCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mpdclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MPDClusters that match those selectors.
func (c *mPDClusters) List(opts metav1.ListOptions) (result *v1.MPDClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MPDClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mpdclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mPDClusters.
func (c *mPDClusters) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mpdclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mPDCluster and creates it.  Returns the server's representation of the mPDCluster, and an error, if there is any.
func (c *mPDClusters) Create(mPDCluster *v1.MPDCluster) (result *v1.MPDCluster, err error) {
	result = &v1.MPDCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mpdclusters").
		Body(mPDCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mPDCluster and updates it. Returns the server's representation of the mPDCluster, and an error, if there is any.
func (c *mPDClusters) Update(mPDCluster *v1.MPDCluster) (result *v1.MPDCluster, err error) {
	result = &v1.MPDCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mpdclusters").
		Name(mPDCluster.Name).
		Body(mPDCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mPDClusters) UpdateStatus(mPDCluster *v1.MPDCluster) (result *v1.MPDCluster, err error) {
	result = &v1.MPDCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mpdclusters").
		Name(mPDCluster.Name).
		SubResource("status").
		Body(mPDCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the mPDCluster and deletes it. Returns an error if one occurs.
func (c *mPDClusters) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mpdclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mPDClusters) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mpdclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mPDCluster.
func (c *mPDClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MPDCluster, err error) {
	result = &v1.MPDCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mpdclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
