// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	mongo "go.mongodb.org/mongo-driver/mongo"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockMongoDatabaseIF is an autogenerated mock type for the MongoDatabaseIF type
type MockMongoDatabaseIF struct {
	mock.Mock
}

type MockMongoDatabaseIF_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMongoDatabaseIF) EXPECT() *MockMongoDatabaseIF_Expecter {
	return &MockMongoDatabaseIF_Expecter{mock: &_m.Mock}
}

// Collection provides a mock function with given fields: name, opts
func (_m *MockMongoDatabaseIF) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.Collection
	if rf, ok := ret.Get(0).(func(string, ...*options.CollectionOptions) *mongo.Collection); ok {
		r0 = rf(name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Collection)
		}
	}

	return r0
}

// MockMongoDatabaseIF_Collection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Collection'
type MockMongoDatabaseIF_Collection_Call struct {
	*mock.Call
}

// Collection is a helper method to define mock.On call
//   - name string
//   - opts ...*options.CollectionOptions
func (_e *MockMongoDatabaseIF_Expecter) Collection(name interface{}, opts ...interface{}) *MockMongoDatabaseIF_Collection_Call {
	return &MockMongoDatabaseIF_Collection_Call{Call: _e.mock.On("Collection",
		append([]interface{}{name}, opts...)...)}
}

func (_c *MockMongoDatabaseIF_Collection_Call) Run(run func(name string, opts ...*options.CollectionOptions)) *MockMongoDatabaseIF_Collection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.CollectionOptions, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*options.CollectionOptions)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockMongoDatabaseIF_Collection_Call) Return(_a0 *mongo.Collection) *MockMongoDatabaseIF_Collection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMongoDatabaseIF_Collection_Call) RunAndReturn(run func(string, ...*options.CollectionOptions) *mongo.Collection) *MockMongoDatabaseIF_Collection_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockMongoDatabaseIF interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMongoDatabaseIF creates a new instance of MockMongoDatabaseIF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMongoDatabaseIF(t mockConstructorTestingTNewMockMongoDatabaseIF) *MockMongoDatabaseIF {
	mock := &MockMongoDatabaseIF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}