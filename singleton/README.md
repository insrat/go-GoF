# 单例设计模式

一个类只允许创建一个对象（或者实例），那这个类就是一个单例类，这种设计模式就叫作单例设计模式，简称单例模式。


## 为什么要使用单例

* 处理资源访问冲突
* 表示全局唯一类

### 如何实现一个单例

* 构造函数需要是 private 访问权限的，这样才能避免外部通过 new 创建实例
* 考虑对象创建时的线程安全问题
* 考虑是否支持延迟加载
* 考虑 getInstance() 性能是否高（是否加锁）

### 实现方法

* 饿汉式

在类加载的时候，静态实例就已经创建并初始化好。

```go
type singleton struct{}

var ins *singleton = &singleton{}

func GetIns() *singleton {
	return ins
}
```

* 懒汉式

懒汉式相对于饿汉式的优势是支持延迟加载，但不支持高并发。

```go
type singleton struct{}

var ins *singleton

func GetIns() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}
```

* 双重检测

一种既支持延迟加载、又支持高并发的单例实现方式。

```go
type singleton struct{}

var ins *singleton
var mu sync.Mutex

func GetIns() *singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			ins = &singleton{}
		}
	}
	return ins
}
```

* sync.Once实现

```go
type singleton struct{}

var ins *singleton
var once sync.Once

func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
```


## 单例存在哪些问题

* 单例对 OOP 特性的支持不友好
* 单例会隐藏类之间的依赖关系
* 单例对代码的扩展性不友好

数据库连接池、线程池这类的资源池，最好还是不要设计成单例类。

* 单例对代码的可测试性不友好
* 单例不支持有参数的构造函数


## 有何替代的解决方案

* 通过工厂模式、IOC 容器来保证
* 由程序员自己来保证（自己在编写代码的时候自己保证不要创建两个类对象）


## 集群环境下的分布式单例模式


### 单例模式中的唯一性

单例模式创建的对象是进程唯一的，单例类中对象的唯一性的作用范围是进程内的，在进程间是不唯一的。

### 实现线程唯一的单例

通过一个 HashMap 来存储对象，其中 key 是线程 ID，value 是对象。

### 实现集群环境下的单例

所谓的集群唯一，即不同的进程间共享同一个对象，不能创建同一个类的多个对象，因此需要把这个单例对象序列化并存储到外部共享存储区（如文件、Redis）。
