package main

import "fmt"

// TODO: 可変長引数を練習しよう
// 1. 可変長引数で受け取った文字列を ", " で結合して返す関数 joinStrings を作成
func joinStrings(strings ...string) string {
	var result string
	for i, char := range strings {
		if i > 0 {
			result += ", "
		}
		result += char
	}
	return result
}

// 2. 可変長引数の整数から平均値を計算する関数 average を作成
func average(nums ...int) float64 {
	var result float64
	for _, num := range nums {
		result += float64(num)
	}

	return result / float64(len(nums))
}

// 3. 最初の引数をセパレータとし、残りの文字列を結合する関数 joinWith を作成
//    例: joinWith("-", "a", "b", "c") → "a-b-c"
func joinWith(strings ...string) string {
	var result string
	separator := strings[0]
	for i, char := range strings {
		if i == 0 {
			continue
		}
		if i == len(strings) - 1 {
			result += char
		} else {
			result += char + separator
		}
	}
	return result
}

func main() {
	joinString := joinStrings("a", "b", "c")
	fmt.Println(joinString)

	average := average(1, 2, 3, 4, 5)
	fmt.Println(average)

	joinWith := joinWith("-", "a", "b", "c")
	fmt.Println(joinWith)
}
