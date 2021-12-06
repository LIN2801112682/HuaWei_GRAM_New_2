package go_dic

type trieTree struct {
	qmin int
	qmax int
	root *trieTreeNode
}

//初始化trieTree
func NewTrieTree(qmin int , qmax int) *trieTree {
	return &trieTree{
		qmin: qmin,
		qmax: qmax,
		root: NewTrieTreeNode(""),
	}
}

//将gram插入trieTree上
//trieTree:待插入的树
//gram:待插入字符串数组
func InsertIntoTrieTree(tree *trieTree, gram *[]string)  {
	//初始化node、qmin
	node := tree.root
	qmin := tree.qmin
	// 孩子节点在childrenlist中的位置
	var childindex = 0;
	for i , char := range *gram{
		childindex = getNode(node.children , (*gram)[i])
		if(childindex == -1){
			// childrenlist里没有该节点
			currentnode := NewTrieTreeNode(char)
			NodeArrayInsertStrategy(&node.children , currentnode)
			node = currentnode
		}else {
			//childrenlist里有该节点
			//childrenindex为该节点在数组中的位置
			node = node.children[childindex]
			node.frequency ++
		}
		if(i >= qmin - 1){
			node.isleaf = true
		}
	}
}

//剪枝
//trieTree:待修剪的树
//T:阈值
func PruneTree(tree *trieTree, T int)  {
	PruneNode(tree.root , T)
}

func PrintTree(tree *trieTree){
	PrintTreeNode(tree.root , 0)
}

//更新root节点的频率
func UpdateRootFrequency(tree *trieTree){
	for _ , child := range tree.root.children{
		tree.root.frequency += child.frequency
	}
	tree.root.frequency --
}
