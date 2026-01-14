# https://codeforces.com/contest/2112/problem/C
# 双指针做法
# 时间复杂度 O(n^2)，python 能过
# 倒序枚举 j 能让 k 递增，递减也可，但是不如递增好写
# 还有一个瓶颈在预处理 max_mx_2i_list上，python 的 max 开销也很大

from bisect import bisect_left

import sys
input = lambda: sys.stdin.readline().split()
si = lambda: input()[0]
ii = lambda: int(si())
lii = lambda: list(map(int, input()))
lsi = lambda: input()

def solve():
    T = ii()
    out = []
    for _ in range(T):
        n = ii()
        nums = lii()
        mx = nums[-1]
        ans = 0
        max_mx_2i_list = [max(mx, 2 * nums[i]) for i in range(n)]
        for i in range(2, n):
            for j in range(1, i):
                k = bisect_left(nums, max_mx_2i_list[i] - nums[i] - nums[j] + 1, 0, j)
                ans += j - k

        out.append(str(ans))
    sys.stdout.buffer.write('\n'.join(out).encode())

if __name__ == "__main__":
    solve()