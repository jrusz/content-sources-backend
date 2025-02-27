// Code generated by mockery v2.20.0. DO NOT EDIT.

package dao

import mock "github.com/stretchr/testify/mock"

// MockSnapshotDao is an autogenerated mock type for the SnapshotDao type
type MockSnapshotDao struct {
	mock.Mock
}

// Create provides a mock function with given fields: snap
func (_m *MockSnapshotDao) Create(snap *Snapshot) error {
	ret := _m.Called(snap)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Snapshot) error); ok {
		r0 = rf(snap)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: orgId, repoConfigUuid
func (_m *MockSnapshotDao) List(repoConfigUuid string) ([]Snapshot, error) {
	ret := _m.Called(repoConfigUuid)

	var r0 []Snapshot
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]Snapshot, error)); ok {
		return rf(repoConfigUuid, repoConfigUuid)
	}
	if rf, ok := ret.Get(0).(func(string, string) []Snapshot); ok {
		r0 = rf(repoConfigUuid, repoConfigUuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Snapshot)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewMockSnapshotDao interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSnapshotDao creates a new instance of MockSnapshotDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSnapshotDao(t mockConstructorTestingTNewMockSnapshotDao) *MockSnapshotDao {
	mock := &MockSnapshotDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
