<div align="center">
    <h3> 
        <strong>Grace</strong>
    </h3>
    <p>
        <strong>Grace</strong> is graceful shutdown implementation for Go.
    </p>
</div>



## Installation
```bash
go get github.com/msrexe/grace
```

### Example
```go
func main() {
	go func() {
		grace.ShutdownWithTimeout(5*time.Second, func() {
			log.Println("Application gracefully stopped")
		})
	}()

	for {
		log.Println("Hello World")
		time.Sleep(2 * time.Second)
	}
}
```
For more examples, look at the [example](github.com/msrexe/grace/tree/master/example) directory.
