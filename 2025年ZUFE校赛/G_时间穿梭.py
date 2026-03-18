# https://cpc.zufe.ac.cn/contest/7/sub/5/problem/G
# 日期问题，用 datetime 库，模拟

from datetime import *

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
    todo_date, todo_time = input().split()
    d0 = datetime.strptime(todo_date, "%Y/%m/%d").date()
    hour, minute, second = map(int, todo_time.split(':'))
    total = hour * 10000 + minute * 100 + second
    d1 = d0 + timedelta(total // 86400)
    total %= 86400
    t1 = time(total // 3600, (total - total // 3600 * 3600) // 60, (total - total // 3600 * 3600) % 60)
    ans = datetime.combine(d1, t1)
    print(ans.strftime("%Y/%m/%d %H:%M:%S"))