package main

import "fmt"

/*
Route Between Nodes
Given a directed graph, design an algorithm to find out whether there is a 
route between two nodes.
*/
type Node struct {
	adjs []*Node
}

func hasRoute(node1, node2 *Node) bool {
	seen := map[*Node]bool{node1: true}
	q := []*Node{node1}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		

		for _, adj := range curr.adjs {
			if adj == node2 {
				return true
			}
			if !seen[adj] {
				q = append(q, adj)
			}
		}
	}
	return false
}


func main() {   
	g := createGraph([][]int{{1, 2}, {2}, {0}}) // 0 connected to 1 & 2, 1 is connected to 2, 2 is connected to 0
	fmt.Println(hasRoute(g[1], g[0]))
	fmt.Println(hasRoute(g[1], g[1]))
}

func createGraph(list [][]int) []*Node {
	g := make([]*Node, len(list))

	for idx, l := range list {
		g[idx] = &Node{adjs: make([]*Node, len(l))}
	}

	for idx, l := range list {
		for i, v := range l {
			g[idx].adjs[i] = g[v]
		}
	}
	return g
}