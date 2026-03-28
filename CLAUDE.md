# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

roadmap.sh の Go言語ロードマップ（`golang.pdf`参照）に沿った学習プロジェクト。番号付きディレクトリが各トピックに対応し、基礎から応用へ段階的に進む構成。

## ディレクトリ構成

`01_introduction/` から `16_advanced/` まで番号順に並び、各ディレクトリ内にトピック別のサブディレクトリがある。
PDFのロードマップに対応する番号は以下の通り:

1. Introduction to Go
2. Language Basics (Variables & Constants, Data Types)
3. Composite Types (Arrays, Slices, Strings, Maps, Structs)
4. Control Flow (Conditionals, Loops)
5. Functions
6. Pointers
7. Methods and Interfaces
8. Generics
9. Error Handling
10. Code Organization
11. Concurrency
12. Standard Library
13. Testing & Benchmarking
14. Ecosystem & Popular Libraries
15. Go Toolchain and Tools
16. Advanced Topics

- 各トピックは独立したサブディレクトリに分ける（例: `02_language_basics/const_iota/`, `02_language_basics/scope_shadowing/`）
- 各サブディレクトリに `package main` の `main.go` を置き、`go run` で単独実行できるようにする
- 1つのサブディレクトリに複数トピックを混ぜない
- Claudeが作成する `main.go` は参考コード。ユーザーが練習用コードを書くために各トピックディレクトリ内に `practice/` サブディレクトリを作成する
  - 例: `02_language_basics/const_iota/practice/main.go`
  - `practice/` も `package main` で `go run` で単独実行できるようにする

## コマンド

```bash
# 特定の例を実行
go run ./03_composite_types/slices/main.go

# 特定トピックのテスト実行
go test ./13_testing/table_driven/...

# 全テスト実行
go test ./...

# ベンチマーク実行
go test -bench=. ./13_testing/benchmarks/...

# コードフォーマット
go fmt ./...

# 静的解析
go vet ./...
```

## 規約

- ルートに単一の `go.mod` を置くモノレポ構成
- 各トピックディレクトリは自己完結的で単独実行可能にする
- サンプルコードは対象の概念に焦点を絞り簡潔に書く
- 学習者は日本語話者のため、コメントは日本語で記述する
