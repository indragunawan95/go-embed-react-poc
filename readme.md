# Embedding React in Go Proof of Concept

## How does it work

Go embed module will serve React App static file from http.
React App will be serve on root or /.
React also can also request to Go Web server.

## Compile to binary

Linux
```
go build main.go
```

Windows
```
env GOOS=windows GOARCH=amd64 go build main.go
```