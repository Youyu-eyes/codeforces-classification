// https://acm.timus.ru/problem.aspx?space=1&num=1521
// 经典题，约瑟夫环

// 线段树，有线段树查找方法，这里采用外部二分，多一个 logn 的时间
// 另有 nlogn 树状数组二分做法

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
        return a + b; // **根据题目修改** 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等
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
            tree[node] = val;
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

void solve() {
    int n, k; cin >> n >> k;
    SegmentTree<ll> segtree(n, 1);
    int cur = 0;
    for (int order = 0; order < n; ++order) {
        cur = (cur + k - 1) % (n - order);
        
        int l = -1, r = n;
        while (l + 1 < r) {
            int mid = (l + r) >> 1;
            if (segtree.query(0, mid) > cur) {
                r = mid;
            } else {
                l = mid;
            }
        }
        segtree.update(r, 0);
        cout << r + 1 << ' ';
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}