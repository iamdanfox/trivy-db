// Code generated by mockery v1.0.0. DO NOT EDIT.

package types

import mock "github.com/stretchr/testify/mock"

// MockVulnSrc is an autogenerated mock type for the VulnSrc type
type MockVulnSrc struct {
	mock.Mock
}

type GetArgs struct {
	Release         string
	ReleaseAnything bool
	PkgName         string
	PkgNameAnything bool
}

type GetReturns struct {
	Advisories []Advisory
	Err        error
}

type GetExpectation struct {
	Args    GetArgs
	Returns GetReturns
}

func (_m *MockVulnSrc) ApplyGetExpectation(e GetExpectation) {
	var args []interface{}
	if e.Args.ReleaseAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Release)
	}
	if e.Args.PkgNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.PkgName)
	}
	_m.On("Get", args...).Return(e.Returns.Advisories, e.Returns.Err)
}

func (_m *MockVulnSrc) ApplyGetExpectations(expectations []GetExpectation) {
	for _, e := range expectations {
		_m.ApplyGetExpectation(e)
	}
}

// Get provides a mock function with given fields: release, pkgName
func (_m *MockVulnSrc) Get(release string, pkgName string) ([]Advisory, error) {
	ret := _m.Called(release, pkgName)

	var r0 []Advisory
	if rf, ok := ret.Get(0).(func(string, string) []Advisory); ok {
		r0 = rf(release, pkgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Advisory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(release, pkgName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type UpdateArgs struct {
	Dir         string
	DirAnything bool
}

type UpdateReturns struct {
	Err error
}

type UpdateExpectation struct {
	Args    UpdateArgs
	Returns UpdateReturns
}

func (_m *MockVulnSrc) ApplyUpdateExpectation(e UpdateExpectation) {
	var args []interface{}
	if e.Args.DirAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Dir)
	}
	_m.On("Update", args...).Return(e.Returns.Err)
}

func (_m *MockVulnSrc) ApplyUpdateExpectations(expectations []UpdateExpectation) {
	for _, e := range expectations {
		_m.ApplyUpdateExpectation(e)
	}
}

// Update provides a mock function with given fields: dir
func (_m *MockVulnSrc) Update(dir string) error {
	ret := _m.Called(dir)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dir)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
