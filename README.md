# swaggerui
contains an embedded and standalone swagger-ui http handler.

## example

```go
	http.HandleFunc("/swaggerui/", Handler("/swaggerui/", mySpec))
	err := http.ListenAndServe(":8081", nil)
	if err != nil{
		panic(err)
	}
```