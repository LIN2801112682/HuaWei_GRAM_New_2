package main

import (
	"build_VGram_index"
	"build_dictionary"
	"fmt"
	"new_precise_query"
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
	TraceMemStats()
	fmt.Println()
	root := build_dictionary.GererateTree("src/resources/5000Dic.txt", 2, 12, 40) //
	fmt.Println()
	TraceMemStats()
	fmt.Println()

	fmt.Println("索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	TraceMemStats()
	fmt.Println()
	_, indexTreeNode := build_VGram_index.GererateIndex("src/resources/index5000.txt", 2, 12, root) //
	fmt.Println()
	TraceMemStats()
	fmt.Println()

	/*fmt.Println("新增索引后的索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	//TraceMemStats()
	fmt.Println()
	index_maintenance.AddIndex("src/resources/add2000.txt", 2, 6, root, indexTree)
	fmt.Println()
	//TraceMemStats()
	fmt.Println()*/

	//resInt := precise_query.MatchSearch(" HTTP/1.1", root, indexTreeNode, 2, 10) //get english venues
	resInt := new_precise_query.MatchSearch("GET / HTTP/1.0", root, indexTreeNode, 2, 12)

	fmt.Println(resInt)
	fmt.Println(len(resInt))
}
