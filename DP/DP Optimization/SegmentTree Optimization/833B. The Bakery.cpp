// https://codeforces.com/problemset/problem/833/B
// 2200
// 划分型 DP + Lazy 线段树优化 DP

// 典型的划分型 DP
// 定义 f[k][i] 为将前 i 个数划分为 k 个数的最大价值
// 思考上一次划分的右端点 j
// f[k][i] = max{j = [k - 1, i - 1]} (f[k - 1][j] + different(j, i))
// different(j, i) 表示 [j, i) 区间内不同元素的个数
// 假设第 i 个数上一次出现的下标为 last[i]
// 如果将 nums[i] 加入后，i 对答案的贡献是 j in [last[i] + 1, i] 时，different(j, i) 多一
// 在查询时需要动态维护区间最值，同时有区间更新操作，用 Lazy 线段树

// 代码实现的时候空间优化，同时内部循环优化
// 时间复杂度 O((n - K) * K * logn)，其中 K 表示划分个数，n 为数组长度
// 空间复杂度 O(nlogn)

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

// Lazy 线段树
// 模板来源 https://github.com/Youyu-eyes/codeforces-classification
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
        return max(a, b); // **根据题目修改**
    }

    // 合并两个懒标记
    F merge_todo(const F& a, const F& b) const {
        return a + b; // **根据题目修改**
    }

    // 把懒标记作用到 node 子树（本例为区间加）
    void apply(int node, int l, int r, F todo) {
        Node& cur = tree[node];
        // 计算 tree[node] 区间的整体变化
        cur.val += todo; // **根据题目修改**
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

void solve() {
    int n, K; cin >> n >> K;
    vector<int> nums(n);
    vector<int> last(n, -1);
    unordered_map<int, int> x2idx;
    for (int i = 0; i < n; ++i) {
        cin >> nums[i];
        if (x2idx.contains(nums[i])) {
            last[i] = x2idx[nums[i]];
        }
        x2idx[nums[i]] = i;
    }

    vector<int> f(n + 1);
    for (int k = 1; k <= K; ++k) {
        vector<int> n_f(n + 1);
        LazySegmentTree<int, int> LsegTree(f);
        LsegTree.update(last[k - 1] + 1, k - 1, 1);
        for (int i = k; i <= n + k - K; ++i) {
            n_f[i] = LsegTree.query(k - 1, i - 1);
            if (i < n) LsegTree.update(last[i] + 1, i, 1);
        }
        f = n_f;
    }
    cout << f[n] << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}