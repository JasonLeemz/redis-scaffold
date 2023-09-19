package gedis

import (
	"context"
)

type StringOperationImpl interface {
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

func (o *StringOperation) Set() {
	//TODO implement me
	panic("implement me")
}
