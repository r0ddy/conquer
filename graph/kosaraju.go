package graph

import "sort"

type StronglyConnectedComponent []Node

func Kosaraju(g Graph) []StronglyConnectedComponent {
	sccs := make([]StronglyConnectedComponent, 0)
	L := make([]Node, 0)
	DepthFirstSearch(g, DepthFirstSearchOptions{
		AfterRecursion: func(n Node) {
			L = append([]Node{n}, L...)
		},
	})

	scc := make(StronglyConnectedComponent, 0)
	DepthFirstSearch(g, DepthFirstSearchOptions{
		CustomOrder:  L,
		ReverseGraph: true,
		BeforeRecursion: func(n Node) {
			scc = append(scc, n)
		},
		AfterSearch: func(n Node) {
			if len(scc) > 0 {
				sccs = append(sccs, scc)
			}
			scc = make([]Node, 0)
		},
	})
	sortSCCs(sccs)
	return sccs
}

func sortSCCs(sccs []StronglyConnectedComponent) {
	for _, scc := range sccs {
		sortNodes(scc)
	}
	sort.Slice(sccs, func(i, j int) bool {
		if len(sccs[i]) != len(sccs[j]) {
			return len(sccs[i]) < len(sccs[j])
		}
		return sccs[i][0].GetID() < sccs[j][0].GetID()
	})
}
