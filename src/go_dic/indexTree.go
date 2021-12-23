package go_dic

type indexTree struct {
	qmin int
	qmax int
	cout int
	root *indexTreeNode
}

//初始化trieTree
func NewIndexTree(qmin int, qmax int) *indexTree {
	return &indexTree{
		qmin: qmin,
		qmax: qmax,
		cout: 0,
		root: NewIndexTreeNode(""),
	}
}

//将gram插入trieTree上
//TrieTree:待插入的树
//gram:待插入字符串数组
//sid:字符串所属sid
//position:字符串在sid中的位置
func InsertIntoIndexTree(tree *indexTree, gram *[]string, sid int, position int) {
	//初始化node、qmin
	node := tree.root
	qmin := tree.qmin
	// 孩子节点在childrenlist中的位置
	var childindex = 0
	for i, char := range *gram {
		childindex = getIndexNode(node.children, (*gram)[i])
		if childindex == -1 {
			// childrenlist里没有该节点
			currentnode := NewIndexTreeNode(char)
			IndexNodeArrayInsertStrategy(&node.children, currentnode)
			node = currentnode
		} else {
			//childrenlist里有该节点
			//childrenindex为该节点在数组中的位置
			node = node.children[childindex]
			node.frequency++
		}
		//从root的孩子节点开始判断，少一层故大于等于 qmin-1 不是qmin
		if i >= qmin-1 {
			node.isleaf = true
		}
		if i == len(*gram)-1 {
			//叶子节点，需要挂倒排链表
			InsertInvertedIndex(node, sid, position)
		}
	}
}

func PrintIndexTree(tree *indexTree) {
	PrintIndexTreeNode(tree.root, 0)
}

//更新root节点的频率
func UpdateIndexRootFrequency(tree *indexTree) {
	for _, child := range tree.root.children {
		tree.root.frequency += child.frequency
	}
	tree.root.frequency--
}
