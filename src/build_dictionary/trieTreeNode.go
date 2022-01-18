package build_dictionary

import (
	"fmt"
	"sort"
)

//字典节点信息
type TrieTreeNode struct {
	Data      string
	frequency int
	Children  []*TrieTreeNode
	isleaf    bool
}

func NewTrieTreeNode(data string) *TrieTreeNode {
	return &TrieTreeNode{
		Data:      data,
		frequency: 1,
		isleaf:    false,
		Children:  make([]*TrieTreeNode, 0),
	}
}

//剪枝
func PruneNode(node *TrieTreeNode, T int) {
	if !node.isleaf {
		for _, child := range node.Children {
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
func PruneStrategyLessT(node *TrieTreeNode) {
	node.Children = make([]*TrieTreeNode, 0)
}

//剪枝策略>T
//剪掉最大子集，若无法剪枝则递归剪子树
func PruneStrategyMoreT(node *TrieTreeNode, T int) {
	arraylength := len(node.Children)
	frequencylist := make([]int, arraylength)
	for i := 0; i < arraylength; i++ {
		frequencylist[i] = node.Children[i].frequency
	}
	sort.Ints(frequencylist)
	totoalsum := 0
	for i := arraylength - 1; i >= 0; i-- {
		//从大到小遍历数组
		if totoalsum+frequencylist[i] <= T {
			totoalsum = totoalsum + frequencylist[i]
			for j, child := range node.Children {
				// TODO 剪枝策略有问题！需要修改
				if child.frequency == frequencylist[i] {
					//找到对应枝条，进行剪枝
					//删除该孩子节点
					NodeArrayRemoveStrategy(&node.Children, j)
					break
				}
			}
		}
	}
	// 不存在最大子集
	for _, child := range node.Children {
		PruneStrategyMoreT(child, T)
	}
}

//删除数组策略
func NodeArrayRemoveStrategy(array *[]*TrieTreeNode, index int) {
	*array = append((*array)[:index], (*array)[index+1:]...)
}

//插入数组策略
func NodeArrayInsertStrategy(array *[]*TrieTreeNode, node *TrieTreeNode) {
	*array = append(*array, node)
}

//判断children有无此节点
func getNode(children []*TrieTreeNode, char string) int {
	for i, child := range children {
		if child.Data == char {
			return i
		}
	}
	return -1
}

//输出以node为根的子树
func PrintTreeNode(node *TrieTreeNode, level int) {
	fmt.Println()
	for i := 0; i < level; i++ {
		fmt.Print("      ")
	}
	fmt.Print(node.Data, " - ", node.frequency, " - ", node.isleaf)
	for _, child := range node.Children {
		PrintTreeNode(child, level+1)
	}
}
