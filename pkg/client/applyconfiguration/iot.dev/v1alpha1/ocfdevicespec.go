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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// OCFDeviceSpecApplyConfiguration represents an declarative configuration of the OCFDeviceSpec type for use
// with apply.
type OCFDeviceSpecApplyConfiguration struct {
	Id            *string                          `json:"id,omitempty"`
	Name          *string                          `json:"name,omitempty"`
	Owned         *bool                            `json:"owned,omitempty"`
	OwnerID       *string                          `json:"ownerId,omitempty"`
	ResourceTypes []ResourceTypeApplyConfiguration `json:"resourceTypes,omitempty"`
}

// OCFDeviceSpecApplyConfiguration constructs an declarative configuration of the OCFDeviceSpec type for use with
// apply.
func OCFDeviceSpec() *OCFDeviceSpecApplyConfiguration {
	return &OCFDeviceSpecApplyConfiguration{}
}

// WithId sets the Id field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Id field is set to the value of the last call.
func (b *OCFDeviceSpecApplyConfiguration) WithId(value string) *OCFDeviceSpecApplyConfiguration {
	b.Id = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *OCFDeviceSpecApplyConfiguration) WithName(value string) *OCFDeviceSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithOwned sets the Owned field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Owned field is set to the value of the last call.
func (b *OCFDeviceSpecApplyConfiguration) WithOwned(value bool) *OCFDeviceSpecApplyConfiguration {
	b.Owned = &value
	return b
}

// WithOwnerID sets the OwnerID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OwnerID field is set to the value of the last call.
func (b *OCFDeviceSpecApplyConfiguration) WithOwnerID(value string) *OCFDeviceSpecApplyConfiguration {
	b.OwnerID = &value
	return b
}

// WithResourceTypes adds the given value to the ResourceTypes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourceTypes field.
func (b *OCFDeviceSpecApplyConfiguration) WithResourceTypes(values ...*ResourceTypeApplyConfiguration) *OCFDeviceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResourceTypes")
		}
		b.ResourceTypes = append(b.ResourceTypes, *values[i])
	}
	return b
}
