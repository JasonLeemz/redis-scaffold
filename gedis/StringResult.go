package gedis

// StringResult 任何错误判断都由这处理
type StringResult struct {
	Result string
	Err    error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

func (r *StringResult) Unwrap() string {
	if r.Err != nil {
		panic(r.Err)
	}

	return r.Result
}

func (r *StringResult) Unwrap_OR(defaultValue string) string {
	if r.Err != nil {
		return defaultValue
	}

	return r.Result
}

type StringResultImpl interface {
	Unwrap() string
	Unwrap_OR(defaultValue string) string
}
