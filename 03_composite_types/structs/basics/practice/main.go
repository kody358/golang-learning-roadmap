package main

import "fmt"

// 構造体の定義
	type Person struct {
		Name string
		Age int
		City string
	}
	
	// コンストラクタ関数
	func NewPerson(name string, age int, city string) Person {
		return Person {
			Name: name,
			Age: age,
			City: city,
		}
	}
	
func main() {
	p1 := Person {
		Name: "田中太郎",
		Age: 30,
		City: "東京",
	}
	fmt.Printf("p1: %+v\n", p1)

	// ゼロ値で初期化
	var p2 Person
	fmt.Printf("p2: %+v\n", p2)

	// コンストラクタ関数で初期化
	p3 := NewPerson("山田次郎", 24, "大阪")
	fmt.Printf("p3: %+v\n", p3)

	// フィールドアクセス
	fmt.Printf("名前: %s, 年齢: %d\n", p1.Name, p1.Age)

	// 構造体のポインタ型作成
	pp1 := new(Person)
	pp1.Name = "佐藤三郎"
	pp1.Age = 31
	fmt.Printf("pp1: %+v\n", *pp1)
}
