// Code generated by mockery v2.20.0. DO NOT EDIT.

package dao

import (
	api "github.com/content-services/content-sources-backend/pkg/api"
	mock "github.com/stretchr/testify/mock"
)

// MockRepositoryConfigDao is an autogenerated mock type for the RepositoryConfigDao type
type MockRepositoryConfigDao struct {
	mock.Mock
}

// BulkCreate provides a mock function with given fields: newRepositories
func (_m *MockRepositoryConfigDao) BulkCreate(newRepositories []api.RepositoryRequest) ([]api.RepositoryResponse, []error) {
	ret := _m.Called(newRepositories)

	var r0 []api.RepositoryResponse
	var r1 []error
	if rf, ok := ret.Get(0).(func([]api.RepositoryRequest) ([]api.RepositoryResponse, []error)); ok {
		return rf(newRepositories)
	}
	if rf, ok := ret.Get(0).(func([]api.RepositoryRequest) []api.RepositoryResponse); ok {
		r0 = rf(newRepositories)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.RepositoryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func([]api.RepositoryRequest) []error); ok {
		r1 = rf(newRepositories)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	return r0, r1
}

// Create provides a mock function with given fields: newRepo
func (_m *MockRepositoryConfigDao) Create(newRepo api.RepositoryRequest) (api.RepositoryResponse, error) {
	ret := _m.Called(newRepo)

	var r0 api.RepositoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(api.RepositoryRequest) (api.RepositoryResponse, error)); ok {
		return rf(newRepo)
	}
	if rf, ok := ret.Get(0).(func(api.RepositoryRequest) api.RepositoryResponse); ok {
		r0 = rf(newRepo)
	} else {
		r0 = ret.Get(0).(api.RepositoryResponse)
	}

	if rf, ok := ret.Get(1).(func(api.RepositoryRequest) error); ok {
		r1 = rf(newRepo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: orgID, uuid
func (_m *MockRepositoryConfigDao) Delete(orgID string, uuid string) error {
	ret := _m.Called(orgID, uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(orgID, uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: orgID, uuid
func (_m *MockRepositoryConfigDao) Fetch(orgID string, uuid string) (api.RepositoryResponse, error) {
	ret := _m.Called(orgID, uuid)

	var r0 api.RepositoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (api.RepositoryResponse, error)); ok {
		return rf(orgID, uuid)
	}
	if rf, ok := ret.Get(0).(func(string, string) api.RepositoryResponse); ok {
		r0 = rf(orgID, uuid)
	} else {
		r0 = ret.Get(0).(api.RepositoryResponse)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(orgID, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: orgID, paginationData, filterData
func (_m *MockRepositoryConfigDao) List(orgID string, paginationData api.PaginationData, filterData api.FilterData) (api.RepositoryCollectionResponse, int64, error) {
	ret := _m.Called(orgID, paginationData, filterData)

	var r0 api.RepositoryCollectionResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(string, api.PaginationData, api.FilterData) (api.RepositoryCollectionResponse, int64, error)); ok {
		return rf(orgID, paginationData, filterData)
	}
	if rf, ok := ret.Get(0).(func(string, api.PaginationData, api.FilterData) api.RepositoryCollectionResponse); ok {
		r0 = rf(orgID, paginationData, filterData)
	} else {
		r0 = ret.Get(0).(api.RepositoryCollectionResponse)
	}

	if rf, ok := ret.Get(1).(func(string, api.PaginationData, api.FilterData) int64); ok {
		r1 = rf(orgID, paginationData, filterData)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(string, api.PaginationData, api.FilterData) error); ok {
		r2 = rf(orgID, paginationData, filterData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SavePublicRepos provides a mock function with given fields: urls
func (_m *MockRepositoryConfigDao) SavePublicRepos(urls []string) error {
	ret := _m.Called(urls)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(urls)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: orgID, uuid, repoParams
func (_m *MockRepositoryConfigDao) Update(orgID string, uuid string, repoParams api.RepositoryRequest) error {
	ret := _m.Called(orgID, uuid, repoParams)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, api.RepositoryRequest) error); ok {
		r0 = rf(orgID, uuid, repoParams)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateParameters provides a mock function with given fields: orgId, params, excludedUUIDS
func (_m *MockRepositoryConfigDao) ValidateParameters(orgId string, params api.RepositoryValidationRequest, excludedUUIDS []string) (api.RepositoryValidationResponse, error) {
	ret := _m.Called(orgId, params, excludedUUIDS)

	var r0 api.RepositoryValidationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, api.RepositoryValidationRequest, []string) (api.RepositoryValidationResponse, error)); ok {
		return rf(orgId, params, excludedUUIDS)
	}
	if rf, ok := ret.Get(0).(func(string, api.RepositoryValidationRequest, []string) api.RepositoryValidationResponse); ok {
		r0 = rf(orgId, params, excludedUUIDS)
	} else {
		r0 = ret.Get(0).(api.RepositoryValidationResponse)
	}

	if rf, ok := ret.Get(1).(func(string, api.RepositoryValidationRequest, []string) error); ok {
		r1 = rf(orgId, params, excludedUUIDS)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRepositoryConfigDao interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepositoryConfigDao creates a new instance of MockRepositoryConfigDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepositoryConfigDao(t mockConstructorTestingTNewMockRepositoryConfigDao) *MockRepositoryConfigDao {
	mock := &MockRepositoryConfigDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
