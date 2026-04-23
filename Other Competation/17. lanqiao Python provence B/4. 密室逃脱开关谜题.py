import sys
from collections import deque

def solve():
    # 使用 sys.stdin.read 实现快速 I/O
    input_data = sys.stdin.read().split()
    if not input_data:
        return
    
    t = int(input_data[0])
    idx = 1
    out = []
    
    for _ in range(t):
        n = int(input_data[idx])
        m = int(input_data[idx+1])
        idx += 2
        
        # 链式前向星建图
        head = [-1] * m
        to = []
        nxt = []
        edge_id = []
        
        deg = [0] * m
        T = [1] * m  # 目标奇偶性：全都需要被翻转1次（奇数次）
        
        def add_edge(u, v, eid):
            to.append(v)
            edge_id.append(eid)
            nxt.append(head[u])
            head[u] = len(to) - 1
            deg[u] += 1
            
            # 如果不是自环，建立双向映射
            if u != v:
                to.append(u)
                edge_id.append(eid)
                nxt.append(head[v])
                head[v] = len(to) - 1
                deg[v] += 1

        # 真正有意义的独立开关只有 min(n, m) 个
        num_edges = min(n, m)
        for i in range(num_edges):
            u = i
            v = (2 * i) % m
            add_edge(u, v, i)
            
        active_edge = [True] * num_edges
        
        # 1. 剥离度数 <= 1 的叶子节点
        q = deque()
        for i in range(m):
            if deg[i] <= 1:
                q.append(i)
                
        ans = 0
        possible = True
        
        while q:
            u = q.popleft()
            if deg[u] == 0:
                if T[u] == 1:
                    possible = False
                    break
                continue
                
            # deg[u] == 1, 找到那条唯一的可用边
            e = -1
            v = -1
            curr = head[u]
            while curr != -1:
                eid = edge_id[curr]
                if active_edge[eid]:
                    e = eid
                    v = to[curr]
                    break
                curr = nxt[curr]
                
            if e == -1:
                continue
                
            # 为了满足 u 的最终要求，边 e 的状态被强制确定为 T[u]
            val = T[u]
            if val == 1:
                ans += 1
                # 翻转被影响的节点
                if u == v:
                    T[u] ^= 1
                else:
                    T[u] ^= 1
                    T[v] ^= 1
                    
            # 移除边，更新拓扑
            active_edge[e] = False
            deg[u] -= 1
            if u != v:
                deg[v] -= 1
                if deg[v] == 1:
                    q.append(v)
                    
        if not possible:
            out.append("-1")
            continue
            
        # 2. 处理剩余的简单环
        for i in range(m):
            if deg[i] > 0:
                curr_node = i
                cycle_nodes = []
                
                # 提取出整个环
                while deg[curr_node] > 0:
                    e = -1
                    v = -1
                    curr = head[curr_node]
                    while curr != -1:
                        eid = edge_id[curr]
                        if active_edge[eid]:
                            e = eid
                            v = to[curr]
                            break
                        curr = nxt[curr]
                        
                    cycle_nodes.append(curr_node)
                    active_edge[e] = False
                    deg[curr_node] -= 1
                    if curr_node != v:
                        deg[v] -= 1
                    curr_node = v
                    
                k = len(cycle_nodes)
                
                # 检查环内的系统一致性
                sum_T = 0
                for node in cycle_nodes:
                    sum_T ^= T[node]
                    
                if sum_T == 1:
                    possible = False
                    break
                    
                # 因为只有两条出路（枚举环第一条边按或不按）
                ones0 = 0
                curr_x = 0
                for j in range(1, k):
                    curr_x ^= T[cycle_nodes[j]]
                    ones0 += curr_x
                    
                ones1 = k - ones0
                ans += min(ones0, ones1)
                
        if not possible:
            out.append("-1")
        else:
            out.append(str(ans))
            
    print('\n'.join(out))

if __name__ == '__main__':
    sys.setrecursionlimit(200000)
    solve()