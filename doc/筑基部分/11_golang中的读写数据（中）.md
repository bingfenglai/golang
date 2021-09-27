# Golang中的读写数据（中）

## 写在前面

在上一篇文章《[Golang中的读写数据（上）](https://code81192.github.io/2021/09/24/golang/10_golang%E4%B8%AD%E7%9A%84%E8%AF%BB%E5%86%99%E6%95%B0%E6%8D%AE(%E4%B8%8A)/)》当中，我们介绍了Golang中一些简单的文件读写、拷贝操作，接下来，我们将继续学习Golang中的读写数据的相关知识点。

## 从命令行读取参数

在一些场景当中，我们在执行一个软件，有时候需要传入一些初始化的信息，例如连接数据库的`username`和`password`等属性。那么，在Golang中是如何读取参数的呢？请看下面的例子:

### 使用`os.Args` 获取参数

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		for i:=1;i<len(os.Args);i++ {
			sayHello(os.Args[i])
		}
	}else {
		fmt.Println("参数为空")
	}

}

func sayHello(name... string) {

	fmt.Println("hello!",name)
}
```

输出：

![image-20210927204103589](https://raw.githubusercontent.com/code81192/art-demo/master/art-img/image-20210927204103589.png)

**注意：**

这个命令行参数会放置在切片 `os.Args[]` 中（以空格分隔），从索引1开始（`os.Args[0]` 放的是程序本身的名字）。函数 `strings.Join` 以空格为间隔连接这些参数。

### 使用`flag.Args`获取参数

```go
package main

import (
	"flag"
	"fmt"
)

func main() {

	flag.Parse()

	for _, arg := range flag.Args() {
		sayHello(arg)
	}


}

func sayHello(name... string) {

	fmt.Println("hello!",name)
}

```

输出：

![image-20210927224440371](https://raw.githubusercontent.com/code81192/art-demo/master/art-img/image-20210927224440371.png)

**注意：**

`flag.Arg(0)` 就是第一个真实的 flag，而不是像 `os.Args(0)` 放置程序的名字。

### 一个例子：使用缓存读取文件与`flag.Args`的综合应用

接触过Linux系统的小伙伴们应该经常会`cat filename`这个命令，接下来，我们将使用Golang实现这个小工具.

#### 需求描述

我们需要一个小工具来打印文件的内容，当文件不存在时给用户一个友好的提示。

#### 功能设计

1. 根据用户输入的文件名读取文件到缓存并分批次打印
2. 当文件不存在时输出"不存在文件 {文件名}"

#### 代码：

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	filename :=flag.Arg(0)

	if filename=="" {
		fmt.Println("Command Usage Format: mycat filename")
		return
	}

	open, err := os.Open(filename)
	defer open.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	myCat(bufio.NewReader(open))
}

func myCat(reader *bufio.Reader) {

	for  {
		buf, err := reader.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s\n", buf)
		if err == io.EOF {
			break
		}
	}
	return
}

```



**测试：**

1. 当文件存在时打印文件内容：

![image-20210927231646066](https://raw.githubusercontent.com/code81192/art-demo/master/art-img/image-20210927231646066.png)

2. 当文件不存在时给出提示：

![image-20210927231918164](https://raw.githubusercontent.com/code81192/art-demo/master/art-img/image-20210927231918164.png)

#### 提个问题：这个程序能满足需求吗？

大家可以思考一下，如果被打印的文件是一个只有一行数据的大文件，会出现什么情况？

因此，为了避免出现这个情况，我们将对`mycat`进行改造。

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	filename :=flag.Arg(0)

	if filename=="" {
		fmt.Println("Command Usage Format: mycat filename")
		return
	}

	file, err := os.Open(filename)
	defer file.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}

	myCatV2(bufio.NewReader(file))
}

func myCatV2(reader *bufio.Reader) {
	buf := make([]byte,512)

	for  {
		n, err := reader.Read(buf)
		fmt.Fprintf(os.Stdout,"%s",buf[0:n])
		if err ==io.EOF {
			break
		}
	}
	return
}
```

在这一版当中，我们指定了缓存区的大小，它是一个512位的`byte`数组。

输出：

![image-20210928002030354](https://raw.githubusercontent.com/code81192/art-demo/master/art-img/image-20210928002030354.png)

我们指定了缓冲区的大小，这样就避免了将整个文件都加载到内存当中。

当然，大家可以继续对程序进行改进，例如引入协程等技术，这便不在本文的讨论当中了。

## 重要的一点：使用`defer`关闭文件

在前面的文章当中，我们介绍了`defer`关键字的作用：他将在函数退出时（return之后）执行其修饰的语句。在这里，我们使用其来在`main`函数退出前关闭文件。

## 写在最后

关于Golang中读取命令行参数以及与文件读写操作的综合应用我们就介绍到这里。本文当中涉及到的例子可以[点击此处下载](https://github.com/code81192/golang)。如果我的学习笔记能够给你带来帮助，还请多多点赞鼓励。文章如有错漏之处还请各位小伙伴帮忙斧正。

在下一篇文章当中我们将一起来学习Golang中的数据格式化以及数据网络传输等内容。





