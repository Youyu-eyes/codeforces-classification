# 位运算 Trick

`res - 1 | 1` 表示奇数不变，偶数减一，相当于
```go
if res & 1 == 0 {
    res--
}
```