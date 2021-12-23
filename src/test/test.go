package main

import (
	"../util"
	"fmt"
	"strconv"
)

func main() {
	//读取不同批量文件数据
	fileName := "src/resources/dics/500Dic.txt"
	qMax := 2
	dicTree := util.BuildDicTree(fileName, 1, qMax)
	// 统计频率
	frequencies := util.GetFrequency(dicTree)

	// 得到T
	T := util.CountDicSize(frequencies)

	// 输出修剪前的树
	//fmt.Println("输出修剪前的树：")
	//util.PrintDicTree(dicTree)

	// 测试当前树所占用的空间
	util.TraceMemStats()
	// 修剪trie tree
	util.PureDicTree(dicTree, T)
	// 测试当前字典占用的空间
	util.TraceMemStats()

	//输出修剪后的树
	//fmt.Println("输出修剪后的树：")
	//util.PrintDicTree(dicTree)

	// 重新统计频率
	frequencies2 := util.GetFrequency(dicTree)

	fmt.Println()
	fmt.Println(*frequencies)
	fmt.Println(util.GetDicSum(frequencies))
	fmt.Println(*frequencies2)
	fmt.Println(util.GetDicSum(frequencies2))

	// 写文件测试
	fileName1 := "src/resources/dics/output/" + "500Dic-qMax" + strconv.Itoa(qMax) + ".txt"
	fileName2 := "src/resources/dics/output/" + "500Dic-qMax" + strconv.Itoa(qMax) + "-T" + strconv.Itoa(T) + ".txt"
	util.WriteFrequency(fileName1, frequencies)
	util.WriteFrequency(fileName2, frequencies2)

}
