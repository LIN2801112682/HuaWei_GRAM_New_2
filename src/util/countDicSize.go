package util

import (
	"../go_dic"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"time"
)

// 建立字典树
func BuildDicTree(fileName string, qMin int, qMax int) *go_dic.TrieTree {
	tree := go_dic.NewTrieTree(qMin, qMax)
	data, err := os.Open(fileName)
	defer data.Close()
	if err != nil {
		fmt.Println(err)
	}
	buff := bufio.NewReader(data)
	// 统计花费时间变量
	var sum = 0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		str := (string)(data)
		start2 := time.Now()
		for i := 0; i < len(str)-qMax; i++ {
			substring := str[i : i+qMax]
			//字符串变字符串数组
			gram := make([]string, qMax)
			for j := 0; j < qMax; j++ {
				gram[j] = substring[j : j+1]
			}
			go_dic.InsertIntoTrieTree(tree, &gram)
		}
		for i := len(str) - qMax; i < len(str)-qMin+1; i++ {
			substring := str[i:len(str)]
			gram := make([]string, len(str)-i)
			for j := 0; j < len(str)-i; j++ {
				gram[j] = substring[j : j+1]
			}
			go_dic.InsertIntoTrieTree(tree, &gram)
		}
		end2 := time.Since(start2).Microseconds()
		sum = int(end2) + sum
	}
	return tree

}

// 对树进行修剪
func PureDicTree(dicTree *go_dic.TrieTree, T int) {
	// 先统计出 不同长度的频率再来求 T
	fmt.Printf("\n根据 T = %d 进行剪枝...\n", T)
	go_dic.PruneTree(dicTree, T)
}

// 处理频率 得到 T
func CountDicSize(dicSize *map[int][]int) int {
	// 保存 所有 频率
	var frequencies []int
	// 计算 T
	for _, v := range *dicSize {
		frequencies = append(frequencies, v...)
	}
	//fmt.Println("排序前：",frequencies)
	// 对所有频率排序
	sort.Ints(frequencies)
	//fmt.Println("排序后：",frequencies)
	size := len(frequencies)

	// 取中位数
	T := frequencies[size/2]
	fmt.Println("频率中位数：", T)

	// 求和取平均数
	sum, err := ArraySum(frequencies)
	if err != nil {
		fmt.Println("求和出现错误：", err)
	}
	Temp := sum / float64(size)
	T2 := int(Temp)
	fmt.Println("平均数：", T2)
	// 得到 T
	fmt.Println("返回：", T)
	return T

}

// 数组切片求和
func ArraySum(input interface{}) (sum float64, err error) {
	list := reflect.ValueOf(input)
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < list.Len(); i++ {
			val := list.Index(i)
			v, err := toFloat64(val.Interface())
			if err != nil {
				return 0, err
			}
			sum += v
		}
	default:
		return 0, errors.New("input must be slice or array")
	}
	return
}

func toFloat64(in interface{}) (f64 float64, err error) {
	switch val := in.(type) {
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case uint:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(val), nil
		}
		return 0, errors.New("convert uint to float64 failed")
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(val), nil
		}
		return 0, errors.New("convert int to float64 failed")
	case bool:
		if val {
			f64 = 1
		}
		return
	case string:
		f64, err = strconv.ParseFloat(val, 64)
		if err == nil {
			return
		}
		return 0, errors.New("convert string to float64 failed")
	default:
		return 0, errors.New("convert to float64 failed")
	}
}

/*
统计不同 gram长度出现的频率 返回T
*/
func GetFrequency(dicTree *go_dic.TrieTree) *map[int][]int {
	dicSize := make(map[int][]int)
	PrintDicTreeNode(go_dic.GetRoot(dicTree), 0, &dicSize)
	return &dicSize
}

func PrintDicTreeNode(node *go_dic.TrieTreeNode, level int, dicSize *map[int][]int) {
	if go_dic.IsLeaf(node) {
		value, ok := (*dicSize)[level]
		if !ok {
			value = make([]int, 0)
		}
		value = append(value, go_dic.GetFrequency(node))
		(*dicSize)[level] = value
	}
	for _, child := range go_dic.GetChildren(node) {
		PrintDicTreeNode(child, level+1, dicSize)
	}

}

func PrintDicTree(dicTree *go_dic.TrieTree) {
	go_dic.PrintTree(dicTree)
}

func WriteFrequency(fileName string, dicSize *map[int][]int) {
	var f *os.File
	var err1 error
	if checkFileIsExist(fileName) { //如果文件存在
		f, err1 = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(fileName) //创建文件
		fmt.Println("文件不存在,创建！", fileName)
	}
	check(err1)
	defer f.Close()
	for k, v := range *dicSize {
		_, err1 := io.WriteString(f, "gram_"+strconv.Itoa(k)+"#"+strconv.Itoa(len(v))+":")
		check(err1)
		for _, value := range v {
			_, err1 := io.WriteString(f, strconv.Itoa(value)+",")
			check(err1)
		}
		_, err1 = io.WriteString(f, "\n")
		check(err1)
	}
	fmt.Printf("写入文件 %s 成功！", fileName)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 统计内存
func TraceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}
