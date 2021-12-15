package go_dic

import (
	"fmt"
	"sort"
)

type trieTreeNode struct {
	data      string
	frequency int
	children  []*trieTreeNode
	isleaf    bool
}

func NewTrieTreeNode(data string) *trieTreeNode {
	return &trieTreeNode{
		data:      data,
		frequency: 1,
		isleaf:    false,
		children:  make([]*trieTreeNode, 0),
	}
}

//剪枝
func PruneNode(node *trieTreeNode, T int) {
	if !node.isleaf {
		for _, child := range node.children {
			PruneNode(child, T)
		}
	} else {
		if node.frequency <= T {
			PruneStrategyLessT(node)
		} else {
			PruneStrategyMoreT(node, T)
		}
	}
}

//剪枝策略<=T
func PruneStrategyLessT(node *trieTreeNode) {
	node.children = make([]*trieTreeNode, 0)
}

//剪枝策略>T
//剪掉最大子集，若无法剪枝则递归剪子树
func PruneStrategyMoreT(node *trieTreeNode, T int) {
	arraylength := len(node.children)
	frequencylist := make([]int, arraylength)
	for i := 0; i < arraylength; i++ {
		frequencylist[i] = node.children[i].frequency
	}
	sort.Ints(frequencylist)
	totoalsum := 0
	for i := arraylength - 1; i >= 0; i-- {
		//从大到小遍历数组
		if totoalsum+frequencylist[i] <= T {
			totoalsum = totoalsum + frequencylist[i]
			for j, child := range node.children {
				if child.frequency == frequencylist[i] {
					//找到对应枝条，进行剪枝
					//删除该孩子节点
					NodeArrayRemoveStrategy(&node.children, j)
					break
				}
			}
		}
	}
	//if(totoalsum == 0){ //
	// 不存在最大子集
	for _, child := range node.children {
		PruneStrategyMoreT(child, T)
	}
	//}
}

//删除数组策略
func NodeArrayRemoveStrategy(array *[]*trieTreeNode, index int) {
	*array = append((*array)[:index], (*array)[index+1:]...)
}

//插入数组策略
func NodeArrayInsertStrategy(array *[]*trieTreeNode, node *trieTreeNode) {
	*array = append(*array, node)
}

//判断children有无此节点
func getNode(children []*trieTreeNode, char string) int {
	for i, child := range children {
		if child.data == char {
			return i
		}
	}
	return -1
}

//输出以node为根的子树
func PrintTreeNode(node *trieTreeNode, level int) {
	fmt.Println()
	for i := 0; i < level; i++ {
		fmt.Print("      ")
	}
	fmt.Print(node.data, " - ", node.frequency, " - ", node.isleaf)
	for _, child := range node.children {
		PrintTreeNode(child, level+1)
	}
}
