package main

import "fmt"

func main() {
	// =====================
	// 1. 基本的な if 文
	// =====================
	x := 10
	if x > 5 {
		fmt.Println("x は 5 より大きい")
	}

	// =====================
	// 2. if-else
	// =====================
	age := 17
	if age >= 18 {
		fmt.Println("成人です")
	} else {
		fmt.Println("未成年です")
	}

	// =====================
	// 3. if-else if-else チェーン
	// =====================
	score := 75
	if score >= 90 {
		fmt.Println("評価: A")
	} else if score >= 80 {
		fmt.Println("評価: B")
	} else if score >= 70 {
		fmt.Println("評価: C")
	} else {
		fmt.Println("評価: D")
	}

	// =====================
	// 4. if の初期化文（short statement）
	// =====================
	// if の条件の前にセミコロンで区切って変数を宣言できる
	// この変数は if-else ブロック内でのみ有効
	if n := 42; n%2 == 0 {
		fmt.Printf("%d は偶数\n", n)
	} else {
		fmt.Printf("%d は奇数\n", n)
	}
	// fmt.Println(n) // コンパイルエラー: n は if ブロック外からアクセスできない

	// 関数の戻り値チェックでよく使うパターン
	if err := doSomething(); err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("成功")
	}

	// =====================
	// 5. 基本的な switch 文
	// =====================
	day := "水曜日"
	switch day {
	case "月曜日":
		fmt.Println("週の始まり")
	case "水曜日":
		fmt.Println("週の半ば")
	case "金曜日":
		fmt.Println("もうすぐ週末")
	case "土曜日", "日曜日": // 複数の値をカンマで列挙できる
		fmt.Println("週末！")
	default:
		fmt.Println("普通の日")
	}
	// Go の switch は自動で break される（C言語と異なり fall through しない）

	// =====================
	// 6. switch の初期化文
	// =====================
	switch os := getOS(); os {
	case "linux":
		fmt.Println("Linux です")
	case "darwin":
		fmt.Println("macOS です")
	case "windows":
		fmt.Println("Windows です")
	default:
		fmt.Printf("その他の OS: %s\n", os)
	}

	// =====================
	// 7. 式なし switch（条件分岐の代替）
	// =====================
	// switch の後に式を書かないと、各 case で bool 式を評価する
	// 長い if-else if チェーンの代わりに使える
	temp := 35
	switch {
	case temp >= 35:
		fmt.Println("猛暑日")
	case temp >= 30:
		fmt.Println("真夏日")
	case temp >= 25:
		fmt.Println("夏日")
	default:
		fmt.Println("過ごしやすい気温")
	}

	// =====================
	// 8. fallthrough
	// =====================
	// 明示的に次の case に処理を続けたい場合は fallthrough を使う
	num := 1
	switch num {
	case 1:
		fmt.Println("1 にマッチ")
		fallthrough // 次の case の処理も実行される
	case 2:
		fmt.Println("2 の処理も実行")
		// fallthrough がないのでここで止まる
	case 3:
		fmt.Println("3 の処理は実行されない")
	}

	// =====================
	// 9. 型 switch（type switch）
	// =====================
	// interface{} の実際の型に基づいて分岐する
	checkType(42)
	checkType("hello")
	checkType(3.14)
	checkType(true)
}

// doSomething はエラーを返す関数の例
func doSomething() error {
	return nil
}

// getOS は OS 名を返す（デモ用）
func getOS() string {
	return "darwin"
}

// checkType は型 switch のデモ
func checkType(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("int 型: %d\n", t)
	case string:
		fmt.Printf("string 型: %s\n", t)
	case float64:
		fmt.Printf("float64 型: %f\n", t)
	default:
		fmt.Printf("未知の型: %T\n", t)
	}
}
