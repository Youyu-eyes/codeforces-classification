// https://codeforces.com/contest/1196/problem/D2
// 1600 感觉也没有，大概 1400
// 定长滑窗

// 对于一个周期性的字符串或数组，假设周期为 T，对于下标 i 的字符，对应在第一个周期里的字符的下标为 i % T
// 本题由于周期很短，我们可以枚举全部不同周期对应的 t 字符串
// 定义偏移量 d，下标 i 对应的字符为 t[(i + d) % T]，本题 T = 3，d <= 3，我们可以枚举 d

// 固定了 s 对应的 t 字符串，我们需要在一个长度为 k 的滑窗里找到对应 t 不同的字符数量
// 有点抽象，举个例子，比如：
//    t = RGBRGBRGB
//    s = RGGBGBRBG
// diff = 001100011
// 问题就变成了，在 diff 中找到长度为 k 的子字符串，使得子字符串的元素和最小
// 很经典的定长滑窗，也可以用前缀和解决，但是滑窗没有额外空间开销
// 具体定长滑窗就不过多赘述，参考 leetcode / bilibili 灵茶山艾府 的讲解

// 复杂度：
// 时间复杂度：O(T*n)，其中 T 为周期，本题 T = 3
// 空间复杂度：O(1)

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;

const string t = "RGB";

void solve() {
    int n, k; cin >> n >> k;
    string s; cin >> s;
    int ans = inf;
    for (int d = 0; d < 3; ++d) {
        int cnt = 0;
        for (int i = 0; i < n; ++i) {
            cnt += (s[i] != t[(i + d) % 3]);
            int l = i - k + 1;
            if (l < 0) {
                continue;
            }
            ans = min(ans, cnt);
            cnt -= (s[l] != t[(l + d) % 3]);
        }
    }
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