# todoApp

Go言語製のシンプルなToDo管理アプリケーションです。  
ユーザー管理・ToDoの作成/取得/更新/削除など、基本的な機能を備えています。

## 特徴

- Go1.13以上対応
- ディレクトリ構成分離（models, handlers, config, utils, router等）
- PostgreSQL対応（`github.com/lib/pq`使用）
- 設定ファイルから環境設定をロード可能
- ログファイル出力機能
- サンプルユーザーによるToDo CRUD実装例あり

## ディレクトリ構成（抜粋）

- `main.go` ... アプリケーションのエントリーポイント
- `models/` ... DBモデル（User, Todo等）の定義やDB操作
- `handlers/` ... HTTPリクエストのハンドリング
- `config/` ... 設定ファイル管理
- `router/` ... ルーティング
- `utils/` ... ユーティリティ（ロギング等）
- `go.mod`, `go.sum` ... Goモジュール管理

## 主要技術・ライブラリ

- Go標準
- [github.com/lib/pq](https://pkg.go.dev/github.com/lib/pq) (PostgreSQLドライバ)
- [github.com/go-ini/ini](https://pkg.go.dev/github.com/go-ini/ini) (設定ファイル)
- [github.com/google/uuid](https://pkg.go.dev/github.com/google/uuid)
- [github.com/stretchr/testify](https://pkg.go.dev/github.com/stretchr/testify) (テスト補助)

## 簡単な使い方（例）

1. 必要な設定ファイル`config/`を用意
2. PostgreSQLのセットアップ
3. 依存パッケージの取得
    ```sh
    go mod tidy
    ```
4. アプリケーションのビルドと実行
    ```sh
    go run main.go
    ```
5. コード例
    - ユーザーのToDo作成や取得、更新、削除が`main.go`に実装例として記載されています
