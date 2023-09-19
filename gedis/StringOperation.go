package gedis

import (
	"context"
	"time"
)

type StringOperationInterface interface {
	Get(key string) *StringResult
	MGet(key ...string) *SliceResult
	Set()
}

type StringOperation struct {
	ctx context.Context
}

func NewStringOperation(ctx context.Context) *StringOperation {
	return &StringOperation{ctx: ctx}
}

// Get key
func (o *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(o.ctx, key).Result())
}

// MGet key1 key2 ...
func (o *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(o.ctx, keys...).Result())
}

func (o *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *InterfaceResult {
	exp := OperationAttrs(attrs).Find(AttrExpire)
	if exp == nil {
		exp = time.Duration(0)
	}

	nx := OperationAttrs(attrs).Find(AttrNX)
	if nx != nil {
		return NewInterfaceResult(Redis().SetNX(o.ctx, key, value, exp.(time.Duration)).Result())
	} else {
		return NewInterfaceResult(Redis().Set(o.ctx, key, value, exp.(time.Duration)).Result())
	}
}
