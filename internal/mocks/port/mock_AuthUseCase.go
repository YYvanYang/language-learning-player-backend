// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"context"

	mock "github.com/stretchr/testify/mock"
	"github.com/yvanyang/language-learning-player-api/internal/domain"
	"github.com/yvanyang/language-learning-player-api/internal/port"
)

// NewMockAuthUseCase creates a new instance of MockAuthUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAuthUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAuthUseCase {
	mock := &MockAuthUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockAuthUseCase is an autogenerated mock type for the AuthUseCase type
type MockAuthUseCase struct {
	mock.Mock
}

type MockAuthUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthUseCase) EXPECT() *MockAuthUseCase_Expecter {
	return &MockAuthUseCase_Expecter{mock: &_m.Mock}
}

// AuthenticateWithGoogle provides a mock function for the type MockAuthUseCase
func (_mock *MockAuthUseCase) AuthenticateWithGoogle(ctx context.Context, googleIdToken string) (port.AuthResult, error) {
	ret := _mock.Called(ctx, googleIdToken)

	if len(ret) == 0 {
		panic("no return value specified for AuthenticateWithGoogle")
	}

	var r0 port.AuthResult
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (port.AuthResult, error)); ok {
		return returnFunc(ctx, googleIdToken)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) port.AuthResult); ok {
		r0 = returnFunc(ctx, googleIdToken)
	} else {
		r0 = ret.Get(0).(port.AuthResult)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = returnFunc(ctx, googleIdToken)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockAuthUseCase_AuthenticateWithGoogle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthenticateWithGoogle'
type MockAuthUseCase_AuthenticateWithGoogle_Call struct {
	*mock.Call
}

// AuthenticateWithGoogle is a helper method to define mock.On call
//   - ctx
//   - googleIdToken
func (_e *MockAuthUseCase_Expecter) AuthenticateWithGoogle(ctx interface{}, googleIdToken interface{}) *MockAuthUseCase_AuthenticateWithGoogle_Call {
	return &MockAuthUseCase_AuthenticateWithGoogle_Call{Call: _e.mock.On("AuthenticateWithGoogle", ctx, googleIdToken)}
}

func (_c *MockAuthUseCase_AuthenticateWithGoogle_Call) Run(run func(ctx context.Context, googleIdToken string)) *MockAuthUseCase_AuthenticateWithGoogle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAuthUseCase_AuthenticateWithGoogle_Call) Return(authResult port.AuthResult, err error) *MockAuthUseCase_AuthenticateWithGoogle_Call {
	_c.Call.Return(authResult, err)
	return _c
}

func (_c *MockAuthUseCase_AuthenticateWithGoogle_Call) RunAndReturn(run func(ctx context.Context, googleIdToken string) (port.AuthResult, error)) *MockAuthUseCase_AuthenticateWithGoogle_Call {
	_c.Call.Return(run)
	return _c
}

// LoginWithPassword provides a mock function for the type MockAuthUseCase
func (_mock *MockAuthUseCase) LoginWithPassword(ctx context.Context, emailStr string, password string) (port.AuthResult, error) {
	ret := _mock.Called(ctx, emailStr, password)

	if len(ret) == 0 {
		panic("no return value specified for LoginWithPassword")
	}

	var r0 port.AuthResult
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) (port.AuthResult, error)); ok {
		return returnFunc(ctx, emailStr, password)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string) port.AuthResult); ok {
		r0 = returnFunc(ctx, emailStr, password)
	} else {
		r0 = ret.Get(0).(port.AuthResult)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = returnFunc(ctx, emailStr, password)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockAuthUseCase_LoginWithPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoginWithPassword'
type MockAuthUseCase_LoginWithPassword_Call struct {
	*mock.Call
}

// LoginWithPassword is a helper method to define mock.On call
//   - ctx
//   - emailStr
//   - password
func (_e *MockAuthUseCase_Expecter) LoginWithPassword(ctx interface{}, emailStr interface{}, password interface{}) *MockAuthUseCase_LoginWithPassword_Call {
	return &MockAuthUseCase_LoginWithPassword_Call{Call: _e.mock.On("LoginWithPassword", ctx, emailStr, password)}
}

func (_c *MockAuthUseCase_LoginWithPassword_Call) Run(run func(ctx context.Context, emailStr string, password string)) *MockAuthUseCase_LoginWithPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockAuthUseCase_LoginWithPassword_Call) Return(authResult port.AuthResult, err error) *MockAuthUseCase_LoginWithPassword_Call {
	_c.Call.Return(authResult, err)
	return _c
}

func (_c *MockAuthUseCase_LoginWithPassword_Call) RunAndReturn(run func(ctx context.Context, emailStr string, password string) (port.AuthResult, error)) *MockAuthUseCase_LoginWithPassword_Call {
	_c.Call.Return(run)
	return _c
}

// Logout provides a mock function for the type MockAuthUseCase
func (_mock *MockAuthUseCase) Logout(ctx context.Context, refreshTokenValue string) error {
	ret := _mock.Called(ctx, refreshTokenValue)

	if len(ret) == 0 {
		panic("no return value specified for Logout")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = returnFunc(ctx, refreshTokenValue)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockAuthUseCase_Logout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Logout'
type MockAuthUseCase_Logout_Call struct {
	*mock.Call
}

// Logout is a helper method to define mock.On call
//   - ctx
//   - refreshTokenValue
func (_e *MockAuthUseCase_Expecter) Logout(ctx interface{}, refreshTokenValue interface{}) *MockAuthUseCase_Logout_Call {
	return &MockAuthUseCase_Logout_Call{Call: _e.mock.On("Logout", ctx, refreshTokenValue)}
}

func (_c *MockAuthUseCase_Logout_Call) Run(run func(ctx context.Context, refreshTokenValue string)) *MockAuthUseCase_Logout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAuthUseCase_Logout_Call) Return(err error) *MockAuthUseCase_Logout_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockAuthUseCase_Logout_Call) RunAndReturn(run func(ctx context.Context, refreshTokenValue string) error) *MockAuthUseCase_Logout_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshAccessToken provides a mock function for the type MockAuthUseCase
func (_mock *MockAuthUseCase) RefreshAccessToken(ctx context.Context, refreshTokenValue string) (port.AuthResult, error) {
	ret := _mock.Called(ctx, refreshTokenValue)

	if len(ret) == 0 {
		panic("no return value specified for RefreshAccessToken")
	}

	var r0 port.AuthResult
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (port.AuthResult, error)); ok {
		return returnFunc(ctx, refreshTokenValue)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) port.AuthResult); ok {
		r0 = returnFunc(ctx, refreshTokenValue)
	} else {
		r0 = ret.Get(0).(port.AuthResult)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = returnFunc(ctx, refreshTokenValue)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockAuthUseCase_RefreshAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshAccessToken'
type MockAuthUseCase_RefreshAccessToken_Call struct {
	*mock.Call
}

// RefreshAccessToken is a helper method to define mock.On call
//   - ctx
//   - refreshTokenValue
func (_e *MockAuthUseCase_Expecter) RefreshAccessToken(ctx interface{}, refreshTokenValue interface{}) *MockAuthUseCase_RefreshAccessToken_Call {
	return &MockAuthUseCase_RefreshAccessToken_Call{Call: _e.mock.On("RefreshAccessToken", ctx, refreshTokenValue)}
}

func (_c *MockAuthUseCase_RefreshAccessToken_Call) Run(run func(ctx context.Context, refreshTokenValue string)) *MockAuthUseCase_RefreshAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAuthUseCase_RefreshAccessToken_Call) Return(authResult port.AuthResult, err error) *MockAuthUseCase_RefreshAccessToken_Call {
	_c.Call.Return(authResult, err)
	return _c
}

func (_c *MockAuthUseCase_RefreshAccessToken_Call) RunAndReturn(run func(ctx context.Context, refreshTokenValue string) (port.AuthResult, error)) *MockAuthUseCase_RefreshAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterWithPassword provides a mock function for the type MockAuthUseCase
func (_mock *MockAuthUseCase) RegisterWithPassword(ctx context.Context, emailStr string, password string, name string) (*domain.User, port.AuthResult, error) {
	ret := _mock.Called(ctx, emailStr, password, name)

	if len(ret) == 0 {
		panic("no return value specified for RegisterWithPassword")
	}

	var r0 *domain.User
	var r1 port.AuthResult
	var r2 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string, string) (*domain.User, port.AuthResult, error)); ok {
		return returnFunc(ctx, emailStr, password, name)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string, string) *domain.User); ok {
		r0 = returnFunc(ctx, emailStr, password, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string, string) port.AuthResult); ok {
		r1 = returnFunc(ctx, emailStr, password, name)
	} else {
		r1 = ret.Get(1).(port.AuthResult)
	}
	if returnFunc, ok := ret.Get(2).(func(context.Context, string, string, string) error); ok {
		r2 = returnFunc(ctx, emailStr, password, name)
	} else {
		r2 = ret.Error(2)
	}
	return r0, r1, r2
}

// MockAuthUseCase_RegisterWithPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterWithPassword'
type MockAuthUseCase_RegisterWithPassword_Call struct {
	*mock.Call
}

// RegisterWithPassword is a helper method to define mock.On call
//   - ctx
//   - emailStr
//   - password
//   - name
func (_e *MockAuthUseCase_Expecter) RegisterWithPassword(ctx interface{}, emailStr interface{}, password interface{}, name interface{}) *MockAuthUseCase_RegisterWithPassword_Call {
	return &MockAuthUseCase_RegisterWithPassword_Call{Call: _e.mock.On("RegisterWithPassword", ctx, emailStr, password, name)}
}

func (_c *MockAuthUseCase_RegisterWithPassword_Call) Run(run func(ctx context.Context, emailStr string, password string, name string)) *MockAuthUseCase_RegisterWithPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockAuthUseCase_RegisterWithPassword_Call) Return(user *domain.User, authResult port.AuthResult, err error) *MockAuthUseCase_RegisterWithPassword_Call {
	_c.Call.Return(user, authResult, err)
	return _c
}

func (_c *MockAuthUseCase_RegisterWithPassword_Call) RunAndReturn(run func(ctx context.Context, emailStr string, password string, name string) (*domain.User, port.AuthResult, error)) *MockAuthUseCase_RegisterWithPassword_Call {
	_c.Call.Return(run)
	return _c
}
