#include <bits/stdc++.h>
using namespace std;

using ll = long long;

int main() {
    int n;
    vector nums(n, vector<int>(n));

    // 以数组第一项升序
    ranges::sort(nums, {}, [](auto& a) { return a[0]; });
    // 降序
    ranges::sort(nums, greater<>(), [](const auto& v) { return v[0]; });
    // 双关键字排序
    ranges::sort(nums, [](const auto& a, const auto& b) {
        return tie(a[0], b[1]) < tie(b[0], a[1]);
    });

    return 0;
}