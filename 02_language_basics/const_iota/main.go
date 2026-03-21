package main

import "fmt"

// ===== const: 定数の定義 =====

const Pi = 3.14159
const AppName = "GoLearning"

// まとめて定義
const (
	StatusOK       = 200
	StatusNotFound = 404
)

// ===== iota: 自動連番 =====

// 基本: 0から始まる連番
type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

// 0をスキップして1から始める
type Role int

const (
	_     Role = iota // 0をスキップ
	Admin             // 1
	Editor            // 2
	Viewer            // 3
)

// ビットフラグ: 権限管理などに使う
type Permission int

const (
	Read    Permission = 1 << iota // 1  (001)
	Write                          // 2  (010)
	Execute                        // 4  (100)
)

// バイト単位: ファイルサイズの表現
const (
	_  = iota
	KB = 1 << (10 * iota) // 1024
	MB                     // 1048576
	GB                     // 1073741824
)

func main() {
	// --- const の確認 ---
	fmt.Println("=== const ===")
	fmt.Println("Pi:", Pi)
	fmt.Println("AppName:", AppName)
	fmt.Println("StatusOK:", StatusOK)
	fmt.Println("StatusNotFound:", StatusNotFound)

	// const は再代入できない（以下はコンパイルエラー）
	// Pi = 3.14

	// --- iota: 基本の連番 ---
	fmt.Println("\n=== iota: 曜日 ===")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Saturday:", Saturday)

	// --- iota: 0スキップ ---
	fmt.Println("\n=== iota: ロール（1から開始） ===")
	fmt.Println("Admin:", Admin)
	fmt.Println("Editor:", Editor)
	fmt.Println("Viewer:", Viewer)

	// --- iota: ビットフラグ ---
	fmt.Println("\n=== iota: ビットフラグ（権限） ===")
	fmt.Printf("Read:    %d (%03b)\n", Read, Read)
	fmt.Printf("Write:   %d (%03b)\n", Write, Write)
	fmt.Printf("Execute: %d (%03b)\n", Execute, Execute)

	// ビット演算で権限を組み合わせる
	readWrite := Read | Write
	fmt.Printf("Read|Write: %d (%03b)\n", readWrite, readWrite)

	// 権限チェック
	fmt.Println("readWriteにReadが含まれる?", readWrite&Read != 0)
	fmt.Println("readWriteにExecuteが含まれる?", readWrite&Execute != 0)

	// --- iota: バイト単位 ---
	fmt.Println("\n=== iota: バイト単位 ===")
	fmt.Println("1 KB =", KB, "bytes")
	fmt.Println("1 MB =", MB, "bytes")
	fmt.Println("1 GB =", GB, "bytes")
}
