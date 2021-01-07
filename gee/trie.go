package gee

import (
	"fmt"
	"strings"
)

// node wraps each part of a url-pattern as a trie node,
// the complete pattern is stored in a leaf
type node struct {
	pattern  string
	part     string
	children []*node
	isParam  bool
}

func (n *node) String() string {
	return fmt.Sprintf(
		"node{pattern=%s, part=%s, isParam=%t}",
		n.pattern, n.part, n.isParam,
	)
}

// insert registers the pattern into the tires, note that
// parts is splitted from pattern
func (n *node) insert(pattern string, parts []string, curLevel int) {
	if len(parts) == curLevel { // leaf
		n.pattern = pattern
		return
	}
	part := parts[curLevel]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:    part,
			isParam: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, curLevel+1)
}

// search returns the leaf that contains the matched pattern
func (n *node) search(parts []string, curLevel int) *node {
	if len(parts) == curLevel || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[curLevel]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, curLevel+1)
		if result != nil {
			return result
		}
	}
	return nil
}

// travels search the whole trie from n,
// store the footprints in nodes
func (n *node) travel(nodes *([]*node)) {
	if n.pattern != "" {
		*nodes = append(*nodes, n)
	}
	for _, child := range n.children {
		child.travel(nodes)
	}
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isParam {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isParam {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
