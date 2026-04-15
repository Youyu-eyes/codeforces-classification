#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

// 线段树
template<typename T>
class SegmentTree {
    // 注：也可以去掉 template<typename T>，改在这里定义 T
    // using T = pair<int, int>;

    int n;
    vector<T> tree;

    T merge_val(T a, T b) const {
        return max(a, b); // **根据题目修改** 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等
    }

    void maintain(int node) {
        tree[node] = merge_val(tree[node << 1], tree[node << 1 | 1]);
    }

    void build(const vector<T>& a, int node, int l, int r) {
        if (l == r) {
            tree[node] = a[l];
            return;
        }
        int m = (l + r) >> 1;
        build(a, node << 1, l, m);
        build(a, node << 1 | 1, m + 1, r);
        maintain(node);
    }

    void update(int node, int l, int r, int i, T val) {
        if (l == r) {
            tree[node] = merge_val(tree[node], val);
            return;
        }
        int m = (l + r) >> 1;
        if (i <= m) {
            update(node << 1, l, m, i, val);
        } else {
            update(node << 1 | 1, m + 1, r, i, val);
        }
        maintain(node);
    }

    T query(int node, int l, int r, int ql, int qr) const {
        if (ql <= l && r <= qr) {
            return tree[node];
        }
        int m = (l + r) >> 1;
        if (qr <= m) {
            return query(node << 1, l, m, ql, qr);
        }
        if (ql > m) {
            return query(node << 1 | 1, m + 1, r, ql, qr);
        }
        T l_res = query(node << 1, l, m, ql, qr);
        T r_res = query(node << 1 | 1, m + 1, r, ql, qr);
        return merge_val(l_res, r_res);
    }

    int _find_first(int node, int l, int r, int ql, int qr, const std::function<bool(const T&)>& f) const {
        if (r < ql || l > qr) return -1;
        if (!f(tree[node])) return -1;
        if (l == r) return l;
        int m = (l + r) >> 1;
        int left_res = _find_first(node << 1, l, m, ql, qr, f);
        if (left_res != -1) return left_res;
        return _find_first(node << 1 | 1, m + 1, r, ql, qr, f);
    }

    int _find_last(int node, int l, int r, int ql, int qr, const std::function<bool(const T&)>& f) const {
        if (r < ql || l > qr) return -1;
        if (!f(tree[node])) return -1;
        if (l == r) return l;
        int m = (l + r) >> 1;
        int right_res = _find_last(node << 1 | 1, m + 1, r, ql, qr, f);
        if (right_res != -1) return right_res;
        return _find_last(node << 1, l, m, ql, qr, f);
    }

public:
    // 线段树维护一个长为 n 的数组（下标从 0 到 n-1），元素初始值为 init_val
    SegmentTree(int n, T init_val) : SegmentTree(vector<T>(n, init_val)) {}

    // 线段树维护数组 a
    SegmentTree(const vector<T>& a) : n(a.size()), tree(2 << bit_width(a.size() - 1)) {
        build(a, 1, 0, n - 1);
    }

    // 单点更新
    void update(int i, T val) {
        update(1, 0, n - 1, i, val);
    }

    // 区间查询 [ql, qr]
    T query(int ql, int qr) const {
        return query(1, 0, n - 1, ql, qr);
    }

    // 单点查询
    T get(int i) const {
        return query(1, 0, n - 1, i, i);
    }

    // 返回 [ql, qr] 内 第一个  满足 f 的下标，如果不存在，返回 -1
    int find_first(int ql, int qr, const std::function<bool(const T&)>& f) const {
        return _find_first(1, 0, n-1, ql, qr, f);
    }

    // 返回 [ql, qr] 内 最后一个满足 f 的下标，如果不存在，返回 -1
    int find_last(int ql, int qr, const std::function<bool(const T&)>& f) const {
        return _find_last(1, 0, n-1, ql, qr, f);
    }
};


// Lazy 线段树
// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
template<typename T, typename F>
class LazySegmentTree {
    // 注：也可以去掉 template<typename T, typename F>，改在这里定义 T 和 F
    // using T = pair<int, int>;
    // using F = pair<int, int>;

    // 懒标记初始值
    const F TODO_INIT = 0; // **根据题目修改**

    struct Node {
        T val;
        F todo;
    };

    int n;
    vector<Node> tree;

    // 合并两个 val
    T merge_val(const T& a, const T& b) const {
        return a + b; // **根据题目修改**
    }

    // 合并两个懒标记
    F merge_todo(const F& a, const F& b) const {
        return a + b; // **根据题目修改**
    }

    // 把懒标记作用到 node 子树（本例为区间加）
    void apply(int node, int l, int r, F todo) {
        Node& cur = tree[node];
        // 计算 tree[node] 区间的整体变化
        cur.val += todo * (r - l + 1); // **根据题目修改**
        cur.todo = merge_todo(todo, cur.todo);
    }

    // 把当前节点的懒标记下传给左右儿子
    void spread(int node, int l, int r) {
        Node& cur = tree[node];
        F todo = cur.todo;
        if (todo == TODO_INIT) { // 没有需要下传的信息
            return;
        }
        int m = (l + r) >> 1;
        apply(node << 1, l, m, todo);
        apply(node << 1 | 1, m + 1, r, todo);
        cur.todo = TODO_INIT; // 下传完毕
    }

    // 合并左右儿子的 val 到当前节点的 val
    void maintain(int node) {
        tree[node].val = merge_val(tree[node << 1].val, tree[node << 1 | 1].val);
    }

    // 用 a 初始化线段树
    // 时间复杂度 O(n)
    void build(const vector<T>& a, int node, int l, int r) {
        Node& cur = tree[node];
        cur.todo = TODO_INIT;
        if (l == r) { // 叶子
            cur.val = a[l]; // 初始化叶节点的值
            return;
        }
        int m = (l + r) >> 1;
        build(a, node << 1, l, m); // 初始化左子树
        build(a, node << 1 | 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    void update(int node, int l, int r, int ql, int qr, F f) {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            apply(node, l, r, f);
            return;
        }
        spread(node, l, r);
        int m = (l + r) >> 1;
        if (ql <= m) { // 更新左子树
            update(node << 1, l, m, ql, qr, f);
        }
        if (qr > m) { // 更新右子树
            update(node << 1 | 1, m + 1, r, ql, qr, f);
        }
        maintain(node);
    }

    T query(int node, int l, int r, int ql, int qr) {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node].val;
        }
        spread(node, l, r);
        int m = (l + r) >> 1;
        if (qr <= m) { // [ql, qr] 在左子树
            return query(node << 1, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 在右子树
            return query(node << 1 | 1, m + 1, r, ql, qr);
        }
        T l_res = query(node << 1, l, m, ql, qr);
        T r_res = query(node << 1 | 1, m + 1, r, ql, qr);
        return merge_val(l_res, r_res);
    }

    // 私有递归：查找第一个满足条件的下标
    int _find_first(int node, int l, int r, int ql, int qr, const std::function<bool(const T&)>& f) {
        if (r < ql || l > qr) return -1;
        if (!f(tree[node].val)) return -1;
        if (l == r) return l;
        spread(node, l, r); // 下传懒标记
        int m = (l + r) >> 1;
        int left_res = _find_first(node << 1, l, m, ql, qr, f);
        if (left_res != -1) return left_res;
        return _find_first(node << 1 | 1, m + 1, r, ql, qr, f);
    }

    // 私有递归：查找最后一个满足条件的下标
    int _find_last(int node, int l, int r, int ql, int qr, const std::function<bool(const T&)>& f) {
        if (r < ql || l > qr) return -1;
        if (!f(tree[node].val)) return -1;
        if (l == r) return l;
        spread(node, l, r);
        int m = (l + r) >> 1;
        int right_res = _find_last(node << 1 | 1, m + 1, r, ql, qr, f);
        if (right_res != -1) return right_res;
        return _find_last(node << 1, l, m, ql, qr, f);
    }

public:
    // 线段树维护一个长为 n 的数组（下标从 0 到 n-1），元素初始值为 init_val
    LazySegmentTree(int n, T init_val = 0) : LazySegmentTree(vector<T>(n, init_val)) {}

    // 线段树维护数组 a
    LazySegmentTree(const vector<T>& a) : n(a.size()), tree(2 << bit_width(a.size() - 1)) {
        build(a, 1, 0, n - 1);
    }

    // 用 f 更新 [ql, qr] 中的每个 a[i]
    // 0 <= ql <= qr <= n-1
    // 时间复杂度 O(log n)
    void update(int ql, int qr, F f) {
        update(1, 0, n - 1, ql, qr, f);
    }

    // 返回用 merge_val 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
    // 0 <= ql <= qr <= n-1
    // 时间复杂度 O(log n)
    T query(int ql, int qr) {
        return query(1, 0, n - 1, ql, qr);
    }
    // 返回 [ql, qr] 内第一个满足 f 的下标，如果不存在，返回 -1
    // 例如查询区间内第一个大于等于 target 的元素下标，需要线段树维护区间最大值
    //     int idx = seg.find_first(l, r, [&](const T& node_max) { return node_max >= target; });
    int find_first(int ql, int qr, const std::function<bool(const T&)>& f) {
        return _find_first(1, 0, n-1, ql, qr, f);
    }

    // 返回 [ql, qr] 内最后一个满足 f 的下标，如果不存在，返回 -1
    // 例如查询区间内最后一个小于等于 target 的元素下标，需要线段树维护区间最小值
    //     int idx = seg.find_last(l, r, [&](const T& node_min) { return node_min <= target; });
    int find_last(int ql, int qr, const std::function<bool(const T&)>& f) {
        return _find_last(1, 0, n-1, ql, qr, f);
    }
};

int main() {
    LazySegmentTree<long long, long long> t(8); // 默认值为 0
    t.update(3, 5, 100);
    t.update(4, 6, 10);
    cout << t.query(0, 7) << endl;

    vector<long long> nums = {3, 1, 4, 1, 5, 9, 2, 6};
    LazySegmentTree<long long, long long> t2(nums);
    t2.update(3, 5, 1);
    t2.update(4, 6, 1);
    cout << t2.query(0, 7) << endl;
    return 0;
}


// 李超线段树
const double eps = 1e-10;

class LiChaoSegmentTree {
private:
    struct Line {
        double k, b;
        int id;
        Line() : k(0), b(ll_inf), id(0) {}
        // 最大值：Line() : k(0), b(-ll_inf), id(0) {}
        Line(double k, double b, int id) : k(k), b(b), id(id) {}
        double calc(int x) const { return k * x + b; }
    };

    struct Node {
        Line line;
        Node() : line(0, ll_inf, 0) {}
        // 最大值：Node() : line(0, -ll_inf, 0) {}
    };

    int n;
    vector<Node> tree;

    Line better(const Line& a, const Line& b, int x) const {
        if (a.id == 0) return b;
        if (b.id == 0) return a;
        double va = a.calc(x), vb = b.calc(x);
        if (va - vb < -eps) return a;
        // 最大值：if (va - vb > eps) return a;
        if (vb - va < -eps) return b;
        // 最大值：if (vb - va > eps) return b;
        return a.id < b.id ? a : b;
    }

    void update(int node, int l, int r, Line new_line) {
        if (l > r) return;
        Line old_line = tree[node].line;
        int mid = (l + r) >> 1;
        bool better_at_mid = new_line.calc(mid) < old_line.calc(mid) - eps;
        // 最大值：bool better_at_mid = new_line.calc(mid) > old_line.calc(mid) + eps;
        if (better_at_mid) swap(tree[node].line, new_line);
        if (l == r) return;
        bool better_at_left = new_line.calc(l) < old_line.calc(l) - eps;
        // 最大值：bool better_at_left = new_line.calc(l) > old_line.calc(l) + eps;
        bool better_at_right = new_line.calc(r) < old_line.calc(r) - eps;
        // 最大值：bool better_at_right = new_line.calc(r) > old_line.calc(r) + eps;
        if (better_at_left != better_at_mid)
            update(node << 1, l, mid, new_line);
        else if (better_at_right != better_at_mid)
            update(node << 1 | 1, mid + 1, r, new_line);
    }

    void insert_line(int node, int l, int r, int ql, int qr, const Line& line) {
        if (l > qr || r < ql) return;
        if (ql <= l && r <= qr) {
            update(node, l, r, line);
            return;
        }
        int mid = (l + r) >> 1;
        insert_line(node << 1, l, mid, ql, qr, line);
        insert_line(node << 1 | 1, mid + 1, r, ql, qr, line);
    }

    Line query(int node, int l, int r, int x) {
        if (l == r) return tree[node].line;
        int mid = (l + r) >> 1;
        Line res = tree[node].line;
        Line child_res = (x <= mid) ? query(node << 1, l, mid, x)
                                    : query(node << 1 | 1, mid + 1, r, x);
        return better(res, child_res, x);
    }

public:
    LiChaoSegmentTree(int range) : n(range) {
        tree.resize(4 * (n + 1));
    }

    void insert(int l, int r, const Line& line) {
        insert_line(1, 0, n, l, r, line);
    }

    double query(int x) {
        Line line = query(1, 0, n, x);
        if (line.id == 0) return ll_inf;
        // 最大值：if (line.id == 0) return -ll_inf;
        return line.calc(x);
    }
};


// 动态开点李超线段树
class DynamicLiChaoTree {
private:
    struct Line {
        ll k, b;
        int id;
        Line() : k(0), b(ll_inf), id(0) {}
        Line(ll k, ll b, int id = 0) : k(k), b(b), id(id) {}
        ll calc(int x) const { return k * x + b; }
    };

    struct Node {
        int lc, rc;      // 左右孩子下标（-1 表示空）
        Line line;       // 该节点存储的直线
        Node() : lc(-1), rc(-1), line() {}
    };

    int minX, maxX;
    vector<Node> nodes;   // 节点池

    // 判断 a 在 x 处是否严格优于 b（求最小值）
    bool better(const Line& a, const Line& b, int x) const {
        if (a.id == 0) return false;
        if (b.id == 0) return true;
        ll va = a.calc(x), vb = b.calc(x);
        if (va != vb) return va < vb;
        return a.id < b.id;
    }

    // 新建节点，返回下标
    int new_node() {
        nodes.emplace_back();
        return nodes.size() - 1;
    }

    // 在节点 p 对应区间 [l, r] 插入直线 new_line
    void insert(int p, int l, int r, Line new_line) {
        if (p == -1) return;
        Line &old = nodes[p].line;
        int mid = l + (r - l) / 2;

        bool better_mid = better(new_line, old, mid);
        if (better_mid) swap(old, new_line);
        if (l == r) return;

        bool better_left = better(new_line, old, l);
        bool better_right = better(new_line, old, r);

        if (better_left != better_mid) {
            if (nodes[p].lc == -1) nodes[p].lc = new_node();
            insert(nodes[p].lc, l, mid, new_line);
        } else if (better_right != better_mid) {
            if (nodes[p].rc == -1) nodes[p].rc = new_node();
            insert(nodes[p].rc, mid + 1, r, new_line);
        }
    }

    // 查询 x 处的最小值
    ll query(int p, int l, int r, int x) const {
        if (p == -1) return ll_inf;
        ll res = (nodes[p].line.id == 0) ? ll_inf : nodes[p].line.calc(x);
        if (l == r) return res;
        int mid = l + (r - l) / 2;
        if (x <= mid) res = min(res, query(nodes[p].lc, l, mid, x));
        else          res = min(res, query(nodes[p].rc, mid + 1, r, x));
        return res;
    }

    // 合并两棵李超树（将 q 合并到 p），返回合并后的根下标
    int merge(int p, int q, int l, int r) {
        if (p == -1) return q;
        if (q == -1) return p;
        if (l == r) {
            if (better(nodes[q].line, nodes[p].line, l))
                nodes[p].line = nodes[q].line;
            return p;
        }
        int mid = l + (r - l) / 2;
        nodes[p].lc = merge(nodes[p].lc, nodes[q].lc, l, mid);
        nodes[p].rc = merge(nodes[p].rc, nodes[q].rc, mid + 1, r);
        if (nodes[q].line.id != 0)
            insert(p, l, r, nodes[q].line);
        return p;
    }

public:
    int root;   // 根节点下标，初始为 -1（空树）

    DynamicLiChaoTree(int minX, int maxX) : minX(minX), maxX(maxX), root(-1) {}

    // 插入一条直线
    void insert(ll k, ll b, int id = 0) {
        if (root == -1) root = new_node();
        insert(root, minX, maxX, Line(k, b, id));
    }

    // 查询 x 处的最小值
    ll query(int x) const {
        return query(root, minX, maxX, x);
    }

    // 将另一棵树合并到当前树（other 中的直线会加入当前树）
    void merge(DynamicLiChaoTree& other) {
        if (other.root == -1) return;
        if (root == -1) root = new_node();
        root = merge(root, other.root, minX, maxX);
    }
};