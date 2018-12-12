ADD - عدد
=========

[![Build Status](https://travis-ci.org/Mazafard/add.svg?branch=master)](https://travis-ci.org/mazafard/add)
[![GoDoc](https://godoc.org/github.com/Mazafard/add?status.svg)](https://godoc.org/https://travis-ci.org/Mazafard/add)

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

## Changelog

See the [changelog] before upgrading.

## License

MIT





[changelog]: https://github.com/Mazafard/add/blob/master/changelog.md
[docs]: http://godoc.org/github.com/Mazafard/add