package main

import "fmt"

// =====================
// 関数の基本
// =====================

// 1. 基本的な関数定義
// func 関数名(引数 型) 戻り値の型 { ... }
func greet(name string) string {
	return "こんにちは、" + name + "さん！"
}

// 2. 引数なし・戻り値なしの関数
func sayHello() {
	fmt.Println("Hello!")
}

// 3. 複数の引数（同じ型はまとめて書ける）
func add(a, b int) int {
	return a + b
}

// 4. 異なる型の引数
func describe(name string, age int) string {
	return fmt.Sprintf("%sは%d歳です", name, age)
}

func main() {
	// =====================
	// 1. 基本的な関数呼び出し
	// =====================
	msg := greet("太郎")
	fmt.Println(msg) // こんにちは、太郎さん！

	// =====================
	// 2. 戻り値なしの関数
	// =====================
	sayHello() // Hello!

	// =====================
	// 3. 引数が複数の関数
	// =====================
	result := add(3, 5)
	fmt.Println("3 + 5 =", result) // 3 + 5 = 8

	// =====================
	// 4. 異なる型の引数
	// =====================
	fmt.Println(describe("花子", 25)) // 花子は25歳です

	// =====================
	// 5. 関数は値（第一級オブジェクト）
	// =====================
	// Go の関数は変数に代入できる
	f := add
	fmt.Println("f(10, 20) =", f(10, 20)) // f(10, 20) = 30

	// 関数の型を確認
	fmt.Printf("fの型: %T\n", f) // fの型: func(int, int) int
}
