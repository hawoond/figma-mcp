package util

import (
	"strings"

	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

func FindNodeByID(root *types.Node, id string) *types.Node {
	if root == nil {
		return nil
	}
	if root.ID == id {
		return root
	}
	for i := range root.Children {
		if found := FindNodeByID(&root.Children[i], id); found != nil {
			return found
		}
	}
	return nil
}

func FindNodesByType(root *types.Node, nodeType string) []*types.Node {
	var results []*types.Node
	walkNodes(root, func(n *types.Node) bool {
		if n.Type == nodeType {
			results = append(results, n)
		}
		return true
	})
	return results
}

func FindNodesByName(root *types.Node, name string, exact bool) []*types.Node {
	var results []*types.Node
	walkNodes(root, func(n *types.Node) bool {
		if exact {
			if n.Name == name {
				results = append(results, n)
			}
		} else {
			if strings.Contains(strings.ToLower(n.Name), strings.ToLower(name)) {
				results = append(results, n)
			}
		}
		return true
	})
	return results
}

func WalkNodes(root *types.Node, fn func(node *types.Node, depth int)) {
	walkNodesWithDepth(root, fn, 0)
}

func walkNodesWithDepth(node *types.Node, fn func(node *types.Node, depth int), depth int) {
	if node == nil {
		return
	}
	fn(node, depth)
	for i := range node.Children {
		walkNodesWithDepth(&node.Children[i], fn, depth+1)
	}
}

func walkNodes(node *types.Node, fn func(node *types.Node) bool) {
	if node == nil {
		return
	}
	if !fn(node) {
		return
	}
	for i := range node.Children {
		walkNodes(&node.Children[i], fn)
	}
}

func CollectAllNodeIDs(root *types.Node) []string {
	var ids []string
	walkNodes(root, func(n *types.Node) bool {
		ids = append(ids, n.ID)
		return true
	})
	return ids
}

func CollectImageNodes(root *types.Node) []*types.Node {
	var results []*types.Node
	walkNodes(root, func(n *types.Node) bool {
		for _, fill := range n.Fills {
			if fill.Type == "IMAGE" {
				results = append(results, n)
				return true
			}
		}
		return true
	})
	return results
}

func CollectTextNodes(root *types.Node) []*types.Node {
	return FindNodesByType(root, "TEXT")
}

func CollectComponentInstances(root *types.Node) []*types.Node {
	return FindNodesByType(root, "INSTANCE")
}

func GetNodePath(root *types.Node, targetID string) []*types.Node {
	var path []*types.Node
	if findPath(root, targetID, &path) {
		return path
	}
	return nil
}

func findPath(node *types.Node, targetID string, path *[]*types.Node) bool {
	if node == nil {
		return false
	}
	*path = append(*path, node)
	if node.ID == targetID {
		return true
	}
	for i := range node.Children {
		if findPath(&node.Children[i], targetID, path) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func FlattenNodes(root *types.Node) []*types.Node {
	var nodes []*types.Node
	walkNodes(root, func(n *types.Node) bool {
		nodes = append(nodes, n)
		return true
	})
	return nodes
}

type NodeSummary struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Visible  bool   `json:"visible"`
	Children int    `json:"children"`
}

func SummarizeNode(n *types.Node) NodeSummary {
	visible := true
	if n.Visible != nil {
		visible = *n.Visible
	}
	return NodeSummary{
		ID:       n.ID,
		Name:     n.Name,
		Type:     n.Type,
		Visible:  visible,
		Children: len(n.Children),
	}
}
