package main

import "fmt"

func main() {
	// ========================================
	// マップの基本
	// ========================================

	// マップの宣言と初期化（リテラル）
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 35,
	}
	fmt.Println("ages:", ages)

	// makeで空マップを作成
	scores := make(map[string]int)
	scores["math"] = 90
	scores["english"] = 85
	fmt.Println("scores:", scores)

	// varで宣言したマップはnil（書き込むとpanicする）
	var nilMap map[string]int
	fmt.Println("nilMap == nil:", nilMap == nil)
	fmt.Println("len(nilMap):", len(nilMap))
	// nilMap["key"] = 1 // panic: assignment to entry in nil map

	// ========================================
	// 値の取得・更新・削除
	// ========================================

	fmt.Println("\n--- 値の操作 ---")
	ages["Alice"] = 31 // 更新
	ages["Dave"] = 28   // 新規追加
	fmt.Println("更新後:", ages)

	delete(ages, "Bob") // 削除
	fmt.Println("Bob削除後:", ages)

	// 存在しないキーを削除してもpanicしない
	delete(ages, "Unknown")

	// 存在しないキーにアクセスするとゼロ値が返る
	fmt.Println("存在しないキー:", ages["Nobody"]) // 0

	// ========================================
	// Comma-Ok Idiom（カンマOKイディオム）
	// ========================================

	fmt.Println("\n--- Comma-Ok Idiom ---")

	// マップから値を取得する際、2番目の戻り値でキーの存在を確認できる
	val, ok := ages["Alice"]
	fmt.Printf("Alice: val=%d, ok=%v\n", val, ok)

	val, ok = ages["Bob"]
	fmt.Printf("Bob: val=%d, ok=%v\n", val, ok) // 削除済みなのでok=false

	// ゼロ値と「キーが存在しない」を区別するために必須
	data := map[string]int{
		"zero":     0,
		"negative": -1,
	}

	// 値だけ見ると0なのでキーがあるのか無いのか分からない
	fmt.Println("\ndata[\"zero\"]:", data["zero"])
	fmt.Println("data[\"missing\"]:", data["missing"]) // どちらも0

	// Comma-Okで区別する
	if v, ok := data["zero"]; ok {
		fmt.Println("\"zero\"キーは存在する、値:", v)
	}
	if _, ok := data["missing"]; !ok {
		fmt.Println("\"missing\"キーは存在しない")
	}

	// ========================================
	// マップのイテレーション
	// ========================================

	fmt.Println("\n--- イテレーション ---")

	// for rangeでキーと値を取得（順序は不定）
	for name, age := range ages {
		fmt.Printf("%s: %d歳\n", name, age)
	}

	// キーだけ取得
	fmt.Print("キー一覧: ")
	for name := range ages {
		fmt.Print(name, " ")
	}
	fmt.Println()

	// ========================================
	// マップは参照型
	// ========================================

	fmt.Println("\n--- 参照型 ---")

	original := map[string]int{"a": 1, "b": 2}
	alias := original    // コピーではなく同じマップを参照
	alias["a"] = 100     // aliasを変更するとoriginalも変わる
	fmt.Println("original:", original)
	fmt.Println("alias:", alias)

	// ========================================
	// マップの活用例：文字の出現回数をカウント
	// ========================================

	fmt.Println("\n--- 活用例: 文字カウント ---")

	word := "hello"
	charCount := make(map[rune]int)
	for _, ch := range word {
		charCount[ch]++ // ゼロ値が0なので初期化不要
	}
	for ch, count := range charCount {
		fmt.Printf("'%c': %d回\n", ch, count)
	}

	// ========================================
	// マップのキーに使える型
	// ========================================

	fmt.Println("\n--- キーに使える型 ---")

	// == で比較可能な型ならキーに使える
	// 構造体もキーにできる
	type Point struct {
		X, Y int
	}
	visited := map[Point]bool{
		{0, 0}: true,
		{1, 2}: true,
	}
	fmt.Println("(1,2)を訪問済み?", visited[Point{1, 2}])
	fmt.Println("(3,4)を訪問済み?", visited[Point{3, 4}]) // false（ゼロ値）

	// スライス・マップ・関数はキーに使えない（==で比較不可）
	// map[[]int]string{} // コンパイルエラー
}
