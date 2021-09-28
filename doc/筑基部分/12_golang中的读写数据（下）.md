# Golang中的读写数据（下）--JSON数据的编码与解码

## 写在前面

在上一篇文章[《Golang中的读写数据（中）》](https://code81192.github.io/2021/09/28/golang/11_golang%E4%B8%AD%E7%9A%84%E8%AF%BB%E5%86%99%E6%95%B0%E6%8D%AE%EF%BC%88%E4%B8%AD%EF%BC%89/)中，我们学习了在Golang中如何读取命令行参数，接下来我们将学习Golang中的数据格式化。

我们都知道数据要在网络当中进行传输，或者是将其保存到文件当中，就要必须对其编码和解码。常见的编码格式有JSON和XML等。

### 一些概念的补充

**编码：** 从特定数据结构到数据流这一过程

**解码：** 解码是编码的逆过程，即从数据流到数据结构这一过程

**序列化：** 将内存当中的数据抓换成指定的格式的过程，例如将一个Java对象转为一个字符串

接下来我们将介绍在Golang中将数据编码为JSON的相关内容，在这一部分内容当中，我们将使用到`encoding`这个库

## JSON数据操作

为了演示这一部分的内容，我将之前我们学习结构体时用到的结构体抄写在下面

```go
// 修仙者
type Immortal struct {
	Name   string
	Age    int
	Gender string
}
```

然后大家请看下面这个例子：

```go
package main

import (
	"encoding/json"
	"fmt"
)

// 修仙者
type Immortal struct {
	Name   string
	Age    int
	Gender string
}

func main() {

	immortal := &Immortal{
		Name:   "韩立",
		Age:    18,
		Gender: "男性",
	}
	// 将修仙者韩立编码为json的[]byte
	jsonByteImmortal, _ := json.Marshal(immortal)

	fmt.Printf("%s\n", jsonByteImmortal)
}

```

输出：

```json
{"Name":"韩立","Age":18,"Gender":"男性"}
```

上面用到的`json.Marshal`函数的函数签名是`func Marshal(v interface{}) ([]byte, error)` 它返回的是byte数组，因此打印时需要指定格式。

JSON 与 Go 类型对应如下：

- bool 对应 JSON 的 boolean
- float64 对应 JSON 的 number
- string 对应 JSON 的 string
- nil 对应 JSON 的 null

不是所有的数据都可以编码为 JSON 类型：只有验证通过的数据结构才能被编码：

- JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map[string]T（T是 `json` 包中支持的任何类型）
- Channel，复杂类型和函数类型不能被编码
- 不支持循环数据结构；它将引起序列化进入一个无限循环
- 指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）

## 反序列化操作

在Golang中如何将一个JSON转换为Golang中的数据结构呢？

请看下面的例子：

```go
package main

import (
	"encoding/json"
	"fmt"
)

// 修仙者
type Immortal struct {
	Name   string
	Age    int
	Gender string
}


func main() {

	immortal := &Immortal{
		Name:   "韩立",
		Age:    18,
		Gender: "男性",
	}

	jsonImmortal, _ := json.Marshal(immortal)


	fmt.Printf("%s\n", jsonImmortal)

	// 1. 事先知道json对应的数据类型时
	 var jsonValue Immortal

	json.Unmarshal(jsonImmortal, &jsonValue)

	fmt.Println("name",jsonValue.Name)
	fmt.Println("age",jsonValue.Age)
	fmt.Println("gender",jsonValue.Gender)

	// 2. 不知道json对应的数据结构
	var m interface{}
	json.Unmarshal(jsonImmortal,&m)

	jsonMap := m.(map[string]interface{})
	for key, value := range jsonMap {
		printJson(key,value)
	}


}

func printJson(key string, value interface{}) {

		switch value.(type) {
		case string:
			fmt.Println(key,"value is a string: ",value)
		case float64:
			fmt.Println(key,"value is int type: ",value)
		case []interface{}:
			fmt.Println(key,"value is a array",value)
		case map[string]interface{}:
			m:= value.(map[string]interface{})
			for k, v := range m {
				printJson(k,v)
			}
			
		}


}
```

输出：

```
{"Name":"韩立","Age":18,"Gender":"男性"}
name 韩立
age 18
gender 男性
Name value is a string:  韩立
Age value is int type:  18
Gender value is a string:  男性
```

在这个例子当中，存在着两种情况：

第一种情况：**我们事先知道JSON数据对应的数据结构**，则调用`json.Unmarshal`函数将其解码（也可以理解为反序列化）并存入该数据结构指针变量指向的内存地址当中；

第二种情况：**我们事先不知道JSON数据对应的数据结构**，则可以使用**类型断言**技术得到JSON数据当中`key: value`对应的值。

## 解码以及编码JSON数据流

json 包提供 Decoder 和 Encoder 类型来支持常用 JSON 数据流读写。NewDecoder 和 NewEncoder 函数分别封装了 io.Reader 和 io.Writer 接口。

```go
func NewDecoder(r io.Reader) *Decoderfunc NewEncoder(w io.Writer) *Encoder
```

要想把 JSON 直接写入文件，可以使用 json.NewEncoder 初始化文件（或者任何实现 io.Writer 的类型），并调用 Encode()；反过来与其对应的是使用 json.NewDecoder 和 Decode() 函数：

```go
func NewDecoder(r io.Reader) *Decoderfunc (dec *Decoder) Decode(v interface{}) error
```

## 写在最后

在Golang中解析转换JSON数据的内容我们就简单介绍到这里。本文当中涉及到的例子可以[点击此处下载](https://github.com/code81192/golang)。如果我的学习笔记能够给你带来帮助，还请多多点赞鼓励。文章如有错漏之处还请各位小伙伴帮忙斧正。

在下一篇文章当中，我们将一起来学习Golang当中一种独有的编码格式`Gob` 

