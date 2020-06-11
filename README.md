### Requirements
Require Go version 1.14!

### Usage
The server directory contains  small HTTP server implementation that simply hosts the files of the current working directory. It also downloads the wasm_exec.js JavaScript bridge from the official golang repository.

```
- go build -o server.bin ./server
- ./server.bin
```

### How to compile
```
- GOARCH=wasm GOOS=js go build -o go.wasm ./components
```

### Screenshot
![Hacker News Example](screenshot.png?raw=true "Hacker News Example")