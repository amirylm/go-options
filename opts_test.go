package options

import "testing"

func TestApply(t *testing.T) {
	opts := Apply(nil, WithName("hello"), WithSize(32))
	if opts == nil {
		t.Fail()
	}
	_ = Apply(opts, WithName("world"))
	if opts == nil {
		t.Fail()
	} else if opts.name != "world" {
		t.Fail()
	} else if opts.size != 32 {
		t.Fail()
	}
}

type myStruct struct {
	size int
	name string
}

func WithSize(size int) Option[myStruct] {
	return func(o *myStruct) {
		o.size = size
	}
}

func WithName(name string) Option[myStruct] {
	return func(o *myStruct) {
		o.name = name
	}
}
