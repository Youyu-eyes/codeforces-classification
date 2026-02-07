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
const double eps = 1e-10;
struct Line {
    double k, b;
    int id;
    
    Line() : k(0), b(0), id(0) {}
    Line(double k, double b, int id) : k(k), b(b), id(id) {}
    
    // 计算在x处的y值
    double calc(int x) const {
        return k * x + b;
    }
    
    // 通过两点创建直线
    void set_by_points(int x0, int y0, int x1, int y1, int id_val) {
        this->id = id_val;
        if (x0 == x1) {
            k = 0;
            b = max(y1, y0);
            return;
        }
        k = 1.0 * (y1 - y0) / (x1 - x0);
        b = 1.0 * y0 - k * x0;
    }
};

class LiChaoSegmentTree {
private:
    struct Node {
        int line_id;  // 当前节点保存的线段id，0表示无线段
        Node() : line_id(0) {}
    };
    
    int n;  // x坐标的范围 [0, n]
    vector<Node> tree;
    vector<Line> lines;  // 存储所有线段，lines[0]为空线段
    
    // 比较两个结果，优先返回y值大的，y值相同时返回id小的
    pair<double, int> better(const pair<double, int>& a, const pair<double, int>& b) const {
        if (a.first - b.first > eps) return a;
        if (b.first - a.first > eps) return b;
        return a.second < b.second ? a : b;
    }
    
    // 递归更新线段
    void update(int node, int l, int r, int ql, int qr, int line_id) {
        if (r < ql || l > qr) return;
        
        int mid = (l + r) >> 1;
        
        if (ql <= l && r <= qr) {
            if (!tree[node].line_id) {
                tree[node].line_id = line_id;
                return;
            }
            
            int cur_id = tree[node].line_id;
            Line& cur = lines[cur_id];
            Line& new_line = lines[line_id];
            
            // 比较中点处的值
            bool better_at_mid = new_line.calc(mid) - cur.calc(mid) > eps;
            
            if (better_at_mid) {
                swap(tree[node].line_id, line_id);
                swap(cur_id, tree[node].line_id);
            }
            
            if (l == r) return;
            
            bool better_at_left = new_line.calc(l) - cur.calc(l) > eps ||
                                  (fabs(new_line.calc(l) - cur.calc(l)) <= eps && line_id < cur_id);
            bool better_at_right = new_line.calc(r) - cur.calc(r) > eps ||
                                   (fabs(new_line.calc(r) - cur.calc(r)) <= eps && line_id < cur_id);
            
            if (better_at_left) {
                update(node << 1, l, mid, ql, qr, line_id);
            }
            if (better_at_right) {
                update(node << 1 | 1, mid + 1, r, ql, qr, line_id);
            }
            return;
        }
        
        update(node << 1, l, mid, ql, qr, line_id);
        update(node << 1 | 1, mid + 1, r, ql, qr, line_id);
    }
    
    // 递归查询
    pair<double, int> query(int node, int l, int r, int x) const {
        if (l > x || r < x) return {0, 0};
        
        pair<double, int> res = {0, 0};
        if (tree[node].line_id) {
            const Line& line = lines[tree[node].line_id];
            res = {line.calc(x), tree[node].line_id};
        }
        
        if (l == r) return res;
        
        int mid = (l + r) >> 1;
        if (x <= mid) {
            return better(res, query(node << 1, l, mid, x));
        } else {
            return better(res, query(node << 1 | 1, mid + 1, r, x));
        }
    }
    
public:
    // 构造函数，n为x坐标的最大值（范围是[0, n]）
    LiChaoSegmentTree(int range) : n(range) {
        // 分配线段树空间，通常4倍足够
        tree.resize(4 * (n + 1));
        lines.emplace_back();  // 第0条是空线段
    }
    
    // 插入一条线段，从x1到x2
    void insert(int x1, int y1, int x2, int y2) {
        int line_id = lines.size();
        lines.emplace_back();
        lines.back().set_by_points(x1, y1, x2, y2, line_id);
        update(1, 0, n, min(x1, x2), max(x1, x2), line_id);
    }
    
    // 插入一条线段，指定线段对象和区间
    void insert_segment(int l, int r, const Line& line) {
        int line_id = lines.size();
        lines.push_back(line);
        lines.back().id = line_id;
        update(1, 0, n, l, r, line_id);
    }
    
    // 在x处查询
    pair<double, int> query(int x) const {
        return query(1, 0, n, x);
    }
    
    // 获取线段数量（包括第0条空线段）
    int get_line_count() const {
        return lines.size();
    }
    
    // 获取指定id的线段
    const Line& get_line(int id) const {
        return lines[id];
    }
};