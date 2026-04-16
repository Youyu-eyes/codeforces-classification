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
    int find_kth(int k) const {
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