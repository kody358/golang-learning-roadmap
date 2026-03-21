package main

import "fmt"

// パッケージスコープ: パッケージ内のどこからでもアクセス可能
var packageVar = "パッケージスコープの変数"

func main() {
	// ===== スコープ =====
	fmt.Println("=== スコープ ===")
	fmt.Println("packageVar:", packageVar)

	// 関数スコープ: この関数内でのみアクセス可能
	funcVar := "関数スコープの変数"
	fmt.Println("funcVar:", funcVar)

	// ブロックスコープ: {}の中だけ
	if true {
		blockVar := "ブロックスコープの変数"
		fmt.Println("blockVar(if内):", blockVar)
		fmt.Println("funcVar(if内):", funcVar) // 外側の変数にはアクセスできる
	}
	// fmt.Println(blockVar) // コンパイルエラー: blockVarはifの外からは見えない

	// forループのスコープ
	for i := 0; i < 3; i++ {
		loopVar := i * 10
		_ = loopVar
	}
	// fmt.Println(i)       // コンパイルエラー: iはforの外からは見えない
	// fmt.Println(loopVar) // コンパイルエラー: loopVarもforの外からは見えない

	// ===== シャドウイング =====
	fmt.Println("\n=== シャドウイング ===")

	x := 10
	fmt.Println("外側のx:", x)

	if true {
		x := 99 // 新しいxが作られる（外側のxとは別物）
		fmt.Println("内側のx:", x)

		x = 50 // 内側のxを更新
		fmt.Println("内側のxを更新:", x)
	}

	fmt.Println("外側のx（変わらない）:", x) // 10のまま

	// ===== シャドウイングを避ける場合 =====
	fmt.Println("\n=== シャドウイングを避ける（=で代入） ===")

	y := 10
	fmt.Println("外側のy:", y)

	if true {
		y = 99 // := ではなく = を使うと外側のyを更新する
		fmt.Println("内側のy:", y)
	}

	fmt.Println("外側のy（更新された）:", y) // 99に変わる

	// ===== よくあるバグ: errのシャドウイング =====
	fmt.Println("\n=== よくあるバグ: errのシャドウイング ===")

	err := doTask("タスク1")
	fmt.Println("タスク1後のerr:", err)

	if err == nil {
		// := で新しいerrが作られてしまう（外側のerrは更新されない）
		result, err := doTaskWithResult("タスク2")
		fmt.Println("タスク2 result:", result)
		fmt.Println("タスク2 err(内側):", err) // "タスク2失敗"
	}

	fmt.Println("タスク2後のerr(外側):", err) // nilのまま！バグの原因になる

	// 正しいやり方: 事前に変数を宣言しておく
	fmt.Println("\n=== 正しいやり方 ===")

	var err2 error
	var result string

	err2 = doTask("タスク1")
	fmt.Println("タスク1後のerr2:", err2)

	if err2 == nil {
		result, err2 = doTaskWithResult("タスク2") // = を使う
		fmt.Println("タスク2 result:", result)
	}

	fmt.Println("タスク2後のerr2(正しく更新):", err2) // "タスク2失敗"
}

func doTask(name string) error {
	fmt.Printf("  %s: 成功\n", name)
	return nil
}

func doTaskWithResult(name string) (string, error) {
	fmt.Printf("  %s: 失敗\n", name)
	return "", fmt.Errorf("%s失敗", name)
}
