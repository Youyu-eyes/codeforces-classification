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


// ------- 二维 ST 表 ------- //
// 初始化：SparseTable2D st(nums)

class SparseTable2D {
    vector<vector<vector<vector<int>>>> st_min;
    vector<vector<vector<vector<int>>>> st_max;
    vector<vector<vector<vector<int>>>> st_gcd;

public:
    SparseTable2D(const vector<vector<int>>& matrix) {
        if (matrix.empty() || matrix[0].empty()) {
            return;
        }

        int m = matrix.size();
        int n = matrix[0].size();
        int log_m = bit_width((uint32_t) m);
        int log_n = bit_width((uint32_t) n);

        st_min.assign(log_m, vector<vector<vector<int>>>(log_n, vector<vector<int>>(m, vector<int>(n))));
        st_max.assign(log_m, vector<vector<vector<int>>>(log_n, vector<vector<int>>(m, vector<int>(n))));
        st_gcd.assign(log_m, vector<vector<vector<int>>>(log_n, vector<vector<int>>(m, vector<int>(n))));

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                st_min[0][0][i][j] = matrix[i][j];
                st_max[0][0][i][j] = matrix[i][j];
                st_gcd[0][0][i][j] = matrix[i][j];
            }
        }

        for (int l = 1; l < log_n; l++) {
            int length = 1 << (l - 1);
            for (int i = 0; i < m; i++) {
                for (int j = 0; j + (1 << l) <= n; j++) {
                    st_min[0][l][i][j] = min(st_min[0][l - 1][i][j], st_min[0][l - 1][i][j + length]);
                    st_max[0][l][i][j] = max(st_max[0][l - 1][i][j], st_max[0][l - 1][i][j + length]);
                    st_gcd[0][l][i][j] = std::gcd(st_gcd[0][l - 1][i][j], st_gcd[0][l - 1][i][j + length]);
                }
            }
        }

        for (int k = 1; k < log_m; k++) {
            int length_r = 1 << (k - 1);
            for (int l = 0; l < log_n; l++) {
                for (int i = 0; i + (1 << k) <= m; i++) {
                    for (int j = 0; j + (1 << l) <= n; j++) {
                        st_min[k][l][i][j] = min(st_min[k - 1][l][i][j], st_min[k - 1][l][i + length_r][j]);
                        st_max[k][l][i][j] = max(st_max[k - 1][l][i][j], st_max[k - 1][l][i + length_r][j]);
                        st_gcd[k][l][i][j] = std::gcd(st_gcd[k - 1][l][i][j], st_gcd[k - 1][l][i + length_r][j]);
                    }
                }
            }
        }
    }

    int query_min(int r1, int r2, int c1, int c2) const {
        int k = bit_width((uint32_t)(r2 - r1)) - 1;
        int l = bit_width((uint32_t)(c2 - c1)) - 1;

        int len_r = 1 << k;
        int len_c = 1 << l;

        int op1 = st_min[k][l][r1][c1];
        int op2 = st_min[k][l][r2 - len_r][c1];
        int op3 = st_min[k][l][r1][c2 - len_c];
        int op4 = st_min[k][l][r2 - len_r][c2 - len_c];

        return min({op1, op2, op3, op4});
    }

    int query_max(int r1, int r2, int c1, int c2) const {
        int k = bit_width((uint32_t)(r2 - r1)) - 1;
        int l = bit_width((uint32_t)(c2 - c1)) - 1;

        int len_r = 1 << k;
        int len_c = 1 << l;

        int op1 = st_max[k][l][r1][c1];
        int op2 = st_max[k][l][r2 - len_r][c1];
        int op3 = st_max[k][l][r1][c2 - len_c];
        int op4 = st_max[k][l][r2 - len_r][c2 - len_c];

        return max({op1, op2, op3, op4});
    }

    int query_gcd(int r1, int r2, int c1, int c2) const {
        int k = bit_width((uint32_t)(r2 - r1)) - 1;
        int l = bit_width((uint32_t)(c2 - c1)) - 1;

        int len_r = 1 << k;
        int len_c = 1 << l;

        int op1 = st_gcd[k][l][r1][c1];
        int op2 = st_gcd[k][l][r2 - len_r][c1];
        int op3 = st_gcd[k][l][r1][c2 - len_c];
        int op4 = st_gcd[k][l][r2 - len_r][c2 - len_c];

        return std::gcd(std::gcd(op1, op2), std::gcd(op3, op4));
    }
};