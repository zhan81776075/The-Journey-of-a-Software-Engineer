## Go语言概念目录


## Go语言问题

#### Q: Go语言中的数组和C语言的数组的区别是什么，优劣势是什么？

#### Q: Go语言的优劣势是什么? 比起C语言呢? 比起Java语言呢?

#### Q: 切片的好处是什么? 主要引用场景呢?

#### Q：如何进行接口完整性检查
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
这样就做到了强验证的方法。时间
## 
