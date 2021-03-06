// Code generated by mockery v1.0.0. DO NOT EDIT.

package github

import context "context"
import mock "github.com/stretchr/testify/mock"

// MockVCSClientInterface is an autogenerated mock type for the VCSClientInterface type
type MockVCSClientInterface struct {
	mock.Mock
}

type UploadReleaseAssetsArgs struct {
	Ctx               context.Context
	CtxAnything       bool
	FilePaths         []string
	FilePathsAnything bool
}

type UploadReleaseAssetsReturns struct {
	Err error
}

type UploadReleaseAssetsExpectation struct {
	Args    UploadReleaseAssetsArgs
	Returns UploadReleaseAssetsReturns
}

func (_m *MockVCSClientInterface) ApplyUploadReleaseAssetsExpectation(e UploadReleaseAssetsExpectation) {
	var args []interface{}
	if e.Args.CtxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ctx)
	}
	if e.Args.FilePathsAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.FilePaths)
	}
	_m.On("UploadReleaseAssets", args...).Return(e.Returns.Err)
}

func (_m *MockVCSClientInterface) ApplyUploadReleaseAssetsExpectations(expectations []UploadReleaseAssetsExpectation) {
	for _, e := range expectations {
		_m.ApplyUploadReleaseAssetsExpectation(e)
	}
}

// UploadReleaseAssets provides a mock function with given fields: ctx, filePaths
func (_m *MockVCSClientInterface) UploadReleaseAssets(ctx context.Context, filePaths []string) error {
	ret := _m.Called(ctx, filePaths)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, filePaths)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
