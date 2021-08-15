# Golang入门学习之map

## 写在前面

之前说好的每周更新一次技术博客，但是由于种种原因一直推迟到现在... ...🐶

有一句话，我一直把它当作我的座右铭，写在这里与诸君共勉：业精于勤荒于嬉，行成于思毁于随。

废话不多说，开始我们今天的学习。

## 什么是map

map是一种无序的、基于键值对<key,value>的数据结构，又称之为字段或关联数组。

在Go语言当中，map是引用类型，必须初始化后才能使用，其内部也是使用散列表来实现的。

## map的定义与使用

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

### 第三种方式 声明时直接赋值

```go
m := map[string]string{
		"小明":"小朋友",
		"小张": "是个大人",
		"小李": "是个司机",
		"小王": "家住隔壁",
	}
```

## map的CRUD

#### 添加元素操作

```go
// 语法
map[key] = value
```



```go
func main() {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"

	fmt.Println(m)
}
```

**注意：** 如果key不存在，就是增加，如果key存在，则是修改操作

### 修改元素操作

```go
func main() {

	var m =  map[string]string {
		"小王":"一听家就住在不远",
	}

	fmt.Println("修改前：",m)
	m["小王"] = "家住在隔壁"

	fmt.Println("修改后", m)
}
```

输出：

```go
修改前： map[小王:一听家就住在不远]
修改后 map[小王:家住在隔壁]
```

### 删除操作

delete(m map[Type]Type1, key Type)是一个内建函数，它将删除map中指定key的元素。当key不存在时，则不进行删除。

```go
func main() {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"

	fmt.Println("删除前： ",m)
    
	// delete(m map[Type]Type1, key Type)是一个内建函数，它将删除map中指定key的元素。当key不存在时，则不进行删除
	delete(m, "小王")

	fmt.Println("删除后： ", m)

}
```

输出：

```go
删除前：  map[小张:一听就是个大人 小明:一听就是个小朋友 小李:一听就是个司机 小王:一听家就住在不远]
删除后：  map[小张:一听就是个大人 小明:一听就是个小朋友 小李:一听就是个司机]
```

### 查询操作

```go
func main () {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"
	
    // 如果map中存在key为“小王”的元素，则ok为true,value为对应的元素 
	value,ok := m["小王"]

	if ok {
		fmt.Println(value)
	}

}
```

输出：

```
一听家就住在不远
```

## map的遍历

#### 使用`for range`遍历`map`的`key`和`value`

```go
func main() {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"

	for key, value := range m {
		fmt.Println(key,value)
	}

}
```

输出：

```
小张 一听就是个大人
小李 一听就是个司机
小王 一听家就住在不远
小明 一听就是个小朋友
```

#### 使用`for range`遍历`map`的`key` 

```go
func traverse2() {
	var m  map[string]string
	m = make(map[string]string,4)

	// 添加元素
	m["小明"] = "一听就是个小朋友"
	m["小张"] = "一听就是个大人"
	m["小李"] = "一听就是个司机"
	m["小王"] = "一听家就住在不远"

	for s := range m {
		fmt.Println(s)
	}

}
```

输出：

```
小王
小明
小张
小李
```

#### 按照一定的顺序遍历map

在前文当中我们说到，map它是无序的，那么，我们应当如何对它进行有序的遍历呢？















