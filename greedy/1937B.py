# https://codeforces.com/problemset/problem/1937/B
# 1300
# 贪心
# 右边和下面的数字相同时，走右边能让自己的选择更多，一定更优

# 本题值得关注的是，在 Python 中，处理列表比字符串的拼接操作开销更小
# 即使最后要额外把列表改成字符串输出，也不要拼接字符串，开销不是 O(1) 的

import sys
input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII(b = 0):
    return list(MII(b))

for _ in range(II()):
    n = II()
    a = input()
    b = input()

    ans = []
    num = 1
    
    for i in range(n - 1):
        ans.append(a[i])
        x = a[i + 1]
        y = b[i]
        
        if x > y:
            ans.extend(b[i:])
            break
        elif x < y:
            num = 1
        else:
            num += 1
    else:
        ans.append(a[-1])
        ans.append(b[-1])
    
    print(''.join(ans))
    print(num)