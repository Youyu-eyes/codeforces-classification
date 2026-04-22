// https://github.com/Youyu-eyes/codeforces-classification
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
            tree[node] = merge_val(tree[node], val);  // 如果想直接覆盖就改成 tree[node] = val;
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

    // 单点更新 将 tree[node] 改成 merge_val(tree[node], val)，需要直接覆盖要修改私有函数
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


// 动态开点线段树
#include <bits/stdc++.h>
using namespace std;

// 线段树节点（定义在类外）
struct SegNode {
    SegNode* left;
    SegNode* right;
    int l, r;
    long long val;

    SegNode(int l_, int r_, long long val_)
        : left(nullptr), right(nullptr), l(l_), r(r_), val(val_) {}
};

class DynamicSegmentTree {
private:
    // 空节点哨兵（自环）
    static SegNode* emptyNode() {
        static SegNode* empty = [](){
            SegNode* e = new SegNode(0, 0, 0);
            e->left = e->right = e;
            return e;
        }();
        return empty;
    }

    SegNode* root;
    int minX_, maxX_;
    long long default_val;  // 默认值（例如求最大值且含负数时传 LLONG_MIN）

    long long merge_val(long long a, long long b) const {
        return max(a, b);  // 可改为 min, +, gcd 等
    }

    void maintain(SegNode* o) {
        o->val = merge_val(o->left->val, o->right->val);
    }

    void update(SegNode* o, int i, long long val) {
        if (o->l == o->r) {
            o->val = merge_val(o->val, val);
            return;
        }
        int m = (o->l + o->r) >> 1;
        if (i <= m) {
            if (o->left == emptyNode()) {
                o->left = new SegNode(o->l, m, default_val);
                o->left->left = o->left->right = emptyNode();
            }
            update(o->left, i, val);
        } else {
            if (o->right == emptyNode()) {
                o->right = new SegNode(m + 1, o->r, default_val);
                o->right->left = o->right->right = emptyNode();
            }
            update(o->right, i, val);
        }
        maintain(o);
    }

    long long query(SegNode* o, int ql, int qr) const {
        if (o == emptyNode() || ql > o->r || qr < o->l)
            return default_val;
        if (ql <= o->l && o->r <= qr)
            return o->val;
        return merge_val(query(o->left, ql, qr), query(o->right, ql, qr));
    }

    int find_first(SegNode* o, int ql, int qr, const function<bool(long long)>& f) const {
        if (o == emptyNode() || o->r < ql || o->l > qr || !f(o->val))
            return -1;
        if (o->l == o->r)
            return o->l;
        int left_res = find_first(o->left, ql, qr, f);
        if (left_res != -1) return left_res;
        return find_first(o->right, ql, qr, f);
    }

    int find_last(SegNode* o, int ql, int qr, const function<bool(long long)>& f) const {
        if (o == emptyNode() || o->r < ql || o->l > qr || !f(o->val))
            return -1;
        if (o->l == o->r)
            return o->l;
        int right_res = find_last(o->right, ql, qr, f);
        if (right_res != -1) return right_res;
        return find_last(o->left, ql, qr, f);
    }

    SegNode* merge_nodes(SegNode* a, SegNode* b) {
        if (a == emptyNode()) return b;
        if (b == emptyNode()) return a;
        if (a->l == a->r) {
            a->val = merge_val(a->val, b->val);
            return a;
        }
        a->left = merge_nodes(a->left, b->left);
        a->right = merge_nodes(a->right, b->right);
        maintain(a);
        return a;
    }

public:
    DynamicSegmentTree(int minX, int maxX, long long def)
        : minX_(minX), maxX_(maxX), default_val(def)
    {
        emptyNode()->val = def;
        root = new SegNode(minX, maxX, def);
        root->left = root->right = emptyNode();
    }

    // 禁用拷贝（避免浅拷贝问题）
    DynamicSegmentTree(const DynamicSegmentTree&) = delete;
    DynamicSegmentTree& operator=(const DynamicSegmentTree&) = delete;

    // 单点更新
    void update(int i, long long val) {
        if (i < minX_ || i > maxX_) return;
        update(root, i, val);
    }

    // 区间查询
    long long query(int ql, int qr) const {
        if (ql > qr || ql > maxX_ || qr < minX_) return default_val;
        ql = max(ql, minX_);
        qr = min(qr, maxX_);
        return query(root, ql, qr);
    }

    // 单点查询
    long long get(int i) const {
        return query(i, i);
    }

    // 在 [ql, qr] 中查找第一个满足 f 的位置，不存在返回 -1
    int find_first(int ql, int qr, const function<bool(long long)>& f) const {
        if (ql > qr || ql > maxX_ || qr < minX_) return -1;
        ql = max(ql, minX_);
        qr = min(qr, maxX_);
        return find_first(root, ql, qr, f);
    }

    // 在 [ql, qr] 中查找最后一个满足 f 的位置，不存在返回 -1
    int find_last(int ql, int qr, const function<bool(long long)>& f) const {
        if (ql > qr || ql > maxX_ || qr < minX_) return -1;
        ql = max(ql, minX_);
        qr = min(qr, maxX_);
        return find_last(root, ql, qr, f);
    }

    // 线段树合并
    void merge(DynamicSegmentTree& other) {
        if (minX_ != other.minX_ || maxX_ != other.maxX_) {
            throw invalid_argument("值域范围不同，无法合并");
        }
        root = merge_nodes(root, other.root);
    }
};


// 动态开点 Lazy 线段树
struct LazySegNode {
    LazySegNode* left;
    LazySegNode* right;
    int l, r;
    long long val;
    long long todo;

    LazySegNode(int l_, int r_, long long val_, long long todo_)
        : left(nullptr), right(nullptr), l(l_), r(r_), val(val_), todo(todo_) {}
};

class DynamicLazySegmentTree {
private:
    static LazySegNode* emptyNode() {
        static LazySegNode* empty = [](){
            LazySegNode* e = new LazySegNode(0, 0, 0, 0);
            e->left = e->right = e;
            return e;
        }();
        return empty;
    }

    LazySegNode* root;
    int minX_, maxX_;          // 值域范围 [minX, maxX]
    long long val_default;     // 值默认值（如区间和时为0）
    long long todo_init;       // 懒标记初始值（如无标记时为0）

    long long merge_val(long long a, long long b) const {
        return a + b;  // 区间求和；可改为 max, min 等
    }

    long long merge_todo(long long a, long long b) const {
        return a + b;  // 懒标记累加；可改为 max, min 等
    }

    void apply(LazySegNode* o, long long f) {
        o->val += f * (o->r - o->l + 1);  // ** 根据题目修改 ** //
        o->todo = merge_todo(o->todo, f);
    }

    void maintain(LazySegNode* o) {
        o->val = merge_val(o->left->val, o->right->val);
    }

    void spread(LazySegNode* o) {
        if (o->todo == todo_init) return;
        int m = (o->l + o->r) >> 1;
        // 动态创建左子节点
        if (o->left == emptyNode()) {
            o->left = new LazySegNode(o->l, m, val_default, todo_init);
            o->left->left = o->left->right = emptyNode();
        }
        // 动态创建右子节点
        if (o->right == emptyNode()) {
            o->right = new LazySegNode(m + 1, o->r, val_default, todo_init);
            o->right->left = o->right->right = emptyNode();
        }
        long long f = o->todo;
        apply(o->left, f);
        apply(o->right, f);
        o->todo = todo_init;
    }

    void update(LazySegNode* o, int ql, int qr, long long f) {
        if (ql <= o->l && o->r <= qr) {
            apply(o, f);
            return;
        }
        spread(o);
        int m = (o->l + o->r) >> 1;
        if (ql <= m) update(o->left, ql, qr, f);
        if (qr > m)  update(o->right, ql, qr, f);
        maintain(o);
    }

    long long query(LazySegNode* o, int ql, int qr) {
        if (o == emptyNode() || ql > o->r || qr < o->l)
            return val_default;
        if (ql <= o->l && o->r <= qr)
            return o->val;
        spread(o);
        return merge_val(query(o->left, ql, qr), query(o->right, ql, qr));
    }

    int find_first(LazySegNode* o, int ql, int qr, const function<bool(long long)>& f) {
        if (o == emptyNode() || o->r < ql || o->l > qr || !f(o->val))
            return -1;
        if (o->l == o->r)
            return o->l;
        spread(o);
        int left_res = find_first(o->left, ql, qr, f);
        if (left_res != -1) return left_res;
        return find_first(o->right, ql, qr, f);
    }

    int find_last(LazySegNode* o, int ql, int qr, const function<bool(long long)>& f) {
        if (o == emptyNode() || o->r < ql || o->l > qr || !f(o->val))
            return -1;
        if (o->l == o->r)
            return o->l;
        spread(o);
        int right_res = find_last(o->right, ql, qr, f);
        if (right_res != -1) return right_res;
        return find_last(o->left, ql, qr, f);
    }

    LazySegNode* merge_nodes(LazySegNode* a, LazySegNode* b) {
        if (a == emptyNode()) return b;
        if (b == emptyNode()) return a;
        spread(a);
        spread(b);
        if (a->l == a->r) {
            a->val = merge_val(a->val, b->val);
            return a;
        }
        a->left = merge_nodes(a->left, b->left);
        a->right = merge_nodes(a->right, b->right);
        maintain(a);
        return a;
    }

public:
    DynamicLazySegmentTree(int minX, int maxX, long long val_def, long long todo_def)
        : minX_(minX), maxX_(maxX), val_default(val_def), todo_init(todo_def)
    {
        emptyNode()->val = val_def;
        emptyNode()->todo = todo_def;
        root = new LazySegNode(minX, maxX, val_def, todo_def);
        root->left = root->right = emptyNode();
    }

    // 禁用拷贝
    DynamicLazySegmentTree(const DynamicLazySegmentTree&) = delete;
    DynamicLazySegmentTree& operator=(const DynamicLazySegmentTree&) = delete;

    // 区间更新
    void update(int ql, int qr, long long f) {
        if (ql > qr || ql > maxX_ || qr < minX_) return;
        ql = max(ql, minX_);
        qr = min(qr, maxX_);
        update(root, ql, qr, f);
    }

    // 区间查询
    long long query(int ql, int qr) {
        if (ql > qr || ql > maxX_ || qr < minX_) return val_default;
        ql = max(ql, minX_);
        qr = min(qr, maxX_);
        return query(root, ql, qr);
    }

    // 单点查询
    long long get(int i) {
        return query(i, i);
    }

    // 二分查找第一个满足 f 的位置
    int find_first(int ql, int qr, const function<bool(long long)>& f) {
        if (ql > qr || ql > maxX_ || qr < minX_) return -1;
        ql = max(ql, minX_);
        qr = min(qr, maxX_);
        return find_first(root, ql, qr, f);
    }

    // 二分查找最后一个满足 f 的位置
    int find_last(int ql, int qr, const function<bool(long long)>& f) {
        if (ql > qr || ql > maxX_ || qr < minX_) return -1;
        ql = max(ql, minX_);
        qr = min(qr, maxX_);
        return find_last(root, ql, qr, f);
    }

    // 线段树合并
    void merge(DynamicLazySegmentTree& other) {
        if (minX_ != other.minX_ || maxX_ != other.maxX_)
            throw invalid_argument("值域范围不同，无法合并");
        root = merge_nodes(root, other.root);
    }
};


// 李超线段树
const double eps = 1e-10;

struct Line {
    double k, b;
    int id;
    Line() : k(0), b(ll_inf), id(0) {}
    // 最大值：Line() : k(0), b(-ll_inf), id(0) {}
    Line(double k, double b, int id) : k(k), b(b), id(id) {}
    double calc(int x) const { return k * x + b; }
};

class LiChaoSegmentTree {
private:
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


// 动态开点李超线段树（值域李超线段树）
const double eps = 1e-10;

struct Line {
    double k, b;
    int id;
    Line() : k(0), b(ll_inf), id(0) {}
    // 最大值：Line() : k(0), b(-ll_inf), id(0) {}
    Line(double k, double b, int id) : k(k), b(b), id(id) {}
    double calc(int x) const { return k * x + b; }
};

class DynamicLiChaoTree {
private:
    struct Node {
        Node* lc;
        Node* rc;
        Line line;
        Node() : lc(nullptr), rc(nullptr), line() {}
    };

    int minX, maxX;
    Node* root;

    bool better(const Line& a, const Line& b, int x) const {
        if (a.id == 0) return false;
        if (b.id == 0) return true;
        double va = a.calc(x), vb = b.calc(x);
        if (va - vb < -eps) return true;
        // 最大值：if (va - vb > eps) return true;
        if (vb - va < -eps) return false;
        // 最大值：if (vb - va > eps) return false;
        return a.id < b.id;
    }

    void update(Node* &p, int l, int r, Line new_line) {
        if (!p) {
            p = new Node();
            p->line = new_line;
            return;
        }
        int mid = (l + r) >> 1;
        bool left_better  = better(new_line, p->line, l);
        bool mid_better   = better(new_line, p->line, mid);

        if (mid_better) {
            swap(p->line, new_line);
            mid_better  = better(new_line, p->line, mid);
            left_better = better(new_line, p->line, l);
        }
        if (l == r) return;
        if (left_better != mid_better) {
            update(p->lc, l, mid, new_line);
        } else {
            update(p->rc, mid + 1, r, new_line);
        }
    }

    void insert_segment(Node* &p, int l, int r, int ql, int qr, const Line& line) {
        if (!p) p = new Node();
        if (ql <= l && r <= qr) {
            update(p, l, r, line);
            return;
        }
        int mid = (l + r) >> 1;
        if (ql <= mid) insert_segment(p->lc, l, mid, ql, qr, line);
        if (qr > mid)  insert_segment(p->rc, mid + 1, r, ql, qr, line);
    }

    double query(Node* p, int l, int r, int x) const {
        if (!p) return ll_inf;
        // 最大值：if (!p) return -ll_inf;
        double res = (p->line.id == 0) ? ll_inf : p->line.calc(x);
        // 最大值：double res = (p->line.id == 0) ? -ll_inf : p->line.calc(x);
        if (l == r) return res;
        int mid = (l + r) >> 1;
        if (x <= mid) return min(res, query(p->lc, l, mid, x));
        else          return min(res, query(p->rc, mid + 1, r, x));
        // 最大值：return max(res, ...)
    }

    Node* merge(Node* p, Node* q, int l, int r) {
        if (!p) return q;
        if (!q) return p;
        if (l == r) {
            if (better(q->line, p->line, l))
                p->line = q->line;
            return p;
        }
        int mid = (l + r) >> 1;
        p->lc = merge(p->lc, q->lc, l, mid);
        p->rc = merge(p->rc, q->rc, mid + 1, r);
        if (q->line.id != 0)
            update(p, l, r, q->line);
        return p;
    }

    void clear(Node* p) {
        if (!p) return;
        clear(p->lc);
        clear(p->rc);
        delete p;
    }

public:
    DynamicLiChaoTree(int minX, int maxX) : minX(minX), maxX(maxX), root(nullptr) {}

    ~DynamicLiChaoTree() { clear(root); }

    void insert(int l, int r, const Line& line) {
        if (l > r) return;
        l = max(l, minX);
        r = min(r, maxX);
        if (l <= r) insert_segment(root, minX, maxX, l, r, line);
    }

    double query(int x) const {
        return query(root, minX, maxX, x);
    }

    void merge(DynamicLiChaoTree& other) {
        if (!other.root) return;
        root = merge(root, other.root, minX, maxX);
        other.root = nullptr;
    }
};

// todo
// https://codeforces.com/problemset/problem/1175/G
// https://codeforces.com/contest/1303/problem/G