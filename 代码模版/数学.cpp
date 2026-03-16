#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = LLONG_MAX;

// 时间复杂度 O(sqrt(n))
bool is__prime(int n) {
    for (int i = 2; i * i <= n; i++) {
        if (n % i == 0) {
            return false;
        }
    }
    return n >= 2; // 1 不是质数
}


// 埃氏筛预处理质数
// 时间复杂度 O(MX * log log MX)
constexpr int MX = 1'000'001;
bool is_prime[MX];
vector<int> primes;

auto init = [] {
    ranges::fill(is_prime, true);
    is_prime[0] = is_prime[1] = false; // 0 和 1 不是质数
    for (int i = 2; i < MX; i++) {
        if (is_prime[i]) {
            primes.push_back(i);
            for (long long j = 1LL * i * i; j < MX; j += i) {
                is_prime[j] = false; // j 是质数 i 的倍数
            }
        }
    }
    return 0;
}();


// 质因数分解
// 例如 prime_factorization(45) = {{3, 2}, {5, 1}}，表示 45 = 3^2 * 5^1
// 时间复杂度 O(log x)
constexpr int MX = 1'000'001;
int lpf[MX];

int init = [] {
    for (int i = 2; i < MX; i++) {
        if (lpf[i] == 0) { // i 是质数
            for (int j = i; j < MX; j += i) {
                if (lpf[j] == 0) { // 首次访问 j
                    lpf[j] = i;
                }
            }
        }
    }
    return 0;
}();

vector<pair<int, int>> prime_factorization(int x) {
    vector<pair<int, int>> res;
    while (x > 1) {
        int p = lpf[x];
        int e = 1;
        for (x /= p; x % p == 0; x /= p) {
            e++;
        }
        res.emplace_back(p, e);
    }
    return res;
}

// 预处理因子
constexpr int MX = 1'000'001; // **根据题目调整**
vector<int> divisors[MX];

int init = [] {
    for (int i = 1; i < MX; i++) {
        for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
            divisors[j].push_back(i); // i 是 j 的因子
        }
    }
    return 0;
}();

// 离散化
vector<int> discretize(vector<int>& arr, int& m) {
    vector<int> sorted_unique = arr;
    sort(sorted_unique.begin(), sorted_unique.end());
    sorted_unique.erase(unique(sorted_unique.begin(), sorted_unique.end()), sorted_unique.end());
    m = sorted_unique.size();
    vector<int> discretized;
    for (int x : arr) {
        discretized.push_back(lower_bound(sorted_unique.begin(), sorted_unique.end(), x) - sorted_unique.begin());
    }
    return discretized;
}
// 初始化
// auto discretized_nums = discretize(nums, m);


// 凸包
struct Vec {
    int x, y;

    Vec sub(const Vec& b) const {
        return {x - b.x, y - b.y};
    }

    long long det(const Vec& b) const {
        return 1LL * x * b.y - 1LL * y * b.x;
    }
};

// Andrew 算法，计算 points 的凸包（逆时针顺序）
// 时间复杂度 O(n log n)，其中 n = points.size()
vector<Vec> convexHull(vector<Vec>& points) {
    ranges::sort(points, {}, [](auto& p) { return pair(p.x, p.y); });

    vector<Vec> q;

    // 计算下凸包（从左到右）
    for (auto& p : points) {
        // 新来的点 p，能否让旧的点变成在凸包内的点？ ->  需要判断向量左右关系  ->  det
        while (q.size() > 1 && q.back().sub(q[q.size() - 2]).det(p.sub(q.back())) <= 0) {
            q.pop_back();
        }
        q.push_back(p);
    }

    // 计算上凸包（从右到左）
    // 注意下凸包的最后一个点，已经是上凸包的（右边）第一个点了，所以从 n-2 开始遍历
    int lower_size = q.size();
    for (int i = (int) points.size() - 2; i >= 0; i--) {
        auto& p = points[i];
        while (q.size() > lower_size && q.back().sub(q[q.size() - 2]).det(p.sub(q.back())) <= 0) {
            q.pop_back();
        }
        q.push_back(p);
    }

    // 此时首尾是同一个点 points[0]，需要去掉
    q.pop_back();

    return q;
}


// 快速幂
double myPow(double x, int N) {
    double ans = 1;
    long long n = N;
    if (n < 0) {
        n = -n;
        x = 1 / x;
    }
    while (n) {
        if (n & 1) {
            ans *= x;
        }
        x *= x;
        n >>= 1;
    }
    return ans;
}


// 矩阵快速幂
using matrix = vector<vector<long long>>;

// 返回矩阵 a 和矩阵 b 相乘的结果
matrix mul(matrix& a, matrix& b) {
    int n = a.size(), m = b[0].size();
    matrix c = matrix(n, vector<long long>(m));
    for (int i = 0; i < n; i++) {
        for (int k = 0; k < a[i].size(); k++) {
            if (a[i][k] == 0) {
                continue;
            }
            for (int j = 0; j < m; j++) {
                c[i][j] += a[i][k] * b[k][j];
            }
        }
    }
    return c;
}

// a^n * f0
matrix pow_mul(matrix a, int n, matrix& f0) {
    matrix res = f0;
    while (n) {
        if (n & 1) {
            res = mul(a, res);
        }
        a = mul(a, a);
        n >>= 1;
    }
    return res;
}