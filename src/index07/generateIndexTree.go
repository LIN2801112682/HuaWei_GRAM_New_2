package index07

import (
	"bufio"
	"dictionary"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

//根据一批日志数据通过字典树划分VG，构建索引项集
func GenerateIndexTree(filename string, qmin int, qmax int, root *dictionary.TrieTreeNode) (*IndexTree, *IndexTreeNode) {
	start := time.Now().UnixMicro()
	indexTree := NewIndexTree(qmin, qmax)
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Print(err)
	}
	buff := bufio.NewReader(data)
	var id int32
	id = 0
	var sum1 int64
	sum1 = 0
	var sum2 int64
	sum2 = 0
	timeStamp := time.Now().Unix()
	for {
		start1 := time.Now().UnixMicro()
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
		VGCons(root, qmin, qmax, str, vgMap)
		var keys = []int{}
		for key := range vgMap {
			keys = append(keys, key)
		}
		//对map中的key进行排序（map遍历是无序的）
		sort.Sort(sort.IntSlice(keys))
		end1 := time.Now().UnixMicro()
		sum1 += (end1 - start1)
		start2 := time.Now().UnixMicro()
		for i := 0; i < len(keys); i++ {
			vgKey := keys[i]
			//字符串变字符串数组
			gram := make([]string, len(vgMap[vgKey]))
			for j := 0; j < len(vgMap[vgKey]); j++ {
				gram[j] = vgMap[vgKey][j : j+1]
			}
			indexTree.InsertIntoIndexTree(&gram, *sid, vgKey)
		}
		end2 := time.Now().UnixMicro()
		sum2 += (end2 - start2)
	}
	indexTree.cout = (int(id))
	indexTree.UpdateIndexRootFrequency()
	elapsed := time.Now().UnixMicro()
	fmt.Println("构建索引项集总花费时间（us）：", elapsed-start)
	fmt.Println("读取日志并划分索引项花费时间（us）：", sum1)
	fmt.Println("插入索引树花费时间（us）：", sum2)
	//indexTree.PrintIndexTree()
	return indexTree, indexTree.root
}

//根据字典D划分日志为VG
func VGCons(root *dictionary.TrieTreeNode, qmin int, qmax int, str string, vgMap map[int]string) {
	len1 := len(str)
	for p := 0; p < len1-qmin+1; p++ {
		tSub = ""
		FindLongestGramFromDic(root, str, p)
		t := tSub
		if t == "" || len(t) < qmin { //t != str[p:p+len(t)]  目前qmin - qmax之间都是叶子节点也就是说FindLongestGramFromDic找到的只要是长度大于qmin就都是VG的gram
			t = str[p : p+qmin]
		}
		if !IsSubStrOfVG(t, vgMap) {
			vgMap[p] = t
		}
	}
}

func IsSubStrOfVG(t string, vgMap map[int]string) bool {
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

func FindLongestGramFromDic(root *dictionary.TrieTreeNode, str string, p int) { //优化：字典序可以二分查找   最多找到qmx长就停止
	if p < len(str) {
		c := str[p : p+1]
		for i := 0; i < len(root.Children()); i++ {
			if root.Children()[i].Data() == c {
				tSub += c
				FindLongestGramFromDic(root.Children()[i], str, p+1)
			}
			if i == len(root.Children()) {
				return
			}
		}
	}
}
