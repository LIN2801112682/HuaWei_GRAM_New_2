package dictionary

import (
	"fmt"
	"sort"
)

//字典节点信息
type TrieTreeNode struct {
	data      string
	frequency int
	children  []*TrieTreeNode
	isleaf    bool
}

func (t *TrieTreeNode) Data() string {
	return t.data
}

func (t *TrieTreeNode) SetData(data string) {
	t.data = data
}

func (t *TrieTreeNode) Frequency() int {
	return t.frequency
}

func (t *TrieTreeNode) SetFrequency(frequency int) {
	t.frequency = frequency
}

func (t *TrieTreeNode) Children() []*TrieTreeNode {
	return t.children
}

func (t *TrieTreeNode) SetChildren(children []*TrieTreeNode) {
	t.children = children
}

func (t *TrieTreeNode) Isleaf() bool {
	return t.isleaf
}

func (t *TrieTreeNode) SetIsleaf(isleaf bool) {
	t.isleaf = isleaf
}

func NewTrieTreeNode(data string) *TrieTreeNode {
	return &TrieTreeNode{
		data:      data,
		frequency: 1,
		isleaf:    false,
		children:  make([]*TrieTreeNode, 0),
	}
}

//剪枝
func (node *TrieTreeNode) PruneNode(T int) {
	if !node.isleaf {
		for _, child := range node.children {
			child.PruneNode(T)
		}
	} else {
		if node.frequency <= T {
			node.PruneStrategyLessT()
		} else {
			node.PruneStrategyMoreT(T)
		}
	}
}

//剪枝策略<=T
func (node *TrieTreeNode) PruneStrategyLessT() {
	node.children = make([]*TrieTreeNode, 0)
}

//剪枝策略>T
//剪掉最大子集，若无法剪枝则递归剪子树
func (node *TrieTreeNode) PruneStrategyMoreT(T int) {
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
				// TODO 剪枝策略有问题！需要修改
				if child.frequency == frequencylist[i] {
					//找到对应枝条，进行剪枝
					//删除该孩子节点
					node.children = append(node.children[:j], node.children[j+1:]...)
					break
				}
			}
		}
	}
	// 不存在最大子集
	for _, child := range node.children {
		child.PruneStrategyMoreT(T)
	}
}

//判断children有无此节点
func GetNode(children []*TrieTreeNode, char string) int {
	for i, child := range children {
		if child.data == char {
			return i
		}
	}
	return -1
}

//输出以node为根的子树
func (node *TrieTreeNode) PrintTreeNode(level int) {
	fmt.Println()
	for i := 0; i < level; i++ {
		fmt.Print("      ")
	}
	fmt.Print(node.data, " - ", node.frequency, " - ", node.isleaf)
	for _, child := range node.children {
		child.PrintTreeNode(level + 1)
	}
}
