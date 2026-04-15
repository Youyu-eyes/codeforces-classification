from math import inf
from typing import List, Tuple, Callable

# 线段树
class SegmentTree:
    def __init__(self, arr, default = 0) -> None:
        if isinstance(arr, int):
            arr = [default] * arr
        n = len(arr)
        self._n: int = n
        self._tree: List[int] = [0 for _ in range(2 << (n - 1).bit_length())]
        self._build(arr, 1, 0, n - 1)

    def _merge_val(self, x: int, y: int) -> int:
        return max(x, y)  # **根据题目修改** 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等

    def _maintain(self, node: int) -> None:
        self._tree[node] = self._merge_val(self._tree[node << 1], self._tree[node << 1 | 1])

    def _build(self, a: List[int], node: int, l: int, r: int) -> None:
        if l == r:
            self._tree[node] = a[l]
            return
        m = (l + r) >> 1
        self._build(a, node << 1, l, m)
        self._build(a, node << 1 | 1, m + 1, r)
        self._maintain(node)

    def _update(self, node: int, l: int, r: int, i: int, x: int) -> None:
        if l == r:
            self._tree[node] = x
            return
        m = (l + r) >> 1
        if i <= m:
            self._update(node << 1, l, m, i, x)
        else:
            self._update(node << 1 | 1, m + 1, r, i, x)
        self._maintain(node)

    def _query(self, node: int, l: int, r: int, ql: int, qr: int) -> int:
        if ql <= l and r <= qr:
            return self._tree[node]
        m = (l + r) >> 1
        if qr <= m:
            return self._query(node << 1, l, m, ql, qr)
        if ql > m:
            return self._query(node << 1 | 1, m + 1, r, ql, qr)
        l_res = self._query(node << 1, l, m, ql, qr)
        r_res = self._query(node << 1 | 1, m + 1, r, ql, qr)
        return self._merge_val(l_res, r_res)

    def _find_first(self, node: int, l: int, r: int, ql: int, qr: int, f: Callable[[int], bool]) -> int:
        if r < ql or l > qr or not f(self._tree[node]):
            return -1
        if l == r:
            return l
        m = (l + r) >> 1
        left_res = self._find_first(node << 1, l, m, ql, qr, f)
        if left_res != -1:
            return left_res
        return self._find_first(node << 1 | 1, m + 1, r, ql, qr, f)

    def _find_last(self, node: int, l: int, r: int, ql: int, qr: int, f: Callable[[int], bool]) -> int:
        if r < ql or l > qr or not f(self._tree[node]):
            return -1
        if l == r:
            return l
        m = (l + r) >> 1
        right_res = self._find_last(node << 1 | 1, m + 1, r, ql, qr, f)
        if right_res != -1:
            return right_res
        return self._find_last(node << 1, l, m, ql, qr, f)

    # ---------- 封装函数 ---------- #
    # 单点更新
    def update(self, i: int, x: int) -> None:
        self._update(1, 0, self._n - 1, i, x)

    # 区间查询
    def query(self, ql: int, qr: int) -> int:
        return self._query(1, 0, self._n - 1, ql, qr)

    # 单点查询
    def get(self, i: int) -> int:
        return self._query(1, 0, self._n - 1, i, i)

    # 返回 [ql, qr] 之间 第一个  满足 f 的数，不存在返回 -1
    def find_first(self, ql: int, qr: int, f: Callable[[int], bool]) -> int:
        return self._find_first(1, 0, self._n - 1, ql, qr, f)

    # 返回 [ql, qr] 之间 最后一个满足 f 的数，不存在返回 -1
    def find_last(self, ql: int, qr: int, f: Callable[[int], bool]) -> int:
        return self._find_last(1, 0, self._n - 1, ql, qr, f)

# Lazy 线段树
class Node:
    __slots__ = 'val', 'todo'

class LazySegmentTree:
    _TODO_INIT = 0   # 懒标记初始值（**根据题目修改**）

    def __init__(self, arr, default = 0):
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

    # 将 todo 作用到节点 node 表示的区间 [l, r]
    def _apply(self, node: int, l: int, r: int, todo: int) -> None:
        cur = self._tree[node]
        cur.val += todo * (r - l + 1)   # **根据题目修改**
        cur.todo = self._merge_todo(todo, cur.todo)

    # 下传懒标记
    def _spread(self, node: int, l: int, r: int) -> None:
        todo = self._tree[node].todo
        if todo == self._TODO_INIT:
            return
        m = (l + r) >> 1
        self._apply(node << 1, l, m, todo)
        self._apply(node << 1 | 1, m + 1, r, todo)
        self._tree[node].todo = self._TODO_INIT

    def _maintain(self, node: int) -> None:
        self._tree[node].val = self._merge_val(self._tree[node << 1].val, self._tree[node << 1 | 1].val)

    def _build(self, a: List[int], node: int, l: int, r: int) -> None:
        self._tree[node].todo = self._TODO_INIT
        if l == r:
            self._tree[node].val = a[l]
            return
        m = (l + r) >> 1
        self._build(a, node << 1, l, m)
        self._build(a, node << 1 | 1, m + 1, r)
        self._maintain(node)

    def _update(self, node: int, l: int, r: int, ql: int, qr: int, f: int) -> None:
        if ql <= l and r <= qr:
            self._apply(node, l, r, f)
            return
        self._spread(node, l, r)
        m = (l + r) >> 1
        if ql <= m:
            self._update(node << 1, l, m, ql, qr, f)
        if qr > m:
            self._update(node << 1 | 1, m + 1, r, ql, qr, f)
        self._maintain(node)

    def _query(self, node: int, l: int, r: int, ql: int, qr: int) -> int:
        if ql <= l and r <= qr:
            return self._tree[node].val
        self._spread(node, l, r)
        m = (l + r) >> 1
        if qr <= m:
            return self._query(node << 1, l, m, ql, qr)
        if ql > m:
            return self._query(node << 1 | 1, m + 1, r, ql, qr)
        l_res = self._query(node << 1, l, m, ql, qr)
        r_res = self._query(node << 1 | 1, m + 1, r, ql, qr)
        return self._merge_val(l_res, r_res)

    def _find_first(self, node: int, l: int, r: int, ql: int, qr: int, f: Callable[[int], bool]) -> int:
        if r < ql or l > qr or not f(self._tree[node].val):
            return -1
        if l == r:
            return l
        self._spread(node, l, r)
        m = (l + r) >> 1
        left_res = self._find_first(node << 1, l, m, ql, qr, f)
        if left_res != -1:
            return left_res
        return self._find_first(node << 1 | 1, m + 1, r, ql, qr, f)

    def _find_last(self, node: int, l: int, r: int, ql: int, qr: int, f: Callable[[int], bool]) -> int:
        if r < ql or l > qr or not f(self._tree[node].val):
            return -1
        if l == r:
            return l
        self._spread(node, l, r)
        m = (l + r) >> 1
        right_res = self._find_last(node << 1 | 1, m + 1, r, ql, qr, f)
        if right_res != -1:
            return right_res
        return self._find_last(node << 1, l, m, ql, qr, f)

    # ---------- 封装函数 ---------- #
    # 区间更新
    def update(self, ql: int, qr: int, f: int) -> None:
        self._update(1, 0, self._n - 1, ql, qr, f)

    # 区间查询
    def query(self, ql: int, qr: int) -> int:
        return self._query(1, 0, self._n - 1, ql, qr)

    # 单点查询
    def get(self, i: int) -> int:
        return self._query(1, 0, self._n - 1, i, i)

    # 返回 [ql, qr] 之间 第一个  满足 f 的数，不存在返回 -1
    def find_first(self, ql: int, qr: int, f: Callable[[int], bool]) -> int:
        return self._find_first(1, 0, self._n - 1, ql, qr, f)

    # 返回 [ql, qr] 之间 最后一个满足 f 的数，不存在返回 -1
    def find_last(self, ql: int, qr: int, f: Callable[[int], bool]) -> int:
        return self._find_last(1, 0, self._n - 1, ql, qr, f)


# 李超线段树
eps = 1e-10

class Line:
    def __init__(self, k=0, b=inf, idx=0):
        # 最大值：def __init__(self, k=0, b=-inf, idx=0):
        self.k = k
        self.b = b
        self.idx = idx
    
    def calc(self, x):
        return self.k * x + self.b

class LiChaoSegmentTree:
    def __init__(self, n):
        self.n = n
        self.tree = [Line() for _ in range(4 * (n + 1))]
    
    def _better(self, a, b, x):
        if a.idx == 0:
            return b
        if b.idx == 0:
            return a
        va = a.calc(x)
        vb = b.calc(x)
        if va - vb < -eps:
            return a
        # 最大值：if va - vb > eps:
        if vb - va < -eps:
            return b
        # 最大值：if vb - va > eps:
        return a if a.idx < b.idx else b
    
    def _update(self, node, l, r, new_line):
        if l > r:
            return
        old_line = self.tree[node]
        mid = (l + r) >> 1
        better_at_mid = new_line.calc(mid) < old_line.calc(mid) - eps
        # 最大值：better_at_mid = new_line.calc(mid) > old_line.calc(mid) + eps
        if better_at_mid:
            self.tree[node], new_line = new_line, old_line
        if l == r:
            return
        better_at_left = new_line.calc(l) < old_line.calc(l) - eps
        # 最大值：better_at_left = new_line.calc(l) > old_line.calc(l) + eps
        better_at_right = new_line.calc(r) < old_line.calc(r) - eps
        # 最大值：better_at_right = new_line.calc(r) > old_line.calc(r) + eps
        if better_at_left != better_at_mid:
            self._update(node << 1, l, mid, new_line)
        elif better_at_right != better_at_mid:
            self._update(node << 1 | 1, mid + 1, r, new_line)
    
    def _insert_line(self, node, l, r, ql, qr, line):
        if l > qr or r < ql:
            return
        if ql <= l and r <= qr:
            self._update(node, l, r, line)
            return
        mid = (l + r) >> 1
        self._insert_line(node << 1, l, mid, ql, qr, line)
        self._insert_line(node << 1 | 1, mid + 1, r, ql, qr, line)
    
    def _query(self, node, l, r, x):
        if l == r:
            return self.tree[node]
        mid = (l + r) >> 1
        res = self.tree[node]
        child_res = self._query(node << 1, l, mid, x) if x <= mid else self._query(node << 1 | 1, mid + 1, r, x)
        return self._better(res, child_res, x)
    
    # ---------- 封装函数 ---------- #
    def insert(self, l, r, line):
        self._insert_line(1, 0, self.n, l, r, line)
    
    def query(self, x):
        line = self._query(1, 0, self.n, x)
        if line.idx == 0:
            return inf
            # 最大值：return -inf
        return line.calc(x)
    

# 动态开点李超线段树
eps = 1e-10

class Line:
    def __init__(self, k=0, b=inf, idx=0):
        # 最大值：def __init__(self, k=0, b=-inf, idx=0):
        self.k = k
        self.b = b
        self.idx = idx

    def calc(self, x):
        return self.k * x + self.b


class Node:
    __slots__ = ('lc', 'rc', 'line')
    def __init__(self):
        self.lc = None
        self.rc = None
        self.line = Line()


class DynamicLiChaoTree:
    def __init__(self, minX, maxX):
        self.minX = minX
        self.maxX = maxX
        self.root = None

    def _better(self, a, b, x):
        if a.idx == 0:
            return False
        if b.idx == 0:
            return True
        va = a.calc(x)
        vb = b.calc(x)
        if va - vb < -eps:
            return True
        # 最大值：if va - vb > eps: return True
        if vb - va < -eps:
            return False
        # 最大值：if vb - va > eps: return False
        return a.idx < b.idx

    def _update(self, p, l, r, new_line):
        if p is None:
            p = Node()
            p.line = new_line
            return p
        mid = (l + r) >> 1
        left_better = self._better(new_line, p.line, l)
        mid_better = self._better(new_line, p.line, mid)

        if mid_better:
            p.line, new_line = new_line, p.line
            mid_better = self._better(new_line, p.line, mid)
            left_better = self._better(new_line, p.line, l)

        if l == r:
            return p

        if left_better != mid_better:
            p.lc = self._update(p.lc, l, mid, new_line)
        else:
            p.rc = self._update(p.rc, mid + 1, r, new_line)
        return p

    def _insert_segment(self, p, l, r, ql, qr, line):
        if p is None:
            p = Node()
        if ql <= l and r <= qr:
            return self._update(p, l, r, line)
        mid = (l + r) >> 1
        if ql <= mid:
            p.lc = self._insert_segment(p.lc, l, mid, ql, qr, line)
        if qr > mid:
            p.rc = self._insert_segment(p.rc, mid + 1, r, ql, qr, line)
        return p

    def _query(self, p, l, r, x):
        if p is None:
            return inf
            # 最大值：return -inf
        res = inf if p.line.idx == 0 else p.line.calc(x)
        # 最大值：res = -inf if p.line.idx == 0 else p.line.calc(x)
        if l == r:
            return res
        mid = (l + r) >> 1
        child_res = self._query(p.lc, l, mid, x) if x <= mid else self._query(p.rc, mid + 1, r, x)
        return min(res, child_res)
        # 最大值：return max(res, child_res)

    def _merge(self, p, q, l, r):
        if p is None:
            return q
        if q is None:
            return p
        if l == r:
            if self._better(q.line, p.line, l):
                p.line = q.line
            return p
        mid = (l + r) >> 1
        p.lc = self._merge(p.lc, q.lc, l, mid)
        p.rc = self._merge(p.rc, q.rc, mid + 1, r)
        if q.line.idx != 0:
            p = self._update(p, l, r, q.line)
        return p

    def insert(self, l, r, line):
        if l > r:
            return
        l = max(l, self.minX)
        r = min(r, self.maxX)
        if l <= r:
            self.root = self._insert_segment(self.root, self.minX, self.maxX, l, r, line)

    def query(self, x):
        return self._query(self.root, self.minX, self.maxX, x)

    def merge(self, other):
        if other.root is None:
            return
        self.root = self._merge(self.root, other.root, self.minX, self.maxX)
        other.root = None