package gedis

type InterfaceResultInterface interface {
	Unwrap() interface{}
	Unwrap_OR(defaultValue interface{}) interface{}
}

// InterfaceResult 任何错误判断都由这处理
type InterfaceResult struct {
	Result interface{}
	Err    error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}

func (r *InterfaceResult) Unwrap() interface{} {
	if r.Err != nil {
		panic(r.Err)
	}

	return r.Result
}

func (r *InterfaceResult) Unwrap_OR(defaultValue interface{}) interface{} {
	if r.Err != nil {
		return defaultValue
	}

	return r.Result
}
