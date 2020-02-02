/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package factory

import (
	"errors"
	"testing"
	"time"

	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/streamer/mconfig/test_protos"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
)

type mockMconfigBuilder struct {
	result map[string]proto.Message
	err    error
}

func (builder *mockMconfigBuilder) Build(networkId string, gatewayId string) (map[string]proto.Message, error) {
	return builder.result, builder.err
}

type mockClock struct {
	now time.Time
}

func (mockClock *mockClock) Now() time.Time {
	return mockClock.now
}

func TestCreateMconfig(t *testing.T) {
	factory.builders = factory.builders[:0]
	factory.clock = &mockClock{now: time.Unix(1551916956, 0)}

	builder1 := &mockMconfigBuilder{
		result: map[string]proto.Message{
			"builder1_1": &test_protos.Message1{Field: "hello"},
			"builder1_2": &test_protos.Message2{Field1: "hello", Field2: "world"},
		},
	}
	builder2 := &mockMconfigBuilder{
		result: map[string]proto.Message{
			"builder2_1": &test_protos.Message1{Field: "foo"},
		},
	}
	RegisterMconfigBuilder(builder1)
	RegisterMconfigBuilder(builder2)

	actual, err := CreateMconfig("foo", "bar")
	assert.NoError(t, err)

	expectedMap := map[string]proto.Message{
		"builder1_1": &test_protos.Message1{Field: "hello"},
		"builder1_2": &test_protos.Message2{Field1: "hello", Field2: "world"},
		"builder2_1": &test_protos.Message1{Field: "foo"},
	}
	expectedAny := make(map[string]*any.Any, len(expectedMap))
	for k, v := range expectedMap {
		anyV, err := ptypes.MarshalAny(v)
		assert.NoError(t, err)
		expectedAny[k] = anyV
	}
	expected := &protos.GatewayConfigs{
		ConfigsByKey: expectedAny,
		Metadata: &protos.GatewayConfigsMetadata{
			CreatedAt: 1551916956,
		},
	}
	assert.Equal(t, *expected, *actual)
}

func TestCreateMconfig_DuplicateKey(t *testing.T) {
	factory.builders = factory.builders[:0]
	factory.clock = &mockClock{now: time.Unix(1551916956, 0)}

	builder1 := &mockMconfigBuilder{
		result: map[string]proto.Message{
			"key": &test_protos.Message1{},
		},
	}
	builder2 := &mockMconfigBuilder{
		result: map[string]proto.Message{
			"key": &test_protos.Message1{},
		},
	}
	RegisterMconfigBuilder(builder1)
	RegisterMconfigBuilder(builder2)

	_, err := CreateMconfig("foo", "bar")
	assert.Error(t, err)
	assert.Equal(t, "mconfig builder returned result for duplicate key key", err.Error())
}

func TestCreateMconfig_BuilderError(t *testing.T) {
	factory.builders = factory.builders[:0]
	factory.clock = &mockClock{now: time.Unix(1551916956, 0)}

	builder := &mockMconfigBuilder{
		err: errors.New("FOO"),
	}
	RegisterMconfigBuilder(builder)

	_, err := CreateMconfig("foo", "bar")
	assert.Error(t, err)
	assert.Equal(t, "FOO", err.Error())
}
