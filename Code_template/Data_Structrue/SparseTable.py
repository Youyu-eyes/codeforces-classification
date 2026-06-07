from typing import List, Callable

# 区间最值 ST 表
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

# 区间最值 ST 表，下标版本
class SparseTableIndex:
    def __init__(self, nums: List[int], op: Callable[[int, int], int]):
        self.nums = nums
        self.op = op
        n = len(nums)
        k = n.bit_length()
        self.st = [[0] * n for _ in range(k)]
        for i in range(n):
            self.st[0][i] = i
        for i in range(1, k):
            step = 1 << (i - 1)
            for j in range(n - (1 << i) + 1):
                left = self.st[i-1][j]
                right = self.st[i-1][j+step]
                if op(nums[left], nums[right]) == nums[left]:
                    self.st[i][j] = left
                else:
                    self.st[i][j] = right

    def query(self, l: int, r: int) -> int:
        length = r - l
        k = length.bit_length() - 1
        left = self.st[k][l]
        right = self.st[k][r - (1 << k)]
        if self.op(self.nums[left], self.nums[right]) == self.nums[left]:
            return left
        else:
            return right

# 用例
if __name__ == "__main__":
    arr = [3, 1, 4, 1, 5, 9, 2, 6]
    
    # 最大值下标ST表
    st_max = SparseTableIndex(arr, max)
    print("最大值下标查询：")
    print("区间[1,5) 最大值下标:", st_max.query(1, 5))  # 期望 4 (值5)
    print("区间[0,3) 最大值下标:", st_max.query(0, 3))  # 期望 2 (值4)
    print("区间[2,7) 最大值下标:", st_max.query(2, 7))  # 期望 5 (值9)
    
    # 最小值下标ST表
    st_min = SparseTableIndex(arr, min)
    print("\n最小值下标查询：")
    print("区间[1,5) 最小值下标:", st_min.query(1, 5))  # 期望 1 (值1，最左)
    print("区间[0,3) 最小值下标:", st_min.query(0, 3))  # 期望 1 (值1)
    print("区间[2,7) 最小值下标:", st_min.query(2, 7))  # 期望 3 (值1)


# ------- 二维 ST 表 ------- #

class SparseTable2D:
    def __init__(self, matrix: List[List[int]], op: Callable[[int, int], int]):
        if not matrix or not matrix[0]:
            self.st = []
            return
            
        self.op = op
        m, n = len(matrix), len(matrix[0])
        log_m = m.bit_length()
        log_n = n.bit_length()
 
        st = [[[[0] * n for _ in range(m)] for _ in range(log_n)] for _ in range(log_m)]

        for i in range(m):
            for j in range(n):
                st[0][0][i][j] = matrix[i][j]

        for l in range(1, log_n):
            length = 1 << (l - 1)
            for i in range(m):
                for j in range(n - (1 << l) + 1):
                    st[0][l][i][j] = self.op(st[0][l - 1][i][j], st[0][l - 1][i][j + length])

        for k in range(1, log_m):
            length_r = 1 << (k - 1)
            for l in range(log_n):
                for i in range(m - (1 << k) + 1):
                    for j in range(n - (1 << l) + 1):
                        st[k][l][i][j] = self.op(st[k - 1][l][i][j], st[k - 1][l][i + length_r][j])
                        
        self.st = st

    def query(self, r1: int, r2: int, c1: int, c2: int) -> int:
        k = (r2 - r1).bit_length() - 1
        l = (c2 - c1).bit_length() - 1
        
        len_r = 1 << k
        len_c = 1 << l

        op1 = self.st[k][l][r1][c1]
        op2 = self.st[k][l][r2 - len_r][c1]
        op3 = self.st[k][l][r1][c2 - len_c]
        op4 = self.st[k][l][r2 - len_r][c2 - len_c]
        
        return self.op(self.op(op1, op2), self.op(op3, op4))