package build_dictionary

type TrieTree struct {
	qmin int
	qmax int
	root *TrieTreeNode
}

func NewTrieTree(qmin int, qmax int) *TrieTree {
	return &TrieTree{
		qmin: qmin,
		qmax: qmax,
		root: NewTrieTreeNode(""),
	}
}

//将gram插入trieTree上
//TrieTree:待插入的树
//gram:待插入字符串数组
func InsertIntoTrieTree(tree *TrieTree, gram *[]string) {
	//初始化node、qmin
	node := tree.root
	qmin := tree.qmin
	// 孩子节点在childrenlist中的位置
	var childindex = 0
	for i, char := range *gram {
		childindex = getNode(node.Children, (*gram)[i])
		if childindex == -1 {
			// childrenlist里没有该节点
			currentnode := NewTrieTreeNode(char)
			NodeArrayInsertStrategy(&node.Children, currentnode)
			node = currentnode
		} else {
			//childrenlist里有该节点
			//childrenindex为该节点在数组中的位置
			node = node.Children[childindex]
			node.frequency++
		}
		if i >= qmin-1 {
			node.isleaf = true
		}
	}
}

//剪枝
//TrieTree:待修剪的树
//T:阈值
func PruneTree(tree *TrieTree, T int) {
	PruneNode(tree.root, T)
}

func PrintTree(tree *TrieTree) {
	PrintTreeNode(tree.root, 0)
}

//更新root节点的频率
func UpdateRootFrequency(tree *TrieTree) {
	for _, child := range tree.root.Children {
		tree.root.frequency += child.frequency
	}
	tree.root.frequency--
}
