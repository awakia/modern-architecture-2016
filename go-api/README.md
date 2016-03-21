# Go-lang based API Server

Go言語でJSONを返すAPIサーバーを作る場合

## 基本設計

- HTTP Handler: [negroni](https://github.com/codegangsta/negroni)
- JSON rendering: [render](https://github.com/unrolled/render)
- Routing: [httptreemux](https://github.com/dimfeld/httptreemux)


## Dockerで使える形式でコンパイルする

```
GOOS=linux GOARCH=amd64 go build -o bin/server_linux_amd64 server.go
```
