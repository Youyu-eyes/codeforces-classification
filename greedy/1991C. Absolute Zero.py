# https://codeforces.com/problemset/problem/1991/C
# 1300
# 跟官解不一样的写法，最多需要 min(len(set(nums)), log(mx)) 次操作，其中 mx = max(nums)
# 最大最小值贪心，有点类似中位数贪心

# 首先证明 nums 中所有数同奇偶
# 最后一次操作的前一次操作，nums[i] = x
# 倒数第二次操作一定是把所有 nums[i] 全部变成同一个数，即 abs(nums[i] - x) = abs(nums[j] - x)
# 分类讨论，令 a = nums[i]，b = nums[j]，若 a，b 同时大于或小于 x，则易知 a = b
# 反之则有 a - x = x - b，即 (a + b) / 2 = x，x 为正整数，则 a + b 为偶数，则 a，b 同奇偶
# 推广到 nums 数组中，任意 i, j 位置上的值同奇偶，即 nums 数组全数同奇偶

# 接下来证明取两个极值的中点为 x 一定能使 nums 全变成同一个数
# 假设 最大值 mx，最小值 mn，mid = (mx + mn) // 2
# abs(a - x) 的几何意义是 a 离 x 的距离，而当 x = mid 时，max(abs(a - mid)) = mx - mid
# 下一个 mx = mx - mid，因此上界变小，而下界最小固定为 0，因此不断减 mid 后一定会减到 0
# 据此我们可以得到一个操作次数 log(mx)

# 最后证明操作次数
# 前面我们证明了一个操作次数可能 log(mx)，但实际上这个数远大于答案，因为没有考虑 mn 的更改
# 当 mx = mx - (mx + mn) // 2，mn = (mx + mn) // 2 - mn，可得 mx = mn = (mx - mn) // 2（几何意义更快）
# 所以每次操作后，数组中都会多一个相等的元素
# 而我们的目标是所有数相等，因此一共需要 len(set(nums)) 次操作，能使所有数相等

# 代码就是每次取极值中点，更新 nums，并维护 mn 即可
# 值得注意的是 mx 不需要再维护 nums 中额外维护，原因见第二点证明
# 时间复杂度 O(len(set(nums)) * n)

# 后记说一下为什么不用中位数贪心
# 中位数贪心的话，我们需要一个有序的数组，才能用 O(1) 时间求得中位数
# 每次再维护完 nums 后，我们需要额外排序，复杂度 O(nlogn)
# 不过感兴趣的同学可以实现一下
# 另外提一嘴，这个题只要 x 取在 [max(mn, 1), mx] 之间的数都可以实现，也是能证明中位数贪心的正确性
# 感兴趣的同学也可以尝试证明一下

from math import inf

import sys
input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII():
    return list(MII())

for _ in range(II()):
    n = II()
    nums = LII()
    mx = 0
    mn = inf
    t = nums[0] % 2
    for x in nums:
        if x % 2 != t:
            print(-1)
            break
        mx = max(mx, x)
        mn = min(mn, x)
    else:
        ans = []
        while mn < mx:
            mid = (mx + mn) // 2
            mn = mx
            mx = mx - mid
            ans.append(str(mid))
            for i, x in enumerate(nums):
                nums[i] = abs(x - mid)
                mx = max(mx, nums[i])
                mn = min(mn, nums[i])
        if mn == mx == 0:
            print(0)
            print()
        else:
            ans.append(str(mx))
            print(len(ans))
            print(' '.join(ans))