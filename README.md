# tablego
## Introduction
A self module for table file process

Reference by golang [bufio][] page

## Requirements
No.

And you need to require css and js of sort_table when you use it in browser.
```golang
fr, err := os.Open("test.tsv")
tabO := tablego.Input(fr)
for dA := range tabO.Iter() {
	fmt.Println(dA)
}
```

## Usage
### input a bufio by tablego

## Change logs
* 0.0.1

	Initiate the project

[bufio]:	https://github.com/golang/go/tree/master/src/bufio
