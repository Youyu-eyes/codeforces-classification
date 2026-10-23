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

a = 0
ans = 0
for b in range(1, 2027):
    l_b = len(str(b))
    a = (a * 10 ** l_b % 26 + b % 26) % 26
    if a == 0:
        ans += 1
print(ans)