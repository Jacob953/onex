// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// EnvVarSourceApplyConfiguration represents an declarative configuration of the EnvVarSource type for use
// with apply.
type EnvVarSourceApplyConfiguration struct {
	FieldRef         *ObjectFieldSelectorApplyConfiguration   `json:"fieldRef,omitempty"`
	ResourceFieldRef *ResourceFieldSelectorApplyConfiguration `json:"resourceFieldRef,omitempty"`
	ConfigMapKeyRef  *ConfigMapKeySelectorApplyConfiguration  `json:"configMapKeyRef,omitempty"`
	SecretKeyRef     *SecretKeySelectorApplyConfiguration     `json:"secretKeyRef,omitempty"`
}

// EnvVarSourceApplyConfiguration constructs an declarative configuration of the EnvVarSource type for use with
// apply.
func EnvVarSource() *EnvVarSourceApplyConfiguration {
	return &EnvVarSourceApplyConfiguration{}
}

// WithFieldRef sets the FieldRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FieldRef field is set to the value of the last call.
func (b *EnvVarSourceApplyConfiguration) WithFieldRef(value *ObjectFieldSelectorApplyConfiguration) *EnvVarSourceApplyConfiguration {
	b.FieldRef = value
	return b
}

// WithResourceFieldRef sets the ResourceFieldRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceFieldRef field is set to the value of the last call.
func (b *EnvVarSourceApplyConfiguration) WithResourceFieldRef(value *ResourceFieldSelectorApplyConfiguration) *EnvVarSourceApplyConfiguration {
	b.ResourceFieldRef = value
	return b
}

// WithConfigMapKeyRef sets the ConfigMapKeyRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigMapKeyRef field is set to the value of the last call.
func (b *EnvVarSourceApplyConfiguration) WithConfigMapKeyRef(value *ConfigMapKeySelectorApplyConfiguration) *EnvVarSourceApplyConfiguration {
	b.ConfigMapKeyRef = value
	return b
}

// WithSecretKeyRef sets the SecretKeyRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretKeyRef field is set to the value of the last call.
func (b *EnvVarSourceApplyConfiguration) WithSecretKeyRef(value *SecretKeySelectorApplyConfiguration) *EnvVarSourceApplyConfiguration {
	b.SecretKeyRef = value
	return b
}
