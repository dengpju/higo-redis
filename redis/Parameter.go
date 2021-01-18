package redis

import (
	"time"
)

const (
	PARAM_EXPIRE = "EX"
	PARAM_PX     = "PX"
	PARAM_NX     = "NX"
	PARAM_XX     = "XX"
	TIMEOUT      = "timeout"
)

type Parameter struct {
	Name  string
	Value interface{}
}

func NewParameter(name string, value interface{}) *Parameter {
	return &Parameter{Name: name, Value: value}
}

type Parameters []*Parameter

func (this Parameters) Find(name string) interface{} {
	for _, p := range this {
		if p.Name == name {
			return p.Value
		}
	}
	return nil
}

func (this Parameters) Handle(args ...interface{}) []interface{} {
	params := make([]interface{}, 0)
	for _, p := range this {
		if p.Name != "" {
			params = append(params, p.Name)
		}
		params = append(params, p.Value)
	}
	return append(args, params...)
}

func WithEvalNumkeys(v interface{}) *Parameter {
	return NewParameter("", v)
}

func WithEvalKey(v interface{}) *Parameter {
	return NewParameter("", v)
}

func WithEvalArg(v interface{}) *Parameter {
	return NewParameter("", v)
}

func WithKey(v interface{}) *Parameter {
	return NewParameter("", v)
}

func WithExpire(t time.Duration) *Parameter {
	return NewParameter(PARAM_EXPIRE, int(t))
}

func WithNx() *Parameter {
	return NewParameter(PARAM_NX, "")
}

func WithXx() *Parameter {
	return NewParameter(PARAM_XX, "")
}

func WithTimeout(t time.Duration) *Parameter {
	return NewParameter(TIMEOUT, int(t))
}
