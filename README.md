# SMGP 3.0
[![Build Status](https://github.com/yedamao/go_smgp/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/yedamao/go_smgp/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/yedamao/go_smgp)](https://goreportcard.com/report/github.com/yedamao/go_smgp)
[![codecov](https://codecov.io/gh/yedamao/go_smgp/branch/master/graph/badge.svg)](https://codecov.io/gh/yedamao/go_smgp)

go_smgp是为SP设计实现的smgp3.0协议开发工具包。包括smgp协议包和命令行工具。

## 安装
```
go get github.com/yedamao/go_smgp/...
cd $GOPATH/src/github.com/yedamao/go_smgp && make
```

## Smgp协议包

### support operation

- [x] Login
- [x] LoginResp
- [x] Submit
- [x] SubmitResp
- [x] Deliver
- [x] DeliverResp
- [x] ActiveTest
- [x] ActiveTestResp
- [x] Exit
- [x] ExitResp

### Example
参照cmd/transmitter/main.go, cmd/receiver/main.go

## 命令行工具

### mockserver
SMGW短信网关模拟器

```
Usage of ./bin/mockserver:
  -addr string
        addr(本地监听地址) (default ":8890")
```

### transmitter
提交单条短信至短信网关

```
Usage of ./bin/transmitter:
  -addr string
        smgw addr(运营商地址) (default ":8890")
  -clientID string
        登陆账号
  -dest-number string
        接收手机号码, 86..., 多个使用，分割
  -msg string
        短信内容
  -secret string
        登陆密码
  -sp-number string
        SP的接入号码
  -spID string
        企业代码
```

### receiver
接收运营商回执状态/上行短信

 ```
 Usage of ./bin/receiver:
  -addr string
        smgw addr(运营商地址) (default ":8890")
  -clientID string
        登陆账号
  -secret string
        登陆密码
 ```
 
 
