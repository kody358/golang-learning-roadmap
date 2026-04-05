package main

import "fmt"

// TODO: 関数の基本を練習しよう
// 1. 2つの整数を受け取り、大きい方を返す関数 max を作成
// 2. 文字列と回数を受け取り、繰り返した文字列を返す関数 repeat を作成
// 3. 摂氏温度を華氏に変換する関数 celsiusToFahrenheit を作成（F = C * 9/5 + 32）

// 関数の定義
func greet(name string) string {
	return "こんにちは" + name + "さん"
}

func main() {
	msg := greet("田中")
	fmt.Println(msg)
}
