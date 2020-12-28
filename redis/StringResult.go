package redis

type StringResult struct {
	*Reply
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{&Reply{Result:result, Err: err}}
}

func (this *StringResult) Unwrap() *Reply {
	return this.Reply
}

func (this *StringResult) Default(v string) *Reply {
	if this.Err != nil {
		this.Result = v
		return this.Reply
	}
	return this.Reply
}
