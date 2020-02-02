/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// Code generated by mockery v1.0.0
package mocks

import (
	"magma/lte/cloud/go/protos"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/mock"
)

// Decoder is an autogenerated mock type for the Decoder type
type Decoder struct {
	mock.Mock
}

// protoFromAttributeMap provides a mock function with given fields: _a0
func (_m *Decoder) ProtoFromAttributeMap(_a0 map[string]*dynamodb.AttributeValue) (*protos.FlowRecord, error) {
	ret := _m.Called(_a0)

	var r0 *protos.FlowRecord
	if rf, ok := ret.Get(0).(func(map[string]*dynamodb.AttributeValue) *protos.FlowRecord); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.FlowRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]*dynamodb.AttributeValue) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
