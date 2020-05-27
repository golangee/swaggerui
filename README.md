# swaggerui [![GoDoc](https://godoc.org/github.com/golangee/swaggerui?status.svg)](http://godoc.org/github.com/golangee/swaggerui)  
contains an embedded and standalone swagger-ui http handler.

## example

```go
	http.HandleFunc("/swaggerui/", Handler("/swaggerui/", mySpec))
	err := http.ListenAndServe(":8081", nil)
	if err != nil{
		panic(err)
	}
```