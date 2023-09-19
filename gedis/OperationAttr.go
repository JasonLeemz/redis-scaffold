package gedis

import "time"

const (
	AttrExpire = "expr"
	AttrNX     = "nx"
)

// OperationAttr 属性
type OperationAttr struct {
	Name  string
	Value interface{}
}

type OperationAttrs []*OperationAttr

func (attrs OperationAttrs) Find(name string) interface{} {
	for _, attr := range attrs {
		if attr.Name == name {
			return attr.Value
		}
	}
	return nil
}

func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{
		Name:  AttrExpire,
		Value: t,
	}
}

func WithNX() *OperationAttr {
	return &OperationAttr{
		Name:  AttrNX,
		Value: struct{}{},
	}
}
