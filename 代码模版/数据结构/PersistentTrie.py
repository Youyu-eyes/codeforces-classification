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