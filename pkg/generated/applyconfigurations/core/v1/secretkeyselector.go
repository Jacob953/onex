// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// SecretKeySelectorApplyConfiguration represents an declarative configuration of the SecretKeySelector type for use
// with apply.
type SecretKeySelectorApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Key                                    *string `json:"key,omitempty"`
	Optional                               *bool   `json:"optional,omitempty"`
}

// SecretKeySelectorApplyConfiguration constructs an declarative configuration of the SecretKeySelector type for use with
// apply.
func SecretKeySelector() *SecretKeySelectorApplyConfiguration {
	return &SecretKeySelectorApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SecretKeySelectorApplyConfiguration) WithName(value string) *SecretKeySelectorApplyConfiguration {
	b.Name = &value
	return b
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *SecretKeySelectorApplyConfiguration) WithKey(value string) *SecretKeySelectorApplyConfiguration {
	b.Key = &value
	return b
}

// WithOptional sets the Optional field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Optional field is set to the value of the last call.
func (b *SecretKeySelectorApplyConfiguration) WithOptional(value bool) *SecretKeySelectorApplyConfiguration {
	b.Optional = &value
	return b
}
