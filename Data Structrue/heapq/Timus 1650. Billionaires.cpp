// https://acm.timus.ru/problem.aspx?space=1&num=1650
// LazyHeap 懒删除堆

// 考验码力的一个题，不难，但是实现复杂，细节多

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

#include <bits/stdc++.h>
using namespace std;

template<typename T, typename Compare = less<T>>
class LazyHeap {
    priority_queue<T, vector<T>, Compare> pq;
    map<T, int> remove_cnt; // 每个元素剩余需要删除的次数
    size_t sz = 0; // 堆的实际大小

    // 正式执行删除操作
    void apply_remove() {
        while (!pq.empty() && remove_cnt[pq.top()] > 0) {
            remove_cnt[pq.top()]--;
            pq.pop();
        }
    }

public:
    // 删除堆中所有应该删除的元素后，堆的实际大小
    size_t size() {
        return sz;
    }

    // 删除
    void remove(T x) {
        remove_cnt[x]++; // 懒删除
        sz--;
    }

    // 查看堆顶
    T top() {
        apply_remove();
        return pq.top(); // 真正的堆顶
    }

    // 出堆
    T pop() {
        apply_remove();
        sz--;
        T x = pq.top();
        pq.pop();
        return x;
    }

    // 入堆
    void push(T x) {
        if (remove_cnt[x] > 0) {
            remove_cnt[x]--; // 抵消之前的删除
        } else {
            pq.push(x);
        }
        sz++;
    }
};

void solve() {
    map<string, string> p2c;
    map<string, ll> p2w;
    map<string, ll> c2w;
    int n; cin >> n;
    for (int i = 0; i < n; ++i) {
        string name, city; cin >> name >> city;
        ll wealth; cin >> wealth;
        p2c[name] = city;
        p2w[name] = wealth;
        c2w[city] += wealth;
    }

    LazyHeap<pair<ll, string>> lpq;
    for (auto& [k, v] : c2w) {
        lpq.push({v, k});
    }

    int m, k; cin >> m >> k;

    vector<vector<pair<string, string>>> travels(m + 1);
    for (int i = 0; i < k; ++i) {
        int day; cin >> day;
        string name, city; cin >> name >> city;
        travels[day].push_back({name, city});
    }

    map<string, int> ans;
    for (int d = 1; d <= m; ++d) {
        if (lpq.size() > 0) {
            // 白天计算财富
            auto&& [top_wealth, top_city] = lpq.pop();
            if (lpq.size() > 0) {
                auto&& [second_wealth, second_city] = lpq.top();
                if (top_wealth > second_wealth) {
                    ans[top_city]++;
                }
            } else {
                ans[top_city]++;
            }
            lpq.push({top_wealth, top_city});
        }

        // 晚上旅行
        for (auto& [name, new_city] : travels[d]) {
            // 原来的城市减少
            ll delta_wealth = p2w[name];
            string old_city = p2c[name];
            ll old_wealth = c2w[old_city];
            lpq.remove({old_wealth, old_city});
            old_wealth -= delta_wealth;
            lpq.push({old_wealth, old_city});
            c2w[old_city] = old_wealth;

            // 现在的城市增加
            ll new_wealth = c2w[new_city];
            lpq.remove({new_wealth, new_city});
            new_wealth += delta_wealth;
            lpq.push({new_wealth, new_city});
            c2w[new_city] = new_wealth;

            // 人走动
            p2c[name] = new_city;
        }
    }

    for (auto& [city, day] : ans) {
        cout << city << ' ' << day << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}