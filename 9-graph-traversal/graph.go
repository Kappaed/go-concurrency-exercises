package main

import "fmt"



func dfs(node int, parent int, neighmap map[int][]int) {
	fmt.Println(node)
	for _,neighbor := range neighmap[node] {
		if (neighbor != parent) {
			dfs(neighbor, node, neighmap)
		}
	}
}

func bfs(node int, neighmap map[int][]int) {
	queue := [][2]int{{node,-1}}
	for len(queue) > 0 {
		popped := queue[0][0]
		parent := queue[0][1]
		queue = queue[1:]
		fmt.Println(popped)
		for _,neigh := range neighmap[popped]  {
			if (neigh != parent) {
				queue = append(queue, [2]int{neigh, popped})
			}
		}
	}
}
	

func main() {
	edges := [][2]int{{1,2},{2,3},{2,4},{4,5},{4,6},{3,7}}
	adjlist := make(map[int][]int)
	for _,edge := range edges {
		n1 := edge[0]
		n2 := edge[1]
		_, ok1 := adjlist[n1]
		_, ok2 := adjlist[n2]
		if (!ok1) {
			adjlist[n1] = []int{n2}
		} else {
			adjlist[n1] = append(adjlist[n1], n2)
		}
		if (!ok2) {
			adjlist[n2] = []int{n1}
		} else {
			adjlist[n2] = append(adjlist[n2], n1)
		}
	}
	
	fmt.Println("DFS:")
	dfs(1,-1,adjlist)
	fmt.Println("BFS:")
	bfs(1, adjlist)
}
