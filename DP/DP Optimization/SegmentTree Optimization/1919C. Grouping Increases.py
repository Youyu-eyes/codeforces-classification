# https://codeforces.com/contest/1919/problem/C
# Lazy 线段树优化DP
# 1800 另有贪心做法

# 这版过不了，是python的运行效率，c++ 能过，请看下面一版
# 能过的为贪心算法，详见贪心题单

from math import inf
from typing import List
import sys

input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def II():
    return int(input())

def MII(b=0):
    return map(lambda x: int(x) - b, input().split())

def LII():
    return list(MII())

class Node:
    __slots__ = 'val', 'todo'

class LazySegmentTree:
    _TODO_INIT = 0

    def __init__(self, arr, default=0):
        if isinstance(arr, int):
            arr = [default] * arr
        n = len(arr)
        self._n = n
        self._tree = [Node() for _ in range(2 << (n - 1).bit_length())]
        self._build(arr, 1, 0, n - 1)

    def _merge_val(self, a: int, b: int) -> int:
        return min(a, b)

    def _merge_todo(self, a: int, b: int) -> int:
        return a + b

    def _apply(self, node: int, l: int, r: int, todo: int) -> None:
        cur = self._tree[node]
        cur.val += todo
        cur.todo = self._merge_todo(todo, cur.todo)

    def _spread(self, node: int, l: int, r: int) -> None:
        todo = self._tree[node].todo
        if todo == self._TODO_INIT:
            return
        m = (l + r) // 2
        self._apply(node * 2, l, m, todo)
        self._apply(node * 2 + 1, m + 1, r, todo)
        self._tree[node].todo = self._TODO_INIT

    def _maintain(self, node: int) -> None:
        self._tree[node].val = self._merge_val(self._tree[node * 2].val, self._tree[node * 2 + 1].val)

    def _build(self, a: List[int], node: int, l: int, r: int) -> None:
        self._tree[node].todo = self._TODO_INIT
        if l == r:
            self._tree[node].val = a[l]
            return
        m = (l + r) // 2
        self._build(a, node * 2, l, m)
        self._build(a, node * 2 + 1, m + 1, r)
        self._maintain(node)

    def _update(self, node: int, l: int, r: int, ql: int, qr: int, f: int) -> None:
        if ql <= l and r <= qr:
            self._apply(node, l, r, f)
            return
        self._spread(node, l, r)
        m = (l + r) // 2
        if ql <= m:
            self._update(node * 2, l, m, ql, qr, f)
        if qr > m:
            self._update(node * 2 + 1, m + 1, r, ql, qr, f)
        self._maintain(node)

    def _query(self, node: int, l: int, r: int, ql: int, qr: int) -> int:
        if ql <= l and r <= qr:
            return self._tree[node].val
        self._spread(node, l, r)
        m = (l + r) // 2
        if qr <= m:
            return self._query(node * 2, l, m, ql, qr)
        if ql > m:
            return self._query(node * 2 + 1, m + 1, r, ql, qr)
        l_res = self._query(node * 2, l, m, ql, qr)
        r_res = self._query(node * 2 + 1, m + 1, r, ql, qr)
        return self._merge_val(l_res, r_res)

    def update(self, ql: int, qr: int, f: int) -> None:
        self._update(1, 0, self._n - 1, ql, qr, f)

    def query(self, ql: int, qr: int) -> int:
        return self._query(1, 0, self._n - 1, ql, qr)
    
    def get(self, i):
        return self.query(i, i)
    
    def set_val(self, i, new):
        cur = self.get(i)
        self.update(i, i, new - cur)

def discretize(arr):
    sorted_unique = sorted(set(arr))
    value_to_index = {value: idx for idx, value in enumerate(sorted_unique)}
    discretized = [value_to_index[x] for x in arr]
    return discretized, len(sorted_unique)

def solve():
    n = II()
    nums = LII()
    
    if n <= 1:
        print(0)
        return
    
    nums, m = discretize(nums)
    
    INF = 10**9
    seg = LazySegmentTree(m + 1, INF)
    seg.set_val(m, 0)
    
    for i in range(1, n):
        val1 = seg.query(0, nums[i] - 1) + 1 if nums[i] > 0 else INF
        val2 = seg.query(nums[i], m)
        ndp = min(val1, val2)

        if nums[i - 1] < nums[i]:
            seg.update(0, m, 1)

        current_val = seg.get(nums[i - 1])
        if ndp < current_val:
            seg.set_val(nums[i - 1], ndp)
    
    print(seg.query(0, m))

if __name__ == "__main__":
    for _ in range(II()):
        solve()