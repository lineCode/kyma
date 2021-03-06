// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import v1beta1 "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"

// KubelessFunctionFinder is an autogenerated mock type for the KubelessFunctionFinder type
type KubelessFunctionFinder struct {
	mock.Mock
}

// FindFunctionModifiedByServiceBindingUsage provides a mock function with given fields: fnNs, usageName
func (_m *KubelessFunctionFinder) FindFunctionModifiedByServiceBindingUsage(fnNs string, usageName string) (*v1beta1.Function, error) {
	ret := _m.Called(fnNs, usageName)

	var r0 *v1beta1.Function
	if rf, ok := ret.Get(0).(func(string, string) *v1beta1.Function); ok {
		r0 = rf(fnNs, usageName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Function)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(fnNs, usageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
