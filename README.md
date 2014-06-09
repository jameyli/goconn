goconn
======

go 语言练习项目

将客户端的并发接入转化为后端的串行传输

初步想法，每个接入一个goroutine()

这里要厘清对象和goroutine的概念


How To Run?
-----------

* 启动 connsvr

    cd connsvr;
    go run *.go 6666 8888

* 启动 后端

    go run client.go 8888

* 启动 客户端

    go run client.go 6666




