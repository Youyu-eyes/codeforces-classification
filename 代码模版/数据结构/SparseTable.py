from typing import List, Callable

class SparseTable:
    # 时间复杂度 O(n * log n)
    def __init__(self, nums: List[int], op: Callable[[int, int], int]):
        n = len(nums)
        w = n.bit_length()
        st = [[0] * n for _ in range(w)]
        st[0] = nums[:]
        for i in range(1, w):
            for j in range(n - (1 << i) + 1):
                st[i][j] = op(st[i - 1][j], st[i - 1][j + (1 << (i - 1))])
        self.st = st
        self.op = op

    # [l, r) 左闭右开，下标从 0 开始
    # 返回 op(nums[l:r])
    # 时间复杂度 O(1)
    def query(self, l: int, r: int) -> int:
        k = (r - l).bit_length() - 1
        return self.op(self.st[k][l], self.st[k][r - (1 << k)])


# 使用方法举例
nums = [3, 1, 4, 1, 5, 9, 2, 6]
st = SparseTable(nums, max)
print(st.query(0, 5))  # max(nums[0:5]) = 5
print(st.query(1, 1))  # 错误：必须保证 l < r