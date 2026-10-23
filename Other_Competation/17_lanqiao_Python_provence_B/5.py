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
    ans = 0
    nums = MII()
    for x in nums:
        if x == 1:
            ans += 1
        else:
            ans += 2
    if ans & 1:
        print('L')
    else:
        print('Q')