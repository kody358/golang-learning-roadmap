package main

import (
	"fmt"
	"strconv"
	"errors"
	"math"
)

// TODO: 複数戻り値を練習しよう
// 1. 整数スライスを受け取り、最小値と最大値を返す関数 minMax を作成
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}
// 2. 文字列を整数に変換し、(変換結果, エラー) を返す関数 parseAge を作成
//    - 負の値の場合もエラーを返す
func parseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("変換できない文字列が含まれています")
	}
	
	if n < 0 {
		return 0, errors.New("負の数です")
	}

	return n, nil
}

// 3. 名前付き戻り値を使って、円の面積と円周を返す関数 circleInfo を作成
func circleInfo(r float64) (area, circumference float64) {
	area = math.Pi * r * r
	circumference = 2 * math.Pi * r

	return
}



func main() {
	nums := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	min, max := minMax(nums) 
	fmt.Printf("min = %d, max = %d\n", min, max)
	
	num, err := parseAge("32")
	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err)
	}
	num, err = parseAge("5a")
	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err)
	}

	area, circumference := circleInfo(4)
	fmt.Printf("area = %f, circumference = %f\n", area, circumference)
}
