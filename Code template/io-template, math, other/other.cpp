#include <bits/stdc++.h>
using namespace std;

using ll = long long;

int main() {
    int n;
    vector nums(n, vector<int>(n));
    vector<int> nums1(n);
    vector<int> nums2(n);

    // 以数组第一项升序
    ranges::sort(nums, {}, [](auto& a) { return a[0]; });
    // 降序
    ranges::sort(nums, greater<>(), [](const auto& v) { return v[0]; });
    // 双关键字排序
    ranges::sort(nums, [](const auto& a, const auto& b) {
        return tie(a[0], b[1]) < tie(b[0], a[1]);
    });
    // nums1 与 nums2 下标一一对应，实现 nums1 升序排序，nums2 保持下标对应排序
    auto zipped = views::zip(nums1, nums2);
    ranges::sort(zipped, std::less{},
                  [](const auto& t) { return std::get<0>(t); });
    
    // 元音字母下标对应
    static constexpr array<int, 128> mp = [] {
        array<int, 128> a;
        a.fill(-1);
        a['a'] = 0;
        a['e'] = 1;
        a['i'] = 2;
        a['o'] = 3;
        a['u'] = 4;
        return a;
    }();

    int cnt[5]{};
    string vowels;
    // 稳定排序，在出现次数相同时，会自动把位置靠前的元音排在前面
    ranges::stable_sort(vowels, {}, [&](char ch) { return -cnt[mp[ch]]; });

    return 0;
}