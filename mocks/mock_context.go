// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/emccode/libstorage/api/types (interfaces: Context)

package mocks

import (
	time "time"

	logrus "github.com/Sirupsen/logrus"
	types "github.com/emccode/libstorage/api/types"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// Mock of Context interface
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *_MockContextRecorder
}

// Recorder for MockContext (not exported)
type _MockContextRecorder struct {
	mock *MockContext
}

func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &_MockContextRecorder{mock}
	return mock
}

func (_m *MockContext) EXPECT() *_MockContextRecorder {
	return _m.recorder
}

func (_m *MockContext) Deadline() (time.Time, bool) {
	ret := _m.ctrl.Call(_m, "Deadline")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockContextRecorder) Deadline() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Deadline")
}

func (_m *MockContext) Debug(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Debug", _s...)
}

func (_mr *_MockContextRecorder) Debug(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Debug", arg0...)
}

func (_m *MockContext) Debugf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Debugf", _s...)
}

func (_mr *_MockContextRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Debugf", _s...)
}

func (_m *MockContext) Debugln(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Debugln", _s...)
}

func (_mr *_MockContextRecorder) Debugln(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Debugln", arg0...)
}

func (_m *MockContext) Done() <-chan struct{} {
	ret := _m.ctrl.Call(_m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

func (_mr *_MockContextRecorder) Done() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Done")
}

func (_m *MockContext) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockContextRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockContext) Error(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Error", _s...)
}

func (_mr *_MockContextRecorder) Error(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Error", arg0...)
}

func (_m *MockContext) Errorf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Errorf", _s...)
}

func (_mr *_MockContextRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Errorf", _s...)
}

func (_m *MockContext) Errorln(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Errorln", _s...)
}

func (_mr *_MockContextRecorder) Errorln(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Errorln", arg0...)
}

func (_m *MockContext) Fatal(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Fatal", _s...)
}

func (_mr *_MockContextRecorder) Fatal(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fatal", arg0...)
}

func (_m *MockContext) Fatalf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Fatalf", _s...)
}

func (_mr *_MockContextRecorder) Fatalf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fatalf", _s...)
}

func (_m *MockContext) Fatalln(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Fatalln", _s...)
}

func (_mr *_MockContextRecorder) Fatalln(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fatalln", arg0...)
}

func (_m *MockContext) Info(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Info", _s...)
}

func (_mr *_MockContextRecorder) Info(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Info", arg0...)
}

func (_m *MockContext) Infof(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Infof", _s...)
}

func (_mr *_MockContextRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Infof", _s...)
}

func (_m *MockContext) Infoln(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Infoln", _s...)
}

func (_mr *_MockContextRecorder) Infoln(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Infoln", arg0...)
}

func (_m *MockContext) Join(_param0 context.Context) types.Context {
	ret := _m.ctrl.Call(_m, "Join", _param0)
	ret0, _ := ret[0].(types.Context)
	return ret0
}

func (_mr *_MockContextRecorder) Join(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Join", arg0)
}

func (_m *MockContext) Panic(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Panic", _s...)
}

func (_mr *_MockContextRecorder) Panic(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Panic", arg0...)
}

func (_m *MockContext) Panicf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Panicf", _s...)
}

func (_mr *_MockContextRecorder) Panicf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Panicf", _s...)
}

func (_m *MockContext) Panicln(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Panicln", _s...)
}

func (_mr *_MockContextRecorder) Panicln(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Panicln", arg0...)
}

func (_m *MockContext) Print(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Print", _s...)
}

func (_mr *_MockContextRecorder) Print(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Print", arg0...)
}

func (_m *MockContext) Printf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Printf", _s...)
}

func (_mr *_MockContextRecorder) Printf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Printf", _s...)
}

func (_m *MockContext) Println(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Println", _s...)
}

func (_mr *_MockContextRecorder) Println(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Println", arg0...)
}

func (_m *MockContext) Value(_param0 interface{}) interface{} {
	ret := _m.ctrl.Call(_m, "Value", _param0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

func (_mr *_MockContextRecorder) Value(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Value", arg0)
}

func (_m *MockContext) Warn(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Warn", _s...)
}

func (_mr *_MockContextRecorder) Warn(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Warn", arg0...)
}

func (_m *MockContext) Warnf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Warnf", _s...)
}

func (_mr *_MockContextRecorder) Warnf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Warnf", _s...)
}

func (_m *MockContext) Warning(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Warning", _s...)
}

func (_mr *_MockContextRecorder) Warning(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Warning", arg0...)
}

func (_m *MockContext) Warningf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Warningf", _s...)
}

func (_mr *_MockContextRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Warningf", _s...)
}

func (_m *MockContext) Warningln(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Warningln", _s...)
}

func (_mr *_MockContextRecorder) Warningln(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Warningln", arg0...)
}

func (_m *MockContext) Warnln(_param0 ...interface{}) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Warnln", _s...)
}

func (_mr *_MockContextRecorder) Warnln(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Warnln", arg0...)
}

func (_m *MockContext) WithError(_param0 error) *logrus.Entry {
	ret := _m.ctrl.Call(_m, "WithError", _param0)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

func (_mr *_MockContextRecorder) WithError(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WithError", arg0)
}

func (_m *MockContext) WithField(_param0 string, _param1 interface{}) *logrus.Entry {
	ret := _m.ctrl.Call(_m, "WithField", _param0, _param1)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

func (_mr *_MockContextRecorder) WithField(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WithField", arg0, arg1)
}

func (_m *MockContext) WithFields(_param0 logrus.Fields) *logrus.Entry {
	ret := _m.ctrl.Call(_m, "WithFields", _param0)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

func (_mr *_MockContextRecorder) WithFields(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WithFields", arg0)
}

func (_m *MockContext) WithValue(_param0 interface{}, _param1 interface{}) types.Context {
	ret := _m.ctrl.Call(_m, "WithValue", _param0, _param1)
	ret0, _ := ret[0].(types.Context)
	return ret0
}

func (_mr *_MockContextRecorder) WithValue(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WithValue", arg0, arg1)
}