class Node:
    __slots__ = 'val', 'todo'

class LasySegmentTree:
    _TODO_INIT = 0

    def __init__(self, arr, default = 0):
        if isinstance(arr, int):
            n = arr
            arr = [default] * n
        n = len(arr)
        self._n = n
        self._tree = [Node() for _ in range(2 << (n - 1).bit_length())]
        self._build(arr, 1, 0, n - 1)
    
    # 值 合并函数
    def _merge_val(self, x, y):
        return max(x, y)    # 根据题目改变
    
    # 懒标记 合并函数
    def _merge_todo(self, x, y):
        return x + y    # 根据题目改变

    # 懒标记应用于当前节点
    def _apply(self, node, todo):
        cur = self._tree[node]
        cur.val += todo
        cur.todo = self._merge_todo(todo, cur.todo)

    # 懒标记传到下面两个节点
    def _spread(self, node):
        todo = self._tree[node].todo
        if todo == self._TODO_INIT:
            return
        self._apply(node * 2, todo)
        self._apply(node * 2 + 1, todo)
        self._tree[node].todo = self._TODO_INIT

    # 节点合并
    def _maintain(self, node):
        self._tree[node].val = self._merge_val(self._tree[node * 2].val, self._tree[node * 2 + 1].val)
    
    # 建树
    def _build(self, a, node, l, r):
        if l == r:
            self._tree[node].val = a[l]
            return
        m = (l + r) // 2
        self._build(a, node * 2, l, m)
        self._build(a, node * 2 + 1, m + 1, r)
        self._maintain(node)
    
    # 更新
    def _update(self, node, l, r, ql, qr, f):
        if ql <= l and r <= qr:
            self._apply(node, f)
            return
        self._spread(node)
        m = (l + r) // 2
        if ql <= m:
            self._update(node * 2, l, m, ql, qr)
        if qr > m:
            self._update(node * 2 + 1, m + 1, r, ql, qr)
        self._maintain(node)
    
    # 查询
    def _query(self, node, l, r, ql, qr):
        if ql <= l and r <= qr:
            return self._tree[node].val
        self._spread(node)
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
    def update(self, ql, qr, f):
        self._update(1, 0, self._n - 1, ql, qr, f)
    
    def query(self, ql, qr):
        return self._query(1, 0, self._n - 1, ql, qr)
    
    def get(self, i):
        return self._query(1, 0, self._n - 1, i, i)