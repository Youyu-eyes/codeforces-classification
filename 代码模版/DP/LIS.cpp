#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        vector<int> g;
        for (int x : nums) {
            auto it = ranges::lower_bound(g, x);
            if (it == g.end()) {
                g.push_back(x); // >=x 的 g[j] 不存在
            } else {
                *it = x;
            }
        }
        return g.size();
    }
};