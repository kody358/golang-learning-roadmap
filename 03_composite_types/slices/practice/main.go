package main

import "fmt"

func main() {
	// ここに練習コードを書いてください
	// 実行: go run ./03_composite_types/slices/practice/

	// --- 練習1: スライスを作成してappendで要素を追加してみよう ---
	slc := []string{"Go", "Python", "Java"}

	fmt.Printf("配列: %v\n", slc);

	// --- 練習2: make()でlen, capを指定してスライスを作り、追加時のcapの変化を観察しよう ---
	ms := make([]int, 3, 5)
	fmt.Printf("%v len=%d cap=%d\n", ms, len(ms), cap(ms))

	// --- 練習3:スライシングで部分取得し、元のスライスと値を共有していることを確認しよう ---
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("元の配列", nums)
	fmt.Println("nums[2:4]", nums[2:4]);

	// --- 練習4: copyで独立したコピーを作り、元に影響しないことを確認しよう ---
	src := []string{"abc", "def", "ghi"}
	dst := make([]string, len(src))
	copy(dst, src)
	fmt.Println("copy", dst)

	// --- 練習5: appendで要素を削除してみよう ---
	del := []string{"aaa", "bbb", "ccc", "ddd"}
	i :=2

	del = append(del[:i], del[i+1:]...) // index2を飛ばしてappendしたものを作成する

	fmt.Println("index 2を削除", del)

}
