# golang中的读写数据之Gob

## 写在前面

在上一篇文章[《Golang中的读写数据（下）》](https://code81192.github.io/2021/09/29/golang/12_golang%E4%B8%AD%E7%9A%84%E8%AF%BB%E5%86%99%E6%95%B0%E6%8D%AE%EF%BC%88%E4%B8%8B%EF%BC%89/)当中，我们学会了Golang当中对于JSON数据的解析，接下来，我们将学习Gob编码方式。

## 什么是Gob

**Gob的定义：** Gob是Go自己的以二进制形式序列化和反序列化程序数据的格式，这种数据格式简称之为**Gob** (Go binary)。

它类似于Java语言当中的`Serialization` 。你可以在`encoding` 包中找到它。

## Gob可以做什么

Gob 通常用于远程方法调用（RPC）参数和结果的传输，以及应用程序和机器之间的数据传输。

 那么，它与我们之前普遍用到的JSON有什么不同呢？

Gob因为是 Go自己的以二进制形式序列化和反序列化程序数据的格式，因此呢只能用于纯Go环境当中，并不适用于异构的环境。例如，它可以用于两个Go程序之间的通信。

## Gob的特点

1. Gob 文件或流是完全自描述的：里面包含的所有类型都有一个对应的描述，并且总是可以用 Go 解码，而不需要了解文件的内容。

2. 只有**可导出**的字段会被编码，零值会被忽略。
3. 在解码结构体的时候，只有**同时匹配名称和可兼容类型**的字段才会被解码。
4. 当源数据类型增加新字段后，Gob 解码客户端仍然可以以这种方式正常工作：解码客户端会继续识别以前存在的字段。



## 使用Gob传输数据

和 JSON 的使用方式一样，Gob 使用通用的 `io.Writer` 接口，通过 `NewEncoder()` 函数创建 `Encoder` 对象并调用 `Encode()`；相反的过程使用通用的 `io.Reader` 接口，通过 `NewDecoder()` 函数创建 `Decoder` 对象并调用 `Decode()`。

请看下面的例子：

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// 修仙者
type Immortal struct {
	Name   string
	Age    int
	Gender string
}

type SimpleImmortal struct {
	Name   string
	Age    int
}

var buf  bytes.Buffer

func main() {
	
	var hanli = Immortal{
		Name:   "韩立",
		Age:    18,
		Gender: "男性",
	}

	fmt.Println("发送数据: ",hanli)
	sendMsg(&hanli)
	fmt.Println("buf中的数据：",buf)
	var i SimpleImmortal
	msg, _ := receiveMsg(i)

	fmt.Println("接收到数据：",msg)
}

func sendMsg(immortal *Immortal) error {
	enc :=gob.NewEncoder(&buf)
	return enc.Encode(immortal)
}

func receiveMsg(immortal SimpleImmortal) (SimpleImmortal,error) {
	dec := gob.NewDecoder(&buf)

	return immortal,dec.Decode(&immortal)

}
```

输出：

```go
发送数据:  {韩立 18 男性}
buf中的数据： {[50 255 129 3 1 1 8 73 109 109 111 114 116 97 108 1 255 130 0 1 3 1 4 78 97 109 101 1 12 0 1 3 65 103 101 1 4 0 1 6 71 101 110 100 101 114 1 12 0 0 0 21 255 130 1 6 233 159 169 231 171 139 1 36 1 6 231 148 183 230 128 167 0] 0 0}
接收到数据： {韩立 18}
```



## 写在最后

关于Gob的内容我们就简单介绍到这里。本文当中涉及到的例子可以[点击此处下载](https://github.com/code81192/golang)。如果我的学习笔记能够给你带来帮助，还请多多点赞鼓励。文章如有错漏之处还请各位小伙伴帮忙斧正。

在下一篇文章当中，我们将一起来学习Golang当中的错误处理相关的内容。

