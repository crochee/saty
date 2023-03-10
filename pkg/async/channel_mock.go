// Code generated by MockGen. DO NOT EDIT.
// Source: ./channel.go

// Package async is a generated GoMock package.
package async

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	amqp "github.com/streadway/amqp"
)

// MockChannel is a mock of Channel interface.
type MockChannel struct {
	ctrl     *gomock.Controller
	recorder *MockChannelMockRecorder
}

// MockChannelMockRecorder is the mock recorder for MockChannel.
type MockChannelMockRecorder struct {
	mock *MockChannel
}

// NewMockChannel creates a new mock instance.
func NewMockChannel(ctrl *gomock.Controller) *MockChannel {
	mock := &MockChannel{ctrl: ctrl}
	mock.recorder = &MockChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannel) EXPECT() *MockChannelMockRecorder {
	return m.recorder
}

// Consume mocks base method.
func (m *MockChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWail bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", queue, consumer, autoAck, exclusive, noLocal, noWail, args)
	ret0, _ := ret[0].(<-chan amqp.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockChannelMockRecorder) Consume(queue, consumer, autoAck, exclusive, noLocal, noWail, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockChannel)(nil).Consume), queue, consumer, autoAck, exclusive, noLocal, noWail, args)
}

// DeclareAndBind mocks base method.
func (m *MockChannel) DeclareAndBind(exchange, kind, queue, key string, args ...map[string]interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{exchange, kind, queue, key}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeclareAndBind", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeclareAndBind indicates an expected call of DeclareAndBind.
func (mr *MockChannelMockRecorder) DeclareAndBind(exchange, kind, queue, key interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{exchange, kind, queue, key}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclareAndBind", reflect.TypeOf((*MockChannel)(nil).DeclareAndBind), varargs...)
}

// Publish mocks base method.
func (m *MockChannel) Publish(exchange, key string, mandatory, immediate bool, msg ...amqp.Publishing) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{exchange, key, mandatory, immediate}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockChannelMockRecorder) Publish(exchange, key, mandatory, immediate interface{}, msg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{exchange, key, mandatory, immediate}, msg...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockChannel)(nil).Publish), varargs...)
}