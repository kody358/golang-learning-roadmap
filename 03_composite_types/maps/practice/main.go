package main

import "fmt"

func main() {
	// 初期化
	ages := map[string]int{
		"Tanaka": 12,
		"Yamada": 20,
		"Yoshida": 9,
	}
	fmt.Println("ages:", ages)
	
	// makeで空マップ作成
	scores := make(map[string]int)
	fmt.Println(scores) // 空

	scores["math"] = 80
	scores["english"] = 60

	fmt.Println("scores:", scores)

	// キーの存在を確認
	val, ok := ages["Yamada"]
	fmt.Printf("Yamada: val=%d, ok=%v\n", val, ok)
}
