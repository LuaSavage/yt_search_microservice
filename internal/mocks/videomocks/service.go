// Code generated by mockery v2.12.2. DO NOT EDIT.

package videomocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateVideo provides a mock function with given fields: ctx, _a1
func (_m *Service) CreateVideo(ctx context.Context, _a1 video.Video) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, video.Video) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetVideoByID provides a mock function with given fields: ctx, id
func (_m *Service) GetVideoByID(ctx context.Context, id string) (*video.Video, error) {
	ret := _m.Called(ctx, id)

	var r0 *video.Video
	if rf, ok := ret.Get(0).(func(context.Context, string) *video.Video); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*video.Video)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t testing.TB) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
