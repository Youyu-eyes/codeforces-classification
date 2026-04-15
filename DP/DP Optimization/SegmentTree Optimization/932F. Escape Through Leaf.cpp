// https://codeforces.com/problemset/problem/932/F
// 2700
// 树形 DP + 动态开点李超线段树（值域李超线段树）
// 官解是 树上启发式合并 + 动态凸包

//      0
//    /   \ 
//   1     2
//  / \   /
// 3   4 5

// 首先我们声明，为了便于建图，节点序号从 0 开始，root = 0

// 我们要计算 0 节点到 3 节点的最小开销，我们可以先跳到 1 再跳到 3，也可以直接跳到 3
// 如果先跳到 1，问题就变成了从 1 节点跳到 3 节点的最小开销，这是一个规模更小的子问题，可以用递归解决

// 现在我们将问题抽象
// 不失一般性地讨论从 i 跳到叶子节点的最小开销
// 假设从 i 跳到 j，问题就变成了从 j 到叶子节点的最小开销
// 定义从 i 跳到叶子节点的最小开销为 f[i]
// 则从 i 跳到 j 的条件是 j in subtree(i)
// 状态转移方程为 f[i] = min{j in subtree(i)} (f[j] + a[i] * b[j])
// 叶子节点没有子树，因此 f[i] = 0
// 显然，我们枚举 j 需要 O(n) 的时间，整体复杂度 O(n^2) 我们需要优化

// 如果我们使用凸包优化
// 对于 1 和 2 节点，显然我们需要两个凸包维护，再合并，这样的思路是官解的树上启发式合并 + 凸包
// 有没有更板子的写法？

// 对于凸包优化 DP 如果因为查询区间问题无法解决的，采用李超线段树优化
// 将 f[j] + a[i] * b[j] 看作直线 y = k * x + b, k = b[j], b = f[j], 对于需要查询的 x = a[i]，需要查询 x = a[i] 直线与所有直线的最低交点
// 这是一个标准的李超线段树解决的问题
// 对于根节点 0，要计算所有的 min{j = [1, 5]} (f[j] + a[0] * b[j])
// 假设我们已经维护了 1, 3, 4 节点形成直线的李超线段树 和 2, 5 节点形成直线的李超线段树
// 我们需要的是维护 [1, 5] 节点的李超线段树，因此我们需要合并这两棵树
// 然后就是正常查询线段树的最值，然后将根节点加入这棵树，最后返回

// 由于本题我们需要 n 棵李超线段树，根据 x = a[i] 的值域大小 2 * 10 ^ 5 来说显然会爆炸，因此我们需要动态开点李超线段树
// 需要动态开点的另一个因素是我们需要合并两棵线段树，静态合并的开销太大

// 时间复杂度 O(n*logU*logU) 其中 U 是值域
// 空间复杂度 O(n*logU)

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

const int MN = -100'000;
const int MX = 100'000;

// 动态开点李超线段树
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

void solve() {
    int n; cin >> n;
    vector<ll> a(n);
    vector<ll> b(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }
    for (int i = 0; i < n; ++i) {
        cin >> b[i];
    }
    vector<vector<int>> g(n);
    for (int i = 0; i < n - 1; ++i) {
        int u, v; cin >> u >> v; u--; v--;
        g[u].push_back(v);
        g[v].push_back(u);
    }

    vector<ll> f(n);

    // dfs 遍历
    auto dfs = [&](this auto&& dfs, int u, int fa) -> DynamicLiChaoTree* {
        DynamicLiChaoTree* tree = new DynamicLiChaoTree(MN, MX);
        for (auto v : g[u]) {
            if (v == fa) continue;
            DynamicLiChaoTree* child = dfs(v, u);
            tree->merge(*child);
            delete child;
        }
        if (g[u].size() > 1 || g[u][0] != fa) {
            f[u] = tree->query(a[u]);
        }

        // 比如我维护的是 3, 5 节点，查询完 1 节点后，需要把 1 节点加到线段树里返回，作为 0 的子树
        Line line(b[u], f[u], u);
        tree->insert(MN, MX, line);
        return tree;
    };

    DynamicLiChaoTree* root_tree = dfs(0, -1);
    // root_tree 只是为了能够正常递归，实际无用可删除
    delete root_tree;
    
    for (int i = 0; i < n; ++i) {
        cout << f[i] << ' ';
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}