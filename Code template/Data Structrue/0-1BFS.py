from collections import deque
from typing import List, Tuple
from math import inf
# 计算从 start 到各个节点的最短路长度
# 如果节点无法到达，最短路为 inf
# 节点编号从 0 到 n-1，边权为 0 or 1

def zero_one_bfs(n: int, g: List[List[Tuple[int, int]]], start: int) -> List[int]:
    dis = [inf] * n
    dis[start] = 0
    q = deque([start])

    while q:
        x = q.popleft()
        for y, w in g[x]:
            if dis[x] + w < dis[y]:
                dis[y] = dis[x] + w
                if w == 0:
                    q.appendleft(y)
                else:
                    q.append(y)
    return dis