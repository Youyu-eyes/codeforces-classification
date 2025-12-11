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