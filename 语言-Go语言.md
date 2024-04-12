## Go语言概念目录


## Go语言问题

### Q: Go语言中的数组和C语言的数组的区别是什么，优劣势是什么？

### Q: Go语言的优劣势是什么? 比起C语言呢? 比起Java语言呢?

### Q: 切片的好处是什么? 主要引用场景呢?

### Q: 如何进行接口完整性检查
Go 语言的编译器并没有严格检查一个对象是否实现了某接口所有的接口方法，如下面这个示例：
```go
type Shape interface {
    Sides() int
    Area() int
}
type Square struct {
    len int
}
func (s* Square) Sides() int {
    return 4
}
func main() {
    s := Square{len: 5}
    fmt.Printf("%d\n",s.Sides())
}
```
可以看到，Square 并没有实现 Shape 接口的所有方法，程序虽然可以跑通，但是这样的编程方式并不严谨，如果我们需要强制实现接口的所有方法，那该怎么办呢？在 Go 语言编程圈里，有一个比较标准的做法：
```go
var _ Shape = (*Square)(nil)
```
声明一个 _ 变量（没人用）会把一个 nil 的空指针从 Square 转成 Shape，这样，如果没有实现完相关的接口方法，编译器就会报错：
```
cannot use (*Square)(nil) (type *Square) as type Shape in assignment: *Square does not implement Shape (missing Area method)
```
这样就做到了强验证的方法。

### Q: 如何处理好时间
对于时间来说，这应该是编程中比较复杂的问题了，相信我，时间是一种非常复杂的事（比如《你确信你了解时间吗？》《关于闰秒》等文章）。而且，时间有时区、格式、精度等问题，其复杂度不是一般人能处理的。所以，一定要重用已有的时间处理，而不是自己干。
在 Go 语言中，你一定要使用 time.Time 和 time.Duration  这两个类型。
- 在命令行上，flag 通过 time.ParseDuration 支持了 time.Duration。
- JSON 中的 encoding/json 中也可以把time.Time 编码成 RFC 3339 的格式。
- 数据库使用的 database/sql 也支持把 DATATIME 或 TIMESTAMP 类型转成 time.Time。
- YAML 也可以使用 gopkg.in/yaml.v2 支持 time.Time 、time.Duration 和 RFC 3339 格式。

如果你要和第三方交互，实在没有办法，也请使用 RFC 3339 的格式。
最后，如果你要做全球化跨时区的应用，一定要把所有服务器和时间全部使用 UTC 时间。

### Q: Slice的扩容机制是什么
GO1.17版本及之前
当新切片需要的容量cap大于两倍扩容的容量，则直接按照新切片需要的容量扩容；
当原 slice 容量 < 1024 的时候，新 slice 容量变成原来的 2 倍；
当原 slice 容量 > 1024，进入一个循环，每次容量变成原来的1.25倍,直到大于期望容量。

GO1.18之后
当新切片需要的容量cap大于两倍扩容的容量，则直接按照新切片需要的容量扩容；
当原 slice 容量 < threshold 的时候，新 slice 容量变成原来的 2 倍；
当原 slice 容量 > threshold，进入一个循环，每次容量增加（旧容量+3*threshold）/4。
