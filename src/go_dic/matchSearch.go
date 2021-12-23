package go_dic

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

func MatchSearch(searchStr string, root *trieTreeNode, indexRoot *indexTreeNode, qmin int, qmax int) []int {
	start2 := time.Now()
	var vgMap map[int]string
	vgMap = make(map[int]string)
	VGCons(root, qmin, qmax, searchStr, vgMap)
	fmt.Println(vgMap)
	var resArr []int
	preSeaPosition := 0
	var preInverPositionDis []int
	var nowInverPositionDis []int
	for i := 0; i < len(searchStr); i++ { // 0 1 3   len(searchStr)
		gramArr := vgMap[i]
		if gramArr != "" {
			nowSeaPosition := i
			invertIndex = nil
			invertIndex2 = nil
			searchIndexTreeFromLeaves(gramArr, indexRoot, 0)
			searchListsTreeFromLeaves(indexNode)
			invertIndex = append(invertIndex, invertIndex2...)
			invertIndex = RemoveSliceInvertIndex(invertIndex)
			sort.SliceStable(invertIndex, func(i, j int) bool {
				if invertIndex[i].sid < invertIndex[j].sid {
					return true
				}
				return false
			})
			if invertIndex == nil {
				return nil
			}
			if i == 0 {
				for j := 0; j < len(invertIndex); j++ {
					sid := invertIndex[j].sid
					preInverPositionDis = append(preInverPositionDis, 0)
					nowInverPositionDis = append(nowInverPositionDis, invertIndex[j].position)
					resArr = append(resArr, sid)
				}
			} else {
				for j := 0; j < len(resArr); j++ { //遍历之前合并好的resArr
					var k int
					for k = 0; k < len(invertIndex); k++ {
						if resArr[j] == invertIndex[k].sid {
							nowInverPositionDis[j] = invertIndex[k].position
							if nowInverPositionDis[j]-preInverPositionDis[j] == nowSeaPosition-preSeaPosition {
								break
							}
						}
					}
					if k == len(invertIndex) { //新的倒排表id不在之前合并好的结果集resArr 把此id从resArr删除
						resArr = append(resArr[:j], resArr[j+1:]...)
						preInverPositionDis = append(preInverPositionDis[:j], preInverPositionDis[j+1:]...)
						nowInverPositionDis = append(nowInverPositionDis[:j], nowInverPositionDis[j+1:]...)
						j-- //删除后重新指向，防止丢失元素判断
					}
				}
			}
			//fmt.Println(resArr)
			preSeaPosition = nowSeaPosition
			fmt.Println(preInverPositionDis)
			fmt.Println(nowInverPositionDis)
			copy(preInverPositionDis, nowInverPositionDis)
		}
	}
	elapsed2 := time.Since(start2)
	fmt.Println("精确查询花费时间（ms）：", elapsed2)
	return resArr
}

var invertIndex []inverted_index
var indexNode *indexTreeNode

//查询当前串对应的倒排表（叶子节点）
func searchIndexTreeFromLeaves(gramArr string, indexRoot *indexTreeNode, i int) {
	if indexRoot == nil {
		return
	}
	for j := 0; j < len(indexRoot.children); j++ {
		if i < len(gramArr)-1 && string(gramArr[i]) == indexRoot.children[j].data {
			searchIndexTreeFromLeaves(gramArr, indexRoot.children[j], i+1)
		}
		if i == len(gramArr)-1 && string(gramArr[i]) == indexRoot.children[j].data { //找到那一层的倒排表
			for k := 0; k < len(indexRoot.children[j].invertedIndexList); k++ {
				invertIndex = append(invertIndex, *indexRoot.children[j].invertedIndexList[k])
			}
			indexNode = indexRoot.children[j]
		}
	}
}

var invertIndex2 []inverted_index

func searchListsTreeFromLeaves(indexNode *indexTreeNode) {
	if indexNode != nil {
		for l := 0; l < len(indexNode.children); l++ {
			if indexNode.children[l].invertedIndexList != nil {
				for k := 0; k < len(indexNode.children[l].invertedIndexList); k++ {
					invertIndex2 = append(invertIndex2, *indexNode.children[l].invertedIndexList[k])
				}
			}
			searchListsTreeFromLeaves(indexNode.children[l])
		}
	}
}

func RemoveSliceInvertIndex(invertIndex []inverted_index) (ret []inverted_index) {
	n := len(invertIndex)
	for i := 0; i < n; i++ {
		state := false
		for j := i + 1; j < n; j++ {
			if j > 0 && reflect.DeepEqual(invertIndex[i], invertIndex[j]) {
				state = true
				break
			}
		}
		if !state {
			ret = append(ret, invertIndex[i])
		}
	}
	return
}
