# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

roadmap.sh の Go言語ロードマップ（`golang.pdf`参照）に沿った学習プロジェクト。番号付きディレクトリが各トピックに対応し、基礎から応用へ段階的に進む構成。

## ディレクトリ構成

`01_introduction/` から `14_advanced/` まで番号順に並び、各ディレクトリ内にトピック別のサブディレクトリがある。各トピックには実行可能な `main.go` や `*_test.go` を配置する。

## コマンド

```bash
# 特定の例を実行
go run ./03_control_flow/loops/main.go

# 特定トピックのテスト実行
go test ./11_testing/table_driven/...

# 全テスト実行
go test ./...

# ベンチマーク実行
go test -bench=. ./11_testing/benchmarks/...

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
