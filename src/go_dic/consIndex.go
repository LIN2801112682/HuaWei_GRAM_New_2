package go_dic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//根据一批日志数据通过字典树划分VG，构建索引项集
func GererateIndex(filename string, qmin int, qmax int, root *trieTreeNode) *indexTree {
	start2 := time.Now()
	indexTree := NewIndexTree(qmin, qmax)
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Print(err)
	}
	buff := bufio.NewReader(data)
	sid := 0
	for {
		data, _, eof := buff.ReadLine()

		if eof == io.EOF {
			break
		}
		var vgMap map[int]string
		vgMap = make(map[int]string)
		sid++
		str := string(data)
		VGCons(root, qmin, qmax, str, vgMap)
		for vgKey := range vgMap {
			//字符串变字符串数组
			gram := make([]string, len(vgMap[vgKey]))
			for j := 0; j < len(vgMap[vgKey]); j++ {
				gram[j] = vgMap[vgKey][j : j+1]
			}
			InsertIntoIndexTree(indexTree, &gram, sid, vgKey)
		}
	}
	indexTree.cout = sid
	UpdateIndexRootFrequency(indexTree)
	elapsed2 := time.Since(start2)
	fmt.Println("构建索引项集花费时间（ms）：", elapsed2)
	PrintIndexTree(indexTree)
	return indexTree
}

//根据字典D划分日志为VG
func VGCons(root *trieTreeNode, qmin int, qmax int, str string, vgMap map[int]string) {
	len := len(str)
	for p := 0; p < len-qmin+1; p++ {
		tSub = ""
		FindLongestGramFromDic(root, str, p)
		t := tSub
		if t == "" {
			t = str[p : p+qmin-1]
		}
		if !isSubStrOfVG(t, vgMap) {
			vgMap[p] = t
		}
	}
}

func isSubStrOfVG(t string, vgMap map[int]string) bool {
	var flag = false
	for vgKey := range vgMap {
		str := vgMap[vgKey]
		if strings.Contains(str, t) {
			flag = true
			break
		}
	}
	return flag
}

var tSub string

func FindLongestGramFromDic(root *trieTreeNode, str string, p int) {
	if p < len(str) {
		c := str[p : p+1]
		for i := 0; i < len(root.children); i++ {
			if root.children[i].data == c {
				tSub += c
				FindLongestGramFromDic(root.children[i], str, p+1)
			}
		}
	}
}
