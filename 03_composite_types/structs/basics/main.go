package main

import "fmt"

// ===== 構造体の基本 =====

// 構造体の定義
type Person struct {
	Name string
	Age  int
	City string
}

// コンストラクタ関数（Goにはコンストラクタ構文がないので関数で代用）
func NewPerson(name string, age int, city string) Person {
	return Person{
		Name: name,
		Age:  age,
		City: city,
	}
}

// 座標を表す構造体
type Point struct {
	X, Y float64 // 同じ型のフィールドはまとめて宣言できる
}

// ネストされた構造体
type Address struct {
	PostalCode string
	Prefecture string
	City       string
}

type Employee struct {
	Name    string
	Age     int
	Address Address // 構造体をフィールドに持つ
}

func main() {
	// --- 構造体の初期化方法 ---
	fmt.Println("=== 構造体の初期化 ===")

	// 1. フィールド名を指定して初期化（推奨）
	p1 := Person{
		Name: "田中太郎",
		Age:  30,
		City: "東京",
	}
	fmt.Printf("p1: %+v\n", p1)

	// 2. フィールド順で初期化（フィールドが変わると壊れるので非推奨）
	p2 := Person{"佐藤花子", 25, "大阪"}
	fmt.Printf("p2: %+v\n", p2)

	// 3. ゼロ値で初期化（全フィールドがゼロ値）
	var p3 Person
	fmt.Printf("p3 (ゼロ値): %+v\n", p3)

	// 4. 一部のフィールドだけ指定（残りはゼロ値）
	p4 := Person{Name: "鈴木一郎"}
	fmt.Printf("p4 (部分指定): %+v\n", p4)

	// 5. コンストラクタ関数を使う
	p5 := NewPerson("高橋次郎", 35, "名古屋")
	fmt.Printf("p5 (コンストラクタ): %+v\n", p5)

	// --- フィールドへのアクセス ---
	fmt.Println("\n=== フィールドアクセス ===")
	fmt.Printf("名前: %s, 年齢: %d\n", p1.Name, p1.Age)

	// フィールドの更新
	p1.Age = 31
	fmt.Printf("誕生日後: %+v\n", p1)

	// --- ポインタと構造体 ---
	fmt.Println("\n=== ポインタと構造体 ===")

	// new()で構造体のポインタを生成
	pp1 := new(Person)
	pp1.Name = "山田三郎" // (*pp1).Name と書かなくてよい（自動参照解決）
	pp1.Age = 40
	fmt.Printf("pp1: %+v\n", *pp1)

	// &で構造体リテラルのポインタを取得
	pp2 := &Person{
		Name: "中村四郎",
		Age:  45,
		City: "福岡",
	}
	fmt.Printf("pp2: %+v\n", *pp2)

	// ポインタでもドットアクセスが使える
	fmt.Printf("pp2の名前: %s\n", pp2.Name)

	// --- 構造体の比較 ---
	fmt.Println("\n=== 構造体の比較 ===")
	a := Point{X: 1.0, Y: 2.0}
	b := Point{X: 1.0, Y: 2.0}
	c := Point{X: 3.0, Y: 4.0}
	fmt.Printf("a == b: %v\n", a == b) // true（全フィールドが等しい）
	fmt.Printf("a == c: %v\n", a == c) // false

	// --- ネストされた構造体 ---
	fmt.Println("\n=== ネストされた構造体 ===")
	emp := Employee{
		Name: "伊藤五郎",
		Age:  28,
		Address: Address{
			PostalCode: "100-0001",
			Prefecture: "東京都",
			City:       "千代田区",
		},
	}
	fmt.Printf("社員: %+v\n", emp)
	fmt.Printf("住所: %s %s\n", emp.Address.Prefecture, emp.Address.City)

	// --- 匿名構造体 ---
	fmt.Println("\n=== 匿名構造体 ===")

	// 一時的なデータ構造に便利（型定義不要）
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("config: %+v\n", config)

	// テストのテーブル駆動テストでもよく使われるパターン
	tests := []struct {
		input    int
		expected int
	}{
		{1, 2},
		{2, 4},
		{3, 6},
	}
	for _, tt := range tests {
		fmt.Printf("  input=%d, expected=%d\n", tt.input, tt.expected)
	}
}
