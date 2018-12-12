ADD - عدد
=========

[![Build Status](https://travis-ci.com/Mazafard/add.svg?branch=master)](https://travis-ci.com/Mazafard/add)
[![GoDoc](https://godoc.org/github.com/Mazafard/add?status.svg)](https://godoc.org/github.com/Mazafard/add)
[![codecov](https://codecov.io/gh/Mazafard/add/branch/master/graph/badge.svg)](https://codecov.io/gh/Mazafard/add)

## ADD Convert Numbers to Words

Extra Package for working with the `Numbers` types in Persian Language in `Go` (Golang).


## API usage

First, `import` package

#### go get
```sh
go get -u  github.com/Mazafard/Add
```
#### gb 
```
gb vendor fetch github.com/Mazafard/Add
```
### glide
```$xslt
glide github.com/Mazafard/Add
```
#### godep
```$xslt
godep github.com/Mazafard/Add
```

## Examples

Only examples of common uses are given below; see the [docs] for the full API.

```go
package main

import(
	"fmt" 
	."github.com/Mazafard/add"
)

func main(){
	
	fmt.Printf(Convert(12))
	
	// outputs دوازده
}
```

## Tests

`go test` is used for testing.

Example:

     go test -v 


## Changelog

See the [changelog] before upgrading.

##### Also, Thank you;
* [Sasan Rose]
* [fzerorubigd]

## License

Copyright (c) 2019 Mazafard.
Licensed under the MIT license.

[changelog]: https://github.com/Mazafard/add/blob/master/changelog.md
[docs]: http://godoc.org/github.com/Mazafard/add
[Sasan Rose]: https://github.com/sasanrose
[fzerorubigd]: https://github.com/fzerorubigd
