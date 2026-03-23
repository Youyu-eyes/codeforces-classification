// https://codeforces.com/contest/1919/problem/C
// Lazy 线段树优化DP
// 1800（估计的，感觉低了） 另有贪心做法

// c++ 版本，依旧是灵神的板子，能过
// 另有贪心算法，详见贪心题单


#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int INF = 1e9;

// 线段树节点类型
struct NodeVal {
    int val;
};

// 懒标记类型
struct LazyTag {
    int add;
    LazyTag(int a = 0) : add(a) {}
};

class LazySegmentTree {
private:
    vector<NodeVal> tree;
    vector<LazyTag> lazy;
    int n;
    
    // 合并两个节点值
    NodeVal merge(const NodeVal& a, const NodeVal& b) {
        return {min(a.val, b.val)};
    }
    
    // 应用懒标记
    void apply(int node, int l, int r, const LazyTag& tag) {
        tree[node].val += tag.add;
        lazy[node].add += tag.add;
    }
    
    // 下传懒标记
    void push(int node, int l, int r) {
        if (lazy[node].add == 0) return;
        int mid = (l + r) >> 1;
        apply(node << 1, l, mid, lazy[node]);
        apply(node << 1 | 1, mid + 1, r, lazy[node]);
        lazy[node].add = 0;
    }
    
    void build(int node, int l, int r, int init_val) {
        if (l == r) {
            tree[node] = {init_val};
            return;
        }
        int mid = (l + r) >> 1;
        build(node << 1, l, mid, init_val);
        build(node << 1 | 1, mid + 1, r, init_val);
        tree[node] = merge(tree[node << 1], tree[node << 1 | 1]);
    }
    
    void update_range(int node, int l, int r, int ql, int qr, int val) {
        if (ql <= l && r <= qr) {
            tree[node].val += val;
            lazy[node].add += val;
            return;
        }
        push(node, l, r);
        int mid = (l + r) >> 1;
        if (ql <= mid) update_range(node << 1, l, mid, ql, qr, val);
        if (qr > mid) update_range(node << 1 | 1, mid + 1, r, ql, qr, val);
        tree[node] = merge(tree[node << 1], tree[node << 1 | 1]);
    }
    
    void update_point(int node, int l, int r, int idx, int val) {
        if (l == r) {
            tree[node].val = val;
            return;
        }
        push(node, l, r);
        int mid = (l + r) >> 1;
        if (idx <= mid) update_point(node << 1, l, mid, idx, val);
        else update_point(node << 1 | 1, mid + 1, r, idx, val);
        tree[node] = merge(tree[node << 1], tree[node << 1 | 1]);
    }
    
    NodeVal query(int node, int l, int r, int ql, int qr) {
        if (ql <= l && r <= qr) return tree[node];
        push(node, l, r);
        int mid = (l + r) >> 1;
        if (qr <= mid) return query(node << 1, l, mid, ql, qr);
        if (ql > mid) return query(node << 1 | 1, mid + 1, r, ql, qr);
        return merge(query(node << 1, l, mid, ql, qr),
                     query(node << 1 | 1, mid + 1, r, ql, qr));
    }
    
public:
    LazySegmentTree(int size, int init_val) : n(size) {
        tree.resize(4 * n);
        lazy.resize(4 * n);
        build(1, 0, n - 1, init_val);
    }
    
    void update(int l, int r, int val) {
        update_range(1, 0, n - 1, l, r, val);
    }
    
    void set(int idx, int val) {
        update_point(1, 0, n - 1, idx, val);
    }
    
    int query(int l, int r) {
        return query(1, 0, n - 1, l, r).val;
    }
    
    int get(int idx) {
        return query(idx, idx);
    }
};

vector<int> discretize(vector<int>& arr, int& m) {
    vector<int> sorted_unique = arr;
    sort(sorted_unique.begin(), sorted_unique.end());
    sorted_unique.erase(unique(sorted_unique.begin(), sorted_unique.end()), sorted_unique.end());
    m = sorted_unique.size();
    vector<int> discretized;
    for (int x : arr) {
        discretized.push_back(lower_bound(sorted_unique.begin(), sorted_unique.end(), x) - sorted_unique.begin());
    }
    return discretized;
}

void solve() {
    int n;
    cin >> n;
    vector<int> nums(n);
    for (int i = 0; i < n; i++) {
        cin >> nums[i];
    }
    
    if (n <= 1) {
        cout << 0 << '\n';
        return;
    }
    
    int m;
    auto discretized_nums = discretize(nums, m);
    
    LazySegmentTree seg(m + 1, INF);
    seg.set(m, 0);
    
    for (int i = 1; i < n; i++) {
        int val1 = (discretized_nums[i] > 0) ? seg.query(0, discretized_nums[i] - 1) + 1 : INF;
        int val2 = seg.query(discretized_nums[i], m);
        int ndp = min(val1, val2);
        
        if (discretized_nums[i - 1] < discretized_nums[i]) {
            seg.update(0, m, 1);
        }
        
        int current_val = seg.get(discretized_nums[i - 1]);
        if (ndp < current_val) {
            seg.set(discretized_nums[i - 1], ndp);
        }
    }
    
    cout << seg.query(0, m) << '\n';
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