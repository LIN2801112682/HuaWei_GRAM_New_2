package indexMaintain

import (
	"bufio"
	"dictionary"
	"fmt"
	"index07"
	"io"
	"os"
	"sort"
	"time"
)

//根据一批日志数据通过字典树划分VG，增加到索引项集中
func AddIndex(filename string, qmin int, qmax int, root *dictionary.TrieTreeNode, indexTree *index07.IndexTree) *index07.IndexTree {
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Print(err)
	}
	buff := bufio.NewReader(data)
	id := indexTree.Cout()
	var sum = 0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		var vgMap map[int]string
		vgMap = make(map[int]string)
		id++
		timeStamp := time.Now().Unix()
		sid := index07.NewSeriesId(int32(id), timeStamp)
		str := string(data)
		start2 := time.Now()
		index07.VGCons(root, qmin, qmax, str, vgMap)
		var keys = []int{}
		for key := range vgMap {
			keys = append(keys, key)
		}
		//对map中的key进行排序（map遍历是无序的）
		sort.Sort(sort.IntSlice(keys))
		for i := 0; i < len(keys); i++ {
			vgKey := keys[i]
			//字符串变字符串数组
			gram := make([]string, len(vgMap[vgKey]))
			for j := 0; j < len(vgMap[vgKey]); j++ {
				gram[j] = vgMap[vgKey][j : j+1]
			}
			indexTree.InsertIntoIndexTree(&gram, *sid, vgKey)
		}
		end2 := time.Since(start2).Microseconds()
		sum = int(end2) + sum
	}
	indexTree.SetCout(id)
	indexTree.Root().SetFrequency(1)
	indexTree.UpdateIndexRootFrequency()
	fmt.Println("新增索引项集花费时间（us）：", sum)
	//indexTree.PrintIndexTree()
	return indexTree
}
