package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 構造体JSONタグ
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func main() {
	user := User {
		ID: 1,
		Name: "山田太郎",
		Email: "yamada@example.com",
		Password: "password12345",
	}
	
	// jsonバイト列変換
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))

	// 整形されたjson
	prettyJson, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(prettyJson))
	

	// json -> 構造体（アンマーシャル）
	jsonStr := `{
		"id": 2,
		"name": "田中太郎",
		"email": "tanaka@example.com",
		"password": "password12345"
	}`

	var user2 User
	err = json.Unmarshal([]byte(jsonStr), &user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user2 %T: %+v\n", user2, user2)
}
