package service

import (
	"container/list"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type NodeState struct {
	u int  // 当前节点
	s uint // 状态压缩值
}

var (
	n = 6 // 节点数量
	// 邻接表（节点从1开始）
	graph = map[int][]int{}

	// 路径记录 [节点][状态] -> 前驱节点集合
	path = make(map[int]map[uint][]NodeState)
	// 最短距离 [节点][状态] -> 距离
	distance = make(map[int]map[uint]int)
)

func init() {
	// 初始化存储结构
	for i := 1; i <= n; i++ {
		path[i] = make(map[uint][]NodeState)
		distance[i] = make(map[uint]int)
	}
}

func bfs(start int) {
	q := list.New()
	initState := uint(1 << start)

	// 初始化起点
	distance[start][initState] = 0
	path[start][initState] = []NodeState{{u: 0, s: 0}} // 虚拟终点
	q.PushBack(NodeState{u: start, s: initState})

	for q.Len() > 0 {
		current := q.Remove(q.Front()).(NodeState)
		u, s := current.u, current.s
		currentDist := distance[u][s]

		// 遍历邻接节点
		for _, v := range graph[u] {
			newState := s | (1 << v)
			newDist := currentDist + 1

			// 更新最短路径
			if dist, exists := distance[v][newState]; !exists || newDist < dist {
				distance[v][newState] = newDist
				path[v][newState] = []NodeState{current}
				q.PushBack(NodeState{u: v, s: newState})
			} else if newDist == dist {
				path[v][newState] = append(path[v][newState], current)
			}
		}
	}
}

func dfs(u int, s uint, cur []int, result *[][]int) {
	if u == 0 { // 到达虚拟终点
		reversed := make([]int, len(cur))
		for i := 0; i < len(cur); i++ {
			reversed[i] = cur[len(cur)-1-i]
		}
		*result = append(*result, reversed)
		return
	}

	for _, state := range path[u][s] {
		dfs(state.u, state.s, append(cur, u), result)
	}
}

func GetPaths(start int) [][]int {
	// 初始化数据结构
	for i := 1; i <= n; i++ {
		for k := range path[i] {
			delete(path[i], k)
		}
		for k := range distance[i] {
			delete(distance[i], k)
		}
	}

	bfs(start)
	targetState := uint((1 << (n + 1)) - 2) // 目标状态：所有节点都被访问过

	var result [][]int
	dfs(start, targetState, []int{}, &result)
	return result
}

func GetRandomPath(x int) []int {
	rand.Seed(time.Now().UnixNano())

	//var x int
	//fmt.Print("输入起始节点: ")
	//fmt.Scan(&x)

	res := GetPaths(x)
	fmt.Printf("最短路径数量: %d\n", len(res))
	if len(res) == 0 {
		return nil
	}
	fmt.Printf("最短路径长度: %d\n", len(res[0])-1)

	// 随机选择一条路径
	index := rand.Intn(len(res))
	fmt.Print("随机最短路径: ")
	for _, node := range res[index] {
		fmt.Printf("%d ", node)
	}
	fmt.Println()
	return res[index]
}

func GetAllPaths() [][]int {
	var res [][]int
	for i := 1; i <= n; i++ {
		// 初始化数据结构
		for i := 1; i <= n; i++ {
			for k := range path[i] {
				delete(path[i], k)
			}
			for k := range distance[i] {
				delete(distance[i], k)
			}
		}
		bfs(i)
		targetState := uint((1 << (n + 1)) - 2) // 目标状态：所有节点都被访问过

		var result [][]int
		dfs(i, targetState, []int{}, &result)
		res = append(res, result...)
		fmt.Println(result)
	}
	return res
}

func InputPath(path string) [][]int {
	pairs := strings.Split(path, " ")
	for key := range graph {
		delete(graph, key)
	}
	for _, pair := range pairs {
		// 按逗号分割每个节点对
		nodes := strings.Split(pair, ",")
		if len(nodes) != 2 {
			continue // 跳过格式不正确的项
		}

		// 将字符串转换为整数
		from, err1 := strconv.Atoi(strings.TrimSpace(nodes[0]))
		to, err2 := strconv.Atoi(strings.TrimSpace(nodes[1]))

		if err1 == nil && err2 == nil {
			// 将关系添加到地图中
			graph[from] = append(graph[from], to)
			graph[to] = append(graph[to], from)
		}
	}
	fmt.Println(graph)
	res := GetAllPaths()
	return res
}
