package main

import "fmt"

// =====================
// 可変長引数（Variadic Functions）
// =====================

// 1. 可変長引数の基本
// 最後の引数に ... を付けると、0個以上の値を受け取れる
// 関数内では nums はスライス（[]int）として扱われる
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 2. 通常の引数と組み合わせる
// 可変長引数は最後の引数でなければならない
func printWithPrefix(prefix string, messages ...string) {
	for _, msg := range messages {
		fmt.Printf("[%s] %s\n", prefix, msg)
	}
}

// 3. 可変長引数で最小1つの引数を要求するパターン
func max(first int, rest ...int) int {
	m := first
	for _, v := range rest {
		if v > m {
			m = v
		}
	}
	return m
}

func main() {
	// =====================
	// 1. 可変長引数の呼び出し
	// =====================
	fmt.Println("sum()     =", sum())           // 0
	fmt.Println("sum(1)    =", sum(1))           // 1
	fmt.Println("sum(1,2,3)=", sum(1, 2, 3))     // 6
	fmt.Println("sum(1..5) =", sum(1, 2, 3, 4, 5)) // 15

	// =====================
	// 2. スライスを可変長引数に展開
	// =====================
	// スライスの後に ... を付けると展開される
	numbers := []int{10, 20, 30}
	fmt.Println("sum(slice...) =", sum(numbers...)) // 60

	// =====================
	// 3. 通常の引数との組み合わせ
	// =====================
	printWithPrefix("INFO", "サーバー起動", "ポート8080でリスン中")
	printWithPrefix("ERROR", "接続失敗")

	// =====================
	// 4. 最低1つの引数を要求
	// =====================
	fmt.Println("max(3,1,4,1,5) =", max(3, 1, 4, 1, 5)) // 5
	fmt.Println("max(42)        =", max(42))               // 42

	// =====================
	// 5. 標準ライブラリでの可変長引数
	// =====================
	// fmt.Println は可変長引数: func Println(a ...any) (n int, err error)
	fmt.Println("複数の", "値を", "渡せる")

	// fmt.Sprintf も可変長引数
	s := fmt.Sprintf("%s は %d 歳で %s に住んでいます", "太郎", 30, "東京")
	fmt.Println(s)

	// append も可変長引数: func append(slice []T, elems ...T) []T
	a := []int{1, 2}
	b := []int{3, 4, 5}
	a = append(a, b...) // スライスの展開
	fmt.Println("append結果:", a) // [1 2 3 4 5]
}
