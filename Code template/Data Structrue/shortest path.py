from heapq import heappop, heappush
from collections import deque
from math import inf
from typing import List, Tuple

# 0-1 BFS
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


# dijkstra：单源最短路

def shortestPathDijkstra(n: int, edges: List[List[int]], start: int) -> List[int]:
    # 注：如果节点编号从 1 开始（而不是从 0 开始），可以把 n 加一
    g = [[] for _ in range(n)]  # 邻接表
    for x, y, wt in edges:
        g[x].append((y, wt))
        # g[y].append((x, wt))  # 无向图加上这行

    dis = [inf] * n
    dis[start] = 0  # 起点到自己的距离是 0
    h = [(0, start)]  # 堆中保存 (起点到节点 x 的最短路长度，节点 x)

    while h:
        dis_x, x = heappop(h)
        if dis_x > dis[x]:  # x 之前出堆过
            continue
        for y, wt in g[x]:
            new_dis_y = dis_x + wt
            if new_dis_y < dis[y]:
                dis[y] = new_dis_y  # 更新 x 的邻居的最短路
                # 懒更新堆：只插入数据，不更新堆中数据
                # 相同节点可能有多个不同的 new_dis_y，除了最小的 new_dis_y，其余值都会触发上面的 continue
                heappush(h, (new_dis_y, y))

    return dis

# 网格图 Dijkstra
def shortestPathDijkstra(m: int, n: int, grid: List[List[int]], start: tuple[int, int]) -> List[int]:
    start_x, start_y = start
    dis = [[inf] * n for _ in range(m)]
    dis[start_x][start_y] = 0
    h = [(0, (start_x, start_y))]

    while h:
        dis_xy, (x, y) = heappop(h)
        if dis_xy > dis[x][y]:  # x 之前出堆过
            continue
        for (i, j) in (x - 1, y), (x + 1, y), (x, y - 1), (x, y + 1):
            if i < 0 or i >= m or j < 0 or j >= n:
                continue
            new_dis_ij = dis_xy + grid[i][j]
            if new_dis_ij < dis[i][j]:
                dis[i][j] = new_dis_ij  # 更新 x 的邻居的最短路
                # 懒更新堆：只插入数据，不更新堆中数据
                # 相同节点可能有多个不同的 new_dis_y，除了最小的 new_dis_y，其余值都会触发上面的 continue
                heappush(h, (new_dis_ij, (i, j)))

    return dis


# Floyd：全源最短路

# 返回一个二维列表，其中 (i,j) 这一项表示从 i 到 j 的最短路长度
# 如果无法从 i 到 j，则最短路长度为 math.inf
# 允许负数边权
# 如果计算完毕后，存在 i，使得从 i 到 i 的最短路长度小于 0，说明图中有负环
# 节点编号从 0 到 n-1
# 时间复杂度 O(n^3 + m)，其中 m 是 edges 的长度
def shortestPathFloyd(self, n: int, edges: List[List[int]]) -> List[List[int]]:
    f = [[inf] * n for _ in range(n)]
    for i in range(n):
        f[i][i] = 0

    for x, y, wt in edges:
        f[x][y] = min(f[x][y], wt)  # 如果有重边，取边权最小值
        f[y][x] = min(f[y][x], wt)  # 无向图

    for k in range(n):
        for i in range(n):
            if f[i][k] == inf:  # 针对稀疏图的优化
                continue
            for j in range(n):
                f[i][j] = min(f[i][j], f[i][k] + f[k][j])
    return f