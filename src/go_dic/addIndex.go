package go_dic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

//根据一批日志数据通过字典树划分VG，增加到索引项集中
func AddIndex(filename string, qmin int, qmax int, root *trieTreeNode, indexTree *indexTree) *indexTree {
	//start2 := time.Now()
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Print(err)
	}
	buff := bufio.NewReader(data)
	sid := indexTree.cout
	var sum = 0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		var vgMap map[int]string
		vgMap = make(map[int]string)
		sid++
		str := string(data)
		start2 := time.Now()
		VGCons(root, qmin, qmax, str, vgMap)
		for vgKey := range vgMap {
			//字符串变字符串数组
			gram := make([]string, len(vgMap[vgKey]))
			for j := 0; j < len(vgMap[vgKey]); j++ {
				gram[j] = vgMap[vgKey][j : j+1]
			}
			InsertIntoIndexTree(indexTree, &gram, sid, vgKey)
		}
		end2 := time.Since(start2).Microseconds()
		sum = int(end2) + sum
	}
	indexTree.cout = sid
	indexTree.root.frequency = 1
	UpdateIndexRootFrequency(indexTree)
	//elapsed2 := time.Since(start2)
	fmt.Println("新增索引项集花费时间（us）：", sum)
	//PrintIndexTree(indexTree)
	return indexTree
}
