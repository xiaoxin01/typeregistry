# a library to create type dynamically

## how to use

```go
// 1. define struct
type Student struct {
	Age  int
	Name string
}

// 2. add reflect type
key := AddType(new(Student))

// 3. create struct from registed key
student := Make(key)
```

## benchmark

```bash
go.exe test -benchmem -run=^$ supperxin/typeregistry -bench ^BenchmarkMake$

goos: windows
goarch: amd64
pkg: supperxin/typeregistry
BenchmarkMake-8   	 8823256	       138 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	supperxin/typeregistry	1.551s
```

```bash
go.exe test -benchmem -run=^$ supperxin/typeregistry -bench ^BenchmarkCreateDirectly$

goos: windows
goarch: amd64
pkg: supperxin/typeregistry
BenchmarkCreateDirectly-8   	1000000000	         0.593 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	supperxin/typeregistry	0.857s
```
