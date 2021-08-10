# Golang序章

## 1. 什么是go语言

**Go**（又称 **Golang**）是 Google 的 Robert Griesemer，Rob Pike 及 Ken Thompson 开发的一种静态 、强类型、编译型语言 。Go 语言语法与 C相近，但功能上有：内存安全，GC（垃圾回收），结构形态及 CSP-style 并发计算。

​														-----------百度百科

Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

Go 语言被设计成一门应用于搭载 Web 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。

它是编译型语言。

它考虑了多核计算机的执行特点。（并行编程）

## 特点

1. 简单易学习
2. 开发效率高
3. 执行效率好（号称21世纪的C语言）

兼具效率、性能、安全、健壮等特性

## 应用领域

1. 服务端开发：日志处理、文件系统、监控服务
2. 容器虚拟化：Docker、k8s、Docker Swarm
3. 存储：etcd、TiDB、GroupCache
4. Web开发：http/net、Gin、Echo
5. 区块链：以太坊、fabric
6. 云平台，目前国外很多云平台在采用Go开发，CloudFoundy的部分组建，前VMare的技术总监自己出来搞的apcera云平台。

## 成功的案例

- nsq：bitly开源的消息队列系统，性能非常高，目前他们每天处理数十亿条的消息
- docker:基于lxc的一个虚拟打包工具，能够实现PAAS平台的组建。
- packer:用来生成不同平台的镜像文件，例如VM、vbox、AWS等，作者是vagrant的作者
- skynet：分布式调度框架
- Doozer：分布式同步工具，类似ZooKeeper
- Heka：mazila开源的日志处理系统
- cbfs：couchbase开源的分布式文件系统
- tsuru：开源的PAAS平台，和SAE实现的功能一模一样
- groupcache：memcahe作者写的用于Google下载系统的缓存系统
- god：类似redis的缓存系统，但是支持分布式和扩展性
- gor：网络流量抓包和重放工具

## 语言环境安装

在线下载：<https://golang.org/dl/>

配置go环境变量；

![60125872124](F:\05_Go语言学习\笔记\01_golang入门.assets\1601258721248.png)



cmd输入go version查看版本

![60125756173](F:\05_Go语言学习\笔记\01_golang入门.assets\1601257561739.png)

## Go语言的Hello World

创建一个go语言的工作空间

![60125860811](F:\05_Go语言学习\笔记\01_golang入门.assets\1601258608119.png)

新建src、pkg、bin三个目录

在src目录下新建demo.go文件

```go
package main

import "fmt"

func main() {
   fmt.Println("Hello World !")

}
```

在该文件夹打开cmd 执行go demo.go

![60125933452](F:\05_Go语言学习\笔记\01_golang入门.assets\1601259334521.png)

## 编译

![60126041401](F:\05_Go语言学习\笔记\01_golang入门.assets\1601260414015.png)

输入dir查看目录文件

![60126053594](F:\05_Go语言学习\笔记\01_golang入门.assets\1601260535940.png)

会发现多了一个exe可执行文件，文件默认的名称是当前文件夹的名称

我们可以指定可执行文件名称编译

![60126073290](F:\05_Go语言学习\笔记\01_golang入门.assets\1601260732908.png)

跨平台编译

默认我们`go build`的可执行文件都是当前操作系统可执行的文件，如果需要编译其他平台的go可执行文件，需要先指定平台再编译

```
SET CGO_ENABLED=0  // 禁用CGO，cgo不支持跨平台
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

![60126188170](F:\05_Go语言学习\笔记\01_golang入门.assets\1601261881709.png)

使用sublime打开可以发现是一个二进制文件，拷贝到linux上即可执行

（在执行之前需要查看文件是否有执行权限，不然需要授权）

![60126242203](F:\05_Go语言学习\笔记\01_golang入门.assets\1601262422038.png)



## Go语言项目目录结构



![60125924738](F:\05_Go语言学习\笔记\01_golang入门.assets\1601259247380.png)

## 开发工具（IDE）

免费的VS Code（安装go插件）

收费的Goland

vim

等等

