// https://codeforces.com/contest/1513/problem/C
// 1600
// 常系数齐次线性DP

// 对于一个整数 n，将 n 的每个数位拆分出来，答案就是每个数位执行 m 次操作后的长度总和
// 本题一个技巧，对于不同的 n，我们最终需要查找的 DP 表都是一样的，因此对于一个测试用例，我们只需要一个DP表
// 具体做法就是在 solve() 函数外预处理 DP 表，操作次数维度可直接开 max(m) = 200005 大小的数组

// 二维做法
// 分解 n 后，问题变成了一个 0~9 的整数 i，操作 m 次后的长度
// 定义 f[i][j] 为整数 i(0 <= i <= 9)，执行 j 次操作后的长度
// 执行 j 次操作，最简单情况是 i 变成 i + j，即 i + j <= 9，长度为 1
// 否则，将 i 加到 10 需要 10-i 次操作，剩下 j - (10 - i) = j + i - 10 次操作
// 然后问题就变成了 对 1 执行 j + i - 10 次操作的长度 + 对 0 执行 j + i - 10 次操作的长度
// 根据定义 f[i][j] = f[1][j + i - 10] + f[0][j + i - 10]
// 综上，状态转移方程为：
// f[i][j] = 1  (if i + j <= 9)
//         = f[1][j + i - 10] + f[0][j + i - 10]  (else)
// 然后把 n 分解后，查表即可
// 需要注意的是，外层循环先遍历 j，内层遍历 i，顺序不能变
// 因为如果先遍历 i = 0 的情况，0 操作 10 次后会变成 10，更新会用到 f[1][j - 10]，但是 f[1] 我们没有赋初值
// 本题赋初值比较麻烦，不如在循环内直接执行来的直接

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007, MX = 200005;
vector<vector<int>> f(10, vector<int>(MX + 1));

void solve() {
    int n, m; cin >> n >> m;
    
    int ans = 0;
    while (n) {
        int k = n % 10;
        ans += f[k][m];
        ans %= MOD;
        n /= 10;
    }
    cout << ans << endl;
}

int main() {
    // 注意遍历顺序，先遍历 j，后遍历 i
    for (int j = 0; j <= MX; ++j) {
        for (int i = 0; i < 10; ++i) {
            if (i + j <= 9) {
                f[i][j] = 1;
            } else {
                f[i][j] = (f[1][j + i - 10] + f[0][j + i - 10]) % MOD;
            }
        }
    }

    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    int T;
    cin >> T;
    
    while (T--) {
        solve();
    }
    
    return 0;
}


// 一维做法
// 由于每次都需要把 i 先拼凑成 10，当 i = 10 时，对 1 凑 10，对 0 凑 10
// 因此实际上在转移的过程中，只有第一次拆分 n 时用到了 0, 1 以外的数，其他状态都是直接从 0, 1 转移的
// 具体来说：
// f[0][j] = f[1][j - 10] + f[0][j - 10] => f[0][j + 1] - f[0][j - 9] = f[1][j - 9] => f[0][j + 10] - f[0][j] = f[1][j]
// f[1][j] = f[1][j -  9] + f[0][j -  9] => f[0][j + 10] - f[0][j] = f[0][j + 1] - f[0][j - 9] + f[0][j - 9]
// => f[0][j + 10] = f[0][j] + f[0][j + 1] 令 j = j - 10
// => f[0][j] = f[0][j - 10] + f[0][j - 9]
// 即 f[j] = f[j - 10] + f[j - 9]
// 相同代换可以得出 f[1] 的化简结果也是相同的
// 这个具体的定义就是 f[j] 为数字 10 操作 j 次后的长度
// 初始化 f[0~8] = 2，f[9] = 3（此时数字为 109）
 

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int MX = 200005;
vector<int> f1(MX + 1);

void solve() {
    int n, m; cin >> n >> m;
    int ans = 0;
    while (n) {
        int k = n % 10;
        ans += k + m <= 9 ? 1 : f1[m + k - 10];
        ans %= MOD;
        n /= 10;
    }
    cout << ans << endl;
}

int main() {    
    for (int i = 0; i < 9; ++i) {
        f1[i] = 2;
    }
    f1[9] = 3;
    for (int i = 10; i <= MX; ++i) {
        f1[i] = (f1[i - 9] + f1[i - 10]) % MOD;
    }


    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    int T;
    cin >> T;
    
    while (T--) {
        solve();
    }
    
    return 0;
}