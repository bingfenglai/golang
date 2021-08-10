# Golang入门学习之map

## 写在前面

之前说好的每周更新一次技术博客，但是由于种种原因一直推迟到现在... ...🐶

有一句话，我一直把它当作我的座右铭，写在这里与诸君共勉：业精于勤荒于嬉，行成于思毁于随。

废话不多说，开始我们今天的学习。

## 什么是map

map是一种无序的、基于键值对<key,value>的数据结构，又称之为字段或关联数组。

在Go语言当中，map是引用类型，必须初始化后才能使用，其内部也是使用散列表来实现的。

## map的定义

### 第一种方式：通过map关键字定义

```go
// var mapname map[keytype]valuetype
var map map[string]string
```

其中：

1. `mapName` 为变量名
2. `keyType` 为key的数据类型。key的数据类型可以是bool、数字、string、指针、channel、接口、结构体和数组
3. `valueType`为值的数据类型。value包含的数据类型与key一样，通常为string、结构体和map

**注意：** 

	1. 声明不会分配内存，需要使用make来初始化后才能使用。
 	2. 对于上述的key是不可以重复的，value则是可以重复的。

#### 举个例子

```go
package main

import "fmt"

func main() {

	// 通过map关键字生命map
	var peopleMap map[string]string

	// 初始化
	peopleMap = make(map[string]string, 4)

	peopleMap["01"] = "喜小乐"
	peopleMap["02"] = "东小贝"
	peopleMap["03"] = "北小楠"
	peopleMap["04"] = "楠小茽"
	
	fmt.Println(peopleMap)
	
}

```

输出：

```
map[01:喜小乐 02:东小贝 03:北小楠 04:楠小茽]
```

### 第二种方式：通过make定义并初始化

```go
package main

import "fmt"

func main() {
    
    // 通过make声明并初始化
	m := make(map[string]string, 4)
	m["no1"] = "向北"
	m["no2"] = "向南"
	m["no3"] = "向东"
	m["no4"] = "向西"

	fmt.Print(m)
	
}
```

输出：

```
map[no1:向北 no2:向南 no3:向东 no4:向西]
```

#### 关于make函数的说明

make（type,size）是一个内置函数，它用于分配并初始化一个类型为切片、映射和通道的对象，其中第一个参数为类型，第二个参数size是分配给type的初始大小，其返回值与实参type的类型一致，而非指针。







