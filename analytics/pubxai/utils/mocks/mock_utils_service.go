// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prebid/prebid-server/v3/analytics/pubxai/utils (interfaces: UtilsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	openrtb2 "github.com/prebid/openrtb/v20/openrtb2"
	utils "github.com/prebid/prebid-server/v3/analytics/pubxai/utils"
)

// MockUtilsService is a mock of UtilsService interface.
type MockUtilsService struct {
	ctrl     *gomock.Controller
	recorder *MockUtilsServiceMockRecorder
}

// MockUtilsServiceMockRecorder is the mock recorder for MockUtilsService.
type MockUtilsServiceMockRecorder struct {
	mock *MockUtilsService
}

// NewMockUtilsService creates a new mock instance.
func NewMockUtilsService(ctrl *gomock.Controller) *MockUtilsService {
	mock := &MockUtilsService{ctrl: ctrl}
	mock.recorder = &MockUtilsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUtilsService) EXPECT() *MockUtilsServiceMockRecorder {
	return m.recorder
}

// AppendTimeoutBids mocks base method.
func (m *MockUtilsService) AppendTimeoutBids(arg0 []utils.Bid, arg1 map[string]openrtb2.Imp, arg2 *utils.LogObject) []utils.Bid {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendTimeoutBids", arg0, arg1, arg2)
	ret0, _ := ret[0].([]utils.Bid)
	return ret0
}

// AppendTimeoutBids indicates an expected call of AppendTimeoutBids.
func (mr *MockUtilsServiceMockRecorder) AppendTimeoutBids(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendTimeoutBids", reflect.TypeOf((*MockUtilsService)(nil).AppendTimeoutBids), arg0, arg1, arg2)
}

// ExtractAdunitCodes mocks base method.
func (m *MockUtilsService) ExtractAdunitCodes(arg0 map[string]interface{}) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractAdunitCodes", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// ExtractAdunitCodes indicates an expected call of ExtractAdunitCodes.
func (mr *MockUtilsServiceMockRecorder) ExtractAdunitCodes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractAdunitCodes", reflect.TypeOf((*MockUtilsService)(nil).ExtractAdunitCodes), arg0)
}

// ExtractConsentTypes mocks base method.
func (m *MockUtilsService) ExtractConsentTypes(arg0 map[string]interface{}) utils.ConsentDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractConsentTypes", arg0)
	ret0, _ := ret[0].(utils.ConsentDetail)
	return ret0
}

// ExtractConsentTypes indicates an expected call of ExtractConsentTypes.
func (mr *MockUtilsServiceMockRecorder) ExtractConsentTypes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractConsentTypes", reflect.TypeOf((*MockUtilsService)(nil).ExtractConsentTypes), arg0)
}

// ExtractDeviceData mocks base method.
func (m *MockUtilsService) ExtractDeviceData(arg0 map[string]interface{}) utils.DeviceDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractDeviceData", arg0)
	ret0, _ := ret[0].(utils.DeviceDetail)
	return ret0
}

// ExtractDeviceData indicates an expected call of ExtractDeviceData.
func (mr *MockUtilsServiceMockRecorder) ExtractDeviceData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractDeviceData", reflect.TypeOf((*MockUtilsService)(nil).ExtractDeviceData), arg0)
}

// ExtractFloorDetail mocks base method.
func (m *MockUtilsService) ExtractFloorDetail(arg0 map[string]interface{}) utils.FloorDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractFloorDetail", arg0)
	ret0, _ := ret[0].(utils.FloorDetail)
	return ret0
}

// ExtractFloorDetail indicates an expected call of ExtractFloorDetail.
func (mr *MockUtilsServiceMockRecorder) ExtractFloorDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractFloorDetail", reflect.TypeOf((*MockUtilsService)(nil).ExtractFloorDetail), arg0)
}

// ExtractPageData mocks base method.
func (m *MockUtilsService) ExtractPageData(arg0 map[string]interface{}) utils.PageDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractPageData", arg0)
	ret0, _ := ret[0].(utils.PageDetail)
	return ret0
}

// ExtractPageData indicates an expected call of ExtractPageData.
func (mr *MockUtilsServiceMockRecorder) ExtractPageData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractPageData", reflect.TypeOf((*MockUtilsService)(nil).ExtractPageData), arg0)
}

// ExtractUserIds mocks base method.
func (m *MockUtilsService) ExtractUserIds(arg0 map[string]interface{}) utils.UserDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractUserIds", arg0)
	ret0, _ := ret[0].(utils.UserDetail)
	return ret0
}

// ExtractUserIds indicates an expected call of ExtractUserIds.
func (mr *MockUtilsServiceMockRecorder) ExtractUserIds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractUserIds", reflect.TypeOf((*MockUtilsService)(nil).ExtractUserIds), arg0)
}

// ProcessBidResponses mocks base method.
func (m *MockUtilsService) ProcessBidResponses(arg0 []map[string]interface{}, arg1 string, arg2 int64, arg3, arg4 map[string]interface{}, arg5 utils.FloorDetail) ([]utils.Bid, []utils.Bid) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessBidResponses", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]utils.Bid)
	ret1, _ := ret[1].([]utils.Bid)
	return ret0, ret1
}

// ProcessBidResponses indicates an expected call of ProcessBidResponses.
func (mr *MockUtilsServiceMockRecorder) ProcessBidResponses(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBidResponses", reflect.TypeOf((*MockUtilsService)(nil).ProcessBidResponses), arg0, arg1, arg2, arg3, arg4, arg5)
}

// UnmarshalExtensions mocks base method.
func (m *MockUtilsService) UnmarshalExtensions(arg0 *utils.LogObject) (map[string]interface{}, map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalExtensions", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UnmarshalExtensions indicates an expected call of UnmarshalExtensions.
func (mr *MockUtilsServiceMockRecorder) UnmarshalExtensions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalExtensions", reflect.TypeOf((*MockUtilsService)(nil).UnmarshalExtensions), arg0)
}
