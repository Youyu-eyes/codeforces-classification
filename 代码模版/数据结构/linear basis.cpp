#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = LLONG_MAX;

class XorBasis {
    vector<int> b;

public:
    // n 为值域最大值 U 的二进制长度，例如 U=1e9 时 n=30
    XorBasis(int n) : b(n) {}

    void insert(int x) {
        // 从高到低遍历，保证计算 max_xor 的时候，参与 XOR 的基的最高位（或者说二进制长度）是互不相同的
        for (int i = b.size() - 1; i >= 0; i--) {
            if (x >> i) { // 由于大于 i 的位都被我们异或成了 0，所以 x >> i 的结果只能是 0 或 1
                if (b[i] == 0) { // x 和之前的基是线性无关的
                    b[i] = x; // 新增一个基，最高位为 i
                    return;
                }
                x ^= b[i]; // 保证每个基的二进制长度互不相同
            }
        }
        // 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
    }

    int max_xor() {
        int res = 0;
        // 从高到低贪心：越高的位，越必须是 1
        // 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
        for (int i = b.size() - 1; i >= 0; i--) {
            res = max(res, res ^ b[i]);
        }
        return res;
    }
};