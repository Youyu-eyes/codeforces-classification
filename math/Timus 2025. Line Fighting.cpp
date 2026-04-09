// https://acm.timus.ru/problem.aspx?space=1&num=2025
// 组合数学

// 首先，我们要计算的是总的打架次数，应当为可能的打架次数 - 实际上没打的次数
// 可能的打架次数为在 n 个人中任选 2 个人，可能的组合方式为 C(2, n)
// 实际上没打的次数为同队的人没打，假设第 i 队 ai 个人，可能的组合方式为 C(2, ai)，ai >= 2
// 组合数展开后 c(2, ai) = ai * (ai - 1) / 2，这个式子 ai = 1 时为 0，可以正常计算
// 总的式子为 n * (n - 1) / 2 - sum{i = [1, k]}(ai * (ai - 1) / 2)
// 前一项是定值，要使答案最大，后一项最小
// 将后一项式子进一步展开
// a1 * (a1 - 1) / 2 + a2 * (a2 - 1) / 2 + …… + am * (am - 1) / 2
// = (a1 * a1 + a2 * a2 + …… + am * am - n) / 2  (m <= k)
// 问题就变成了，在 sum(ai) 不变的情况下，怎么使平方和最小

// 先考虑 ai 为实数的情况
// 由柯西不等式
// (a1^2 + a2^2 + …… + am^2) * (1^2 + 1^2 + …… + 1^2) >= (a1 + a2 + …… + am)^2 = n ^ 2
// sum(ai^2) >= n^2 / m  当且仅当 a1 = a2 = a3 = …… = am 时取等号
// 假设 a2, a3, a4, ……, a(m-1) 为定值，则 a1 + am 为定值，假设 a1 为自变量，sum(ai^2) 为因变量
// 当 a1 = am 时，sum 最小；a1 增大，sum 增大；a1 减小，sum 增大
// 可见这个级数是下凸的
// 且 m 越大，和越小，原式关于 m 单调
// 离散是连续的一般情况，当每个 ai 越接近的时候和式越小
// 如果存在 ai = aj + 2，那么我们一定可以分成两个 aj + 1
// ai^2 + aj^2 = 2*aj^2 + 4*aj + 4
// (aj + 1) ^ 2 = 2*aj^2 + 4*aj + 2
// 可见下式一定小于上式，与 aj 取值无关
// 所以当两数之差大于 1 时，我们一定可以对差值重新分配使得和式满足帕累托最优

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

void solve() {
    int n, k; cin >> n >> k;
    ll ans = (n * (n - 1) / 2) - ((n / k) * (n / k) * (k - n % k) + (n / k + 1) * (n / k + 1) * (n % k) - n) / 2;
    cout << ans << endl;
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