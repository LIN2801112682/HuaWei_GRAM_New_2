package index07

import (
	"fmt"
)

type IndexTreeNode struct {
	data          string
	frequency     int
	children      []*IndexTreeNode
	isleaf        bool
	invertedIndex Inverted_index
}

func (i *IndexTreeNode) Data() string {
	return i.data
}

func (i *IndexTreeNode) SetData(data string) {
	i.data = data
}

func (i *IndexTreeNode) Frequency() int {
	return i.frequency
}

func (i *IndexTreeNode) SetFrequency(frequency int) {
	i.frequency = frequency
}

func (i *IndexTreeNode) Children() []*IndexTreeNode {
	return i.children
}

func (i *IndexTreeNode) SetChildren(children []*IndexTreeNode) {
	i.children = children
}

func (i *IndexTreeNode) Isleaf() bool {
	return i.isleaf
}

func (i *IndexTreeNode) SetIsleaf(isleaf bool) {
	i.isleaf = isleaf
}

func (i *IndexTreeNode) InvertedIndex() Inverted_index {
	return i.invertedIndex
}

func (i *IndexTreeNode) SetInvertedIndex(invertedIndex Inverted_index) {
	i.invertedIndex = invertedIndex
}

func NewIndexTreeNode(data string) *IndexTreeNode {
	return &IndexTreeNode{
		data:          data,
		frequency:     1,
		isleaf:        false,
		children:      make([]*IndexTreeNode, 0),
		invertedIndex: make(map[SeriesId][]int),
	}
}

//判断children有无此节点
func GetIndexNode(children []*IndexTreeNode, char string) int {
	for i, child := range children {
		if child.data == char {
			return i
		}
	}
	return -1
}

func (node *IndexTreeNode) InsertPosArrToInvertedIndexMap(sid SeriesId, position int) {
	//倒排列表数组中找到sid的invertedIndex，把position加入到invertedIndex中的posArray中去
	node.invertedIndex[sid] = append(node.invertedIndex[sid], position)
}

//插入倒排
func (node *IndexTreeNode) InsertSidAndPosArrToInvertedIndexMap(sid SeriesId, position int) {
	posArray := []int{}
	posArray = append(posArray, position)
	node.invertedIndex[sid] = posArray
}

//输出以node为根的子树
func (node *IndexTreeNode) PrintIndexTreeNode(level int) {
	fmt.Println()
	for i := 0; i < level; i++ {
		fmt.Print("      ")
	}
	fmt.Print(node.data, " - ", node.frequency, " - ", node.isleaf, " - ", node.invertedIndex)
	for _, child := range node.children {
		child.PrintIndexTreeNode(level + 1)
	}
}
