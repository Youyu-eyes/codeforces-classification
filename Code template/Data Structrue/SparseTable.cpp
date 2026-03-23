#include <bits/stdc++.h>
using namespace std;

// 区间最值 ST 表
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

// 初始化
// SparseTable st(nums);


// 区间最值 ST 表，下标版本
class SparseTableIndex {
    vector<vector<int>> st_min; // 存储区间最小值下标
    vector<vector<int>> st_max; // 存储区间最大值下标
    vector<int> arr;            // 原数组，用于比较

public:
    // 构造函数，预处理 O(n log n)
    SparseTableIndex(const vector<int>& nums) : arr(nums) {
        int n = nums.size();
        int w = bit_width((unsigned int)n);          // 需要的层数
        st_min.resize(w, vector<int>(n));
        st_max.resize(w, vector<int>(n));

        // 长度为 1 的区间
        for (int i = 0; i < n; ++i) {
            st_min[0][i] = i;
            st_max[0][i] = i;
        }

        // 倍增构建
        for (int i = 1; i < w; ++i) {
            int step = 1 << (i - 1);
            for (int j = 0; j + (1 << i) <= n; ++j) {
                // 最小值下标：取两个子区间中值较小的下标，相等时取左侧
                int left_min = st_min[i-1][j];
                int right_min = st_min[i-1][j+step];
                if (nums[left_min] <= nums[right_min]) {
                    st_min[i][j] = left_min;
                } else {
                    st_min[i][j] = right_min;
                }

                // 最大值下标：取两个子区间中值较大的下标，相等时取左侧
                int left_max = st_max[i-1][j];
                int right_max = st_max[i-1][j+step];
                if (nums[left_max] >= nums[right_max]) {
                    st_max[i][j] = left_max;
                } else {
                    st_max[i][j] = right_max;
                }
            }
        }
    }

    // 查询区间 [l, r) 的最小值下标
    int query_min(int l, int r) const {
        int len = r - l;
        int k = bit_width((unsigned int)len) - 1;    // 最大的 2^k ≤ len
        int left_idx = st_min[k][l];
        int right_idx = st_min[k][r - (1 << k)];
        return (arr[left_idx] <= arr[right_idx]) ? left_idx : right_idx;
    }

    // 查询区间 [l, r) 的最大值下标
    int query_max(int l, int r) const {
        int len = r - l;
        int k = bit_width((unsigned int)len) - 1;
        int left_idx = st_max[k][l];
        int right_idx = st_max[k][r - (1 << k)];
        return (arr[left_idx] >= arr[right_idx]) ? left_idx : right_idx;
    }
};

// 测试用例
int main() {
    vector<int> nums = {3, 1, 4, 1, 5, 9, 2, 6};
    SparseTableIndex st(nums);

    cout << "最小值下标查询：" << endl;
    cout << "区间[1,5) : " << st.query_min(1, 5) << endl;   // 期望 1
    cout << "区间[0,3) : " << st.query_min(0, 3) << endl;   // 期望 1
    cout << "区间[2,7) : " << st.query_min(2, 7) << endl;   // 期望 3

    cout << "最大值下标查询：" << endl;
    cout << "区间[1,5) : " << st.query_max(1, 5) << endl;   // 期望 4
    cout << "区间[0,3) : " << st.query_max(0, 3) << endl;   // 期望 2
    cout << "区间[2,7) : " << st.query_max(2, 7) << endl;   // 期望 5

    return 0;
}