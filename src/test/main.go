package main

import (
	"../go_dic"
	"fmt"
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
	root := go_dic.GererateTree("src/resources/500Dic.txt", 2, 6, 50) //
	fmt.Println()
	//TraceMemStats()
	fmt.Println()

	fmt.Println("索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	TraceMemStats()
	fmt.Println()
	go_dic.GererateIndex("src/resources/5000Index.txt", 2, 6, root) //indexTree :=
	fmt.Println()
	TraceMemStats()
	fmt.Println()

	/*fmt.Println("新增索引后的索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	//TraceMemStats()
	fmt.Println()
	go_dic.AddIndex("src/resources/add2000.txt", 2, 6, root, indexTree)
	fmt.Println()
	//TraceMemStats()
	fmt.Println()*/
}
