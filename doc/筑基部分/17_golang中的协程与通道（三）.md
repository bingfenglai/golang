# golang中的协程与通道（三）

## 写在前面

在前一篇文章《[golang中的协程与通道（二） ](https://bingfenglai.github.io/2021/12/15/golang/16_golang中的协程与通道（二）/)》当中，我们初步学习了通道的相关内容，并结合协程进行了简单的应用，接下来，我们将进一步学习该知识点。



## 使用 select 切换协程

从不同的并发执行的协程中获取值可以通过关键字`select`来完成，它和`switch`控制语句非常相似，也被称作通信开关。

```go
package main

import "fmt"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	select {
	case i := <-ch1:
		fmt.Println(i)
	case j := <-ch2:
		fmt.Println(j)
	default:
		fmt.Println(" default")

	}
}
```

关于`select` 使用的说明：

1. `default` 语句是可选的；

2. 在任何一个 case 中执行 `break` 或者 `return`，select 就结束了；

3. `select` 做的就是：选择处理列出的多个通信情况中的一个：

   3.1 如果通道都属于阻塞状态，则会等待直到其中一个可以处理（没有default的情况下）

   3.2 如果多个可以处理，则随机选择其中一个

   3.3 如果没有通道操作可以处理了并且写了`default` 语句，则会执行`default`语句：`default` 永远是可运行的。

