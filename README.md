# wechat SDK for fasthttp
fork自老司机的项目，把server部分对于net/http的依赖改为使用fasthttp了，方便iris的用户调用，绝大多数用法跟原先一样

## 简介
| 模块  | 描述                     |
|-----:|:-------------------------|
| mp   | 微信公众平台 SDK           |
| corp | 微信企业号 SDK             |
| mch  | 微信商户平台(微信支付) SDK   |

## 联系方式
QQ群: 297489459

## 文档
代码下载下来后，放入 GOPATH 路径的 src 下面，可以在shell(windows 下面是 cmd) 里运行
```sh
godoc -http=:8080
```

然后在浏览器里地址栏输入
```sh
http://localhost:8080/
```
即可查看文档


## 授权(LICENSE)
[wechat is licensed under the Apache Licence, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)
