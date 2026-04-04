package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// ===== 構造体タグとJSON =====

// 構造体タグはバッククォートで囲んだメタデータ
// encoding/json パッケージがタグを読み取ってフィールド名をマッピングする

// 基本的なJSONタグ
type User struct {
	ID       int    `json:"id"`                 // JSONでは "id" というキー名になる
	Name     string `json:"name"`               // JSONでは "name"
	Email    string `json:"email"`              // JSONでは "email"
	Password string `json:"-"`                  // "-" はJSONに含めない（機密情報）
	Age      int    `json:"age,omitempty"`       // omitempty: ゼロ値なら省略
	Nickname string `json:"nickname,omitempty"` // ゼロ値（空文字列）なら省略
}

// ネストされた構造体のJSON
type Article struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// スライスやマップを含む構造体
type Response struct {
	Status  string            `json:"status"`
	Code    int               `json:"code"`
	Data    []string          `json:"data"`
	Headers map[string]string `json:"headers,omitempty"`
}

func main() {
	// --- 構造体 → JSON（マーシャリング） ---
	fmt.Println("=== Marshal（構造体 → JSON） ===")

	user := User{
		ID:       1,
		Name:     "田中太郎",
		Email:    "tanaka@example.com",
		Password: "secret123", // json:"-" なのでJSONに含まれない
		Age:      30,
		Nickname: "",          // omitemptyなのでゼロ値は省略される
	}

	// json.Marshal: 構造体をJSONバイト列に変換
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))

	// json.MarshalIndent: 整形されたJSON
	prettyJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(prettyJSON))

	// --- JSON → 構造体（アンマーシャリング） ---
	fmt.Println("\n=== Unmarshal（JSON → 構造体） ===")

	jsonStr := `{
		"id": 2,
		"name": "佐藤花子",
		"email": "sato@example.com",
		"age": 25,
		"nickname": "はなちゃん"
	}`

	var user2 User
	err = json.Unmarshal([]byte(jsonStr), &user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user2: %+v\n", user2)
	fmt.Printf("Password（JSONにないのでゼロ値）: %q\n", user2.Password)

	// --- ネストされた構造体のJSON ---
	fmt.Println("\n=== ネストされた構造体 ===")

	article := Article{
		Title: "Go言語入門",
		Body:  "Goは素晴らしい言語です。",
		Author: Author{
			Name:  "山田三郎",
			Email: "yamada@example.com",
		},
	}

	articleJSON, _ := json.MarshalIndent(article, "", "  ")
	fmt.Println(string(articleJSON))

	// --- 未知のJSONフィールドの扱い ---
	fmt.Println("\n=== 未知のフィールド ===")

	// 構造体にないフィールドは無視される
	unknownJSON := `{"id": 3, "name": "鈴木", "email": "suzuki@example.com", "unknown_field": "無視される"}`
	var user3 User
	json.Unmarshal([]byte(unknownJSON), &user3)
	fmt.Printf("user3: %+v\n", user3)

	// --- map[string]interface{} で動的JSON ---
	fmt.Println("\n=== 動的JSON（map使用） ===")

	dynamicJSON := `{"key1": "文字列", "key2": 42, "key3": true, "key4": [1, 2, 3]}`
	var result map[string]interface{}
	json.Unmarshal([]byte(dynamicJSON), &result)
	for k, v := range result {
		fmt.Printf("  %s: %v (型: %T)\n", k, v, v)
	}

	// --- スライス・マップを含む構造体 ---
	fmt.Println("\n=== スライス・マップを含む構造体 ===")

	resp := Response{
		Status: "ok",
		Code:   200,
		Data:   []string{"item1", "item2", "item3"},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(respJSON))
}
