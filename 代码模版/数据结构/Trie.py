# 字典树
class Node:
    __slots__ = "son", "end"
    def __init__(self):
        self.son = {}
        self.end = False

class Trie:

    def __init__(self):
        self.root = Node()

    def insert(self, word: str) -> None:
        cur = self.root
        for c in word:
            if c not in cur.son:
                cur.son[c] = Node()
            cur = cur.son[c]
        cur.end = True

    def find(self, word: str) -> int:
        cur = self.root
        for c in word:
            if c not in cur.son:
                return 0
            cur = cur.son[c]
        return 2 if cur.end else 1

    def search(self, word: str) -> bool:
        return self.find(word) == 2

    def startsWith(self, prefix: str) -> bool:
        return self.find(prefix) > 0
    
# 0-1 字典树
class Node:
    __slots__ = 'children', 'cnt'

    def __init__(self):
        self.children = [None, None]
        self.cnt = 0  # 子树大小

class Trie:
    def __init__(self, HIGH_BIT = 30):
        self.root = Node()
        self.HIGH_BIT = HIGH_BIT

    # 添加 val
    def insert(self, val: int) -> None:
        cur = self.root
        for i in range(self.HIGH_BIT, -1, -1):
            bit = (val >> i) & 1
            if cur.children[bit] is None:
                cur.children[bit] = Node()
            cur = cur.children[bit]
            cur.cnt += 1  # 维护子树大小
        return cur

    # 删除 val，但不删除节点
    # 要求 val 必须在 trie 中
    def remove(self, val: int) -> None:
        cur = self.root
        for i in range(self.HIGH_BIT, -1, -1):
            cur = cur.children[(val >> i) & 1]
            cur.cnt -= 1  # 维护子树大小
        return cur

    # 返回 val 与 trie 中一个元素的最大异或和
    # 要求 trie 中至少有一个元素
    def max_xor(self, val: int) -> int:
        cur = self.root
        ans = 0
        for i in range(self.HIGH_BIT, -1, -1):
            bit = (val >> i) & 1
            # 如果 cur.children[bit^1].cnt == 0，视作空节点
            if cur.children[bit ^ 1] and cur.children[bit ^ 1].cnt:
                ans |= 1 << i
                bit ^= 1
            cur = cur.children[bit]
        return ans

# 可持久化字典树
class PersistentTrie:
    def __init__(self, max_bit=30):
        self.max_bit = max_bit
        self.ch0 = [0]
        self.ch1 = [0]
        self.cnt = [0]
        self.roots = [0]

    def _new_node(self):
        self.ch0.append(0)
        self.ch1.append(0)
        self.cnt.append(0)
        return len(self.cnt) - 1

    def insert(self, prev_root, value):
        new_root = self._new_node()
        cur_new = new_root
        cur_old = prev_root
        self.cnt[cur_new] = self.cnt[cur_old] + 1

        for bit in range(self.max_bit - 1, -1, -1):
            b = (value >> bit) & 1
            if b == 0:
                self.ch1[cur_new] = self.ch1[cur_old]
                left_new = self._new_node()
                self.ch0[cur_new] = left_new
                old_left = self.ch0[cur_old]
                self.cnt[left_new] = self.cnt[old_left] + 1
                cur_new = left_new
                cur_old = old_left
            else:
                self.ch0[cur_new] = self.ch0[cur_old]
                right_new = self._new_node()
                self.ch1[cur_new] = right_new
                old_right = self.ch1[cur_old]
                self.cnt[right_new] = self.cnt[old_right] + 1
                cur_new = right_new
                cur_old = old_right
        return new_root

    def query_max_xor(self, root1, root2, value):
        node1 = root1
        node2 = root2
        res = 0
        for bit in range(self.max_bit - 1, -1, -1):
            b = (value >> bit) & 1
            if b == 0:
                if self.cnt[self.ch1[node1]] - self.cnt[self.ch1[node2]] > 0:
                    res |= (1 << bit)
                    node1 = self.ch1[node1]
                    node2 = self.ch1[node2]
                else:
                    node1 = self.ch0[node1]
                    node2 = self.ch0[node2]
            else:
                if self.cnt[self.ch0[node1]] - self.cnt[self.ch0[node2]] > 0:
                    res |= (1 << bit)
                    node1 = self.ch0[node1]
                    node2 = self.ch0[node2]
                else:
                    node1 = self.ch1[node1]
                    node2 = self.ch1[node2]
        return res