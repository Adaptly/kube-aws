// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-k8s/pkg/grpcwrapper (interfaces: GRPC)

// Package mock_grpcwrapper is a generated GoMock package.
package mock_grpcwrapper

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockGRPC is a mock of GRPC interface
type MockGRPC struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCMockRecorder
}

// MockGRPCMockRecorder is the mock recorder for MockGRPC
type MockGRPCMockRecorder struct {
	mock *MockGRPC
}

// NewMockGRPC creates a new mock instance
func NewMockGRPC(ctrl *gomock.Controller) *MockGRPC {
	mock := &MockGRPC{ctrl: ctrl}
	mock.recorder = &MockGRPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGRPC) EXPECT() *MockGRPCMockRecorder {
	return m.recorder
}

// Dial mocks base method
func (m *MockGRPC) Dial(arg0 string, arg1 ...grpc.DialOption) (*grpc.ClientConn, error) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Dial", varargs...)
	ret0, _ := ret[0].(*grpc.ClientConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dial indicates an expected call of Dial
func (mr *MockGRPCMockRecorder) Dial(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dial", reflect.TypeOf((*MockGRPC)(nil).Dial), varargs...)
}