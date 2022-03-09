package build_VGram_index

import (
	"bufio"
	"build_dictionary"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

//根据一批日志数据通过字典树划分VG，构建索引项集
func GererateIndex(filename string, qmin int, qmax int, root *build_dictionary.TrieTreeNode) (*IndexTree, *IndexTreeNode) {
	indexTree := NewIndexTree(qmin, qmax)
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Print(err)
	}
	buff := bufio.NewReader(data)
	var id int32
	id = 0
	var sum = 0
	timeStamp := time.Now().Unix()
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		var vgMap map[int]string
		vgMap = make(map[int]string)
		id++
		timeStamp++
		sid := NewSeriesId(id, timeStamp)
		str := string(data)
		start2 := time.Now()
		VGCons(root, qmin, qmax, str, vgMap)
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
			InsertIntoIndexTree(indexTree, &gram, *sid, vgKey)
		}
		end2 := time.Since(start2).Microseconds()
		sum = int(end2) + sum
	}
	indexTree.Cout = (int(id))
	UpdateIndexRootFrequency(indexTree)
	fmt.Println("构建索引项集花费时间（us）：", sum)
	PrintIndexTree(indexTree)
	return indexTree, indexTree.Root
}

//根据字典D划分日志为VG
func VGCons(root *build_dictionary.TrieTreeNode, qmin int, qmax int, str string, vgMap map[int]string) {
	len1 := len(str)
	for p := 0; p < len1-qmin+1; p++ {
		tSub = ""
		FindLongestGramFromDic(root, str, p)
		t := tSub
		if t == "" || len(t) < qmin { //t != str[p:p+len(t)]  目前qmin - qmax之间都是叶子节点也就是说FindLongestGramFromDic找到的只要是长度大于qmin就都是VG的gram
			t = str[p : p+qmin]
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
		if str == t {
			return false
		} else if strings.Contains(str, t) {
			flag = true
			break
		}
	}
	return flag
}

var tSub string

func FindLongestGramFromDic(root *build_dictionary.TrieTreeNode, str string, p int) { //优化：字典序可以二分查找   最多找到qmx长就停止
	if p < len(str) {
		c := str[p : p+1]
		for i := 0; i < len(root.Children); i++ {
			if root.Children[i].Data == c {
				tSub += c
				FindLongestGramFromDic(root.Children[i], str, p+1)
			}
			if i == len(root.Children) {
				return
			}
		}
	}
}
