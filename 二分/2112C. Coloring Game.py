# https://codeforces.com/contest/2112/problem/C
# 1300
# 二分方法 python3 过不了，交 PyPy3
# 复杂度 O(n^2logn)

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