#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const ll INF = 1e18;

// 线段树
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

// 李超线段树
class LiChaoTree {
private:
    ll x_left, x_right;
    LiChaoTree* left = nullptr;
    LiChaoTree* right = nullptr;
    pair<ll, ll> line; // (斜率, 截距)
    
    ll f(pair<ll, ll>& line, ll x) {
        return line.first * x + line.second;
    }
    
public:
    LiChaoTree(ll l, ll r) : x_left(l), x_right(r) {
        line = {0, INF};
    }
    
    void add_line(pair<ll, ll> new_line, ll l, ll r) {
        ll mid = (l + r) / 2;
        
        bool left_new_better = f(new_line, l) < f(line, l);
        bool mid_new_better = f(new_line, mid) < f(line, mid);
        
        if (mid_new_better) {
            swap(line, new_line);
        }
        
        if (r - l == 1) {
            return;
        }
        
        if (left_new_better != mid_new_better) {
            if (!left) {
                left = new LiChaoTree(l, mid);
            }
            left->add_line(new_line, l, mid);
        } else {
            if (!right) {
                right = new LiChaoTree(mid, r);
            }
            right->add_line(new_line, mid, r);
        }
    }
    
    void add_line_range(pair<ll, ll> new_line, ll seg_l, ll seg_r, ll l, ll r) {
        if (seg_r <= l || r <= seg_l) {
            return;
        }
        if (seg_l <= l && r <= seg_r) {
            add_line(new_line, l, r);
            return;
        }
        
        ll mid = (l + r) / 2;
        if (!left) {
            left = new LiChaoTree(l, mid);
        }
        left->add_line_range(new_line, seg_l, seg_r, l, mid);
        
        if (!right) {
            right = new LiChaoTree(mid, r);
        }
        right->add_line_range(new_line, seg_l, seg_r, mid, r);
    }
    
    ll query(ll x, ll l, ll r) {
        ll res = f(line, x);
        if (r - l == 1) {
            return res;
        }
        
        ll mid = (l + r) / 2;
        if (x < mid) {
            if (left) {
                res = min(res, left->query(x, l, mid));
            }
        } else {
            if (right) {
                res = min(res, right->query(x, mid, r));
            }
        }
        return res;
    }
    
    ~LiChaoTree() {
        delete left;
        delete right;
    }
};