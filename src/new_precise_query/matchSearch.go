package new_precise_query

import (
	"build_VGram_index"
	"build_dictionary"
	"fmt"
	"reflect"
	"sort"
	"time"
)

func MatchSearch(searchStr string, root *build_dictionary.TrieTreeNode, indexRoot *build_VGram_index.IndexTreeNode, qmin int, qmax int) []build_VGram_index.SeriesId {
	var vgMap map[int]string
	vgMap = make(map[int]string)
	build_VGram_index.VGCons(root, qmin, qmax, searchStr, vgMap)
	fmt.Println(vgMap)
	var keys = []int{}
	for key := range vgMap {
		keys = append(keys, key)
	}
	//fmt.Println(keys)
	//对map中的key进行排序（map遍历是无序的）
	sort.Sort(sort.IntSlice(keys))
	//fmt.Println(keys)
	var resArr []build_VGram_index.SeriesId
	preSeaPosition := 0
	var preInverPositionDis []PosList
	var nowInverPositionDis []PosList
	start2 := time.Now()
	for m := 0; m < len(keys); m++ {
		i := keys[m]
		gramArr := vgMap[i]
		if gramArr != "" {
			nowSeaPosition := i
			invertIndex = nil
			invertIndex2 = nil
			searchIndexTreeFromLeaves(gramArr, indexRoot, 0)
			searchListsTreeFromLeaves(indexNode)
			invertIndex = append(invertIndex, invertIndex2...)
			invertIndex = RemoveSliceInvertIndex(invertIndex) //去重
			/*invertIndex = MergeInvertIndex(invertIndex) //合并
			fmt.Println(invertIndex)*/
			/*sort.SliceStable(invertIndex, func(i, j int) bool { //排序
				if invertIndex[i].Sid.Id < invertIndex[j].Sid.Id && invertIndex[i].Sid.Time < invertIndex[j].Sid.Time {
					return true
				}
				return false
			})*/
			if invertIndex == nil {
				return nil
			}
			if i == 0 {
				for j := 0; j < len(invertIndex); j++ {
					sid := invertIndex[j].Sid
					preInverPositionDis = append(preInverPositionDis, NewPosList(j, make([]int, len(invertIndex[j].PosArray), len(invertIndex[j].PosArray))))
					nowInverPositionDis = append(nowInverPositionDis, NewPosList(j, invertIndex[j].PosArray))
					resArr = append(resArr, sid)
				}
			} else {
				pos := 0
				for j := 0; j < len(resArr); j++ { //遍历之前合并好的resArr
					var k int
					findFlag := false
					for k = pos; k < len(invertIndex); k++ {
						if resArr[j] == invertIndex[k].Sid { //
							nowInverPositionDis[j] = NewPosList(j, invertIndex[k].PosArray)
							//需要对posList查询从而进一步判断
							for z1 := 0; z1 < len(preInverPositionDis[j].PosArray); z1++ {
								z1Pos := preInverPositionDis[j].PosArray[z1]
								for z2 := 0; z2 < len(nowInverPositionDis[j].PosArray); z2++ {
									z2Pos := nowInverPositionDis[j].PosArray[z2]
									if nowSeaPosition-preSeaPosition == z2Pos-z1Pos {
										findFlag = true
										break
									}
									if findFlag { //后续求解一行日志 那些地方有get 时需更改 这时记录每个出现get的位置
										break
									}
								}
							}
							if findFlag {
								pos = k
								break
							}
						}
						if resArr[j].Id < invertIndex[k].Sid.Id && resArr[j].Time < invertIndex[k].Sid.Time {
							pos = k
							break
						}
					}
					if findFlag == false { //没找到并且候选集的sid比resArr大，删除resArr[j]
						resArr = append(resArr[:j], resArr[j+1:]...)
						preInverPositionDis = append(preInverPositionDis[:j], preInverPositionDis[j+1:]...)
						nowInverPositionDis = append(nowInverPositionDis[:j], nowInverPositionDis[j+1:]...)
						j-- //删除后重新指向，防止丢失元素判断
					}
				}
			}
			preSeaPosition = nowSeaPosition
			copy(preInverPositionDis, nowInverPositionDis)
			//preInverPositionDis = nowInverPositionDis
		}
	}

	elapsed2 := time.Since(start2).Microseconds()
	fmt.Println("精确查询花费时间（us）：", elapsed2)
	return resArr
}

var invertIndex []build_VGram_index.Inverted_index
var indexNode *build_VGram_index.IndexTreeNode

//查询当前串对应的倒排表（叶子节点）
func searchIndexTreeFromLeaves(gramArr string, indexRoot *build_VGram_index.IndexTreeNode, i int) {
	if indexRoot == nil {
		return
	}
	for j := 0; j < len(indexRoot.Children); j++ {
		if i < len(gramArr)-1 && string(gramArr[i]) == indexRoot.Children[j].Data {
			searchIndexTreeFromLeaves(gramArr, indexRoot.Children[j], i+1)
		}
		if i == len(gramArr)-1 && string(gramArr[i]) == indexRoot.Children[j].Data { //找到那一层的倒排表
			for k := 0; k < len(indexRoot.Children[j].InvertedIndexList); k++ {
				invertIndex = append(invertIndex, *indexRoot.Children[j].InvertedIndexList[k])
			}
			indexNode = indexRoot.Children[j]
		}
	}
}

var invertIndex2 []build_VGram_index.Inverted_index

func searchListsTreeFromLeaves(indexNode *build_VGram_index.IndexTreeNode) {
	if indexNode != nil {
		for l := 0; l < len(indexNode.Children); l++ {
			if indexNode.Children[l].InvertedIndexList != nil {
				for k := 0; k < len(indexNode.Children[l].InvertedIndexList); k++ {
					invertIndex2 = append(invertIndex2, *indexNode.Children[l].InvertedIndexList[k])
				}
			}
			searchListsTreeFromLeaves(indexNode.Children[l])
		}
	}
}

//去重：1.Sid和posArray完全相同去重 2.Sid相同，合并posArray
func RemoveSliceInvertIndex(invertIndex []build_VGram_index.Inverted_index) (ret []build_VGram_index.Inverted_index) {
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

func MergeInvertIndex(invertIndex []build_VGram_index.Inverted_index) (ret []build_VGram_index.Inverted_index) {
	for i := 0; i < len(invertIndex); i++ {
		for j := i + 1; j < len(invertIndex); j++ {
			if invertIndex[i].Sid.Id == invertIndex[j].Sid.Id && invertIndex[i].Sid.Time == invertIndex[j].Sid.Time {
				invertIndex[i].PosArray = append(invertIndex[i].PosArray, invertIndex[j].PosArray...)
			}
		}
	}
	return
}
