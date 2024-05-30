// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"

	ssz "github.com/ferranbt/fastssz"

	types "github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
)

// BeaconBlock is an autogenerated mock type for the BeaconBlock type
type BeaconBlock struct {
	mock.Mock
}

type BeaconBlock_Expecter struct {
	mock *mock.Mock
}

func (_m *BeaconBlock) EXPECT() *BeaconBlock_Expecter {
	return &BeaconBlock_Expecter{mock: &_m.Mock}
}

// GetBody provides a mock function with given fields:
func (_m *BeaconBlock) GetBody() types.BeaconBlockBody {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 types.BeaconBlockBody
	if rf, ok := ret.Get(0).(func() types.BeaconBlockBody); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.BeaconBlockBody)
		}
	}

	return r0
}

// BeaconBlock_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type BeaconBlock_GetBody_Call struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetBody() *BeaconBlock_GetBody_Call {
	return &BeaconBlock_GetBody_Call{Call: _e.mock.On("GetBody")}
}

func (_c *BeaconBlock_GetBody_Call) Run(run func()) *BeaconBlock_GetBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetBody_Call) Return(_a0 types.BeaconBlockBody) *BeaconBlock_GetBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetBody_Call) RunAndReturn(run func() types.BeaconBlockBody) *BeaconBlock_GetBody_Call {
	_c.Call.Return(run)
	return _c
}

// GetHeader provides a mock function with given fields:
func (_m *BeaconBlock) GetHeader() *types.BeaconBlockHeader {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetHeader")
	}

	var r0 *types.BeaconBlockHeader
	if rf, ok := ret.Get(0).(func() *types.BeaconBlockHeader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BeaconBlockHeader)
		}
	}

	return r0
}

// BeaconBlock_GetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHeader'
type BeaconBlock_GetHeader_Call struct {
	*mock.Call
}

// GetHeader is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetHeader() *BeaconBlock_GetHeader_Call {
	return &BeaconBlock_GetHeader_Call{Call: _e.mock.On("GetHeader")}
}

func (_c *BeaconBlock_GetHeader_Call) Run(run func()) *BeaconBlock_GetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetHeader_Call) Return(_a0 *types.BeaconBlockHeader) *BeaconBlock_GetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetHeader_Call) RunAndReturn(run func() *types.BeaconBlockHeader) *BeaconBlock_GetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// GetParentBlockRoot provides a mock function with given fields:
func (_m *BeaconBlock) GetParentBlockRoot() bytes.B32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetParentBlockRoot")
	}

	var r0 bytes.B32
	if rf, ok := ret.Get(0).(func() bytes.B32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bytes.B32)
		}
	}

	return r0
}

// BeaconBlock_GetParentBlockRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetParentBlockRoot'
type BeaconBlock_GetParentBlockRoot_Call struct {
	*mock.Call
}

// GetParentBlockRoot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetParentBlockRoot() *BeaconBlock_GetParentBlockRoot_Call {
	return &BeaconBlock_GetParentBlockRoot_Call{Call: _e.mock.On("GetParentBlockRoot")}
}

func (_c *BeaconBlock_GetParentBlockRoot_Call) Run(run func()) *BeaconBlock_GetParentBlockRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetParentBlockRoot_Call) Return(_a0 bytes.B32) *BeaconBlock_GetParentBlockRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetParentBlockRoot_Call) RunAndReturn(run func() bytes.B32) *BeaconBlock_GetParentBlockRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetProposerIndex provides a mock function with given fields:
func (_m *BeaconBlock) GetProposerIndex() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProposerIndex")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlock_GetProposerIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProposerIndex'
type BeaconBlock_GetProposerIndex_Call struct {
	*mock.Call
}

// GetProposerIndex is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetProposerIndex() *BeaconBlock_GetProposerIndex_Call {
	return &BeaconBlock_GetProposerIndex_Call{Call: _e.mock.On("GetProposerIndex")}
}

func (_c *BeaconBlock_GetProposerIndex_Call) Run(run func()) *BeaconBlock_GetProposerIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetProposerIndex_Call) Return(_a0 math.U64) *BeaconBlock_GetProposerIndex_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetProposerIndex_Call) RunAndReturn(run func() math.U64) *BeaconBlock_GetProposerIndex_Call {
	_c.Call.Return(run)
	return _c
}

// GetSlot provides a mock function with given fields:
func (_m *BeaconBlock) GetSlot() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlock_GetSlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlot'
type BeaconBlock_GetSlot_Call struct {
	*mock.Call
}

// GetSlot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetSlot() *BeaconBlock_GetSlot_Call {
	return &BeaconBlock_GetSlot_Call{Call: _e.mock.On("GetSlot")}
}

func (_c *BeaconBlock_GetSlot_Call) Run(run func()) *BeaconBlock_GetSlot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetSlot_Call) Return(_a0 math.U64) *BeaconBlock_GetSlot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetSlot_Call) RunAndReturn(run func() math.U64) *BeaconBlock_GetSlot_Call {
	_c.Call.Return(run)
	return _c
}

// GetStateRoot provides a mock function with given fields:
func (_m *BeaconBlock) GetStateRoot() bytes.B32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStateRoot")
	}

	var r0 bytes.B32
	if rf, ok := ret.Get(0).(func() bytes.B32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bytes.B32)
		}
	}

	return r0
}

// BeaconBlock_GetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStateRoot'
type BeaconBlock_GetStateRoot_Call struct {
	*mock.Call
}

// GetStateRoot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetStateRoot() *BeaconBlock_GetStateRoot_Call {
	return &BeaconBlock_GetStateRoot_Call{Call: _e.mock.On("GetStateRoot")}
}

func (_c *BeaconBlock_GetStateRoot_Call) Run(run func()) *BeaconBlock_GetStateRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetStateRoot_Call) Return(_a0 bytes.B32) *BeaconBlock_GetStateRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetStateRoot_Call) RunAndReturn(run func() bytes.B32) *BeaconBlock_GetStateRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetTree provides a mock function with given fields:
func (_m *BeaconBlock) GetTree() (*ssz.Node, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTree")
	}

	var r0 *ssz.Node
	var r1 error
	if rf, ok := ret.Get(0).(func() (*ssz.Node, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *ssz.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssz.Node)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeaconBlock_GetTree_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTree'
type BeaconBlock_GetTree_Call struct {
	*mock.Call
}

// GetTree is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) GetTree() *BeaconBlock_GetTree_Call {
	return &BeaconBlock_GetTree_Call{Call: _e.mock.On("GetTree")}
}

func (_c *BeaconBlock_GetTree_Call) Run(run func()) *BeaconBlock_GetTree_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetTree_Call) Return(_a0 *ssz.Node, _a1 error) *BeaconBlock_GetTree_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BeaconBlock_GetTree_Call) RunAndReturn(run func() (*ssz.Node, error)) *BeaconBlock_GetTree_Call {
	_c.Call.Return(run)
	return _c
}

// HashTreeRoot provides a mock function with given fields:
func (_m *BeaconBlock) HashTreeRoot() ([32]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HashTreeRoot")
	}

	var r0 [32]byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([32]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() [32]byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeaconBlock_HashTreeRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashTreeRoot'
type BeaconBlock_HashTreeRoot_Call struct {
	*mock.Call
}

// HashTreeRoot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) HashTreeRoot() *BeaconBlock_HashTreeRoot_Call {
	return &BeaconBlock_HashTreeRoot_Call{Call: _e.mock.On("HashTreeRoot")}
}

func (_c *BeaconBlock_HashTreeRoot_Call) Run(run func()) *BeaconBlock_HashTreeRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_HashTreeRoot_Call) Return(_a0 [32]byte, _a1 error) *BeaconBlock_HashTreeRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BeaconBlock_HashTreeRoot_Call) RunAndReturn(run func() ([32]byte, error)) *BeaconBlock_HashTreeRoot_Call {
	_c.Call.Return(run)
	return _c
}

// HashTreeRootWith provides a mock function with given fields: hh
func (_m *BeaconBlock) HashTreeRootWith(hh ssz.HashWalker) error {
	ret := _m.Called(hh)

	if len(ret) == 0 {
		panic("no return value specified for HashTreeRootWith")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(ssz.HashWalker) error); ok {
		r0 = rf(hh)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeaconBlock_HashTreeRootWith_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashTreeRootWith'
type BeaconBlock_HashTreeRootWith_Call struct {
	*mock.Call
}

// HashTreeRootWith is a helper method to define mock.On call
//   - hh ssz.HashWalker
func (_e *BeaconBlock_Expecter) HashTreeRootWith(hh interface{}) *BeaconBlock_HashTreeRootWith_Call {
	return &BeaconBlock_HashTreeRootWith_Call{Call: _e.mock.On("HashTreeRootWith", hh)}
}

func (_c *BeaconBlock_HashTreeRootWith_Call) Run(run func(hh ssz.HashWalker)) *BeaconBlock_HashTreeRootWith_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ssz.HashWalker))
	})
	return _c
}

func (_c *BeaconBlock_HashTreeRootWith_Call) Return(_a0 error) *BeaconBlock_HashTreeRootWith_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_HashTreeRootWith_Call) RunAndReturn(run func(ssz.HashWalker) error) *BeaconBlock_HashTreeRootWith_Call {
	_c.Call.Return(run)
	return _c
}

// IsNil provides a mock function with given fields:
func (_m *BeaconBlock) IsNil() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsNil")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BeaconBlock_IsNil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNil'
type BeaconBlock_IsNil_Call struct {
	*mock.Call
}

// IsNil is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) IsNil() *BeaconBlock_IsNil_Call {
	return &BeaconBlock_IsNil_Call{Call: _e.mock.On("IsNil")}
}

func (_c *BeaconBlock_IsNil_Call) Run(run func()) *BeaconBlock_IsNil_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_IsNil_Call) Return(_a0 bool) *BeaconBlock_IsNil_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_IsNil_Call) RunAndReturn(run func() bool) *BeaconBlock_IsNil_Call {
	_c.Call.Return(run)
	return _c
}

// MarshalSSZ provides a mock function with given fields:
func (_m *BeaconBlock) MarshalSSZ() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalSSZ")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeaconBlock_MarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalSSZ'
type BeaconBlock_MarshalSSZ_Call struct {
	*mock.Call
}

// MarshalSSZ is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) MarshalSSZ() *BeaconBlock_MarshalSSZ_Call {
	return &BeaconBlock_MarshalSSZ_Call{Call: _e.mock.On("MarshalSSZ")}
}

func (_c *BeaconBlock_MarshalSSZ_Call) Run(run func()) *BeaconBlock_MarshalSSZ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_MarshalSSZ_Call) Return(_a0 []byte, _a1 error) *BeaconBlock_MarshalSSZ_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BeaconBlock_MarshalSSZ_Call) RunAndReturn(run func() ([]byte, error)) *BeaconBlock_MarshalSSZ_Call {
	_c.Call.Return(run)
	return _c
}

// MarshalSSZTo provides a mock function with given fields: dst
func (_m *BeaconBlock) MarshalSSZTo(dst []byte) ([]byte, error) {
	ret := _m.Called(dst)

	if len(ret) == 0 {
		panic("no return value specified for MarshalSSZTo")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(dst)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(dst)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(dst)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeaconBlock_MarshalSSZTo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalSSZTo'
type BeaconBlock_MarshalSSZTo_Call struct {
	*mock.Call
}

// MarshalSSZTo is a helper method to define mock.On call
//   - dst []byte
func (_e *BeaconBlock_Expecter) MarshalSSZTo(dst interface{}) *BeaconBlock_MarshalSSZTo_Call {
	return &BeaconBlock_MarshalSSZTo_Call{Call: _e.mock.On("MarshalSSZTo", dst)}
}

func (_c *BeaconBlock_MarshalSSZTo_Call) Run(run func(dst []byte)) *BeaconBlock_MarshalSSZTo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *BeaconBlock_MarshalSSZTo_Call) Return(_a0 []byte, _a1 error) *BeaconBlock_MarshalSSZTo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BeaconBlock_MarshalSSZTo_Call) RunAndReturn(run func([]byte) ([]byte, error)) *BeaconBlock_MarshalSSZTo_Call {
	_c.Call.Return(run)
	return _c
}

// SetStateRoot provides a mock function with given fields: _a0
func (_m *BeaconBlock) SetStateRoot(_a0 bytes.B32) {
	_m.Called(_a0)
}

// BeaconBlock_SetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStateRoot'
type BeaconBlock_SetStateRoot_Call struct {
	*mock.Call
}

// SetStateRoot is a helper method to define mock.On call
//   - _a0 bytes.B32
func (_e *BeaconBlock_Expecter) SetStateRoot(_a0 interface{}) *BeaconBlock_SetStateRoot_Call {
	return &BeaconBlock_SetStateRoot_Call{Call: _e.mock.On("SetStateRoot", _a0)}
}

func (_c *BeaconBlock_SetStateRoot_Call) Run(run func(_a0 bytes.B32)) *BeaconBlock_SetStateRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bytes.B32))
	})
	return _c
}

func (_c *BeaconBlock_SetStateRoot_Call) Return() *BeaconBlock_SetStateRoot_Call {
	_c.Call.Return()
	return _c
}

func (_c *BeaconBlock_SetStateRoot_Call) RunAndReturn(run func(bytes.B32)) *BeaconBlock_SetStateRoot_Call {
	_c.Call.Return(run)
	return _c
}

// SizeSSZ provides a mock function with given fields:
func (_m *BeaconBlock) SizeSSZ() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SizeSSZ")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// BeaconBlock_SizeSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SizeSSZ'
type BeaconBlock_SizeSSZ_Call struct {
	*mock.Call
}

// SizeSSZ is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) SizeSSZ() *BeaconBlock_SizeSSZ_Call {
	return &BeaconBlock_SizeSSZ_Call{Call: _e.mock.On("SizeSSZ")}
}

func (_c *BeaconBlock_SizeSSZ_Call) Run(run func()) *BeaconBlock_SizeSSZ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_SizeSSZ_Call) Return(_a0 int) *BeaconBlock_SizeSSZ_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_SizeSSZ_Call) RunAndReturn(run func() int) *BeaconBlock_SizeSSZ_Call {
	_c.Call.Return(run)
	return _c
}

// UnmarshalSSZ provides a mock function with given fields: buf
func (_m *BeaconBlock) UnmarshalSSZ(buf []byte) error {
	ret := _m.Called(buf)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalSSZ")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(buf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeaconBlock_UnmarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalSSZ'
type BeaconBlock_UnmarshalSSZ_Call struct {
	*mock.Call
}

// UnmarshalSSZ is a helper method to define mock.On call
//   - buf []byte
func (_e *BeaconBlock_Expecter) UnmarshalSSZ(buf interface{}) *BeaconBlock_UnmarshalSSZ_Call {
	return &BeaconBlock_UnmarshalSSZ_Call{Call: _e.mock.On("UnmarshalSSZ", buf)}
}

func (_c *BeaconBlock_UnmarshalSSZ_Call) Run(run func(buf []byte)) *BeaconBlock_UnmarshalSSZ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *BeaconBlock_UnmarshalSSZ_Call) Return(_a0 error) *BeaconBlock_UnmarshalSSZ_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_UnmarshalSSZ_Call) RunAndReturn(run func([]byte) error) *BeaconBlock_UnmarshalSSZ_Call {
	_c.Call.Return(run)
	return _c
}

// Version provides a mock function with given fields:
func (_m *BeaconBlock) Version() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Version")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// BeaconBlock_Version_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Version'
type BeaconBlock_Version_Call struct {
	*mock.Call
}

// Version is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter) Version() *BeaconBlock_Version_Call {
	return &BeaconBlock_Version_Call{Call: _e.mock.On("Version")}
}

func (_c *BeaconBlock_Version_Call) Run(run func()) *BeaconBlock_Version_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_Version_Call) Return(_a0 uint32) *BeaconBlock_Version_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_Version_Call) RunAndReturn(run func() uint32) *BeaconBlock_Version_Call {
	_c.Call.Return(run)
	return _c
}

// NewBeaconBlock creates a new instance of BeaconBlock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBeaconBlock(t interface {
	mock.TestingT
	Cleanup(func())
}) *BeaconBlock {
	mock := &BeaconBlock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
