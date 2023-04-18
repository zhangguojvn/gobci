# go-obci

<!-- [![GoDoc Reference](https://godoc.org/github.com/mattn/go-oci8?status.svg)](http://godoc.org/github.com/mattn/go-oci8)
[![Build Status](https://travis-ci.org/mattn/go-oci8.svg?branch=master)](https://travis-ci.org/mattn/go-oci8)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattn/go-oci8)](https://goreportcard.com/report/github.com/mattn/go-oci8) -->


## 介绍

Go 的 Oceanbase Oracle 模式驱动

## 安装

将官方提供的  obci-*.rpm 包中的 libobci.a 放入 库安装目录 /lib 目录

如不知道相关目录,可以在编译报错中查看到相关目录

```
[root@test ping]# /usr/local/go/bin/go build .
# github.com/zhangguojvn/gobci/_example/ping
/usr/local/go/pkg/tool/linux_amd64/link: running gcc failed: exit status 1
gcc: error: /go/src/lib/libobci.a: No such file or directory
```

在以上例子中复制 libobci.a 到 /go/src/lib/libobci.a 即可

安装 GCC 后即可拉取驱动

```
go get github.com/zhangguojun/gobci
```

## SQL 例子

SQL 例子可以在 GoDoc 中找到:

https://godoc.org/github.com/zhangguojun/gobci

同时也存在于 _example 中:

https://github.com/zhangguojun/gobci/tree/master/_example

## 特别感谢

Yasuhiro Matsumoto (a.k.a mattn)

Jamil Djadala
