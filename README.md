# 运行go

1. 编译：`go build 文件名`
2. 运行: `go run 文件名`

# 初始化项目

1. `go mod init 项目名`
2. `go.mod`: 项目信息,包依赖说明
3. `go.sum`: 校验每一个引入包的哈希值


# 变量与常量

## 变量

```go
var v1 init = 1
var v2 = 1 
v2 := 1 //自动类型推断
```

## 常量

```go
const c1 = 1
```

> iota : 一个常量生成器，用于创建一系列常量。在 const 声明中，iota 的值从 0 开始，并在每行常量声明中自动递增。它通常用于枚举类型或生成一系列相关的常量。

- 重置 iota：在每个 const 声明块中，iota 会被重置为 0。
- 空白标识符：使用 _ = iota 可以跳过某些常量的值。

```go
package main

import (
    "fmt"
)

const (
    _ = iota // 忽略第一个 iota 值
    Red
    Green
    Blue
)

func main() {
    fmt.Println(Red)   // 输出 1
    fmt.Println(Green) // 输出 2
    fmt.Println(Blue)  // 输出 3
}


const (
    Readable = 1 << iota // 1 << 0 = 1
    Writable              // 1 << 1 = 2
    Executable            // 1 << 2 = 4
)

func main() {
    fmt.Println(Readable)   // 输出 1
    fmt.Println(Writable)   // 输出 2
    fmt.Println(Executable) // 输出 4
}
```