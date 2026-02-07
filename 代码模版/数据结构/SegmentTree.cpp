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


// 李超线段树 - max_LichaoSegmentTree
const double eps = 1e-10;
const long long INF = 1e18;

struct Line {
    double k, b;
    int id;
    
    Line() : k(0), b(-INF), id(0) {}
    Line(double k, double b, int id) : k(k), b(b), id(id) {}
    
    double calc(int x) const {
        return k * x + b;
    }
};

class LiChaoSegmentTree {
private:
    struct Node {
        Line line;
        Node() : line(0, -INF, 0) {}
    };
    
    int n;
    vector<Node> tree;

    Line better(const Line& a, const Line& b, int x) const {
        if (a.id == 0) return b;
        if (b.id == 0) return a;
        
        double va = a.calc(x);
        double vb = b.calc(x);
        
        if (va - vb > eps) return a;
        if (vb - va > eps) return b;
        return a.id < b.id ? a : b;
    }
    
    void update(int node, int l, int r, Line new_line) {
        if (l > r) return;
        
        Line old_line = tree[node].line;
        
        int mid = (l + r) / 2;

        bool better_at_mid = new_line.calc(mid) > old_line.calc(mid) + eps;
        
        if (better_at_mid) {
            swap(tree[node].line, new_line);
        }
        
        if (l == r) return;
        
        bool better_at_left = new_line.calc(l) > old_line.calc(l) + eps;
        bool better_at_right = new_line.calc(r) > old_line.calc(r) + eps;
        
        if (better_at_left != better_at_mid) {
            update(node * 2, l, mid, new_line);
        } else if (better_at_right != better_at_mid) {
            update(node * 2 + 1, mid + 1, r, new_line);
        }
    }
    
    void insert_line(int node, int l, int r, int ql, int qr, const Line& line) {
        if (l > qr || r < ql) return;
        
        if (ql <= l && r <= qr) {
            update(node, l, r, line);
            return;
        }
        
        int mid = (l + r) / 2;
        insert_line(node * 2, l, mid, ql, qr, line);
        insert_line(node * 2 + 1, mid + 1, r, ql, qr, line);
    }
    
    Line query(int node, int l, int r, int x) {
        if (l == r) {
            return tree[node].line;
        }
        
        int mid = (l + r) / 2;
        Line res = tree[node].line;
        
        Line child_res;
        if (x <= mid) {
            child_res = query(node * 2, l, mid, x);
        } else {
            child_res = query(node * 2 + 1, mid + 1, r, x);
        }
        
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
        if (line.id == 0) return -INF;
        return line.calc(x);
    }
};


// 李超线段树 - min_LichaoSegmentTree
const double eps = 1e-10;
const long long INF = 1e18;

struct Line {
    double k, b;
    int id;
    
    Line() : k(0), b(INF), id(0) {}
    Line(double k, double b, int id) : k(k), b(b), id(id) {}
    
    double calc(int x) const {
        return k * x + b;
    }
};

class LiChaoSegmentTree {
private:
    struct Node {
        Line line;
        Node() : line(0, INF, 0) {}
    };
    
    int n;
    vector<Node> tree;

    Line better(const Line& a, const Line& b, int x) const {
        if (a.id == 0) return b;
        if (b.id == 0) return a;
        
        double va = a.calc(x);
        double vb = b.calc(x);
        
        if (va - vb < -eps) return a;
        if (vb - va < -eps) return b;
        return a.id < b.id ? a : b;
    }
    
    void update(int node, int l, int r, Line new_line) {
        if (l > r) return;
        
        Line old_line = tree[node].line;

        int mid = (l + r) / 2;
  
        bool better_at_mid = new_line.calc(mid) < old_line.calc(mid) - eps;
        
        if (better_at_mid) {
            swap(tree[node].line, new_line);
        }

        if (l == r) return;
        
        bool better_at_left = new_line.calc(l) < old_line.calc(l) - eps;
        bool better_at_right = new_line.calc(r) < old_line.calc(r) - eps;
        
        if (better_at_left != better_at_mid) {
            update(node * 2, l, mid, new_line);
        } else if (better_at_right != better_at_mid) {
            update(node * 2 + 1, mid + 1, r, new_line);
        }
    }
    
    void insert_line(int node, int l, int r, int ql, int qr, const Line& line) {
        if (l > qr || r < ql) return;
        
        if (ql <= l && r <= qr) {
            update(node, l, r, line);
            return;
        }
        
        int mid = (l + r) / 2;
        insert_line(node * 2, l, mid, ql, qr, line);
        insert_line(node * 2 + 1, mid + 1, r, ql, qr, line);
    }
    
    Line query(int node, int l, int r, int x) {
        if (l == r) {
            return tree[node].line;
        }
        
        int mid = (l + r) / 2;
        Line res = tree[node].line;
        
        Line child_res;
        if (x <= mid) {
            child_res = query(node * 2, l, mid, x);
        } else {
            child_res = query(node * 2 + 1, mid + 1, r, x);
        }
        
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
        if (line.id == 0) return INF;
        return line.calc(x);
    }
};