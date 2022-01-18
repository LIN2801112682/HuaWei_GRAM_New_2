package build_VGram_index

import (
	"fmt"
)

type IndexTreeNode struct {
	Data              string
	Frequency         int
	Children          []*IndexTreeNode
	isleaf            bool
	InvertedIndexList []*Inverted_index
}

func NewIndexTreeNode(data string) *IndexTreeNode {
	return &IndexTreeNode{
		Data:              data,
		Frequency:         1,
		isleaf:            false,
		Children:          make([]*IndexTreeNode, 0),
		InvertedIndexList: make([]*Inverted_index, 0),
	}
}

//插入倒排
func InsertInvertedIndex(node *IndexTreeNode, sid int, position int) {
	// 倒排列表数组中创建新inverted_index，并加入到invertedIndexList中
	newInverted := NewInverted_index(sid, position)
	invertedIndexArrayInsertStrategy(&node.InvertedIndexList, newInverted)
}

//插入数组策略
func IndexNodeArrayInsertStrategy(array *[]*IndexTreeNode, node *IndexTreeNode) {
	*array = append(*array, node)
}

//插入倒排链表策略
func invertedIndexArrayInsertStrategy(array *[]*Inverted_index, invertedindex *Inverted_index) {
	*array = append(*array, invertedindex)
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
	fmt.Print(node.Data, " - ", node.Frequency, " - ", node.isleaf)
	for _, invertedIndex := range node.InvertedIndexList {
		fmt.Print("  /  sid : ", invertedIndex.Sid, " position : ", invertedIndex.Position)
	}
	for _, child := range node.Children {
		PrintIndexTreeNode(child, level+1)
	}
}
