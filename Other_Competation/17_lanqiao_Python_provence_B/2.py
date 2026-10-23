MOD = 998244353

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

ans = 0
A = 20269876543210
B = 20260123456789
for target in range(0, 7*10**6 + 1):
    target *= target
    if target > (A + B):
        break
    ans += (min(target, A) % MOD - target % MOD + min(target, B) % MOD + 1) % MOD
    ans %= MOD
print(ans)