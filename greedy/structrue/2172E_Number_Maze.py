# https://codeforces.com/contest/2172/problem/E
# 1200

import sys
input = lambda: sys.stdin.readline().strip()

def permute(nums: list[int]) -> list[list[int]]:
    n = len(nums)
    check = [False] * n
    ans = []
    path = []
    def backtrack(i):
        if i == n:
            ans.append(path.copy())
            return
        for j in range(n):
            if not check[j]:
                check[j] = True
                path.append(nums[j])
                backtrack(i+1)
                check[j] = False
                path.pop()
    backtrack(0)
    return ans

p1 = [[1, 2], [2, 1]]
p2 = [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]
p3 = permute([1, 2, 3, 4])

T = int(input())
for _ in range(T):
    num, j, k = map(int, input().split())
    if num == 12:
        nums = p1
    elif num == 123:
        nums = p2
    else:
        nums = p3
    a = b = 0
    l_x, l_y = nums[j - 1], nums[k - 1]
    for x, y in zip(l_x, l_y):
        if x == y:
            a += 1
        else:
            b += 1
    print(f'{a}A{b}B')