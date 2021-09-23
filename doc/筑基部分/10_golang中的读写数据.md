# golang中的读写数据

## 写在前面

在前面的文章当中，我们学会了如何去搭建Golang开发环境、学习了Golang当中常见的数据结构、学习了基本流程控制语句、学习了函数和结构体等内容，接下来我们将开始学习Golang当中的文件读写。

## 读取用户在控制台的输入

在Golang当中，如何读取用户在控制台的输入呢？在这里，我们可以使用`fmt`包当中`Scan`开头的函数。

请看下面的例子：

```go
package main

import "fmt"

func main() {
	var firstname,lastname string
	fmt.Println("请输入您的姓名：")
	_, _ = fmt.Scanln(&firstname, &lastname)
	fmt.Printf("你好！%s · %s\n", lastname, firstname)

}
```

输出：

```
请输入您的姓名：
韩 立
你好！立 · 韩
```

`Scanln` 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。`Scanf` 与其类似，除了 `Scanf` 的第一个参数用作格式字符串，用来决定如何读取。`Sscan` 和以 `Sscan` 开头的函数则是从字符串读取，除此之外，与 `Scanf` 相同。如果这些函数读取到的结果与您预想的不同，你可以检查成功读入数据的个数和返回的错误。

除此之外，我们也可以使用 `bufio` 包提供的缓冲读取（buffered reader）来读取数据

请看下面的例子：

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	fmt.Println("请输入一段文字")
	inputReader := bufio.NewReader(os.Stdin)
	s,err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("你输入的是：")
		fmt.Println(s)
	}

}
```

输出：

```
请输入一段文字
去年今日此门中 人面桃花相映红。人面不知何处在 桃花依旧笑春风。
你输入的是：
去年今日此门中 人面桃花相映红。人面不知何处在 桃花依旧笑春风。
```

## 文件读写

### 文件读操作

在 Go 语言中，文件使用指向 `os.File` 类型的指针来表示的，也叫做文件句柄。

#### 按行读取

请看下面的例子：

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("G:\\06_golangProject\\golang\\doc\\筑基部分\\10_golang中的读写数据.md")
	defer inputFile.Close()
	if err!=nil {
		fmt.Println(err)
	}
    
	input :=bufio.NewReader(inputFile)
	for{
		readString, err := input.ReadString('\n')
		fmt.Println(readString)
		if err==io.EOF {
			fmt.Println(err)
			return
		}
	}


}
```

在这个例子当中，我们使用`os.Open`打开一个文件，并在循环当中逐行地打印该文件，直到打印完该文件。

#### 带缓冲的文件读取

很不幸的是，在很多情况下，文件的内容不是按行划分的，甚至有时候文件是一个二进制文件。这时，我们应当如何去读取它呢？

请看下面的这个例子：

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	buf := make([]byte,1024)

	inputFile, err := os.Open("G:\\06_golangProject\\golang\\doc\\筑基部分\\10_golang中的读写数据.md")
	defer inputFile.Close()
	if err!=nil {
		fmt.Println(err)
	}
	for  {
		_, err := inputFile.Read(buf)

		if err==io.EOF {
			return
		}
		fmt.Println(string(buf))
	}


}
```

我们定义了一个`[]byte`类型的缓存，在读取文件时，将读到的内容存入这个缓存中并进行打印。这样，我们就不需要去在意文件当中内容是如何划分的了。

