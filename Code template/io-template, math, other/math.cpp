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
    ll x, y;

    Vec operator-(const Vec& b) {
        return Vec(x - b.x, y - b.y);
    }

    // a.det(b) > 0 => a 到 b 逆时针
    // a.det(b) < 0 => a 到 b 顺时针
    __int128 det(const Vec& b) {
        return (__int128) x * b.y - (__int128) y * b.x;
    }

    ll dot(const Vec& b) {
        return x * b.x + y * b.y;
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
        while (q.size() > 1 && (p - q.back()).det(q.back() - q[q.size() - 2]) >= 0) {
            q.pop_back();
        }
        q.push_back(p);
    }

    // 计算上凸包（从右到左）
    // 注意下凸包的最后一个点，已经是上凸包的（右边）第一个点了，所以从 n-2 开始遍历
    int lower_size = q.size();
    for (int i = (int) points.size() - 2; i >= 0; i--) {
        auto& p = points[i];
        while (q.size() > lower_size && (p - q[q.size() - 2]).det(q.back() - q[q.size() - 2]) <= 0) {
            q.pop_back();
        }
        q.push_back(p);
    }

    // 此时首尾是同一个点 points[0]，需要去掉
    q.pop_back();

    return q;
}


// 矩阵快速幂模板
using namespace std;
using matrix = vector<vector<long long>>;

// 返回矩阵 a 和矩阵 b 相乘的结果，若 mod > 0 则取模
matrix mul(matrix& a, matrix& b, long long mod = 0) {
    int n = a.size(), m = b[0].size();
    matrix c(n, vector<long long>(m, 0));
    for (int i = 0; i < n; i++) {
        for (int k = 0; k < a[i].size(); k++) {
            if (a[i][k] == 0) continue;
            long long aik = a[i][k];
            for (int j = 0; j < m; j++) {
                c[i][j] += aik * b[k][j];
                if (mod > 0) c[i][j] %= mod;
            }
        }
    }
    return c;
}

// 计算 A^n * f0，若 mod > 0 则所有乘法取模
matrix pow_mul(matrix a, int n, matrix& f0, long long mod = 0) {
    matrix res = f0;   // 初始结果即为 f0
    while (n) {
        if (n & 1) {
            res = mul(a, res, mod);   // 左乘 A
        }
        a = mul(a, a, mod);           // A 自乘
        n >>= 1;
    }
    return res;
}


// 预处理阶乘及逆元
const int MOD = 1'000'000'007;
const int MX = 100'001; // 根据题目数据范围修改

long long F[MX]; // F[i] = i!
long long INV_F[MX]; // INV_F[i] = i!^-1 = qpow(i!, MOD-2)

// 快速幂模板，支持负指数
template<typename T = long long>
T qpow(long long x, long long n, long long mod = 0) {
    long long ans = 1;
    long long base = x;
    long long exp = n;

    if (mod) base %= mod;

    if (exp < 0) {
        if (!mod) return T(0);
        // 模质数下求逆元（费马小定理）
        long long inv = 1, b = base, p = mod - 2;
        while (p) {
            if (p & 1) inv = (inv * b) % mod;
            b = (b * b) % mod;
            p >>= 1;
        }
        base = inv;
        exp = -exp;
    }

    while (exp) {
        if (exp & 1) {
            ans *= base;
            if (mod) ans %= mod;
        }
        base *= base;
        if (mod) base %= mod;
        exp >>= 1;
    }
    return T(ans);
}

// 预处理阶乘和逆元
auto init = [] {
    F[0] = 1;
    for (int i = 1; i < MX; i++) {
        F[i] = F[i - 1] * i % MOD;
    }

    INV_F[MX - 1] = qpow(F[MX - 1], MOD - 2);
    for (int i = MX - 1; i; i--) {
        INV_F[i - 1] = INV_F[i] * i % MOD;
    }
    return 0;
}();

// 从 n 个数中选 m 个数的方案数
long long comb(int n, int m) {
    return m < 0 || m > n ? 0 : F[n] * INV_F[m] % MOD * INV_F[n - m] % MOD;
}

// 返回第 n 个卡特兰数 (C_n) 模 MOD
long long catalan(int n) {
    assert(0 <= n && n <= MX);
    if (n == 0) return 1;
    // C_n = C(2n, n) / (n+1)
    long long C_2n_n = comb(2 * n, n);
    long long inv_n1 = qpow<long long>(n + 1, -1, MOD);   // 求 (n+1) 的逆元
    return C_2n_n * inv_n1 % MOD;
}