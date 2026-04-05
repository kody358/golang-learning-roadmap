package main

import "fmt"

func main() {
	// =====================
	// 1. 基本的な for ループ（C言語スタイル）
	// =====================
	// Go のループは for のみ（while や do-while はない）
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// =====================
	// 2. while スタイルの for
	// =====================
	// 条件式だけを書くと while のように動作する
	count := 0
	for count < 3 {
		fmt.Printf("count = %d\n", count)
		count++
	}

	// =====================
	// 3. 無限ループ
	// =====================
	// 条件を省略すると無限ループになる
	sum := 0
	for {
		sum++
		if sum >= 5 {
			break // break で抜ける
		}
	}
	fmt.Println("sum =", sum)

	// =====================
	// 4. continue
	// =====================
	// continue で残りの処理をスキップし、次の反復に進む
	fmt.Print("奇数: ")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 偶数はスキップ
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// =====================
	// 5. for-range（スライス）
	// =====================
	// range でスライスのインデックスと値を取得
	fruits := []string{"りんご", "みかん", "ぶどう"}
	for i, fruit := range fruits {
		fmt.Printf("[%d] %s\n", i, fruit)
	}

	// インデックスが不要な場合は _ で捨てる
	for _, fruit := range fruits {
		fmt.Println("果物:", fruit)
	}

	// 値が不要な場合はインデックスだけ受け取る
	for i := range fruits {
		fmt.Println("インデックス:", i)
	}

	// =====================
	// 6. for-range（マップ）
	// =====================
	colors := map[string]string{
		"red":   "赤",
		"blue":  "青",
		"green": "緑",
	}
	for key, value := range colors {
		fmt.Printf("%s = %s\n", key, value)
	}
	// 注意: マップの反復順序は保証されない

	// =====================
	// 7. for-range（文字列）
	// =====================
	// 文字列の range は rune（Unicodeコードポイント）単位で反復する
	for i, r := range "Go言語" {
		fmt.Printf("バイト位置 %d: %c (U+%04X)\n", i, r, r)
	}

	// =====================
	// 8. for-range（チャネル）
	// =====================
	// チャネルが閉じられるまで値を受信し続ける
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)

	for v := range ch {
		fmt.Println("受信:", v)
	}

	// =====================
	// 9. ラベル付き break / continue
	// =====================
	// ネストしたループで外側のループを制御したい場合に使う
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("(1,1) で外側ループを break")
				break outer // 外側の for を抜ける
			}
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}

	// =====================
	// 10. for-range と整数（Go 1.22+）
	// =====================
	// Go 1.22 以降、整数を range で直接使える
	fmt.Print("0〜4: ")
	for i := range 5 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
