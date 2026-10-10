# 带修莫队
> 确保你有普通莫队的基础

带修莫队顾名思义，在一些询问中有单点修改数组的操作

我们在普通莫队的基础上，增加一个时间戳维度，假设时间戳为 $b.t$，表示这个询问已经经过了 $b.t$ 次操作，注意由于我们的操作起始编号为 $0$，因此第 $b.t$ 次操作还没有经过。

```go
curQ := 0
for qi := range m {
    op := IB()
    if op == 'Q' {
        l, r := II() - 1, II() - 1
        qs = append(qs, query{l / B, r / B, l - 1, r + 1, qi - curQ, curQ})
        curQ++
    } else {
        idx, val := II() - 1, II()
        vs = append(vs, version{idx, val})
    }
}
```

将问题和修改操作编号

```go
slices.SortFunc(qs, func(a, b query) int {
    if a.lid != b.lid {
        return cmp.Compare(a.lid, b.lid)
    }
    if a.rid != b.rid {
        return cmp.Compare(a.rid, b.rid)
    }
    return cmp.Compare(a.t, b.t)
})
```

将所有询问按照左端点分块的编号升序
将所有左端点在同一个块中的询问，按照右端点所在块升序
将左右端点都在同一个块中的询问，按照时间戳升序

```go
upd := func(q query, t int) {
    pos, color := vs[t].pos, vs[t].color
    if q.l < pos && pos < q.r {
        del(nums[pos])
        add(color)
    }
    nums[pos], vs[t].color = vs[t].color, nums[pos]
}
```

经过排序后，时间戳是连续的，因此更新操作也是有序的，不会出现跳过一个更新操作的情况

因此，我们可以直接将旧的数据~~挂转转上回收了~~保存在 $vs[t].color$ 中

```go
for ; t < b.t; t++ {
    upd(b, t)
}

for t > b.t {
    t--
    upd(b, t)
}
```

此时 $t$ 表示我们已经经过了 $t$ 次更新，下一次正好是第 $t$ 次（$t$ 从 $0$ 开始），因此更新时先更新第 $t$ 次，再移动指针

回退时，第 $t$ 次我们还没有经过，最后一次的更新为 $t-1$ 次，因此先移动指针