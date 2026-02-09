# 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
class UnionFind:
    def __init__(self, n: int):
        # 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        # 集合 i 的代表元是自己，自己到自己的距离是 0
        self.fa = list(range(n))  # 代表元
        self.dis = [0] * n  # dis[x] 表示 x 到（x 所在集合的）代表元的距离

    # 返回 x 所在集合的代表元
    # 同时做路径压缩
    def find(self, x: int) -> int:
        fa = self.fa
        if fa[x] != x:
            root = self.find(fa[x])
            self.dis[x] += self.dis[fa[x]]  # 递归更新 x 到其代表元的距离
            fa[x] = root
        return fa[x]

    # 判断 x 和 y 是否在同一个集合（同普通并查集）
    def same(self, x: int, y: int) -> bool:
        return self.find(x) == self.find(y)

    # 计算从 from 到 to 的相对距离
    # 调用时需保证 from 和 to 在同一个集合中，否则返回值无意义
    def get_relative_distance(self, from_: int, to: int) -> int:
        self.find(from_)
        self.find(to)
        # to-from = (x-from) - (x-to) = dis[from] - dis[to]
        return self.dis[from_] - self.dis[to]

    # 合并 from 和 to，新增信息 to - from = value
    # 其中 to 和 from 表示未知量，下文的 x 和 y 也表示未知量
    # 如果 from 和 to 不在同一个集合，返回 True，否则返回是否与已知信息矛盾
    def merge(self, from_: int, to: int, value: int) -> bool:
        x, y = self.find(from_), self.find(to)
        dis = self.dis
        if x == y:  # from 和 to 在同一个集合，不做合并
            # to-from = (x-from) - (x-to) = dis[from] - dis[to] = value
            return dis[from_] - dis[to] == value
        #    x --------- y
        #   /           /
        # from ------- to
        # 已知 x-from = dis[from] 和 y-to = dis[to]，现在合并 from 和 to，新增信息 to-from = value
        # 由于 y-from = (y-x) + (x-from) = (y-to) + (to-from)
        # 所以 y-x = (to-from) + (y-to) - (x-from) = value + dis[to] - dis[from]
        dis[x] = value + dis[to] - dis[from_]  # 计算 x 到其代表元 y 的距离
        self.fa[x] = y
        return True