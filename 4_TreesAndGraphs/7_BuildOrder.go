package main

import "fmt"

/*
Build Order
You are given a list of projects and a list of dependencies (which is a list of pairs
of projects, where the second project is dependent on the first project). Ail of a project's
dependencies must be built before the project is. Find a build order that will allow
the projects to be built. If there is no valid build order, return an error.
*/
const (
	Initial = "initial"
	Partial = "partial"
	Completed = "completed"
)

type Node struct {
	name string
	adjs []*Node
	deps int
	state string
}

func findBuildOrderByToposort(projects []string, deps [][]string) []string {
	g := buildGraph(projects, deps)
	q := findStartProjects(g)
	var order []string
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		order = append(order, curr.name)
		for _, adj := range curr.adjs {
			adj.deps--
			if adj.deps == 0 {
				q = append(q, adj)
			}
		}
	}
	if len(order) != len(g) {
		return nil
	}
	return order
}

func findBuildOrderByDFS(projects []string, deps([][]string)) []string {
	g := buildGraph(projects, deps)
	var order []string
	for _, node := range g {
		if !doDFS(node, &order) {
			return nil
		}
	}
	return rev(order)
}

func buildGraph(projects []string, deps [][]string) map[string]*Node {
	g := map[string]*Node{}
	for _, p := range projects {
		g[p] = &Node{name: p, state: Initial}
	}
	for _, dep := range deps {
		g[dep[0]].adjs = append(g[dep[0]].adjs, g[dep[1]])
		g[dep[1]].deps++
	}
	return g
}

func findStartProjects(g map[string]*Node) []*Node {
	var nodes []*Node
	for _, node := range g {
		if node.deps == 0 {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func doDFS(node *Node, order *[]string) bool {
	if node.state == Partial {
		return false
	}
	if node.state == Completed {
		return true
	}
	node.state = Partial
	for _, adj := range node.adjs {
		if !doDFS(adj, order) {
			return false
		}
	}
	node.state = Completed
	*order = append(*order, node.name)
	return true
}

func rev(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	projects := []string{"a", "b", "c", "d", "e","f"}
	deps := [][]string{{"a", "d"}, {"f", "b"}, {"b", "d"}, {"f", "a"}, {"d", "c"}}
	fmt.Println(findBuildOrderByToposort(projects, deps))
	fmt.Println(findBuildOrderByDFS(projects, deps))

	deps = [][]string{{"a", "d"}, {"d", "a"}}
	fmt.Println(findBuildOrderByToposort(projects, deps))
	fmt.Println(findBuildOrderByDFS(projects, deps))
}