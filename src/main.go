package main

import (
	"dictionary"
	"fmt"
	"index07"
	_ "matchQuery1"
	"runtime"
)

func TraceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func main() {
	fmt.Println("字典树D：===============================================================")
	fmt.Println("字典树D内存占用大小：")
	//TraceMemStats()
	fmt.Println()
	root := dictionary.GenerateDictionaryTree("src/resources/50000Dic.txt", 2, 12, 200) //
	fmt.Println()
	//TraceMemStats()
	fmt.Println()

	fmt.Println("索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	TraceMemStats()
	fmt.Println()
	index07.GenerateIndexTree("src/resources/500Index.txt", 2, 12, root) //_, indexTreeNode :=
	fmt.Println()
	TraceMemStats()
	fmt.Println()

	/*indexTreeNode.FixInvertedIndexSize()
	sort.SliceStable(index07.Res, func(i, j int) bool {
		if index07.Res[i] < index07.Res[j]  {
			return true
		}
		return false
	})
	fmt.Println(index07.Res)
	fmt.Println(len(index07.Res))
	sum := 0
	for _,val := range index07.Res{
		sum += val
	}
	fmt.Println(index07.Res[0])
	fmt.Println(index07.Res[len(index07.Res)-1])
	fmt.Println(index07.Res[len(index07.Res)/2])
	fmt.Println(sum/len(index07.Res))*/

	/*indexTreeNode.SearchGramsFromIndexTree()
	//fmt.Println(index07.Grams)
	fmt.Println(len(index07.Grams))
	var numsOfgrams2_12 [13]int
	for _,val := range index07.Grams{
		numsOfgrams2_12[len(val)]++
	}
	fmt.Println(numsOfgrams2_12)*/

	/*fmt.Println("新增索引后的索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	//TraceMemStats()
	fmt.Println()
	indexMaintain.AddIndex("src/resources/add2000.txt", 2, 6, root, indexTree)
	fmt.Println()
	//TraceMemStats()
	fmt.Println()*/

	//resInt := matchQuery2.MatchSearch("GET /english", root, indexTreeNode, 2, 12)
	/*var searchQuery = [9]string{"GET","GET /english","GET /english/images/","GET /images/","GET /english/images/team_hm_header_shad.gif HTTP/1.0","GET /images/s102325.gif HTTP/1.0","GET /english/history/history_of/images/cup/","/images/space.gif","GET / HTTP/1.0"}
	for i := 0; i < 9; i++ {
		resInt := matchQuery2.MatchSearch(searchQuery[i], root, indexTreeNode, 2, 12) //get english venues
		//fmt.Println(resInt)
		fmt.Println(len(resInt))
		fmt.Println("==================================================")
	}*/
	//fmt.Println(len(resInt))
}
