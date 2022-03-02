package leetcode

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type SpaArray struct {
	row int
	col int
	val int
}

func SparseArray() {
	var original_array [11][11]int
	original_array[0][5] = 1
	original_array[2][7] = 2
	fmt.Println("====原始数组====")
	for _, v := range original_array {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	fmt.Println("====稀疏数组====")
	spa_Array := []SpaArray{}
	val_node := SpaArray{
		row: 11,
		col: 11,
		val: 0,
	}
	spa_Array = append(spa_Array, val_node)
	for i, v := range original_array {
		for j, v2 := range v {
			if v2 != 0 {
				val_node := SpaArray{
					row: i,
					col: j,
					val: v2,
				}
				spa_Array = append(spa_Array, val_node)
			}
		}
	}
	fmt.Println("===稀疏数组写入文件===")
	var data string
	for _, v := range spa_Array {
		fmt.Printf("%d %d %d\t\n", v.row, v.col, v.val)
		data = data + " " + strconv.Itoa(v.row) + " " + strconv.Itoa(v.col) + " " + strconv.Itoa(v.val)

		// data = data + " " + strconv.Itoa(v.row) + " " + strconv.Itoa(v.col) + " " + strconv.Itoa(v.val) + "\n"
	}
	fmt.Println(data)

	fmt.Println("===稀疏矩阵存入文件===")
	err := ioutil.WriteFile("D:/MyCodeGroup/Go/src/leetcode/datas.data", []byte(data), 0666)
	if err != nil {
		fmt.Println("写入错误")
		return
	}
	fmt.Println("===从文件读取并表示为二维数组===")
	databack, err := ioutil.ReadFile("D:/MyCodeGroup/Go/src/leetcode/datas.data")
	if err != nil {
		fmt.Println("读取错误")
		return
	}
	/*1.databack []byte
	2.string(databack)  []byte =>string
	3.databack_list_string string =>[]string
	4.databack_list_int []string => []int
	*/
	fmt.Println("databack:", databack)
	fmt.Println("len(string(databack)),string(databack)", len(string(databack)), string(databack))
	databack_list_string := strings.Fields(string(databack))
	fmt.Println("databack_list_string", databack_list_string)
	databack_list_int := String2Int(databack_list_string)
	var Back_Array [11][11]int
	Back_Array[databack_list_int[3]][databack_list_int[4]] = databack_list_int[5]
	Back_Array[databack_list_int[6]][databack_list_int[7]] = databack_list_int[8]
	for _, v := range Back_Array {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}
