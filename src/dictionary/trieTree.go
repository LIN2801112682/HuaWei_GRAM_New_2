package dictionary

type TrieTree struct {
	qmin int
	qmax int
	root *TrieTreeNode
}

func (tree *TrieTree) Qmin() int {
	return tree.qmin
}

func (tree *TrieTree) SetQmin(qmin int) {
	tree.qmin = qmin
}

func (tree *TrieTree) Qmax() int {
	return tree.qmax
}

func (tree *TrieTree) SetQmax(qmax int) {
	tree.qmax = qmax
}

func (tree *TrieTree) Root() *TrieTreeNode {
	return tree.root
}

func (tree *TrieTree) SetRoot(root *TrieTreeNode) {
	tree.root = root
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
func (tree *TrieTree) InsertIntoTrieTree(gram *[]string) {
	//初始化node、qmin
	node := tree.root
	qmin := tree.qmin
	// 孩子节点在childrenlist中的位置
	var childindex = 0
	for i, char := range *gram {
		childindex = GetNode(node.children, (*gram)[i])
		if childindex == -1 {
			// childrenlist里没有该节点
			currentnode := NewTrieTreeNode(char)
			node.children = append(node.children, currentnode)
			node = currentnode
		} else {
			//childrenlist里有该节点
			//childrenindex为该节点在数组中的位置
			node = node.children[childindex]
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
func (tree *TrieTree) PruneTree(T int) {
	tree.root.PruneNode(T)
}

func (tree *TrieTree) PrintTree() {
	tree.root.PrintTreeNode(0)
}

//更新root节点的频率
func (tree *TrieTree) UpdateRootFrequency() {
	for _, child := range tree.root.children {
		tree.root.frequency += child.frequency
	}
	tree.root.frequency--
}
