// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
#include <bits/stdc++.h>
using namespace std;

class UnionFind {
    vector<int> fa; // 代表元
    vector<int> sz; // 集合大小

public:
    int cc; // 连通块个数

    UnionFind(int n) : fa(n), sz(n, 1), cc(n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己，大小为 1
        ranges::iota(fa, 0); // iota(fa.begin(), fa.end(), 0);
    }

    // 返回 x 所在集合的代表元
    // 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
    int find(int x) {
        // 如果 fa[x] == x，则表示 x 是代表元
        if (fa[x] != x) {
            fa[x] = find(fa[x]); // fa 改成代表元
        }
        return fa[x];
    }

    // 判断 x 和 y 是否在同一个集合
    bool is_same(int x, int y) {
        // 如果 x 的代表元和 y 的代表元相同，那么 x 和 y 就在同一个集合
        // 这就是代表元的作用：用来快速判断两个元素是否在同一个集合
        return find(x) == find(y);
    }

    // 把 from 所在集合合并到 to 所在集合中
    // 返回是否合并成功
    bool merge(int from, int to) {
        int x = find(from), y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            return false;
        }
        fa[x] = y; // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
        sz[y] += sz[x]; // 更新集合大小（注意集合大小保存在代表元上）
        // 无需更新 sz[x]，因为我们不用 sz[x] 而是用 sz[find(x)] 获取集合大小，但 find(x) == y，我们不会再访问 sz[x]
        cc--; // 成功合并，连通块个数减一
        return true;
    }

    // 返回 x 所在集合的大小
    int get_size(int x) {
        return sz[find(x)]; // 集合大小保存在代表元上
    }
};

// 初始化
// UnionFind uf(n);

// 带权并查集
// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
// 根据题目用 UnionFindWithWight<int> uf(n) 或者 UnionFindWithWight<long long> uf(n) 初始化
template<typename T>
class UnionFindWithWight {
public:
    vector<int> fa; // 代表元
    vector<T> dis; // dis[x] 表示 x 到（x 所在集合的）代表元的距离

    UnionFindWithWight(int n) : fa(n), dis(n) {
        // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
        // 集合 i 的代表元是自己，自己到自己的距离是 0
        ranges::iota(fa, 0); // iota(fa.begin(), fa.end(), 0);
    }

    // 返回 x 所在集合的代表元
    // 同时做路径压缩
    int find(int x) {
        if (fa[x] != x) {
            int root = find(fa[x]);
            dis[x] += dis[fa[x]]; // 递归更新 x 到其代表元的距离
            fa[x] = root;
        }
        return fa[x];
    }

    // 判断 x 和 y 是否在同一个集合（同普通并查集）
    bool same(int x, int y) {
        return find(x) == find(y);
    }

    // 计算从 from 到 to 的相对距离
    // 调用时需保证 from 和 to 在同一个集合中，否则返回值无意义
    T get_relative_distance(int from, int to) {
        find(from);
        find(to);
        // to-from = (x-from) - (x-to) = dis[from] - dis[to]
        return dis[from] - dis[to];
    }

    // 合并 from 和 to，新增信息 to - from = value
    // 其中 to 和 from 表示未知量，下文的 x 和 y 也表示未知量
    // 如果 from 和 to 不在同一个集合，返回 true，否则返回是否与已知信息矛盾
    bool merge(int from, int to, T value) {
        int x = find(from), y = find(to);
        if (x == y) { // from 和 to 在同一个集合，不做合并
            // to-from = (x-from) - (x-to) = dis[from] - dis[to] = value
            return dis[from] - dis[to] == value;
        }
        //    x --------- y
        //   /           /
        // from ------- to
        // 已知 x-from = dis[from] 和 y-to = dis[to]，现在合并 from 和 to，新增信息 to-from = value
        // 由于 y-from = (y-x) + (x-from) = (y-to) + (to-from)
        // 所以 y-x = (to-from) + (y-to) - (x-from) = value + dis[to] - dis[from]
        dis[x] = value + dis[to] - dis[from]; // 计算 x 到其代表元 y 的距离
        fa[x] = y;
        return true;
    }
};