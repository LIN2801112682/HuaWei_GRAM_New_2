package go_dic

import (
	"fmt"
	"sort"
)

type TrieTreeNode struct {
	data      string
	frequency int
	children  []*TrieTreeNode
	isleaf    bool
}

func GetData(node *TrieTreeNode) string {
	return node.data
}

func GetFrequency(node *TrieTreeNode) int {
	return node.frequency
}

func IsLeaf(node *TrieTreeNode) bool {
	return node.isleaf
}

func GetChildren(node *TrieTreeNode) []*TrieTreeNode {
	return node.children
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
func PruneNode(node *TrieTreeNode, T int) {
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
func PruneStrategyLessT(node *TrieTreeNode) {
	node.children = make([]*TrieTreeNode, 0)
}

//剪枝策略>T
//剪掉最大子集，若无法剪枝则递归剪子树
func PruneStrategyMoreT(node *TrieTreeNode, T int) {
	//arraylength := len(node.children)
	//frequencylist := make([]int, arraylength)

	// TODO 检验新的剪枝策略
	// 将每个节点的频率和节点本身保存到一个 map 中
	frequencyList := make(map[int][]map[int]*TrieTreeNode)
	for idx, child := range node.children {
		freq := child.frequency
		value, ok := frequencyList[freq]
		if !ok {
			value = make([]map[int]*TrieTreeNode, 0)
		}
		node := make(map[int]*TrieTreeNode)
		node[idx] = child
		value = append(value, node)
		frequencyList[freq] = value
	}
	// 统计map 的所有key ,即频率
	var keys []int
	for k, _ := range frequencyList {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var totalSum = 0
	for i := len(keys) - 1; i >= 0; i-- {
		if totalSum+keys[i] <= T {
			value, _ := frequencyList[keys[i]]
			for _, v := range value {
				if totalSum+keys[i] <= T {
					totalSum += keys[i]
					for idx, node := range v {
						NodeArrayRemoveStrategy(&node.children, idx)
					}
				} else {
					break
				}
			}
		}
	}

	/* 之前的剪枝策略
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
				// TODO 剪枝策略有问题
				if child.frequency == frequencylist[i] {
					//找到对应枝条，进行剪枝
					//删除该孩子节点
					NodeArrayRemoveStrategy(&node.children, j)
					break
				}
			}
		}
	}’
	*/
	//if(totoalsum == 0){ //
	// 不存在最大子集
	for _, child := range node.children {
		PruneStrategyMoreT(child, T)
	}
	//}
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
		if child.data == char {
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
	fmt.Print(node.data, " - ", node.frequency, " - ", node.isleaf)
	for _, child := range node.children {
		PrintTreeNode(child, level+1)
	}
}
