#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const ll INF = 1e18;

// 字典树
struct Node {
    Node* son[26]{};
    bool end = false;
};

class Trie {
    Node* root = new Node();

    int find(string word) {
        Node* cur = root;
        for (char c : word) {
            c -= 'a';
            if (cur->son[c] == nullptr) { // 道不同，不相为谋
                return 0;
            }
            cur = cur->son[c];
        }
        // 走过同样的路（2=完全匹配，1=前缀匹配）
        return cur->end ? 2 : 1;
    }

    void destroy(Node* node) {
        if (node == nullptr) {
            return;
        }
        for (Node* son : node->son) {
            destroy(son);
        }
        delete node;
    }

public:
    ~Trie() {
        destroy(root);
    }

    void insert(string word) {
        Node* cur = root;
        for (char c : word) {
            c -= 'a';
            if (cur->son[c] == nullptr) { // 无路可走？
                cur->son[c] = new Node(); // new 出来！
            }
            cur = cur->son[c];
        }
        cur->end = true;
    }

    bool search(string word) {
        return find(word) == 2;
    }

    bool startsWith(string prefix) {
        return find(prefix) != 0;
    }
};

// 0-1 字典树
constexpr int WIDTH = 30; // nums[i] 二进制长度的最大值

struct TrieNode {
    TrieNode* son[2]{};
    int leaf = 0; // 子树叶子个数
};

class Trie {
    // 注：我没有写析构函数，想写的同学可以自己补上
    TrieNode* root = new TrieNode();

public:
    void insert(int val) {
        TrieNode* cur = root;
        for (int i = WIDTH - 1; i >= 0; i--) {
            int bit = val >> i & 1;
            if (!cur->son[bit]) {
                cur->son[bit] = new TrieNode();
            }
            cur = cur->son[bit];
            cur->leaf++;
        }
    }

    void remove(int val) {
        TrieNode* cur = root;
        for (int i = WIDTH - 1; i >= 0; i--) {
            int bit = val >> i & 1;
            cur = cur->son[bit];
            cur->leaf--; // 如果减成 0 了，说明子树是空的，可以理解成 cur is None
        }
    }

    int max_xor(int val) {
        TrieNode* cur = root;
        int ans = 0;
        for (int i = WIDTH - 1; i >= 0; i--) {
            int bit = val >> i & 1;
            if (cur->son[bit ^ 1] && cur->son[bit ^ 1]->leaf) {
                ans |= 1 << i;
                bit ^= 1;
            }
            cur = cur->son[bit];
        }
        return ans;
    }
};