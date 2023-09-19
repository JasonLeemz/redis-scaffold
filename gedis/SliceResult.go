package gedis

// SliceResult 任何错误判断都由这处理
type SliceResult struct {
	Result []interface{}
	Err    error
}

type Iterator struct {
	data  []interface{}
	index int
}

func (it *Iterator) HasNext() bool {
	if it.data == nil || len(it.data) == 0 {
		return false
	}

	return it.index < len(it.data)
}

func (it *Iterator) Next() interface{} {
	ret := it.data[it.index]
	it.index++
	return ret
}

type IteratorInterface interface {
	HasNext() bool
	Next() interface{}
}

func (r *SliceResult) Unwrap() []interface{} {
	if r.Err != nil {
		panic(r.Err)
	}

	return r.Result
}

func (r *SliceResult) Iterator() *Iterator {
	return newIterator(r.Result)
}

func newIterator(data []interface{}) *Iterator {
	return &Iterator{
		data:  data,
		index: 0,
	}
}

type SliceResultInterface interface {
	Unwrap() []interface{}
	Iterator() *Iterator
}

func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}
