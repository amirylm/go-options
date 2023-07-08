# Golang Options 

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/amirylm/go-options?tab=doc)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/amirylm/go-options/blob/main/LICENSE)
![Go version](https://img.shields.io/badge/go-1.20-blue.svg)
![Github Actions](https://github.com/amirylm/go-options/actions/workflows/lint.yml/badge.svg?branch=main)
![Github Actions](https://github.com/amirylm/go-options/actions/workflows/test.yml/badge.svg?branch=main)

## Overview

This repository contains just a few lines of code,
created for reusing cross multiple projects but you can easily implment these lines in your own code.

Options pattern allows decoupling and encapsulation of object creation.

Golang's variadic functions are used to accept multiple functions or none.
For example, if we want to create a new object:
```golang
o := New(WithSize(32), WithID("0x111"))
```
Using this approach, we can easily change the internal structure of our options, 
while preserving a stable interface.

## Usage

```shell
go get github.com/amirylm/go-options
```

```golang

import (
    "github.com/amirylm/go-options"
)

func main() {
    hello := New(WithSize(4), WithName("hello"))

    world := New(WithName("world"))

    fmt.Println(hello.name) // hello
	fmt.Println(world.name) // world 
	fmt.Println(world.size) // 0 
}

type myStruct struct {
	size int
	name string
}

func WithSize(size int) options.Option[myStruct] {
	return func(o *myStruct) {
		o.size = size
	}
}

func WithName(name string) options.Option[myStruct] {
	return func(o *myStruct) {
		o.name = name
	}
}

func New(o ...options.Option[myStruct]) *myStruct  {
	return options.Apply(nil, o...)
}
```

## Contributing

Contributions to this repository are welcomed and encouraged! If you find a bug or have an idea for an improvement, please open an issue or submit a pull request.

## License

go-options is an open-source software licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute this library as permitted by the license.
