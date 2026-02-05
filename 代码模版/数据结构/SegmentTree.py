from math import inf
from typing import List, Tuple

# 线段树
class SegmentTree:
    def __init__(self, arr, default = 0):
        if isinstance(arr, int):
            n = arr
            arr = [default] * n
        n = len(arr)
        self._n = n
        self._tree = [0 for _ in range(2 << (n - 1).bit_length())]
        self._build(arr, 1, 0, n - 1)
    
    # 合并函数
    def _merge_val(self, x, y):
        return max(x, y)    # 根据题目改变

    # 节点合并
    def _maintain(self, node):
        self._tree[node] = self._merge_val(self._tree[node * 2], self._tree[node * 2 + 1])
    
    # 建树
    def _build(self, a, node, l, r):
        if l == r:
            self._tree[node] = a[l]
            return
        m = (l + r) // 2
        self._build(a, node * 2, l, m)
        self._build(a, node * 2 + 1, m + 1, r)
        self._maintain(node)
    
    # 更新
    def _update(self, node, l, r, i, x):
        if l == r:
            self._tree[node] = x
            return
        m = (l + r) // 2
        if i <= m:
            self._update(node * 2, l, m, i, x)
        else:
            self._update(node * 2 + 1, m + 1, r, i, x)
        self._maintain(node)
    
    # 查询
    def _query(self, node, l, r, ql, qr):
        if ql <= l and r <= qr:
            return self._tree[node]
        m = (l + r) // 2

        # 一般是最先出现，如果题目是最后出现的，要先递归右子树
        if qr <= m:
            return self._query(node * 2, l, m, ql, qr)
        if ql > m:
            return self._query(node * 2 + 1, m + 1, r, ql, qr)
        
        l_res = self._query(node * 2, l, m, ql, qr)
        r_res = self._query(node * 2 + 1, m + 1, r, ql, qr)
        return self._merge_val(l_res, r_res)

    # 封装函数，减少接口
    def update(self, i, x):
        self._update(1, 0, self._n - 1, i, x)
    
    def query(self, ql, qr):
        return self._query(1, 0, self._n - 1, ql, qr)
    
    def get(self, i):
        return self._query(1, 0, self._n - 1, i, i)


# Lazy 线段树    
class Node:
    __slots__ = 'val', 'todo'

class LazySegmentTree:
    # 懒标记初始值
    _TODO_INIT = 0  # **根据题目修改**

    def __init__(self, arr, default=0):
        # 线段树维护一个长为 n 的数组（下标从 0 到 n-1）
        # arr 可以是 list 或者 int
        # 如果 arr 是 int，视作数组大小，默认值为 default
        if isinstance(arr, int):
            arr = [default] * arr
        n = len(arr)
        self._n = n
        self._tree = [Node() for _ in range(2 << (n - 1).bit_length())]
        self._build(arr, 1, 0, n - 1)

    # 合并两个 val
    def _merge_val(self, a: int, b: int) -> int:
        return a + b  # **根据题目修改**

    # 合并两个懒标记
    def _merge_todo(self, a: int, b: int) -> int:
        return a + b  # **根据题目修改**

    # 把懒标记作用到 node 子树（本例为区间加）
    def _apply(self, node: int, l: int, r: int, todo: int) -> None:
        cur = self._tree[node]
        # 计算 tree[node] 区间的整体变化
        cur.val += todo * (r - l + 1)  # **根据题目修改**
        cur.todo = self._merge_todo(todo, cur.todo)

    # 把当前节点的懒标记下传给左右儿子
    def _spread(self, node: int, l: int, r: int) -> None:
        todo = self._tree[node].todo
        if todo == self._TODO_INIT:  # 没有需要下传的信息
            return
        m = (l + r) // 2
        self._apply(node << 1, l, m, todo)
        self._apply(node << 1 | 1, m + 1, r, todo)
        self._tree[node].todo = self._TODO_INIT  # 下传完毕

    # 合并左右儿子的 val 到当前节点的 val
    def _maintain(self, node: int) -> None:
        self._tree[node].val = self._merge_val(self._tree[node * 2].val, self._tree[node * 2 + 1].val)

    # 用 a 初始化线段树
    # 时间复杂度 O(n)
    def _build(self, a: List[int], node: int, l: int, r: int) -> None:
        self._tree[node].todo = self._TODO_INIT
        if l == r:  # 叶子
            self._tree[node].val = a[l]  # 初始化叶节点的值
            return
        m = (l + r) // 2
        self._build(a, node << 1, l, m)  # 初始化左子树
        self._build(a, node << 1 | 1, m + 1, r)  # 初始化右子树
        self._maintain(node)

    def _update(self, node: int, l: int, r: int, ql: int, qr: int, f: int) -> None:
        if ql <= l and r <= qr:  # 当前子树完全在 [ql, qr] 内
            self._apply(node, l, r, f)
            return
        self._spread(node, l, r)
        m = (l + r) // 2
        if ql <= m:  # 更新左子树
            self._update(node << 1, l, m, ql, qr, f)
        if qr > m:  # 更新右子树
            self._update(node << 1 | 1, m + 1, r, ql, qr, f)
        self._maintain(node)

    def _query(self, node: int, l: int, r: int, ql: int, qr: int) -> int:
        if ql <= l and r <= qr:  # 当前子树完全在 [ql, qr] 内
            return self._tree[node].val
        self._spread(node, l, r)
        m = (l + r) // 2
        if qr <= m:  # [ql, qr] 在左子树
            return self._query(node << 1, l, m, ql, qr)
        if ql > m:  # [ql, qr] 在右子树
            return self._query(node << 1 | 1, m + 1, r, ql, qr)
        l_res = self._query(node << 1, l, m, ql, qr)
        r_res = self._query(node << 1 | 1, m + 1, r, ql, qr)
        return self._merge_val(l_res, r_res)

    # 用 f 更新 [ql, qr] 中的每个 a[i]
    # 0 <= ql <= qr <= n-1
    # 时间复杂度 O(log n)
    def update(self, ql: int, qr: int, f: int) -> None:
        self._update(1, 0, self._n - 1, ql, qr, f)

    # 返回用 _merge_val 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
    # 0 <= ql <= qr <= n-1
    # 时间复杂度 O(log n)
    def query(self, ql: int, qr: int) -> int:
        return self._query(1, 0, self._n - 1, ql, qr)


# 李超线段树
class LiChaoTree:
    def __init__(self, x_left, x_right):
        self.x_left = x_left
        self.x_right = x_right
        self.left = None
        self.right = None
        self.line = (0, inf)
    
    def f(self, line, x):
        return line[0] * x + line[1]
    
    def add_line(self, new_line, l, r):
        mid = (l + r) // 2
        
        left_new_better = self.f(new_line, l) < self.f(self.line, l)
        mid_new_better = self.f(new_line, mid) < self.f(self.line, mid)
        
        if mid_new_better:
            self.line, new_line = new_line, self.line
        
        if r - l == 1:
            return
        
        if left_new_better != mid_new_better:
            if self.left is None:
                self.left = LiChaoTree(l, mid)
            self.left.add_line(new_line, l, mid)
        else:
            if self.right is None:
                self.right = LiChaoTree(mid, r)
            self.right.add_line(new_line, mid, r)
    
    def add_line_range(self, new_line, seg_l, seg_r, l, r):
        if seg_r <= l or r <= seg_l:
            return
        if seg_l <= l and r <= seg_r:
            self.add_line(new_line, l, r)
            return
        
        mid = (l + r) // 2
        if self.left is None:
            self.left = LiChaoTree(l, mid)
        self.left.add_line_range(new_line, seg_l, seg_r, l, mid)
        
        if self.right is None:
            self.right = LiChaoTree(mid, r)
        self.right.add_line_range(new_line, seg_l, seg_r, mid, r)
    
    def query(self, x, l, r):
        res = self.f(self.line, x)
        if r - l == 1:
            return res
        
        mid = (l + r) // 2
        if x < mid:
            if self.left:
                res = min(res, self.left.query(x, l, mid))
        else:
            if self.right:
                res = min(res, self.right.query(x, mid, r))
        return res