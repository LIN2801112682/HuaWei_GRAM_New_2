package build_VGram_index

import (
	"fmt"
)

type IndexTreeNode struct {
	Data          string
	Frequency     int
	Children      []*IndexTreeNode
	isleaf        bool
	InvertedIndex map[SeriesId][]int
}

func NewIndexTreeNode(data string) *IndexTreeNode {
	return &IndexTreeNode{
		Data:          data,
		Frequency:     1,
		isleaf:        false,
		Children:      make([]*IndexTreeNode, 0),
		InvertedIndex: make(map[SeriesId][]int),
	}
}

//插入数组策略
func IndexNodeArrayInsertStrategy(array *[]*IndexTreeNode, node *IndexTreeNode) {
	*array = append(*array, node)
}

//判断children有无此节点
func getIndexNode(children []*IndexTreeNode, char string) int {
	for i, child := range children {
		if child.Data == char {
			return i
		}
	}
	return -1
}

//输出以node为根的子树
func PrintIndexTreeNode(node *IndexTreeNode, level int) {
	fmt.Println()
	for i := 0; i < level; i++ {
		fmt.Print("      ")
	}
	fmt.Print(node.Data, " - ", node.Frequency, " - ", node.isleaf, " - ", node.InvertedIndex)
	for _, child := range node.Children {
		PrintIndexTreeNode(child, level+1)
	}
}
