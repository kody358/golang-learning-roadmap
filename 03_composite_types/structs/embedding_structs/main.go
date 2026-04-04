package main

import "fmt"

// ===== 構造体の埋め込み（Embedding） =====

// Goには継承がない代わりに「埋め込み（Embedding）」で構造体を合成する
// 埋め込まれた構造体のフィールドやメソッドが昇格（promote）される

// --- 基本的な埋め込み ---

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Speak() string {
	return fmt.Sprintf("%sは鳴きます", a.Name)
}

func (a Animal) Info() string {
	return fmt.Sprintf("名前: %s, 年齢: %d", a.Name, a.Age)
}

// Dog は Animal を埋め込む（フィールド名を省略して型名だけ書く）
type Dog struct {
	Animal // 埋め込み（匿名フィールド）
	Breed  string
}

// Dog独自のメソッド
func (d Dog) Fetch() string {
	return fmt.Sprintf("%sはボールを取ってきます", d.Name) // d.Animal.Name と同じ
}

// メソッドのオーバーライド（同名メソッドを定義すると上書きできる）
func (d Dog) Speak() string {
	return fmt.Sprintf("%sはワンワンと鳴きます", d.Name)
}

// --- 複数の構造体を埋め込む ---

type Engine struct {
	Horsepower int
	Type       string
}

func (e Engine) Start() string {
	return fmt.Sprintf("%dHPの%sエンジンが始動しました", e.Horsepower, e.Type)
}

type Chassis struct {
	Material string
	Weight   int // kg
}

func (c Chassis) Description() string {
	return fmt.Sprintf("%s製シャーシ（%dkg）", c.Material, c.Weight)
}

// Car は Engine と Chassis の両方を埋め込む
type Car struct {
	Engine
	Chassis
	Brand string
}

// --- フィールド名の衝突 ---

type A struct {
	Value string
}

type B struct {
	Value string // A と同じフィールド名
}

type C struct {
	A
	B
}

// --- インターフェースの暗黙的な実装 ---

type Stringer interface {
	String() string
}

type Product struct {
	Name  string
	Price int
}

func (p Product) String() string {
	return fmt.Sprintf("%s（¥%d）", p.Name, p.Price)
}

// GiftBox は Product を埋め込むことで Stringer を暗黙的に満たす
type GiftBox struct {
	Product
	Wrapping string
}

func printStringer(s Stringer) {
	fmt.Printf("  Stringer: %s\n", s.String())
}

func main() {
	// --- 基本的な埋め込み ---
	fmt.Println("=== 基本的な埋め込み ===")

	dog := Dog{
		Animal: Animal{Name: "ポチ", Age: 3},
		Breed:  "柴犬",
	}

	// 埋め込まれたフィールドに直接アクセス（昇格）
	fmt.Printf("名前: %s\n", dog.Name)       // dog.Animal.Name と同じ
	fmt.Printf("年齢: %d\n", dog.Age)         // dog.Animal.Age と同じ
	fmt.Printf("犬種: %s\n", dog.Breed)

	// 埋め込まれたメソッドの呼び出し
	fmt.Println(dog.Info())   // Animal.Info() が昇格
	fmt.Println(dog.Fetch())  // Dog独自のメソッド

	// メソッドのオーバーライド
	fmt.Println(dog.Speak())         // Dog.Speak() が呼ばれる
	fmt.Println(dog.Animal.Speak())  // 明示的に Animal.Speak() を呼ぶ

	// --- 複数の埋め込み ---
	fmt.Println("\n=== 複数の埋め込み ===")

	car := Car{
		Engine:  Engine{Horsepower: 200, Type: "ハイブリッド"},
		Chassis: Chassis{Material: "アルミニウム", Weight: 1500},
		Brand:   "トヨタ",
	}

	fmt.Printf("ブランド: %s\n", car.Brand)
	fmt.Println(car.Start())       // Engine.Start() が昇格
	fmt.Println(car.Description()) // Chassis.Description() が昇格

	// 明示的にアクセスすることもできる
	fmt.Printf("馬力: %dHP\n", car.Engine.Horsepower)
	fmt.Printf("重量: %dkg\n", car.Chassis.Weight)

	// --- フィールド名の衝突 ---
	fmt.Println("\n=== フィールド名の衝突 ===")

	c := C{
		A: A{Value: "Aの値"},
		B: B{Value: "Bの値"},
	}

	// c.Value はあいまいなのでコンパイルエラーになる
	// fmt.Println(c.Value) // エラー: ambiguous selector c.Value

	// 明示的に指定すればOK
	fmt.Printf("A.Value: %s\n", c.A.Value)
	fmt.Printf("B.Value: %s\n", c.B.Value)

	// --- 埋め込みによるインターフェースの実装 ---
	fmt.Println("\n=== 埋め込みによるインターフェース実装 ===")

	gift := GiftBox{
		Product:  Product{Name: "チョコレート", Price: 3000},
		Wrapping: "赤いリボン",
	}

	// GiftBox は Product の String() を昇格で持つので Stringer を満たす
	printStringer(gift)
	fmt.Printf("ラッピング: %s\n", gift.Wrapping)
	fmt.Printf("商品名: %s, 価格: ¥%d\n", gift.Name, gift.Price)

	// --- ポインタの埋め込み ---
	fmt.Println("\n=== ポインタの埋め込み ===")

	type Base struct {
		ID int
	}

	type Derived struct {
		*Base // ポインタで埋め込む
		Name  string
	}

	base := &Base{ID: 42}
	d1 := Derived{Base: base, Name: "d1"}
	d2 := Derived{Base: base, Name: "d2"}

	fmt.Printf("d1.ID: %d, d2.ID: %d\n", d1.ID, d2.ID)

	// ポインタ埋め込みなので、同じBaseを共有している
	base.ID = 100
	fmt.Printf("Base変更後 → d1.ID: %d, d2.ID: %d\n", d1.ID, d2.ID)
}
