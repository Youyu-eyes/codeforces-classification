// https://codeforces.com/problemset/problem/2167/G
// 1600 复杂度 O(n^2) 这里采用线段树优化 O(nlogn) 的复杂度，大概 1800
// 线段树优化DP
// 原题在子序列DP中，具体转移方程推到见子序列DP

// 状态转移方程 f[i] = max(f[j] for j in range(i) if a[j] <= a[i]) + c[i]
// 本题中采用离散化 + 线段树优化DP
// 这个转移方程相当于两个维度限制
// j < i; a[j] <= a[i]
// 第一个维度表示遍历顺序先后，在顺序遍历时已经满足
// 对 a 离散化后
// 第二个维度表示 rank[j] <= rank[i]
// 对于 a[j] <= a[i]，f[i] 是分散的，非连续
// 对于 离散化 rank 数组，rank 是连续的
// 对于离散化 rank 数组建立线段树 

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const ll inf = 1e18;

template<typename T>
class SegmentTree {
    // 注：也可以去掉 template<typename T>，改在这里定义 T
    // using T = pair<int, int>;

    int n;
    vector<T> tree;

    // 合并两个 val
    T merge_val(T a, T b) const {
        return max(a, b); // **根据题目修改**
    }

    // 合并左右儿子的 val 到当前节点的 val
    void maintain(int node) {
        tree[node] = merge_val(tree[node * 2], tree[node * 2 + 1]);
    }

    // 用 a 初始化线段树
    // 时间复杂度 O(n)
    void build(const vector<T>& a, int node, int l, int r) {
        if (l == r) { // 叶子
            tree[node] = a[l]; // 初始化叶节点的值
            return;
        }
        int m = (l + r) / 2;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    void update(int node, int l, int r, int i, T val) {
        if (l == r) { // 叶子（到达目标）
            // 如果想直接替换的话，可以写 tree[node] = val
            tree[node] = merge_val(tree[node], val);
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) { // i 在左子树
            update(node * 2, l, m, i, val);
        } else { // i 在右子树
            update(node * 2 + 1, m + 1, r, i, val);
        }
        maintain(node);
    }

    T query(int node, int l, int r, int ql, int qr) const {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node];
        }
        int m = (l + r) / 2;
        if (qr <= m) { // [ql, qr] 在左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 在右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        T l_res = query(node * 2, l, m, ql, qr);
        T r_res = query(node * 2 + 1, m + 1, r, ql, qr);
        return merge_val(l_res, r_res);
    }

public:
    // 线段树维护一个长为 n 的数组（下标从 0 到 n-1），元素初始值为 init_val
    SegmentTree(int n, T init_val) : SegmentTree(vector<T>(n, init_val)) {}

    // 线段树维护数组 a
    SegmentTree(const vector<T>& a) : n(a.size()), tree(2 << bit_width(a.size() - 1)) {
        build(a, 1, 0, n - 1);
    }

    // 更新 a[i] 为 merge_val(a[i], val)
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        update(1, 0, n - 1, i, val);
    }

    // 返回用 merge_val 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
    // 时间复杂度 O(log n)
    T query(int ql, int qr) const {
        return query(1, 0, n - 1, ql, qr);
    }

    // 获取 a[i] 的值
    // 时间复杂度 O(log n)
    T get(int i) const {
        return query(1, 0, n - 1, i, i);
    }
};

vector<int> discretize(vector<int>& arr, int& m) {
    vector<int> sorted_unique = arr;
    ranges::sort(sorted_unique);
    sorted_unique.erase(unique(sorted_unique.begin(), sorted_unique.end()), sorted_unique.end());
    m = sorted_unique.size();
    vector<int> discretized;
    for (int x : arr) {
        discretized.push_back(lower_bound(sorted_unique.begin(), sorted_unique.end(), x) - sorted_unique.begin());
    }
    return discretized;
}

void solve() {
    int n; cin >> n;
    vector<int> a(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }
    vector<int> c(n);
    ll ans = 0;
    for (int i = 0; i < n; ++i) {
        cin >> c[i];
        ans += c[i];
    }

    int m;
    vector<int> rank = discretize(a, m);

    vector<ll> f(n);
    ll sum_max = 0;

    SegmentTree<ll> t(m, 0);

    for (int i = 0; i < n; ++i) {
        f[i] = c[i];
        int r = rank[i];
        ll mx = t.query(0, r);
        f[i] += mx;
        t.update(r, f[i]);

        sum_max = max(sum_max, f[i]);
    }
    cout << ans - sum_max << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    int T;
    cin >> T;
    
    while (T--) {
        solve();
    }
    
    return 0;
}