你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# adoc


`adoc`是一个文档生成工具，根据golang代码注释，生成README格式的API文档，只能识别特定的API注释语法。

在前后端分离的开发模式下，后台人员往往要提供一个接口文档给前端人员。
虽然已经有了一些开源工具，可以根据注释生成文档，但这些工具往往都很臃肿，很难根据自己的实际开发体验去做定制。

所以我专门定制一个API文档生成工具，以满足自己的日常开发需求。

## 安装

进入你的gopath: `cd ~/go/src`

拉取代码：`git clone git@github.com:day-dreams/adoc.git adoc.zhangnan.xyz/
`

进入目录: `cd adoc.zhangnan.xyz`

构建: `make`

安装: `make install`

卸载: `make clean`

## 教程

* 新建一个golang文件,如`doc.go`

* 在`doc.go`中写好api注释，格式示例如下：
```golang
package main

//@api
//api name: 根据id获取用户详细信息
//api param: id`必填`string`任意合法id`用户id
//api param: pretty`必填`bool`{true,false}`是否美化输出
//api method: GET
//api path:     /api/v1/user
//api return: {id:"1232131",name:"kakaxi"}
//@api end

//@api
//api name: 用户登陆
//api param: email`必填`string`例如xyz.zhangnan@gmail.com`用户的邮箱
//api param: password`必填`string`例如helloworld`用户的密码
//api method: POST
//api path:     /api/v1/user/login
//api return: {code:0,msg:"登陆成功"}
//@api end


```

* 使用adoc生成README格式的api文档,`adoc doc.go`
