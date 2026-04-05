package main

import (
	"errors"
	"fmt"
	"strconv"
)

// =====================
// 複数戻り値と名前付き戻り値
// =====================

// 1. 複数の戻り値
// Go では関数が複数の値を返せる（タプルのように使う）
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ゼロで割ることはできません")
	}
	return a / b, nil
}

// 2. 名前付き戻り値（named return values）
// 戻り値に名前を付けると、関数内で変数として使える
func rectangleInfo(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return: 名前付き戻り値がそのまま返される
}

// 3. 複数戻り値のイディオム: value, ok パターン
// マップのアクセスや型アサーションでよく使われるパターン
func findUser(id int) (string, bool) {
	users := map[int]string{
		1: "Alice",
		2: "Bob",
		3: "Charlie",
	}
	name, ok := users[id]
	return name, ok
}

func main() {
	// =====================
	// 1. 複数戻り値（value, error パターン）
	// =====================
	// Go で最も重要なイディオム: エラーを戻り値で返す
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	// エラーが発生するケース
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("エラー:", err) // エラー: ゼロで割ることはできません
	}

	// =====================
	// 2. 名前付き戻り値
	// =====================
	area, perimeter := rectangleInfo(5, 3)
	fmt.Printf("面積: %.1f, 周囲: %.1f\n", area, perimeter) // 面積: 15.0, 周囲: 16.0

	// =====================
	// 3. value, ok パターン
	// =====================
	name, ok := findUser(1)
	if ok {
		fmt.Println("ユーザー発見:", name) // ユーザー発見: Alice
	}

	name, ok = findUser(99)
	if !ok {
		fmt.Println("ユーザーが見つかりません")
	}

	// =====================
	// 4. 不要な戻り値は _ で無視
	// =====================
	_, err = strconv.Atoi("abc") // 変換結果は不要、エラーだけ確認
	if err != nil {
		fmt.Println("変換エラー:", err)
	}

	// =====================
	// 5. 標準ライブラリでの複数戻り値の例
	// =====================
	n, err := fmt.Println("この行は標準出力に書き込まれます")
	fmt.Printf("書き込みバイト数: %d, エラー: %v\n", n, err)
}
