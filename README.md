# todoApp

Go言語製のシンプルなToDo管理アプリケーションです。  
本アプリは、APIサーバーとして高パフォーマンスかつ保守性の高いGo言語を採用し、信頼性・拡張性に優れたPostgreSQLをデータストアとして利用しています。
Goは並行処理やAPI開発に強みがあり、PostgreSQLはトランザクション管理や多様なデータ型に対応しているため、今後の機能拡張や実運用にも耐えうる設計となっています

---

## 概要・目的

- RESTfulなAPI設計でToDoおよびユーザー情報の管理を実現します
- シンプルな構成で拡張・カスタマイズしやすい設計を意識しています。

---

## 主な機能

- ユーザー作成・取得・更新・削除
- ToDo作成・取得・更新・削除（ユーザー単位で管理）
- CORS（クロスオリジンリクエスト）対応（フロントエンドと連携しやすい）
- PostgreSQLによるデータ永続化
- 設定ファイルによる環境切替
- ログファイル出力
- ディレクトリ分割による保守性・可読性の向上

---

## インストール方法

1. リポジトリをクローン

    ```sh
    git clone https://github.com/taatolu/todoApp.git
    cd todoApp
    ```

2. 設定ファイルの準備  
   `config/`ディレクトリに`.ini`形式の設定ファイル（例：`product.ini`）を用意します。

3. PostgreSQLのセットアップ  
   - データベースとユーザーを作成
   - 設定ファイルに接続情報を記載

4. 依存パッケージの取得

    ```sh
    go mod tidy
    ```

5. アプリケーションのビルドと起動

    ```sh
    go run main.go
    ```
    デフォルトでは `:8080` ポートでAPIサーバーが起動します。

---

## 🐳 Docker Composeによる開発環境構築 
本アプリは、Docker Composeを利用してGo APIサーバー・PostgreSQL・フロントエンドを一括で立ち上げることができます。

### 1. .envファイルの作成
プロジェクトルートに.envファイルを作成し、以下の内容を記載してください（値は任意で変更してください）。
```:env
DB_PASSWORD=your_db_password
DB_USER=your_db_user
DB_DBNAME=your_db_name
```
※.envファイルは機密情報を含むため、必ず.gitignoreに追加してください。

### 2. サービスの起動
以下のコマンドで、全サービス（API, DB, Web）が一括で起動します。
```
//sh
docker compose up --build
```
初回はイメージのビルドやDBセットアップのため、数分かかる場合があります。

### 3. サービスへのアクセス

| サービス | アクセスURL | 備考 |
|:---|:---|:---|
| フロントエンド | http://localhost:3000 | React等の場合 |
| APIサーバー | http://localhost:8080 | Go製APIエンドポイント |
| データベース | localhost:5432 | PostgreSQL（外部接続用） |

### 4. サービスの停止
```
//sh
docker compose down
```

### 5. ディレクトリ構成例
```Code
todoApp/
├── api/         # Go APIアプリ
├── web/         # フロントエンド
├── docker-compose.yml
├── .env
└── ...
```
### 6. 注意事項
DB_HOSTはdbで固定されています（Docker Compose内のサービス名で自動的に名前解決されます）。
DBデータはdb_dataボリュームに永続化されます。
.envファイルが無い場合、DB接続に失敗します。

## 使い方

### API エンドポイント例

#### ユーザー関連

- ユーザー新規作成  
  `POST /api/v1/users`  
  ```json
  {
    "username": "yourname",
    "email": "your@email.com",
    "password": "yourpassword"
  }
  ```

- ユーザー情報取得・更新・削除  
  `GET/PUT/DELETE /api/v1/users/{user_id}`

#### ToDo関連

- ToDo新規作成  
  `POST /api/v1/todos?user_id={user_id}`  
  ```json
  {
    "content": "やること"
  }
  ```

- ToDo一覧取得  
  `GET /api/v1/todos?user_id={user_id}`

- ToDo個別取得・更新  
  `GET/PUT /api/v1/todos/{todo_id}`

- ToDo削除  
  `DELETE /api/v1/todos/{todo_id}`

### CORS対応について

- `router/cors.go` にて `Access-Control-Allow-Origin` などを設定し、`localhost:3000` からのAPI呼び出しを許可しています。  
  フロントエンド（React等）との連携時も追加設定不要で利用可能です。

---

## 開発環境・依存ライブラリ

- Go 1.13以上
- PostgreSQL
- 主要パッケージ
    - [github.com/lib/pq](https://pkg.go.dev/github.com/lib/pq)（PostgreSQLドライバ）
    - [github.com/go-ini/ini](https://pkg.go.dev/github.com/go-ini/ini)（設定ファイル操作）
    - [github.com/google/uuid](https://pkg.go.dev/github.com/google/uuid)（UUID生成・ハッシュ化）
    - [github.com/stretchr/testify](https://pkg.go.dev/github.com/stretchr/testify)（テスト補助）

---

## ディレクトリ構成（抜粋）

- `main.go` ... エントリーポイント
- `models/` ... DBモデル＆操作（User, Todo等）
- `handlers/` ... 各APIハンドラ
- `config/` ... 設定ファイル管理
- `router/` ... ルーティング・CORSミドルウェア等
- `utils/` ... ロギング等ユーティリティ
- `go.mod`, `go.sum` ... Goモジュール管理

---

## 補足

- サンプルユーザーやToDoのCRUD例も `main.go` 内に記載
- APIの詳細やカスタマイズ方法はソースコードを参照ください

---

## ライセンス

MIT
