# golang中的错误处理

## 写在前面

在前面的几篇文章当中，我们主要是学习了Golang当中文件的读写以及数据的编码方式相关的知识。接下来，我们将开始来学习Golang中的**错误处理**。

## Golang的错误处理模式

Go并没有像Java那样的一套try/catch异常处理机制，它不能执行抛异常操作。它使用的是一套defer-panic-and-recover机制。

那么，Golang是怎么处理错误的呢？它的处理方式是这样的：？通过在函数和方法中返回错误对象作为它们的唯一或最后一个返回值——如果返回 nil，则没有错误发生——并且主调（calling）函数总是应该检查收到的错误。而上面提到的`panic and recover` 是用来处理真正的异常（无法预测的错误）而不是普通的错误。

## 定义一个错误

在Go中有一个预先定义好的error类型的接口

```go
type error interface {
    Error() string
}
```

在`errors`这个包当中有一个errorString的结构体实现了这个接口

```go
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

错误值s用来表示异常状态,当你需要一个新的错误类型，都可以用 `errors`包的 `errors.New` 函数接收合适的错误信息来创建。

请看下面这个例子：

```go
package main

import (
	"errors"
	"fmt"
)

func main() {

	hello, err := sayHello("")
	if err!=nil {
		log.Default().Println(err)
		return
	}
	fmt.Println(hello)
}

func sayHello(name string) (string,error) {

	if ""==name {
		return "", errors.New("name 不能是一个空字符串")
	}

	return "Hello "+name,nil
}

```

输出：

```
2021/10/02 14:53:16 name 不能是一个空字符串
```



在这个例子当中，如果函数`sayHello`收到的参数`name`是一个空串的话,函数将返回一个错误，主调函数当中必须对错误进行处理。如果有不同错误条件可能发生，那么可以对实际的错误使用类型断言或类型判断（type-switch），然后根据错误场景做一些补救和恢复操作。

在Golang当中，对于错误类型以及错误变量有以下的命名规范：错误类型以 “Error” 结尾，错误变量以 “err” 或 “Err” 开头。

## 运行时异常与`panic`  

在Golang中，当发生了类似于数组下标越界或者是类型断言失败这样的运行时错误时，Go将会触发运行时`panic` (程序崩溃)。

这里需要注意的一点是，**`panic`的使用条件应当是相当严苛的并且是不能由程序自行恢复的**。发生这样的错误时，意味着程序将不能为我们提供服务，而需要终止程序。

请看下面的例子：

```go
package main

import "fmt"

func main() {

	fmt.Println("发生程序崩溃前")
	panic("程序崩溃了")
    // 下面这条语句将不会被执行
	fmt.Println("已产生panic...")

}
```

运行这个程序将导致它抛出`panic`，它打印出错误信息和`goroutine`痕迹，并以非零状态(非正常状态)退出。

输出：

```
发生程序崩溃前
panic: 程序崩溃了

goroutine 1 [running]:
main.main()
	G:/06_golangProject/golang/src/go_code/err/panic_demo/main.go:8 +0xa5

```

### panic使用场景与panicking

在前面我们介绍了panic。在这里，我们来探讨一个问题：**我们应该给在什么时候使用panic呢？**

我个人理解，panic适用于这样的场景：当发生的错误是我们不知道要怎么处理时（不打算优雅处理时），这种错误必须要中止程序运行了，那么，我们将使用到panic。

例如在我们学习Golang读写数据时，编写的小程序mycat（它用于打印文件内容，[博文点击此处传送](https://code81192.github.io/2021/09/28/golang/11_golang%E4%B8%AD%E7%9A%84%E8%AF%BB%E5%86%99%E6%95%B0%E6%8D%AE%EF%BC%88%E4%B8%AD%EF%BC%89/#more)）当中，程序根据用户输入的参数打开一个指定的文件，当文件不存在时，就可以抛出`panic`.

在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：这个终止过程就是 *panicking*。

**注意：** 不要随意地用 panic 中止程序，应当尽力补救错误让程序能继续执行。 

## 从 panic 中恢复（Recover）

学会了panic的基本使用，我们需要思考一个问题，那就是，对于panic，我们只能中止程序运行吗？有没有从panic中恢复的方法？

这时，我们将使用到内建函数`recover()` 。它用于从panic或错误场景当中恢复，从而使得程序可以从panicking重新获得控制权，停止终止过程进而恢复程序的正常运行。

请看下面的例子：

```go
package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("done...")
		if err:=recover();err!=nil{
			fmt.Println(err)
			func() {
				fmt.Println("end...")
			}()
		}
	}()

	fmt.Println("start...")
	panic("this is a error")
}

```

输出：

```
start...
done...
this is a error
end...
```

**注意：** 

1. `recover` 只能在 defer 修饰的函数中使用：用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 `recover` 会返回 nil，且没有其它效果。
2. **panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止。**

以上就是defer-panic-and-recover机制。你也可以将其理解为像if、for一样的流程控制。它类似于Java当中的`try...catch` 机制。

我们在使用panic时，可以遵循Go 库的原则：即使在包的内部使用了 panic，在它的对外接口（API）中也必须用 recover 处理成返回显式的错误。

## 自定义包中的错误处理和 panicking

在自定义包中的错误处理时，我们遵循以下原则：

1. 在包内部，总是应该和`panic` 中`recover`: 不应该显式的超出包范围的`panic` ()
2. 向包的调用者返回错误值

请看下面的例子：

```go
package main

import (
	"fmt"
)

func main() {
	var name []string
	name = append(name, "韩立","","南宫婉")

	for i, _ := range name {

		err := SayHello(name[i])
		if err != nil {
			fmt.Println(err)
		}
		continue
	}

}

func SayHello(name string) (err error) {
	defer func()  {
		if r := recover(); r != nil {
			//var ok bool
				err = fmt.Errorf("%v",r)
		}
	}()
	// 注意： 抛出panic的函数必须在defer之后调用
	doSayHello(name)
	return nil

}

func doSayHello(name string) {
	if len(name)==0 {
		panic("名字不能是一个空字符串")

	}
	fmt.Printf("hello %s\n", name)


}

```

```go
hello 韩立
名字不能是一个空字符串
hello 南宫婉
```



在这个例子当中，包内从`panic`中`recover`，并返回给调用者错误提示，使得程序可以继续往下执行。重要的事情多说几遍，**panic**的使用应当严格地限制其场景，尽可能地使程序从**panic**中**recover**

## 使用闭包优雅地处理错误

像上面的代码一样，每当调用函数时，必须检查错误是否发生，这将增加代码的重复率，到处充斥着错误检查，这一点都不优雅。那么，在Golang中有没有机制像Java当中一样，可以统一地对错误进行处理呢？

请看下面的例子：













