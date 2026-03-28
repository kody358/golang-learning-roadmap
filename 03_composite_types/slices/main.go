package main

import "fmt"

func main() {
	// ===== 配列 vs スライス =====
	fmt.Println("=== 配列 vs スライス ===")

	arr := [3]int{1, 2, 3} // 配列: サイズ固定、型に長さが含まれる
	slc := []int{1, 2, 3}  // スライス: サイズ可変

	fmt.Printf("配列: %v (型: %T)\n", arr, arr)
	fmt.Printf("スライス: %v (型: %T)\n", slc, slc)

	// ===== スライスの作成方法 =====
	fmt.Println("\n=== スライスの作成方法 ===")

	// リテラルで作成
	s1 := []int{10, 20, 30}
	fmt.Println("リテラル:", s1)

	// make()で作成: make(型, 長さ, 容量)
	s2 := make([]int, 3, 5) // 長さ3、容量5のスライス
	fmt.Printf("make([]int, 3, 5): %v len=%d cap=%d\n", s2, len(s2), cap(s2))

	// make()で長さだけ指定（容量=長さ）
	s3 := make([]int, 3)
	fmt.Printf("make([]int, 3):    %v len=%d cap=%d\n", s3, len(s3), cap(s3))

	// 配列からスライスを作成
	array := [5]int{1, 2, 3, 4, 5}
	s4 := array[1:4] // index 1〜3 を取り出す
	fmt.Printf("array[1:4]: %v\n", s4)

	// ===== append: 要素の追加 =====
	fmt.Println("\n=== append ===")

	s := []int{1, 2, 3}
	fmt.Printf("追加前: %v len=%d cap=%d\n", s, len(s), cap(s))

	s = append(s, 4)
	fmt.Printf("1つ追加: %v len=%d cap=%d\n", s, len(s), cap(s))

	s = append(s, 5, 6, 7)
	fmt.Printf("3つ追加: %v len=%d cap=%d\n", s, len(s), cap(s))

	// スライス同士を結合（...で展開）
	other := []int{8, 9}
	s = append(s, other...)
	fmt.Printf("結合後: %v\n", s)

	// ===== Capacity and Growth =====
	fmt.Println("\n=== 容量の自動拡張 ===")

	g := make([]int, 0)
	for i := 0; i < 10; i++ {
		g = append(g, i)
		fmt.Printf("len=%2d cap=%2d %v\n", len(g), cap(g), g)
	}
	// 容量が足りなくなると、約2倍に拡張される（小さいスライスの場合）

	// ===== スライシング（部分取得） =====
	fmt.Println("\n=== スライシング ===")

	nums := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("元:", nums)
	fmt.Println("nums[2:4]:", nums[2:4]) // index 2〜3
	fmt.Println("nums[:3]:", nums[:3])    // 先頭〜index 2
	fmt.Println("nums[3:]:", nums[3:])    // index 3〜末尾

	// ===== 背後の配列を共有する注意点 =====
	fmt.Println("\n=== 背後の配列の共有（重要） ===")

	original := []int{1, 2, 3, 4, 5}
	sub := original[1:3] // [2, 3]

	fmt.Println("original:", original)
	fmt.Println("sub:", sub)

	// subを変更するとoriginalも変わる！（同じ配列を指しているため）
	sub[0] = 99
	fmt.Println("sub[0]=99 に変更後:")
	fmt.Println("original:", original) // [1, 99, 3, 4, 5] に変わる
	fmt.Println("sub:", sub)

	// 独立したコピーを作りたい場合
	fmt.Println("\n=== copy: 独立したコピー ===")

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)

	dst[0] = 99
	fmt.Println("src:", src) // 変わらない
	fmt.Println("dst:", dst) // [99, 2, 3, 4, 5]

	// ===== 要素の削除 =====
	fmt.Println("\n=== 要素の削除 ===")

	del := []int{0, 1, 2, 3, 4}
	i := 2 // index 2 を削除したい

	del = append(del[:i], del[i+1:]...)
	fmt.Println("index 2を削除:", del) // [0, 1, 3, 4]

	// ===== nilスライスと空スライス =====
	fmt.Println("\n=== nilスライスと空スライス ===")

	var nilSlice []int          // nil（宣言のみ）
	emptySlice := []int{}       // 空（初期化済み）
	makeSlice := make([]int, 0) // 空（make）

	fmt.Printf("nil:   %v len=%d cap=%d nil?=%t\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("empty: %v len=%d cap=%d nil?=%t\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("make:  %v len=%d cap=%d nil?=%t\n", makeSlice, len(makeSlice), cap(makeSlice), makeSlice == nil)

	// どれもappendできる（nilスライスでも安全）
	nilSlice = append(nilSlice, 1)
	fmt.Println("nilにappend:", nilSlice)
}
