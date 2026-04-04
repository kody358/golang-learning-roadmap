package main

import "fmt"

type Animal struct {
	Name string
	Age int
}

func (a Animal) Speak() string {
	return fmt.Sprintf("%sは鳴きます", a.Name)
}

func (a Animal) Info() string {
	return fmt.Sprintf("名前: %s, 年齢: %d", a.Name, a.Age)
}

type Dog struct {
	Animal // 埋め込み
	Breed string
}

// Dog独自のメソッド
func (d Dog) Fetch() string {
	return fmt.Sprintf("%sはボールをとってきます", d.Name)
}

// メソッドのオーバーライド
func (d Dog) Speak() string {
	return fmt.Sprintf("%sはワンワンと鳴きます", d.Name)
}

func main() {
	dog := Dog {
		Animal: Animal{
			Name: "ポチ",
			Age: 3,
		},
		Breed: "柴犬",
	}

	// 昇格
	fmt.Printf("名前: %s\n", dog.Name)
	fmt.Printf("年齢: %d\n", dog.Age)
	fmt.Printf("犬種: %s\n", dog.Breed)
	
	// 昇格で埋め込まれたAnimalのメソッド呼び出し
	fmt.Println(dog.Info())
}
