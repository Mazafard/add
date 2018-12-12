# ADD Convert Numbers to Words

Extra Package for working with the `Numbers` types in Persian Language.


## Use

### go get
```sh
go get -u  github.com/Mazafard/Add
```
### gb 
```
gb vendor fetch github.com/Mazafard/Add
```
### glide
```$xslt
glide github.com/Mazafard/Add
```
###godep
```$xslt
godep github.com/Mazafard/Add
```

## Changelog

See the [changelog] before upgrading.


## Examples

Only examples of common uses are given below; see the [docs] for the full API.

### Convert(int)

Create literal string from numbers.

```GO
package main

import(
	"fmt",
	."github.com/Mazafard/add",
)

fun main(){
	
	fmt.Printf(Convert(12))
	
	// دوازده
}
```



[changelog]: https://github.com/Mazafard/add/blob/master/changelog.md
[docs]: http://godoc.org/github.com/Mazafard/add
