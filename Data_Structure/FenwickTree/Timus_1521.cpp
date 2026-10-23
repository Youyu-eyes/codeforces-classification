// https://acm.timus.ru/problem.aspx?space=1&num=1521
// 经典题，约瑟夫环

// 树状数组二分
// 线段树做法见线段树题单

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
// 根据题目用 FenwickTree<int> t(n) 或者 FenwickTree<long long> t(n) 初始化
template<typename T>
class FenwickTree {
    vector<T> tree;

public:
    // 使用下标 1 到 n
    FenwickTree(int n) : tree(n + 1) {}

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        for (; i < tree.size(); i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    T pre(int i) const {
        T res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }

    // 求区间和 a[l] + ... + a[r]
    // 1 <= l <= r <= n
    // 时间复杂度 O(log n)
    T query(int l, int r) const {
        if (r < l) {
            return 0;
        }
        return pre(r) - pre(l - 1);
    }

    // 返回满足前缀和 >= k 的最小下标（1-based）
    // 假设 1 <= k <= 总元素个数，且序列元素非负
    int lower_bound(int k) const {
        int pos = 0;
        int n = tree.size() - 1;
        int bit_mask = 1 << (bit_width((unsigned)n) - 1);
        T s = 0;
        while (bit_mask) {
            int nxt = pos + bit_mask;
            if (nxt <= n and s + tree[nxt] < k) {
                pos = nxt;
                s += tree[nxt];
            }
            bit_mask >>= 1;
        }
        return pos + 1;
    }
};

void solve() {
    int n, k; cin >> n >> k;
    FenwickTree<ll> BIT(n);
    for (int i = 1; i <= n; ++i) {
        BIT.update(i, 1);
    }
    int cur = 0;
    for (int i = 0; i < n; ++i) {
        cur = (cur + k - 1) % (n - i);
        int ans = BIT.lower_bound(cur + 1);
        cout << ans << ' ';
        BIT.update(ans, -1);
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}