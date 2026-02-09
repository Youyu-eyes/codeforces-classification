#include <bits/stdc++.h>
using namespace std;

class SparseTable {
    vector<vector<int>> st_min;
    vector<vector<int>> st_max;

public:
    // 时间复杂度 O(n * log n)
    SparseTable(const vector<int>& nums) {
        int n = nums.size();
        int w = bit_width((uint32_t) n);
        st_min.resize(w, vector<int>(n));
        st_max.resize(w, vector<int>(n));

        for (int j = 0; j < n; j++) {
            st_min[0][j] = nums[j];
            st_max[0][j] = nums[j];
        }

        for (int i = 1; i < w; i++) {
            for (int j = 0; j + (1 << i) <= n; j++) {
                st_min[i][j] = min(st_min[i - 1][j], st_min[i - 1][j + (1 << (i - 1))]);
                st_max[i][j] = max(st_max[i - 1][j], st_max[i - 1][j + (1 << (i - 1))]);
            }
        }
    }

    // [l, r) 左闭右开，下标从 0 开始
    // 必须保证 l < r
    // 时间复杂度 O(1)
    int query_min(int l, int r) const {
        int k = bit_width((uint32_t) r - l) - 1;
        return min(st_min[k][l], st_min[k][r - (1 << k)]);
    }

    // [l, r) 左闭右开，下标从 0 开始
    // 必须保证 l < r
    // 时间复杂度 O(1)
    int query_max(int l, int r) const {
        int k = bit_width((uint32_t) r - l) - 1;
        return max(st_max[k][l], st_max[k][r - (1 << k)]);
    }
};