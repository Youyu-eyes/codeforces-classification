# dict 版本
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

# list 版本
class Node:
    __slots__ = "son", "end"

    def __init__(self):
        self.son = []
        self.end = False

class Trie:

    def __init__(self):
        self.root = Node()
    
    def insert(self, word: str) -> None:
        cur = self.root
        for c in word:
            c = ord(c) - ord('a')
            if c not in cur.son:
                cur.son[c] = Node()
            cur = cur.son
        cur.end = True
    
    def find(self, word: str) -> int:
        cur = self.root
        for c in word:
            c = ord(c) - ord('a')
            if c not in cur.son:
                return 0
            cur = cur.son
        return 2 if cur.end else 0

    def search(self, word: str) -> bool:
        return self.find(word) == 2

    def startsWith(self, prefix: str) -> bool:
        return self.find(prefix) > 0