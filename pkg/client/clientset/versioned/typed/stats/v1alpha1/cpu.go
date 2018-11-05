/*
Copyright 2018 interma.

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
	v1alpha1 "github.com/interma/programming-k8s/pkg/apis/stats/v1alpha1"
	scheme "github.com/interma/programming-k8s/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CpusGetter has a method to return a CpuInterface.
// A group's client should implement this interface.
type CpusGetter interface {
	Cpus(namespace string) CpuInterface
}

// CpuInterface has methods to work with Cpu resources.
type CpuInterface interface {
	Create(*v1alpha1.Cpu) (*v1alpha1.Cpu, error)
	Update(*v1alpha1.Cpu) (*v1alpha1.Cpu, error)
	UpdateStatus(*v1alpha1.Cpu) (*v1alpha1.Cpu, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Cpu, error)
	List(opts v1.ListOptions) (*v1alpha1.CpuList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Cpu, err error)
	CpuExpansion
}

// cpus implements CpuInterface
type cpus struct {
	client rest.Interface
	ns     string
}

// newCpus returns a Cpus
func newCpus(c *StatsV1alpha1Client, namespace string) *cpus {
	return &cpus{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cpu, and returns the corresponding cpu object, and an error if there is any.
func (c *cpus) Get(name string, options v1.GetOptions) (result *v1alpha1.Cpu, err error) {
	result = &v1alpha1.Cpu{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cpus").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Cpus that match those selectors.
func (c *cpus) List(opts v1.ListOptions) (result *v1alpha1.CpuList, err error) {
	result = &v1alpha1.CpuList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cpus").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cpus.
func (c *cpus) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cpus").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a cpu and creates it.  Returns the server's representation of the cpu, and an error, if there is any.
func (c *cpus) Create(cpu *v1alpha1.Cpu) (result *v1alpha1.Cpu, err error) {
	result = &v1alpha1.Cpu{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cpus").
		Body(cpu).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cpu and updates it. Returns the server's representation of the cpu, and an error, if there is any.
func (c *cpus) Update(cpu *v1alpha1.Cpu) (result *v1alpha1.Cpu, err error) {
	result = &v1alpha1.Cpu{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cpus").
		Name(cpu.Name).
		Body(cpu).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cpus) UpdateStatus(cpu *v1alpha1.Cpu) (result *v1alpha1.Cpu, err error) {
	result = &v1alpha1.Cpu{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cpus").
		Name(cpu.Name).
		SubResource("status").
		Body(cpu).
		Do().
		Into(result)
	return
}

// Delete takes name of the cpu and deletes it. Returns an error if one occurs.
func (c *cpus) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cpus").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cpus) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cpus").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cpu.
func (c *cpus) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Cpu, err error) {
	result = &v1alpha1.Cpu{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cpus").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}