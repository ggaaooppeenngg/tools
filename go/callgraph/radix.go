package callgraph

import (
	"sort"
	"strings"
)

// RadixNode is a node of radix tree.
type RadixNode struct {
	Prefix string
	// 没有 Edges 的就是叶子节点
	Edges RadixEdges
}

// IsLeaf returns true if node is leaf node (without edges).
func (node *RadixNode) IsLeaf() bool { return len(node.Edges) == 0 }

// AddEdge adds a edge in node.
func (node *RadixNode) AddEdge(e RadixEdge) {
	node.Edges = append(node.Edges, e)
	sort.Sort(node.Edges)
}

func (node *RadixNode) ReplaceEdge(e RadixEdge) {
	num := len(node.Edges)
	idx := sort.Search(num, func(i int) bool {
		return node.Edges[i].Label >= e.Label
	})
	if idx < num && node.Edges[idx].Label == e.Label {
		node.Edges[idx] = e
		return
	}
	panic("replace missing edge")
}

func (n *RadixNode) GetEdge(label byte) (RadixEdge, bool) {
	num := len(n.Edges)
	idx := sort.Search(num, func(i int) bool {
		return n.Edges[i].Label >= label
	})
	if idx < num && n.Edges[idx].Label == label {
		return n.Edges[idx], true
	}
	return RadixEdge{}, false
}

type RadixEdges []RadixEdge
type RadixEdge struct {
	Label      byte
	TargetNode *RadixNode
}

func (r RadixEdges) Less(i, j int) bool { return r[i].Label <= r[j].Label }
func (r RadixEdges) Swap(i, j int)      { r[i], r[j] = r[j], r[i]; return }
func (r RadixEdges) Len() int           { return len(r) }

type RadixTree struct {
	Root *RadixNode
}

// longestPrefix finds the length of the shared prefix
// of two strings
func longestPrefix(k1, k2 string) int {
	var max = len(k1)
	if l := len(k2); l < max {
		max = l
	}
	var i int
	for i = 0; i < max; i++ {
		if k1[i] != k2[i] {
			break
		}
	}
	return i
}

func (tree *RadixTree) Add(s string) {
	n := tree.Root
	search := s
	for {
		if len(search) == 0 {
			break
		}
		parent := n
		e, ok := n.GetEdge(search[0])
		n = e.TargetNode
		// 没有匹配的前缀
		if !ok {
			e := RadixEdge{
				Label: search[0],
				TargetNode: &RadixNode{
					Prefix: search,
				},
			}
			parent.AddEdge(e)
			break
		}
		// 匹配当前前缀
		i := longestPrefix(search, n.Prefix)
		if i == len(n.Prefix) {
			search = search[i:]
			continue
		}
		// 有共同前缀，一种是共同前缀就是search,直接分裂依次
		// 一种是,两者都比共同前缀多一点，就分裂出两个.
		// 是当前前缀的前缀
		child := &RadixNode{
			Prefix: search[:i], // i 后面的部分和 prefix 相同或者不同
		}
		parent.ReplaceEdge(RadixEdge{
			Label:      search[0],
			TargetNode: child,
		})
		child.AddEdge(RadixEdge{
			Label:      n.Prefix[i],
			TargetNode: n,
		})
		n.Prefix = n.Prefix[i:] // 前缀节点，原节点前缀之后的部分
		// 如果search比前缀长还要再插入一个节点
		if len(search[i:]) != 0 {
			child.AddEdge(RadixEdge{
				Label: search[i],
				TargetNode: &RadixNode{
					Prefix: search[i:],
				},
			})
		}
	}
}

// 按最长匹配能匹配到自己就可以
func (tree *RadixTree) LongestPrefix(s string) string {
	n := tree.Root
	search := s
	for {
		if n.IsLeaf() || len(search) == 0 {
			break
		}
		e, ok := n.GetEdge(search[0])
		if !ok {
			break
		}
		n = e.TargetNode
		if strings.HasPrefix(search, n.Prefix) {
			search = search[len(n.Prefix):]
		} else {
			break
		}
	}
	return s[:len(s)-len(search)]
}
