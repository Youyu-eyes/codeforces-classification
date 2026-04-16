class FenwickTree:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)  # 使用下标 1 到 n

    # a[i] 增加 val
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def update(self, i: int, val: int) -> None:
        t = self.tree
        while i < len(t):
            t[i] += val
            i += i & -i

    # 计算前缀和 a[1] + ... + a[i]
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def pre(self, i: int) -> int:
        t = self.tree
        res = 0
        while i > 0:
            res += t[i]
            i &= i - 1
        return res

    # 计算区间和 a[l] + ... + a[r]
    # 1 <= l <= r <= n
    # 时间复杂度 O(log n)
    def query(self, l: int, r: int) -> int:
        if r < l:
            return 0
        return self.pre(r) - self.pre(l - 1)
    
    # 返回满足前缀和 >= k 的最小下标（1-based）
    # 假设 1 <= k <= 总元素个数，且序列元素非负
    def find_kth(self, k: int) -> int:
        pos = 0
        t = self.tree
        n = len(t) - 1
        bit_mask = 1 << (n.bit_length() - 1)
        s = 0
        
        while bit_mask:
            nxt = pos + bit_mask
            if nxt <= n and s + t[nxt] < k:
                pos = nxt
                s += t[nxt]
            bit_mask >>= 1
        return pos + 1